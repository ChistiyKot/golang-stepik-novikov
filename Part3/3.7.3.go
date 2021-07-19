/*
На стандартный ввод подается строковое представление даты и времени в следующем формате:
1986-04-16T05:20:00+06:00
Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:
Wed Apr 16 05:20:00 +0600 1986
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var input string
	fmt.Scan(&input)

	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.Format(time.UnixDate))
}
