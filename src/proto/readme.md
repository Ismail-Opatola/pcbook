# gRPC, Golang, Java

## To generate Golang codes from your proto

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

### Create a Makefile to simplify generating proto codes

- define task `gen:` `<generate protocode code>`
- define task `clean:` `<the build/gen file whenever we want>`
- define task `run:` to run the `main.go` file as well

### NOTE: IMPORT ERROR ISSUE

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

## Java Setup

### Install Java Dev Kit (JDK)

check to see if you have java installed. We would require at least v8 to use all features of gRPC. Run `javac` or `javac --version`

If you don't have it installed, you would be prompted to install the LTS version

**NEXT:**

### Install IntelliJ IDEA

google _intelliJ IDEA for java_ and download the free community version

<https://github.com/google/protobuf-gradle-plugin>

copy the below snippet on the github page and paste in our `build.gradle` file

      plugins {
        id "com.google.protobuf" version "0.8.13"
        id "java"
      }

**NEXT:** Add dependency, search for `maven protobuf java`

Maven repository <https://mvnrepository.com/artifact/com.google.protobuf/protobuf-java>

Under the `General` tab click on the latest version > click on the `gradle` tab and copy the code. Paste it into the dependency block in your `build.gradle` file.

**NEXT:** We would need a package to work with `gRPC`.

Go back to the maven repository, search for `grpc-all`.

Select the latest version, copy the gradle config and paste in your `build.gradle` dependency block.

<https://github.com/google/protobuf-gradle-plugin>
copy the below snippet on the github page and paste in our `build.gradle` file, after the dependency block

Customizing Protobuf compilation
The plugin adds a protobuf block to the project. It provides all the configuration knobs.

Locate external executables
By default the plugin will search for the protoc executable in the system search path. We recommend you to take the advantage of pre-compiled protoc that we have published on Maven Central:

      dependencies {
            ...
      }

      protobuf {
            ...
            // Configure the protoc executable
            protoc {
                  // Download from repositories
                  // artifact = 'com.google.protobuf:protoc:<enter-lts-version>'
                  // goto https://mvnrepository.com/artifact/com.google.protobuf/protoc, copy the lts version and use it here
                  artifact = 'com.google.protobuf:protoc:4.0.0'
            }
            ...
      }

**NEXT:** Tell `protobuf` to use `protoc-gen-grpc-java` plugin, copy snippet from same source as above and check for its lst version on maven repository

**NEXT:** in the protobuf block, add this snippet

    generateProtoTasks {
        all()*.plugins {
            grpc {}
        }
    }

**NEXT:** Tell intelij where our genrated code would be located, so it can correctly perform code analyses and succession later

    sourceSets {
        main {
            java {
                srcDirs 'build/generated/source/proto/main/grpc'
                srcDirs 'build/generated/source/proto/main/java'
            }
        }
    }

### Setup a gradle project

#### Install plugins for protobuf and gRPC

install protocol gradle plugin from google. google `protobuf java gradle`

#### Config automatic code generation upon build

### Use protobuf options to customise generated java codes

**NETX:** In your **Java project**, go to `src/main` folder > create a `proto` folder | copy-over all our `proto/*.proto` files in the Golang project to the folder

## create/initialize a go module

run `go mod init github.com/techschool/pcbook`

create a smaple package `mkdir sample` and `touch sample/generator.go`
