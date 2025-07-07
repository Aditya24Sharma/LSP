package rpc_test

import (
	"educationalsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := ("Content-Length: 16\r\n\r\n{\"Testing\":true}")
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incoming := ("Content-Length: 16\r\n\r\n{\"Testing\":true}")
	contentLength, err := rpc.DecodeMessage([]byte(incoming))
	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 16 {
		t.Fatalf("Content length should be 16 but got %d", contentLength)
	}
}
