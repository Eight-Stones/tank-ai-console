PB_REL="https://github.com/protocolbuffers/protobuf/releases"
PB_VERSION="25.1"
PB_NAME="protoc-$PB_VERSION-linux-x86_64"

curl -LO "$PB_REL/download/v$PB_VERSION/$PB_NAME.zip"
unzip -o "$PB_NAME.zip" -d "$HOME/.local"
rm "$PB_NAME.zip"

go install github.com/bufbuild/buf/cmd/buf@v1.42.0

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

go install github.com/go-swagger/go-swagger/cmd/swagger@latest

go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

go install github.com/pressly/goose/v3/cmd/goose@latest

go install github.com/fe3dback/go-arch-lint@latest