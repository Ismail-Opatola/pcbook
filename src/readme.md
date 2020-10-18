# What the heck is protocol buffers

Protocol buffers are representations of structured data.

## Advantage

- Serializtion and Deserialization of data
- Efficient size compared to JSON
- Strongly typed
- Language agnostic

## Definning protocol buffers

    // employees.proto file

    syntax = "proto3"

    message Employee {
    int32 id = 1;
    string name = 2;
    float slalary = 3;
    }

    message Employees {
    repeated Employee employees = 1;
    }

You're responsible base on the language of choice to convert this proto file into any language of choice (_which supports protocol buffers_) that literaly builds the **_Class:_** `Employee` and its **_properties_**.

You can use `protoc`, a protocol compiler built by google. You feed it the proto file and specify what language you want the output as, and `protoc` is going to output the equivalent language.

**Resources:**

_Golang Protocol Buffer Tutorial
<https://tutorialedge.net/golang/go-protocol-buffer-tutorial/>_

_Protocol Buffer Basics: Golang
<https://developers.google.com/protocol-buffers/docs/gotutorial>_

_Golang Generated Code
<https://developers.google.com/protocol-buffers/docs/reference/go-generated>_
