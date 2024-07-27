package ScrapingDataFromWeb

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
	"time"
)

var (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorPurple = "\033[35m"
	colorRed    = "\033[31m"
)

func ScrapeDataFromDepartment(logChan chan<- string) []string {
	c := colly.NewCollector()
	c.SetRequestTimeout(60 * time.Minute)

	var arrURL []string
	c.OnHTML("div.catalog-menu__sidebar", func(e *colly.HTMLElement) {
		e.ForEach("ul#catalog_menu_XEVOpk_list", func(i int, element *colly.HTMLElement) {
			element.ForEach("li", func(i int, element *colly.HTMLElement) {
				tempURL := element.ChildAttr("a.catalog-menu__nav-item", "href")
				if (strings.HasPrefix(tempURL, "/catalog") ||
					strings.HasPrefix(tempURL, "/podarochnye-/") ||
					strings.HasPrefix(tempURL, "/press")) && !strings.HasPrefix(tempURL, "/press") {

					tempURL = "https://shop.kz" + tempURL
					arrURL = append(arrURL, tempURL)
					//wg.Add(1)
					//go func() {
					//	defer wg.Done()
					//	ScrapeDataFromCategory(ctx, tempURL)
					//	time.Sleep(2 * time.Minute)
					//}()

				}
			})
		})
	})

	c.OnScraped(func(r *colly.Response) {
		logMessage := "\n\nFinished Scrape Data From the Department " + r.Request.URL.String() + "\n\n"
		logChan <- logMessage
		log.Println(colorCyan, logMessage, colorReset)
	})

	err := c.Visit("https://shop.kz/")
	if err != nil {
		log.Println(colorGreen, "ScrapeDataFromDepartment", colorRed, err, colorReset)
	}

	return arrURL
}
