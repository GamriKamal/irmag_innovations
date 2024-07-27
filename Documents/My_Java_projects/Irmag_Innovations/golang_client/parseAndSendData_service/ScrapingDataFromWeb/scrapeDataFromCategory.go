package ScrapingDataFromWeb

import (
	"context"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"sync"
	"time"
)

func ScrapeDataFromCategory(ctx context.Context, departmentURL string, logChan chan<- string) {
	var wg sync.WaitGroup
	c := colly.NewCollector()
	c.SetRequestTimeout(60 * time.Minute)

	c.OnHTML("div.bx-content.col-xs-12", func(e *colly.HTMLElement) {
		e.ForEach("ul.bx_catalog_tile_ul", func(i int, element *colly.HTMLElement) {
			element.ForEach("li", func(i int, element *colly.HTMLElement) {
				categoryName := element.ChildText("h2")
				categoryURL := element.ChildAttr("a", "href")

				if strings.HasPrefix(categoryURL, "/catalog/") {
					categoryURL = "https://shop.kz" + categoryURL
				}

				wg.Add(1)
				go func() {
					defer wg.Done()
					if checkAllowedCategoryToParse(categoryURL) {
						ScrapeDataFromCategory(ctx, categoryURL, logChan)
					}
				}()

				if strings.Contains(categoryURL, "/offers") {

					wg.Add(1)
					go func() {
						defer wg.Done()
						if checkAllowedCategoryToParse(categoryURL) {
							ScrapeDataFromWebPage(ctx, categoryName, categoryURL, logChan)
						}
					}()
				}
			})
		})
	})

	err := c.Visit(departmentURL)
	if err != nil {
		log.Println(colorGreen, "ScrapeDataFromCategory, departmentURL: ", departmentURL, colorRed, err, colorReset)
		time.Sleep(3 * time.Second)
		c.Visit(departmentURL)
	}

	c.OnScraped(func(r *colly.Response) {
		logMessage := "Finished Scrape Data From the Category " + r.Request.URL.String()
		logChan <- logMessage
		log.Println(colorCyan, logMessage, colorReset)
	})

	wg.Wait()
}

func checkAllowedCategoryToParse(url string) bool {
	if url != "https://shop.kz/catalog/avtotovary/" ||
		url != "https://shop.kz/catalog/kantstovary/" ||
		url != "https://shop.kz/offers/podarochnye-sertifikaty/" ||
		url != "https://shop.kz/catalog/uslugi/" {
		return true
	} else {
		return false
	}
}
