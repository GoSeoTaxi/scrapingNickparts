package importData

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"scrapingNickparts/internal/structures"
)

func taskMaker(b []byte) (tasks []structures.Task, err error) {

	var jsonVar structures.JsonImport

	err = json.Unmarshal(b, &jsonVar)
	if err != nil {
		log.Println(`Ошибка Анмаршал`)

		return nil, err
	}

	for _, row := range jsonVar.Result {
		tasks = append(tasks, struct {
			Old structures.JsonOld
			Url string
		}{Old: struct {
			OldNum   string
			OldBrand string
		}{OldNum: row.Num, OldBrand: row.Brand}, Url: urlMaker(row.Num, row.Brand)})
	}

	return tasks, nil
}

func urlMaker(number, brand string) (url string) {
	return fmt.Sprintf("https://nickparts.ru/search.html?article=%v&brand=%v&withAnalogs=1", cleanNumber(number), cleanBrand(brand))
}

func cleanNumber(number string) string {
	return strings.Replace(strings.ToUpper(number), " ", "", -1)
}

func cleanBrand(brand string) string {
	return strings.ToUpper(brand)
}
