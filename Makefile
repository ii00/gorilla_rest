gen:
	protoc --proto_path=./ --proto_path=proto --go_out=pb proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go

tidy:
	go mod tidy