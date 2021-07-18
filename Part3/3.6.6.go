/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON.. https://stepik.org/lesson/353243/step/6?unit=337227.
*/
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type StudentGroup struct {
	Students []Student
}

type Student struct {
	Rating []int
}

func main() {
	var studentGroup StudentGroup

	jsonBlob, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("invalid input.")
	}

	if err := json.Unmarshal(jsonBlob, &studentGroup); err != nil {
		log.Fatal("unmarshall error: ", err)
	}

	count := 0
	for _, student := range studentGroup.Students {
		count += len(student.Rating)
	}

	average := float64(count) / float64(len(studentGroup.Students))
	resultJsonBlob, _ := json.MarshalIndent(map[string]float64{"Average": average}, "", "    ")

	io.WriteString(os.Stdout, string(resultJsonBlob))
}
