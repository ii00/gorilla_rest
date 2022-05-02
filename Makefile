gen:
	protoc --proto_path=./ --proto_path=proto --go_out=pb proto/*.proto

clean:
	rm pb/*.go

tidy:
	go mod tidy	

run:
	go run main.go

test:
	go test -cover -race ./...