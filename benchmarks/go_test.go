package benchmarks

import (
	"testing"

	fastlz "github.com/fananchong/fastlz-go"
	"github.com/fananchong/fastlz-go/tests"
)

const RAND_MAX_COUNT = 200000 * 4

var randValue [RAND_MAX_COUNT]byte

const count = 10000

var out [count][8192]byte
var outLen [count]int
var dec [count][8192]byte

func Benchmark_Go_Compress_128(t *testing.B) {
	t.N = count
	for i := 0; i < t.N; i++ {
		outLen[i] = fastlz.Fastlz_compress(randValue[i:], 128, out[i][:])
		tests.Assert(outLen[i] != 0)
	}
}

func Benchmark_GO_Decompress_128(t *testing.B) {
	t.N = count
	for i := 0; i < t.N; i++ {
		ln := fastlz.Fastlz_decompress(out[i][:], outLen[i], dec[i][:], 8192)
		tests.Assert(ln != 0)
	}
}

func Benchmark_Go_Compress_512(t *testing.B) {
	t.N = count
	for i := 0; i < t.N; i++ {
		outLen[i] = fastlz.Fastlz_compress(randValue[i:], 512, out[i][:])
		tests.Assert(outLen[i] != 0)
	}
}

func Benchmark_GO_Decompress_512(t *testing.B) {
	t.N = count
	for i := 0; i < t.N; i++ {
		ln := fastlz.Fastlz_decompress(out[i][:], outLen[i], dec[i][:], 8192)
		tests.Assert(ln != 0)
	}
}

func Benchmark_Go_Compress_1024(t *testing.B) {
	t.N = count
	for i := 0; i < t.N; i++ {
		outLen[i] = fastlz.Fastlz_compress(randValue[i:], 1024, out[i][:])
		tests.Assert(outLen[i] != 0)
	}
}

func Benchmark_GO_Decompress_1024(t *testing.B) {
	t.N = count
	for i := 0; i < t.N; i++ {
		ln := fastlz.Fastlz_decompress(out[i][:], outLen[i], dec[i][:], 8192)
		tests.Assert(ln != 0)
	}
}

func Benchmark_Go_Compress_4096(t *testing.B) {
	t.N = count
	for i := 0; i < t.N; i++ {
		outLen[i] = fastlz.Fastlz_compress(randValue[i:], 4096, out[i][:])
		tests.Assert(outLen[i] != 0)
	}
}

func Benchmark_GO_Decompress_4096(t *testing.B) {
	t.N = count
	for i := 0; i < t.N; i++ {
		ln := fastlz.Fastlz_decompress(out[i][:], outLen[i], dec[i][:], 8192)
		tests.Assert(ln != 0)
	}
}
