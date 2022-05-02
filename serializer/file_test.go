package serializer_test

import (
	"grpc_app/pb"
	"grpc_app/sample"
	"grpc_app/serializer"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	// Test proto message to binary.
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	// Check if the file is successfully written with testify.
	require.NoError(t, err)

	// Test binary to proto message.
	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	// Test compare data.
	require.True(t, proto.Equal(laptop1, laptop2))

	// Test proto message to JSON.
	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
