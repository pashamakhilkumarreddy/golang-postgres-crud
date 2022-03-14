package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func LogError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func GetInput() string {
	var data string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}
	return data
}
