package unittest

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type storage interface {
	store()
}

type storageImpl struct{}

func (s *storageImpl) store() {
	fmt.Println("store in storageImpl")
}

func createStorageImpl() *storageImpl {
	return nil
}

func createStorageImpl2() storage {
	return nil
}

// 测试interface与nil的比较
var s storage // s被定义为接口storage
func TestNilInterface(t *testing.T) {
	var p *int = nil
	var i interface{} = p
	fmt.Println(i == nil)

	fmt.Printf("接口初始时为零值： %v, 地址%p \n", asInterfaceStructure(s), &s)
	// *storageImpl 实现了 storage 接口，createInterface声明的返回值是一个 *storageImpl，
	// 返回的是一个具体实现，返回值被赋值到了接口上
	s = createStorageImpl()
	// 创建出来的实例的类型指针为<storageImpl>对应的地址，而 nil 对应的 类型指针为零值。所以两者不相等
	fmt.Printf("具体实现由接口接收：%v, %v, %v, %v, %p\n", asInterfaceStructure(s), reflect.TypeOf(s),
		asInterfaceStructure(nil), s == nil, &s)

	// 声明返回值为接口，此时返回nil时，没有具体的实现类型信息，所以类型指针为零值
	s = createStorageImpl2()
	fmt.Printf("返回值声明为接口：  %v, %v, %v, %p\n", asInterfaceStructure(s), reflect.TypeOf(s), s == nil, &s)
	// 以上的所有操作，s对应的都是同一块内存，赋值时发生值拷贝
}

// Go语言中，一个interface{}类型的变量包含了2个指针，一个指针指向值对应的类型，另外一个指针指向实际的值。
type InterfaceStructure struct {
	pt uintptr // 值类型的指针
	pv uintptr // 值内容的指针
}

func asInterfaceStructure(i interface{}) InterfaceStructure {
	return *(*InterfaceStructure)(unsafe.Pointer(&i))
}
