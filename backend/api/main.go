package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	readJSONToken("./movies.json", func(data map[string]any) bool {
		return data["year"].(string) == "2011"
	})

}

func readJSONToken(fileName string, filter func(map[string]any) bool) []map[string]any {
	fileContent, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileContent.Close()

	decoder := json.NewDecoder(fileContent)

	filteredData := []map[string]any{}

	decoder.Token()

	data := map[string]any{}

	for decoder.More() {
		decoder.Decode(&data)
		if filter(data) {
			log.Println(data)
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}
