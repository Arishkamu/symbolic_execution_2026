package main

import (
	"fmt"
	"symbolic-execution-course/internal"
)

func main() {
	source := `
package main

func testFunction(x int) int {
	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}

	return x * 2
}
`

	result := internal.Analyse(source, "testFunction")
	for _, interpreter := range result {
		fmt.Println(interpreter.String())

	}
}
