package sub_package

import (
	"fmt"
)

var You struct {
	name string
}

func Sub_write_2() {
	You.name = "okayu"
	fmt.Println("this is sub package2.")
	fmt.Println(You)
}
