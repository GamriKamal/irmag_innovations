package ScrapingDataFromWeb

import (
	"context"
	"github.com/gocolly/colly"
	"github.com/google/uuid"
	"log"
	"parseAndSendData_service/OperationsWithMongoDB"
	"regexp"
	"strings"
	"sync"
	"time"
)

func ScrapeDataFromPageOfProduct(ctx context.Context, categoryName string, productURL string, logChan chan<- string) {
	c := colly.NewCollector()
	c.SetRequestTimeout(60 * time.Minute)

	item := make(map[string]interface{})
	item["id_item"] = uuid.New().ID()
	item["category"] = categoryName
	item["url"] = productURL

	var wg sync.WaitGroup
	var imageUrls []string
	wg.Add(4)
	go func() {
		defer wg.Done()

		c.OnHTML("div.col-lg-4.col-sm-6.col-xs-12", func(e *colly.HTMLElement) {
			bigImage := e.ChildAttr("div.bx_bigimages_imgcontainer img", "data-src")
			if bigImage != "" {
				bigImage = "https:" + bigImage
				imageUrls = append(imageUrls, bigImage)
			}

			e.ForEach("div.bx_slider_conteiner ul li span.cnt_item", func(_ int, el *colly.HTMLElement) {
				smallImage := el.Attr("style")
				if smallImage != "" {
					re := regexp.MustCompile(`url\(['"]?(.*?)['"]?\)`)
					matches := re.FindStringSubmatch(smallImage)
					if len(matches) > 1 {
						url := matches[1]
						smallImage = strings.Trim(url, "'\"")
					}
					smallImage = "https:" + smallImage
					smallImage = strings.Replace(smallImage, "100_100_1", "450_450_1", 1)
					imageUrls = append(imageUrls, smallImage)
				}
			})
			item["photos"] = imageUrls
		})

	}()

	go func() {
		defer wg.Done()
		c.OnHTML("div.bx-title__container.col-lg-12", func(e *colly.HTMLElement) {
			item["title"] = e.ChildText("h1")
		})
	}()

	go func() {
		defer wg.Done()
		c.OnHTML("div.item_current_price", func(e *colly.HTMLElement) {
			item["price"] = e.Text
		})
	}()

	go func() {
		defer wg.Done()
		c.OnHTML("div.bxe-tabs__content", func(e *colly.HTMLElement) {
			if e.Attr("data-tab") == "description" {
				item["description"] = e.Text
				item["linkForPhoto"] = "https://shop.kz" + e.ChildAttr("img", "data-src")
			}

			if e.Attr("data-tab") == "characteristics" {
				e.ForEach("div.bx_detail_chars_group", func(i int, element *colly.HTMLElement) {
					element.ForEach("div.bx_detail_chars_i", func(j int, element *colly.HTMLElement) {
						key := element.ChildText("dt.bx_detail_chars_i_title")
						value := element.ChildText("dd.bx_detail_chars_i_field")
						item[key] = value
					})
				})

				e.ForEach("div.bx_detail_chars_i", func(i int, element *colly.HTMLElement) {
					key := element.ChildText("dt.bx_detail_chars_i_title")
					value := element.ChildText("dd.bx_detail_chars_i_field")
					if key != "" && value != "" {
						item[key] = value
					}
				})
			}
		})
	}()

	c.OnScraped(func(r *colly.Response) {
		logMessage := "\n\nFinished Scrape Data From Product Page " + r.Request.URL.String() + "\n\n"
		logChan <- logMessage
	})

	err := c.Visit(productURL)
	if err != nil {
		log.Println(colorGreen, "Error visiting product URL:", productURL, colorRed, err, colorReset)
		time.Sleep(3 * time.Second)
		c.Visit(productURL)
	}

	wg.Wait()

	if item["title"] != nil && item["title"] != "" {
		OperationsWithMongoDB.InsertDataToMongoDB(ctx, item)
	}
}
