package importData

import (
	"encoding/json"
	"fmt"
	"scrapingNickparts/internal/structures"
)

func taskMaker(b []byte) (tasks []string, err error) {

	var jsonVar structures.JsonImport

	err = json.Unmarshal(b, &jsonVar)
	if err != nil {
		return nil, err
	}

	for _, row := range jsonVar.Result {
		tasks = append(tasks, urlMaker(row.Num, row.Brand))
	}

	return tasks, nil
}

func urlMaker(number, brand string) (url string) {
	return fmt.Sprintf("https://nickparts.ru/search.html?article=%v&brand=%v&withAnalogs=1", number, brand)
}
