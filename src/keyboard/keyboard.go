//Package keyboard is created for taking user input for string to float64 conversion.
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//GetInput is a function used for taking user input.
func GetInput() (float64, error) {
	fmt.Print("Provide Input : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, fmt.Errorf("invalid input")
	}

	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, fmt.Errorf("string to float conversion failed")
	}

	return number, nil
}
