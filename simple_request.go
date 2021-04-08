package main

import (
    "net/http"
    "io/ioutil"
		"log"
		"encoding/json"
		"os"
)


func getDataFromCEX(method string, option string) string {
	request := "https://cex.io/api/" + method + "/" + option
	response, error := http.Get(request)
	if error != nil {
		log.Fatal(error)
	}

	data, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if error != nil {
		log.Fatal(error)
	}
	values := map[string]string{}
	json.Unmarshal(data, &values)
	str := "<" + "1" + ">" + "/" + "<" + values["high"] + ">"
	return str
}

func main() {
	result := getDataFromCEX("ticker", "BTC/USD")

	f, error := os.Create(os.Args[1])
	if error != nil {
		log.Fatal(error)
	}
  f.WriteString(result)
	f.Close()
}
