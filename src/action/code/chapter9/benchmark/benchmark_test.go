package benchmark
// 基准测试源文件名称也必须以_test结尾
import (
	"fmt"
	"strconv"
	"testing"
)
// 基准测试函数名必须以Benchmark开头，函数没有返回值
/**
	int转字符串操作的基准测试，测试了三个函数，具体见一下三个基准测试方法
	输出：
		goos: windows
		goarch: amd64
		BenchmarkSprintf-4   	 5000000	       367 ns/op
		BenchmarkFormat-4    	100000000	        18.1 ns/op
		BenchmarkItoa-4      	100000000	        18.9 ns/op
	上面的输出         ^ 4表示运行时对应的GOMAXPROCS的值，后面的数字表示执行的循环次数，18.9 ns/op 表示每次循环花费的时间
	想让测试运行的时间更长，可以通过-benchtime指定，比如3秒 go test -bench=. -benchtime=3s
	1 ns = 10^-9 s
	-test.benchmem 参数可以提供每次操作分配内存的次数:
		BenchmarkSprintf-4   	 3000000	       357 ns/op	      16 B/op	       2 allocs/op
		BenchmarkFormat-4    	100000000	        16.7 ns/op	       0 B/op	       0 allocs/op
		BenchmarkItoa-4      	100000000	        19.7 ns/op	       0 B/op	       0 allocs/op
	增加的两列输出表示：每次操作会分配多少内存，每次操作会进行几次内存分配。
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
