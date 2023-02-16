package pointer

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("module pointer")
	setup()
}

func setup() {
	var v = 100

	var pt1 *int = &v
	var pt2 **int = &pt1

	fmt.Println("The value of variable v is =", v)
	fmt.Println("Address of variable v is =", &v)

	fmt.Println("The value of pt1 is = ", pt1)
	fmt.Println("Address of pt1 is = ", &pt1)

	fmt.Println("The value of pt2 is = ", pt2)

	fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)

	fmt.Println("*(Value at the address of pt2 is) or **pt2 =", **pt2)
}
