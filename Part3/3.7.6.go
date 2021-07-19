/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).
Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
*/
package main

import (
	"fmt"
	"time"
)

const now = 1589570165

func main() {
	var minutes, seconds time.Duration
	fmt.Scanf("%d мин. %d сек.", &minutes, &seconds)

	duration, err := time.ParseDuration(fmt.Sprintf("%dm%ds", minutes, seconds))
	if err != nil {
		panic(err)
	}

	output := time.Unix(now, 0).Add(duration).In(time.UTC).Format(time.UnixDate)

	fmt.Println(output)
}
