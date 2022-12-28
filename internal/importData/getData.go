package importData

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"scrapingNickparts/internal/constData"
	"time"
)

func GetDataJSON(url string) (tasks []string, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println(`Не Создали запрос`)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
		return nil, err
	}
	req.Header.Set("User-Agent", constData.UserAgert)
	req.Header.Set("Cache-Control", "no-cache")

	fmt.Println(`Делаем запрос к серверу за задачами`)

	client := &http.Client{Timeout: (time.Second * constData.TimeOutRequestGeter)}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(`Не считали данные`)
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
		fmt.Println(`Error encoding JSON. `)
		fmt.Println(err)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
		return nil, err
	}

	return tasks, nil
}
