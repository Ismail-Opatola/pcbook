gen:
	protoc --proto_path=src/proto --go_out=build/gen --go_opt=paths=source_relative src/proto/*.proto

run:
	go run src/main.go

clean:
	powershell clear
	powershell rm build/gen/*.go	

# Read me
# https://dev.to/deciduously/how-to-make-a-makefile-1dep