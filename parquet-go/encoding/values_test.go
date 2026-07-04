package encoding_test

import (
	"testing"
	"unsafe"

	"parquet_compress/parquet-go/encoding"
)

func TestValuesSize(t *testing.T) {
	t.Log(unsafe.Sizeof(encoding.Values{}))
}
