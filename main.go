package main

import (
	"fmt"
	"strings"

	HelloWorld "github.com/jka236/TDDGo/hello_world"
)

func main() {
	HelloWorld.Hello("Jonghyeok", "English")
	fmt.Println("ba" + strings.Repeat("na", 2))
}
