package variadictest

import (
	"fmt"
)

func TestVariadic(text ...string) {
	//for _, value := range text {
	//	fmt.Println("Content :",value)
	//}

	fmt.Println(text)
}
