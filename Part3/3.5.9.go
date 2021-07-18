/*
bufio vs fmt https://stepik.org/lesson/351892/step/9?unit=335849
run:
echo "33\n47\n12\n79\n15" | go run 3.5.9.go
*/
package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		sum += x
	}

	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
