# To generate go codes from your proto

- Intall protocol buffer compiler -- (protobuf) or (protoc)

      choco install protoc

- goto <https://grpc.io/docs/languages/go/quickstart/>

  - Ensure Prerequisites are installed
  - Install go plugins for the protocol compiler

        set GO111MODULE=on  # Enable module mode
        $ go get github.com/golang/protobuf/protoc-gen-go
        $ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
        $ set PATH="$PATH:$(go env GOPATH)/bin"

  - To compile try example:

        protoc --proto_path=src --go_out=build/gen --go_opt=paths=source_relative src/foo.proto src/bar/baz.proto

  - Run it like this from root dir

        mkdir build
        mkdir build/gen

        protoc --proto_path=src/proto --go_out=build/gen --go_opt=paths=source_relative src/proto/*.proto

  - Or Like this from `src`

        mkdir build
        mkdir build/gen
        protoc --proto_path=proto proto/_.proto --go_out=plugins=grpc:build/gen

  - Specify `go_package` option in `processor_message.proto`

        option go_package = "example.com/foo/bar";

  **_source: <https://developers.google.com/protocol-buffers/docs/reference/go-generated#package>_**

## Create a Makefile to simplify generating proto codes

- define task `gen:` `<generate protocode code>`
- define task `clean:` `<the build/gen file whenever we want>`
- define task `run:` to run the `main.go` file as well

## NOTE: IMPORT ERROR ISSUE

Tried to import `memory_message.proto` in `processor_message.proto`

Throws error warning in the editor when using absolute path `"memory_message.proto"` but compiles succesfully

Error fixed with relative path `"pcbook/src/proto/memory_message.proto"` but fails when compiled.

**SOLUTION:**
vscode > extension > vscode-proto3 | scroll-down and copy config snippet

      {
      "protoc": {
            "path": "/path/to/protoc",
            "compile_on_save": false,
            "options": [
                  "--proto_path=protos/v3",
                  "--proto_path=protos/v2",
                  "--proto_path=${workspaceRoot}/proto",
                  "--proto_path=${env.GOPATH}/src",
                  "--java_out=gen/java"
            ]
      }
      }

**NEXT:**
vscode - preference > settings | enter search "protoc", click `Edit in Setting.json` paste the config to get the path, run `which protoc` in terminal. copy the path and paste in the setting.json `protoc {path: ""}`

snippet

      "protoc": {
      "path": "C:\\ProgramData\\chocolatey\\bin\\protoc.EXE",
      "compile_on_save": false,
      "options": ["--proto_path=proto"]
      }

**NEXT:**
Make sure you have clang-format extension installed in vscode, then install clang-format on pc

      choco install llvm

In proto file, right click and select `Format document` option, set clang-format as document formatter
