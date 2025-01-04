package service

import (
	"context"
	"os"
	"testing"
)

func getOpenAiApiKey() string {
	return os.Getenv("CHATBOT_OPENAI_API_KEY")
}

func getFakeInstanceLogs() string {
	return `
(gdb) run $(python3 -c 'print("A" * 32)')
Starting program: ./vulnerable_program AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

Program received signal SIGSEGV, Segmentation fault.
0x41414141 in ?? ()

(gdb) info func
All defined functions:

Non-debugging symbols:
0x08048350  _init
0x08048380  printf@plt
0x08048390  __libc_start_main@plt
0x080483a0  strcpy@plt
0x08048420  _start
0x08048450  main
0x08048490  vulnerable_function
0x08048500  _fini
0x08048510  __data_start
0x08048510  __bss_start
0x08048514  _edata
0x08048518  _end
0x08048518  __bss_end__

(gdb) disas _start
Dump of assembler code for function _start:
   0x08048420 <+0>:   xor    ebp,ebp
   0x08048422 <+2>:   mov    ecx,esp
   0x08048424 <+4>:   and    esp,0xfffffff0
   0x08048427 <+7>:   push   eax
   0x08048428 <+8>:   push   esp
   0x08048429 <+9>:   call   0x8048390 <__libc_start_main@plt>
   0x0804842e <+14>:  hlt    
End of assembler dump.

(gdb) bt
#0  0x41414141 in ?? ()
#1  0x00000000 in ?? ()

(gdb) info registers
eax            0x0          0
ebx            0x0          0
ecx            0xbffff0c0   0xbffff0c0
edx            0x0          0
esi            0x0          0
edi            0x0          0
ebp            0x41414141   0x41414141
esp            0xbffff0c0   0xbffff0c0
eip            0x41414141   0x41414141
eflags         0x10206      [ PF IF RF ]
cs             0x23         35
ss             0x2b         43
ds             0x2b         43
es             0x2b         43
fs             0x0          0
gs             0x63         99

(gdb) x/16x $esp
0xbffff0c0: 0x41414141  0x41414141  0x41414141  0x41414141
0xbffff0d0: 0x41414141  0x41414141  0x41414141  0x41414141
0xbffff0e0: 0x00000000  0x00000000  0x00000000  0x00000000
0xbffff0f0: 0x08048450  0xb7e2b370  0xbffff118  0x00000000
`
}

func TestChatbotService(t *testing.T) {
	openAiApiKey := getOpenAiApiKey()
	chatbotService := NewChatbotService(context.Background(), openAiApiKey, nil)

	t.Run("SuccessfulCompletion", func(t *testing.T) {
		response, err := chatbotService.ExecuteCompletion(context.Background(), getFakeInstanceLogs(), "What is the vulnerability?")
		if err != nil {
			t.Errorf("ExecuteCompletion() error = %v", err)
		}

		if response == "" {
			t.Error("ExecuteCompletion() returned empty response")
		}

		t.Logf("ExecuteCompletion() response = %v", response)
	})
}
