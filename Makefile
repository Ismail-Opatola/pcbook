gen:
	protoc --proto_path=src/proto --go_out=pb --go_opt=paths=source_relative src/proto/*.proto

clean:
	powershell clear
	powershell rm pb/*.go	

run:
	go run src/main.go

# -cover _messure the code coverage of our test_
# -race _detect any racing condition_
test:
	go test -cover -race ./...

sizeof:
	dir src\\tmp\\*

# Read me
# https://dev.to/deciduously/how-to-make-a-makefile-1dep