LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-auth-api
	make generate-chat-server-api

generate-auth-api:
	mkdir -p auth/pkg/auth
	protoc --proto_path auth/api/auth \
	--go_out=auth/pkg/auth --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=auth/pkg/auth --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	auth/api/auth/auth.proto

generate-chat-server-api:
	mkdir -p chat_server/pkg/chat_server
	protoc --proto_path chat_server/api/chat_server \
	--go_out=chat_server/pkg/chat_server --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=chat_server/pkg/chat_server --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	chat_server/api/chat_server/chat_server.proto