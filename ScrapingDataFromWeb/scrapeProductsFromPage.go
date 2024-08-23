package ScrapingDataFromWeb

import (
	"context"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"sync"
	"time"
)

func ScrapeDataFromWebPage(ctx context.Context, categoryName string, categoryURL string, logChan chan<- string) {
	var wg sync.WaitGroup
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 4,
		RandomDelay: 3 * time.Minute,
	})

	c.SetRequestTimeout(60 * time.Minute)

	c.OnHTML("div.bx_catalog_list_home.list.bx_blue", func(e *colly.HTMLElement) {
		e.ForEach("div.bx_catalog_item.gtm-impression-product", func(i int, e *colly.HTMLElement) {
			productURL := "https://shop.kz" + e.ChildAttr("a", "href")

			wg.Add(1)
			go func() {
				defer wg.Done()
				ScrapeDataFromPageOfProduct(ctx, categoryName, productURL, logChan)
			}()
		})
	})

	c.OnHTML("div.bx-pagination.bx-blue", func(e *colly.HTMLElement) {
		e.ForEach("ul", func(i int, e *colly.HTMLElement) {
			e.ForEach("li.bx-pag-next", func(i int, e *colly.HTMLElement) {
				urlFotNextPage := "https://shop.kz" + e.ChildAttr("a", "href")
				e.ForEach("div.bx_catalog_item_controls", func(i int, e *colly.HTMLElement) {
					if e.ChildText("span") != "Нет в наличии" {
						wg.Add(1)
						go func() {
							defer wg.Done()
							ScrapeDataFromWebPage(ctx, categoryName, urlFotNextPage, logChan)
						}()
					}
				})
			})
		})
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == http.StatusMovedPermanently || r.StatusCode == http.StatusFound {
			location := r.Headers.Get("Location")
			if location != "" {
				if isProductPage(location) {
					c.Visit(location)
				}
				return
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		logMessage := "\n\nFinished Scrape Data From the Category " + r.Request.URL.String() + "\n\n"
		logChan <- logMessage
		log.Println(colorCyan, logMessage, colorReset)
	})

	err := c.Visit(categoryURL)
	if err != nil {
		log.Println(colorGreen, "scrapeProductsFromPage categoryURL: ", categoryURL, colorRed, err, colorReset)
		time.Sleep(3 * time.Second)
		c.Visit(categoryURL)
	}

	wg.Wait()
}

func isProductPage(url string) bool {
	return !isMainMenu(url)
}

func isMainMenu(url string) bool {
	return url == "https://shop.kz/catalog/"
}
