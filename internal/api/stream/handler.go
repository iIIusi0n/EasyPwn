package stream

import (
	"context"
	"easypwn/internal/data"
	"easypwn/internal/pkg/instance"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocketSession(c *gin.Context, ins *instance.Instance, logging bool, command ...string) {
	ctx := context.Background()
	db := data.GetDB()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}
	defer conn.Close()

	inout, err := ins.Execute(ctx, command...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute command"})
		return
	}
	defer inout.Writer.Write([]byte("exit\n"))

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		buf := make([]byte, 1024)
		for {
			n, err := inout.Reader.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Println("Failed to read from gdb:", err)
				}
				return
			}

			if logging {
				ins.WriteLog(ctx, db, string(buf[:n]))
			}

			if n > 0 {
				if err := conn.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
					log.Println("Failed to write to websocket:", err)
					return
				}
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("WebSocket error: %v", err)
				}
				return
			}

			if messageType == websocket.BinaryMessage {
				if _, err := inout.Writer.Write(message); err != nil {
					log.Printf("Failed to write to shell: %v", err)
					return
				}
			} else if messageType == websocket.TextMessage {
				type ResizeMessage struct {
					Type string `json:"type"`
					Cols uint   `json:"cols"`
					Rows uint   `json:"rows"`
				}
				var resizeMsg ResizeMessage
				if err := json.Unmarshal(message, &resizeMsg); err == nil && resizeMsg.Type == "resize" {
					ins.ResizeTTY(ctx, inout.ExecID, resizeMsg.Rows, resizeMsg.Cols)
				}
			}
		}
	}()

	wg.Wait()

	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	conn.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Session closed"})
}

func GetDebuggerSessionHandler() gin.HandlerFunc {
	db := data.GetDB()
	ctx := context.Background()

	return func(c *gin.Context) {
		ins, err := instance.GetInstance(ctx, db, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get instance"})
			return
		}
		handleWebSocketSession(c, ins, true, "gdb", c.MustGet("full_path").(string))
	}
}

func GetShellSessionHandler() gin.HandlerFunc {
	db := data.GetDB()
	ctx := context.Background()

	return func(c *gin.Context) {
		ins, err := instance.GetInstance(ctx, db, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get instance"})
			return
		}
		handleWebSocketSession(c, ins, false, "/bin/bash")
	}
}
