package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"github.com/techschool/pcbook/pb"
	"github.com/techschool/pcbook/src/sample"
	"github.com/techschool/pcbook/src/serializer"
)

/**
* Unit Test
 */

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{} // create new obj
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}

// In Go, `:=` is for declaration `+` assignment, whereas `=` is for assignment only. For example, `var foo int = 10` is the same as `foo := 10`
