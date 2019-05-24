package benchmark

import (
	"fmt"
	"strconv"
	"testing"
)

/**
	int转字符串操作的基准测试，测试了三个函数，具体见一下三个基准测试方法
  */
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i ++ {
		strconv.Itoa(number)
	}
}
