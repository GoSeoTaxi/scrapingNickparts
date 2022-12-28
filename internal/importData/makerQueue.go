package importData

import (
	"encoding/json"
	"fmt"
	"scrapingNickparts/internal/Scraping"
	"scrapingNickparts/internal/browserWindow"
	"scrapingNickparts/internal/structures"
)

func MakerQueue(tasks []string) (err error) {

	var result []structures.JsonExport

	for _, task := range tasks {
		bodyBytes := browserWindow.GetReq(task)
		//fmt.Println(string(bodyBytes))
		out := Scraping.Filling(bodyBytes)
		result = append(result, out)
	}

	resMarchal, err := json.Marshal(result)
	if err != nil {
		fmt.Println("err Marchal")
		return err
	}

	err = sendJson(&resMarchal)
	//Тут готовый json

	return err
}
