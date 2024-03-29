package importData

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"scrapingNickparts/internal/structures"
)

func sendJson(resMarchal *[]byte, debugLog structures.DebugLog) (err error) {

	if debugLog.Debug {
		fmt.Println(debugLog.NumberTrade+"_Trade URL:>", debugLog.UrlE)
	}

	req, err := http.NewRequest("POST", debugLog.UrlE, bytes.NewBuffer(*resMarchal))
	//	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if debugLog.Debug {
		fmt.Println(debugLog.NumberTrade+"_trade response Status:", resp.Status)
		fmt.Println(debugLog.NumberTrade+"_trade response Headers:", resp.Header)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if debugLog.Debug {
		fmt.Println(debugLog.NumberTrade+"_trade response Body:", string(body))
	}

	return err
}
