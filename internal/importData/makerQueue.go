package importData

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"scrapingNickparts/internal/Scraping"
	"scrapingNickparts/internal/browserWindow"
	"scrapingNickparts/internal/structures"
)

func MakerQueue(tasks []structures.Task, pathChrome string, debugLog structures.DebugLog) (err error) {

	var result []structures.JsonExport

	for _, task := range tasks {
		bodyBytes := browserWindow.GetReq(task.Url, pathChrome, debugLog)

		out := Scraping.Filling(bodyBytes, task, debugLog)
		if debugLog.Debug {
			log.Println(`+++ out`)
			fmt.Printf("%+v\n", out)
			time.Sleep(3 * time.Second)
			log.Println("+++2")
		}
		result = append(result, out)
	}

	log.Printf("Поток %v_t собирается отправить %v заданий", debugLog.NumberTrade, len(result))

	if debugLog.Debug {
		log.Println(result)
		log.Println(`+++ MAKER QUEUE`)
	}

	resMarchal, err := json.Marshal(result)
	if err != nil {
		log.Println("err Marchal")
		return err
	}

	if debugLog.Debug {
		log.Println(`++++`)
		log.Println(string(resMarchal))
	}

	err = sendJson(&resMarchal, debugLog)
	// Тут готовый json

	return err
}
