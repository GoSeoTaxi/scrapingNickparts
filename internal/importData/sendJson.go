package importData

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"scrapingNickparts/internal/constData"
)

func sendJson(resMarchal *[]byte) (err error) {

	fmt.Println("URL:>", constData.UrlExportRequest)

	req, err := http.NewRequest("POST", constData.UrlExportRequest, bytes.NewBuffer(*resMarchal))
	//	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return err
}
