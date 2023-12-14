package main

import (
	"log"
	"strconv"
	"time"

	"scrapingNickparts/internal/config"
	"scrapingNickparts/internal/constData"
	"scrapingNickparts/internal/importData"
	"scrapingNickparts/internal/structures"
)

func main() {

	log.Println("Starting app...")
	cfg, err := config.InitConfig()
	if err != nil {
		log.Printf("No load config %v", err)
	}

	// Количество потоков
	// log.Println(cfg.StartCopy)

	iConf, err := strconv.Atoi(cfg.StartCopy)
	if err != nil {
		log.Printf("No load config %v", err)
		iConf = 1
	}

	for i := 1; i <= iConf; i++ {

		go cirkle(cfg.Path+strconv.Itoa(i), i, *cfg)
		time.Sleep(30 * time.Second)
	}

	for {
		time.Sleep(180 * time.Second)
		log.Println(`Я работаю, не выключай меня`)
	}

	/*
		browserWindow.GetReq("https://nickparts.ru/search.html?article=KE90299945&brand=NISSAN&withAnalogs=1")
		browserWindow.GetReq("https://nickparts.ru/search.html?article=AGA001Z&brand=AGA&withAnalogs=1")
		browserWindow.GetReq("https://nickparts.ru/search.html?article=5412400&brand=UFI&withAnalogs=1")
	*/
}

func cirkle(path string, numTrade int, cfg config.Config) {

	debugLog := structures.DebugLog{
		Debug:       cfg.Debug,
		NumberTrade: strconv.Itoa(numTrade),
		UrlI:        cfg.URLImport,
		UrlE:        cfg.URLExport,
	}

	//	var debugLog structures.DebugLog{debugLog.Debug:cfg.Debug, }
	//	debugLog.Debug = cfg.Debug
	//	debugLog.NumberTrade = strconv.Itoa(numTrade)

	log.Println(`Запускаю поток ` + debugLog.NumberTrade)

	for {
		tasks, err := importData.GetDataJSON(cfg.URLImport, debugLog)
		if err != nil {
			log.Println(`Ошибка получения задач. Поток ` + debugLog.NumberTrade)
			log.Println(err)
			time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
			continue
		}

		//	var err error
		//	var tasks []structures.Task
		//	tasks = append(tasks, structures.Task{
		//	Old: structures.JsonOld{},
		//		Url: "https://nickparts.ru/search.html?article=AW06J01150105AG&brand=DOMINANT&withAnalogs=1",
		//		})

		err = importData.MakerQueue(tasks, path, debugLog)
		if err != nil {
			log.Println(`Ошибка получения данных. Поток ` + debugLog.NumberTrade)
			time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
			continue
		}

		if debugLog.Debug {
			log.Println(`Я обработал задачу и жду повтора цикла`)
		}
		log.Println(debugLog.NumberTrade + "_t Завершил отправку")

		time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)

	}

	return
}
