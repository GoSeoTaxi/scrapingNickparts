package importData

import (
	"encoding/json"
	"fmt"
	"scrapingNickparts/internal/Scraping"
	"scrapingNickparts/internal/browserWindow"
	"scrapingNickparts/internal/structures"
)

func MakerQueue(tasks []structures.Task, pathChrome string, debugLog structures.DebugLog) (err error) {

	var result []structures.JsonExport

	for _, task := range tasks {
		bodyBytes := browserWindow.GetReq(task.Url, pathChrome, debugLog)
		out := Scraping.Filling(bodyBytes, task)
		result = append(result, out)
	}

	resMarchal, err := json.Marshal(result)
	if err != nil {
		fmt.Println("err Marchal")
		return err
	}

	if debugLog.Debug {
		fmt.Println(`++++`)
		fmt.Println(string(resMarchal))
	}

	err = sendJson(&resMarchal, debugLog)
	//Тут готовый json

	return err
}
