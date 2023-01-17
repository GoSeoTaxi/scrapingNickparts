package importData

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"scrapingNickparts/internal/constData"
	"scrapingNickparts/internal/structures"
	"time"
)

func GetDataJSON(url string, debugLog structures.DebugLog) (tasks []structures.Task, err error) {
	fmt.Println(debugLog.NumberTrade + `_t. Делаем запрос к серверу за задачами...`)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println(debugLog.NumberTrade + `_t. Не Создали запрос`)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
		return nil, err
	}
	req.Header.Set("User-Agent", constData.UserAgert)
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{Timeout: (time.Second * constData.TimeOutRequestGeter)}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(debugLog.NumberTrade + `_t. Не считали данные`)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(`Error reading body. `)
		fmt.Println(err)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
		res.Body.Close()
		return nil, err
	}
	res.Body.Close()

	tasks, err = taskMaker(body)
	if err != nil {
		fmt.Println(debugLog.NumberTrade + `_t. Error encoding JSON. `)
		fmt.Println(err)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)

	}

	if len(tasks) < 1 {
		return nil, fmt.Errorf(`Нет задач`)
	}

	return tasks, nil
}
