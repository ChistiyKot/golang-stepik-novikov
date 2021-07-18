/*
Данная задача ориентирована на реальную работу с данными в формате JSON. Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности), который был размещен на web-портале data.gov.ru... https://stepik.org/lesson/353243/step/9?unit=337227
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Id int `json:"global_id"`
}

func main() {
	data := make([]Item, 0)
	const url = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("github response error: ", err)
	}

	defer response.Body.Close()

	jsonBlob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("response read error: ", err)
	}
	log.Printf("response body: %s (...truncated)", string(jsonBlob)[0:100])

	if err := json.Unmarshal(jsonBlob, &data); err != nil {
		log.Fatal("unmarshall error: ", err)
	}

	total := 0
	for _, item := range data {
		total += item.Id
	}

	log.Printf("total calculated: %d", total)
}
