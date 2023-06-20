package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Dpt struct {
	Name string
}

type User struct {
	Name string
	Age  int
	Dpt  Dpt
}

func (m User) GetUserName() string {
	fmt.Println("Invoke GetUserName")
	return m.Name
}

func (m *User) SetUserName(stName string) {
	fmt.Println("Invoke SetUserName")
	m.Name = stName
}

func TestReflect(t *testing.T) {
	iUser := User{
		Name: "张三",
		Age:  28,
		Dpt:  Dpt{"技术部"},
	}

	// =========================================================================
	// = 执行值方法
	iUserTypes := reflect.TypeOf(iUser)
	nUserMethodCount := iUserTypes.NumMethod()
	for i := 0; i < nUserMethodCount; i++ {
		ptrMethod := iUserTypes.Method(i)
		iResult := ptrMethod.Func.Call([]reflect.Value{reflect.ValueOf(iUser)})
		fmt.Println(iResult[0])
	}

	// =========================================================================
	// = 执行指针方法
	iUserPtrTypes := reflect.TypeOf(&iUser)
	nPtrUserMethodCount := iUserPtrTypes.NumMethod()
	for i := 0; i < nPtrUserMethodCount; i++ {
		ptrMethod := iUserPtrTypes.Method(i)
		if "SetUserName" == ptrMethod.Name {
			ptrMethod.Func.Call([]reflect.Value{reflect.ValueOf(&iUser), reflect.ValueOf("更新名称后: 李四")})
			fmt.Println(iUser.GetUserName())
		}
	}
}
