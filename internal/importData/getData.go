package importData

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"scrapingNickparts/internal/constData"
	"scrapingNickparts/internal/structures"
)

func GetDataJSON(url string, debugLog structures.DebugLog) (tasks []structures.Task, err error) {
	log.Println(debugLog.NumberTrade + `_t. Делаем запрос к серверу за задачами...`)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		log.Println(debugLog.NumberTrade + `_t. Не Создали запрос`)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
		return nil, err
	}
	req.Header.Set("User-Agent", constData.UserAgert)
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{Timeout: (time.Second * constData.TimeOutRequestGeter)}
	res, err := client.Do(req)
	if err != nil {
		log.Println(debugLog.NumberTrade + `_t. Не считали данные`)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(`Error reading body. `)
		log.Println(err)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
		res.Body.Close()
		return nil, err
	}
	res.Body.Close()

	tasks, err = taskMaker(body)
	if err != nil {
		log.Println(debugLog.NumberTrade + `_t. Error encoding JSON. `)
		log.Println(err)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)

	}

	if len(tasks) < 1 {
		return nil, fmt.Errorf(`Нет задач`)
	}

	log.Printf("Поток %v_t - Получил %v задач \n", debugLog.NumberTrade, len(tasks))

	return tasks, nil
}
