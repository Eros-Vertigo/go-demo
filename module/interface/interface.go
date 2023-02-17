package _interface

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

/**
反射三定律
1. 反射可以将接口类型变量 转换为"反射类型对象"；
2. 反射可以将"反射类型对象"转换为 接口类型变量；
3. 如果要修改"反射类型对象"其类型必须是 可写的；
*/
func init() {
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T, 值为 %v \n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象: Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象: Value对象的类型为 %T \n", v)

	// 从反射对象到接口变量
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量: 新对象的类型为 %T 值为 %v \n", i, i)

	// 可写性
	var name string = "Hello Word"
	v1 := reflect.ValueOf(&name)

	fmt.Println("v1 可写性为 :", v1.CanSet())

	v2 := v1.Elem()
	fmt.Println("v2 可写性为 :", v2.CanSet())
}

func test() {
	// 声明一个空接口实例
	var i interface{}

	// 存 int
	i = 1
	log.Info(reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Kind(), reflect.ValueOf(i).CanSet())

	// 存字符串
	i = "hello"
	log.Info(reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Kind(), reflect.ValueOf(i).CanSet())

	// 存布尔值
	i = false
	log.Info(reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Kind(), reflect.ValueOf(i).CanSet())
}
