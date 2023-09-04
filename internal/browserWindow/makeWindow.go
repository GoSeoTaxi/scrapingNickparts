package browserWindow

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"scrapingNickparts/internal/ChangeData"
	"scrapingNickparts/internal/constData"
	"scrapingNickparts/internal/structures"
)

func GetReq(sURL string, pathChrome string, debugLog structures.DebugLog) (body []byte) {
	time.Sleep(constData.TimeOutRequest * time.Second)
	sURL = ChangeData.Replacer(sURL)

	var res string

	for rep := 0; rep < constData.ReplyGetRequest; rep++ {

		if rep > 0 {
			if debugLog.Debug {
				fmt.Printf("Ошибка сети. Запуск цил повтора. Попытка номер - %v \nURL - %v \n", rep, sURL)
			}
		}

		var err1 error
		func() {

			selectorClick := ""

			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("user-data-dir", pathChrome),
				//	chromedp.Flag("no-sandbox", true),
				chromedp.Flag("no-first-run", true),
				//	chromedp.Flag("headless", !debugLog.Debug),
				chromedp.Flag("headless", false),
				chromedp.Flag("disable-gpu", false),
				chromedp.Flag("enable-automation", false),
				chromedp.Flag("disable-extensions", false),
				//	chromedp.WindowSize(912, 1368),
				chromedp.WindowSize(1366, 768),
			)

			// create context.allocCtx
			allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
			defer cancel()

			// create context.ctx
			ctx, cancel2 := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
			defer cancel2()

			time.AfterFunc(120*time.Second, func() {
				if debugLog.Debug {
					fmt.Println("Завершаем зависший процесс")
				}
				cancel2()
			})

			err := chromedp.Run(ctx,
				chromedp.Navigate(sURL),
				wait5S(debugLog),
				//	chromedp.WaitNotVisible("div.loader"),
			)

			// run task list
			err = chromedp.Run(ctx,
				// clickAccept(),
				wait5S(debugLog),
			)

			err = chromedp.Run(ctx,
				clickMoreCheck(&selectorClick, debugLog),
			)

			if selectorClick == constData.ClassMoreClick {
				err = chromedp.Run(ctx,
					chromedp.Click(constData.ClassMoreClick),
				)
			}

			err = chromedp.Run(ctx,
				wait5S(debugLog),
				clickParameters(debugLog),
				wait5S(debugLog),
				scrapIt(&res),
			)
			if err != nil {
				fmt.Println(err)
				err1 = err
			}

			return

		}()

		if debugLog.Debug {
			fmt.Println(`Создаём запрос из браузера`)
			fmt.Println(sURL)
			fmt.Println("Запрошенная страница")
		}

		if err1 != nil {
			continue
		}

		body = []byte(res)
		break

	}

	return body
}
