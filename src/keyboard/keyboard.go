//Package keyboard responsible for collecting user inputs from command line!
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//GetFloat takes user input as String and then converts it into Float64.
//It also takes care of invalid inputs.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("double is required as an Input")
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, fmt.Errorf("problem occurred during string to double conversion")
	}

	return number, nil
}
