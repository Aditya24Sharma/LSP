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
	incoming := ("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}")
	method, content, err := rpc.DecodeMessage([]byte(incoming))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}

	if method != "hi" {
		t.Fatalf("Expected method 'hi', but got: %s", method)
	}

	if contentLength != 15 {
		t.Fatalf("Content length should be 15 but got %d", contentLength)
	}
}
