package test

import "fmt"

func init() {
	fmt.Println(Hello(""))
}

func Hello(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %s", name)
	}
	return "Hello, world"
}
