package iteration

import (
	"fmt"
	"strings"
)

func Repeat(character string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	fmt.Println("ba" + strings.Repeat("na", 2))
}
