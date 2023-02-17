package tag

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

type PersonP struct {
	Name   string `label:"Name is:"`
	Age    int    `label:"Age is:"`
	Gender string `label:"Gender is:" default:"unknown"`
}

func init() {
	person := PersonP{
		Name:   "YuanTong",
		Age:    25,
		Gender: "",
	}
	Print(person)
}

func Print(obj interface{}) error {
	// 取value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {
		// 取 tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}
	return nil
}
func tag() {
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data1)

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data2)
}
