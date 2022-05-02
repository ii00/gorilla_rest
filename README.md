# grpc_app

> protoc --proto_path=proto proto/*.proto --go_out=pb

> go test -timeout 30s -run ^TestFileSerializer$ grpc_app/serializer