.PHONY: proto images test 

# generate proto files
proto:
	protoc --go_out=internal \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal \
		--go-grpc_opt=paths=source_relative \
		api/*.proto

# build images for docker compose deployment
images:
	for dir in cmd/*; do \
		if [ -d "$$dir" ]; then \
			service=$$(basename $$dir); \
			docker build -f cmd/$$service/Dockerfile -t easypwn/$$service .; \
		fi \
	done
	docker build -f web/Dockerfile -t easypwn/web .

# run tests
test:
	godotenv -f .env.local go test -v ./... -coverprofile=coverage.out

# run tests for a specific function
testfunc:
	@if [ -z "$(func)" ]; then \
		echo "Usage: make testfunc func=TestFunctionName"; \
		exit 1; \
	fi
	go test -v ./... -run $(func)