/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.
*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const DateFormat = "02.01.2006 15:04:05"

func main() {
	dates, err := csv.NewReader(os.Stdin).Read()

	if err != nil {
		panic(err)
	}

	date1 := parseDate(dates[0])
	date2 := parseDate(dates[1])

	diff := date2.Sub(date1)

	if diff < 0 {
		diff = -diff
	}

	fmt.Println(diff)
}

func parseDate(input string) time.Time {
	date, err := time.Parse(DateFormat, input)
	if err != nil {
		panic(err)
	}

	return date
}
