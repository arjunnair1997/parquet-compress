//go:build !s390x

package parquet

import (
	"github.com/parquet-go/bitpack/unsafecast"
	"parquet_compress/parquet-go/deprecated"
)

func unsafecastInt96ToBytes(src []deprecated.Int96) []byte {
	return unsafecast.Slice[byte](src)
}
