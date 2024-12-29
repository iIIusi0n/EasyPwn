.PHONY: proto images test

proto:
	protoc --go_out=internal \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal \
		--go-grpc_opt=paths=source_relative \
		api/*.proto

images:
	for dir in cmd/*; do \
		if [ -d "$$dir" ]; then \
			service=$$(basename $$dir); \
			docker build -f cmd/$$service/Dockerfile -t easypwn/$$service .; \
		fi \
	done

test:
	go test -v ./...