gen:
	protoc --proto_path=src/proto --go_out=build/gen --go_opt=paths=source_relative src/proto/*.proto

run:
	go run src/main.go

clean:
	rm build/gen/*.go