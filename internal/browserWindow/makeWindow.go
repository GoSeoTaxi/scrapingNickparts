package browserWindow

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"scrapingNickparts/internal/ChangeData"
	"scrapingNickparts/internal/constData"
	"time"
)

func GetReq(sURL string) (body []byte) {
	time.Sleep(constData.TimeOutRequest * time.Second)
	sURL = ChangeData.Replacer(sURL)

	var res string

	for rep := 0; rep < constData.ReplyGetRequest; rep++ {

		if rep > 0 {
			fmt.Printf("Ошибка сети. Запуск цил повтора. Попытка номер - %v \nURL - %v \n", rep, sURL)
		}

		var err1 error
		func() {

			selectorClick := ""

			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("no-first-run", true),
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
				fmt.Println("Завершаем зависший процесс")
				cancel2()
			})

			err := chromedp.Run(ctx,
				chromedp.Navigate(sURL),
				wait5S(),
				//	chromedp.WaitNotVisible("div.loader"),
			)

			// run task list
			err = chromedp.Run(ctx,
				clickAccept(),
				wait5S(),
			)

			err = chromedp.Run(ctx,
				clickMoreCheck(&selectorClick),
			)

			if selectorClick == constData.ClassMoreClick {
				err = chromedp.Run(ctx,
					chromedp.Click(constData.ClassMoreClick),
				)
			}

			err = chromedp.Run(ctx,
				wait5S(),
				clickParameters(),
				wait5S(),
				scrapIt(&res),
			)
			if err != nil {
				fmt.Println(err)
				err1 = err
			}

		}()

		fmt.Println(`*`)
		if err1 != nil {
			continue
		}

		body = []byte(res)
		break

	}

	return body
}
