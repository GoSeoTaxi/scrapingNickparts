package browserWindow

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"scrapingNickparts/internal/constData"
	"time"
)

func wait5S() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("start waiting")
			return nil
		}),
		chromedp.Sleep(5 * time.Second),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("end waiting")
			return nil
		}),
	}
}

func clickAccept() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitEnabled("button.cookies-window__btn_accept"),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Нажимаю Принять куки")
			return nil
		}),
		chromedp.Click("button.cookies-window__btn_accept", chromedp.NodeVisible),
	}
}

func clickMoreCheck(selectorClick *string) chromedp.Tasks {

	nodes := []*cdp.Node{}

	return chromedp.Tasks{
		chromedp.Sleep(1 * time.Second),
		chromedp.ActionFunc(func(context.Context) error {
			//		log.Printf("Обработка показать ещё")
			return nil
		}),

		chromedp.ActionFunc(func(ctx context.Context) error {
			//		log.Printf("++++ показать ещё")
			return nil
		}),

		chromedp.Nodes("#main_inner_wrapper > div:nth-child(5) > div.search-parallel > div > div > button > span", &nodes, chromedp.AtLeast(0)),

		chromedp.ActionFunc(func(ctx context.Context) error {
			if len(nodes) > 0 {
				*selectorClick = constData.ClassMoreClick

				//for _, node := range nodes {
				//	fmt.Printf("node name: %s, node value: %s", node.NodeName, node.NodeValue)
				//}

			}

			return nil
		}),
		//	chromedp.Click("//*[@id=\"main_inner_wrapper\"]/div[4]/div[1]/div/div/button/span"),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Обработка показать ещё = Успех")
			return nil
		}),
	}
}

func clickParameters() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Нажимаем парам")
			return nil
		}),
		chromedp.Sleep(3 * time.Second),
		chromedp.Click("parameters__alternatives"),
		chromedp.WaitEnabled("div.modal-content"),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Дождались парам")
			return nil
		}),
	}
}
