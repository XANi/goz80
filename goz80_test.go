package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStrings []string



var farr  [4]func (x int)( int)

func Square(x int) int {
	return x * x
}

func f1 (x int) int {
	return x * x
}

func f2 (x int, y *int) int {
	return x * *y

}

var	a = 3
var b = 3
func init() {
	farr[0] = f1
	farr[1] = func(x int)(int) { return f2(x, &a )}
}



func TestExample(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, 1, "test")
}

func BenchmarkConv(b *testing.B) {
	    var reg [2]byte
        for n := 0; n < b.N; n++ {
        	_ = farr[0](n)
        }
        _ = reg
}

func BenchmarkConv2(b *testing.B) {
	    var reg uint16
        for n := 0; n < b.N; n++ {
        	_ = farr[1](n)
        }
        _ = reg
}