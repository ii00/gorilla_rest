gen:
	protoc --proto_path=./ --go_out=pb --go-grpc_out=pb proto/*.proto

clean:
	rm pb/*.go

tidy:
	go mod tidy	

run:
	go run main.go

test:
	go test -cover -race ./...