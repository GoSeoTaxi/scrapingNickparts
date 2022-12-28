package main

import (
	"fmt"
	"scrapingNickparts/internal/constData"
	"scrapingNickparts/internal/importData"
	"time"
)

func main() {

	for {

		tasks, err := importData.GetDataJSON(constData.UrlImportRequest)
		if err != nil {
			time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
			continue
		}

		err = importData.MakerQueue(tasks)
		if err != nil {
			time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
			continue
		}

		fmt.Println(`Я обработал задачу и жду повтора цикла`)
		time.Sleep(9999999 * time.Second)
		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)

	}

	/*
		browserWindow.GetReq("https://nickparts.ru/search.html?article=KE90299945&brand=NISSAN&withAnalogs=1")
		browserWindow.GetReq("https://nickparts.ru/search.html?article=AGA001Z&brand=AGA&withAnalogs=1")
		browserWindow.GetReq("https://nickparts.ru/search.html?article=5412400&brand=UFI&withAnalogs=1")
	*/
}
