package Scraping

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"scrapingNickparts/internal/structures"
)

func Filling(b []byte) (out structures.JsonExport) {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		fmt.Println(`Не верные данные`)
		log.Fatal(err)
	}

	//	doc1 := doc

	//Собираем основной товар
	doc.Find("div.search__block_type_request").Each(func(i int, sd1 *goquery.Selection) {
		//Левый блок
		sd1.Find("div.search-spare__first-row-wrap").Each(func(i int, sd1_first *goquery.Selection) {
			out.RequestItemData.Pic, _ = sd1_first.Find(`img.search-spare__image-preview_clickable`).Attr("src")
			out.RequestItemData.Brand = sd1_first.Find(`span.search-spare__brand`).Text()
			out.RequestItemData.Num = sd1_first.Find(`span.search-spare__article`).Text()
			out.RequestItemData.Text = sd1_first.Find(`p.search-spare__name`).Text()
		})
		//правый блок
		sd1.Find("div.search-spare__offers").Each(func(i int, sd1_offer *goquery.Selection) {
			out.RequestItemData.Price = sd1_offer.Find(`div.search-offer__price-wrap`).First().Text()
		})

	})

	//Собираем параметры
	doc.Find("div.w-modal__window").Each(func(i int, s *goquery.Selection) {

		s.Find("div.detail-parameters__wrapper").Each(func(i3 int, l1 *goquery.Selection) {

			var param structures.RequestItemparameter
			param.Key, _ = l1.Children().Html()
			param.Value, _ = l1.Children().Next().Html()
			/*fmt.Println(`Key=`)
			fmt.Print(k)
			fmt.Println(`Value=`)
			fmt.Print(v)
			fmt.Println(`+++++`)*/

			out.RequestItemData.Parameters = append(out.RequestItemData.Parameters, param)

		})

	})

	//Собираем Оригинальные аналоги
	doc.Find("div.search__block_type_originalAnalog").Each(func(i int, sd0 *goquery.Selection) {
		sd0.Find("div.search__spare").Each(func(i int, sd1 *goquery.Selection) {

			var oAnalog structures.OriginalAnalog
			//Левый блок
			sd1.Find("div.search-spare__first-row-wrap").Each(func(i int, sd1_orig *goquery.Selection) {
				oAnalog.Pic, _ = sd1_orig.Find(`img.search-spare__image-preview_clickable`).Attr("src")
				oAnalog.Brand = sd1_orig.Find(`span.search-spare__brand`).Text()
				oAnalog.Num = sd1_orig.Find(`a.search-spare__article`).Text()
				oAnalog.Text = sd1_orig.Find(`p.search-spare__name`).Text()
			})
			//правый блок
			sd1.Find("div.search-spare__offers").Each(func(i int, sd2_orig *goquery.Selection) {
				oAnalog.Price = sd2_orig.Find(`div.search-offer__price-wrap`).First().Text()
			})
			out.OriginalAnalogs = append(out.OriginalAnalogs, oAnalog)
			//Тут реализация аппенда
		})

	})

	//Собираем НеОригинальные аналоги
	doc.Find("div.search__block_type_nonOriginalAnalog").Each(func(i int, sd0 *goquery.Selection) {
		sd0.Find("div.search__spare").Each(func(i int, sd1 *goquery.Selection) {

			var NoAnalog structures.NoOriginalAnalog
			//Левый блок
			sd1.Find("div.search-spare__first-row-wrap").Each(func(i int, sd1_orig *goquery.Selection) {
				NoAnalog.Pic, _ = sd1_orig.Find(`img.search-spare__image-preview_clickable`).Attr("src")
				NoAnalog.Brand = sd1_orig.Find(`span.search-spare__brand`).Text()
				NoAnalog.Num = sd1_orig.Find(`a.search-spare__article`).Text()
				NoAnalog.Text = sd1_orig.Find(`p.search-spare__name`).Text()
			})
			//правый блок
			sd1.Find("div.search-spare__offers").Each(func(i int, sd2_orig *goquery.Selection) {
				NoAnalog.Price = sd2_orig.Find(`div.search-offer__price-wrap`).First().Text()
			})
			out.NoOriginalAnalogs = append(out.NoOriginalAnalogs, NoAnalog)
		})

	})

	return out
}
