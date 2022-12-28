package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		//	chromedp.WindowSize(912, 1368),
		chromedp.WindowSize(1366, 768),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var nodes []*cdp.Node
	selector := "#main ul li a"
	pageURL := "https://notepad-plus-plus.org/downloads/"
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(pageURL),
		chromedp.WaitReady(selector),
		chromedp.Nodes(selector, &nodes),
	}); err != nil {
		fmt.Println(`++`)
		panic(err)
	}
	f := func(ctx context.Context, url string) {
		clone, cancel := chromedp.NewContext(ctx)
		defer cancel()
		fmt.Printf("%s is opening in a new tab\n", url)

		if err := chromedp.Run(clone,
			chromedp.Navigate(url),
			chromedp.Sleep(1*time.Second),
		); err != nil {
			// do something nice with you errors!
			fmt.Println(`====`)
			panic(err)
		}
		cancel()
	}
	for _, n := range nodes {
		u := n.AttributeValue("href")
		time.Sleep(10 * time.Second)
		go f(ctx, u)
	}

	time.Sleep(40 * time.Second)
}
