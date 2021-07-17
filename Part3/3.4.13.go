/*
Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект, удовлетворяющий интерфейсу fmt.Stringer. Назовем его "Батарейка"... https://stepik.org/lesson/350788/step/13?unit=334666
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Scan(&input)

	charge := strings.Count(input, "1")
	batteryForTest := Battery{charge: charge}
	fmt.Println(batteryForTest) // comment this line
}

type Battery struct {
	charge int
}

func (b Battery) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", b.charge))
}
