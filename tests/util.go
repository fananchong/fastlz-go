package tests

import (
	"io/ioutil"
	"math"
	"reflect"
	"unsafe"
)

func IsEquals(a, b float32) bool {
	return math.Abs(float64(a-b)) < 0.00001
}

func Assert(cond bool) {
	if cond == false {
		panic("")
	}
}

///  rand data (type float)
var randValueF []float32
var randIndexF = 0

func frand() float32 {
	v := randValueF[randIndexF]
	randIndexF++
	return v
}

func initRandDataF() {
	tempdata, err := ioutil.ReadFile("rand.bin")
	Assert(err == nil)
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&randValueF)))
	sliceHeader.Cap = int(len(tempdata) / int(unsafe.Sizeof(float32(1.0))))
	sliceHeader.Len = int(len(tempdata) / int(unsafe.Sizeof(float32(1.0))))
	sliceHeader.Data = uintptr(unsafe.Pointer(&(tempdata[0])))
}

///  rand data (type byte)
var randValueByte []byte
var randIndexByte = 0

func brand() byte {
	v := randValueByte[randIndexByte]
	randIndexByte++
	return v
}

func initRandDataByte() {
	tempdata, err := ioutil.ReadFile("rand.bin")
	if err == nil {
		sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&randValueByte)))
		sliceHeader.Cap = int(len(tempdata) / int(unsafe.Sizeof(byte(1))))
		sliceHeader.Len = int(len(tempdata) / int(unsafe.Sizeof(byte(1))))
		sliceHeader.Data = uintptr(unsafe.Pointer(&(tempdata[0])))
	}
}

///  result data (type float)
var resultValue []float32
var resultIndex = 0

func initResultData() {
	tempdata, err := ioutil.ReadFile("result.bin")
	if err == nil {
		sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&resultValue)))
		sliceHeader.Cap = int(len(tempdata) / int(unsafe.Sizeof(float32(1.0))))
		sliceHeader.Len = int(len(tempdata) / int(unsafe.Sizeof(float32(1.0))))
		sliceHeader.Data = uintptr(unsafe.Pointer(&(tempdata[0])))
	}
}

func init() {
	initRandDataByte()
	initResultData()
}
