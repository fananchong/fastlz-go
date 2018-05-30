package tests

import (
	"testing"
	"unsafe"

	fastlz "github.com/fananchong/fastlz-go"
)

const FloatSize = int(unsafe.Sizeof(float32(1.0)))
const RAND_MAX_COUNT = 200000

func Test_main(t *testing.T) {

	Assert(len(randValueByte) == RAND_MAX_COUNT*FloatSize)

	const size = int(RAND_MAX_COUNT * FloatSize * 2)
	{
		var out [size]byte
		ln := fastlz.Fastlz_compress(randValueByte[:], len(randValueByte), out[:])
		Assert(ln != 0)
		v := 0
		for i := 0; i < ln; i++ {
			v += int(out[i])
		}
		Assert(IsEquals(float32(v), resultValue[resultIndex]))
		resultIndex++

		dec := make([]byte, size)
		ln2 := fastlz.Fastlz_decompress(out[:], ln, dec, size)
		Assert(ln2 == RAND_MAX_COUNT*FloatSize)
		for i := 0; i < RAND_MAX_COUNT*FloatSize; i++ {
			Assert(dec[i] == randValueByte[i])
		}
	}
	{
		var out [size]byte
		ln := fastlz.Fastlz_compress_level(1, randValueByte[:], len(randValueByte), out[:])
		Assert(ln != 0)
		v := 0
		for i := 0; i < ln; i++ {
			v += int(out[i])
		}
		Assert(IsEquals(float32(v), resultValue[resultIndex]))
		resultIndex++

		dec := make([]byte, size)
		ln2 := fastlz.Fastlz_decompress(out[:], ln, dec, size)
		Assert(ln2 == RAND_MAX_COUNT*FloatSize)
		for i := 0; i < RAND_MAX_COUNT*FloatSize; i++ {
			Assert(dec[i] == randValueByte[i])
		}
	}
	{
		var out [size]byte
		ln := fastlz.Fastlz_compress_level(2, randValueByte[:], len(randValueByte), out[:])
		Assert(ln != 0)
		v := 0
		for i := 0; i < ln; i++ {
			v += int(out[i])
		}
		Assert(IsEquals(float32(v), resultValue[resultIndex]))
		resultIndex++

		dec := make([]byte, size)
		ln2 := fastlz.Fastlz_decompress(out[:], ln, dec, size)
		Assert(ln2 == RAND_MAX_COUNT*FloatSize)
		for i := 0; i < RAND_MAX_COUNT*FloatSize; i++ {
			Assert(dec[i] == randValueByte[i])
		}
	}
}
