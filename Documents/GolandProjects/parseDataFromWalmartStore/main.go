package main

import (
	"github.com/gocolly/colly"
	"strings"
)

func main() {

}

func proccesCategory(url string) {
	c := colly.NewCollector()
}

func retrieveLinkOfAllDepartments() map[string]string {
	c := colly.NewCollector()
	retrievedCategories := make(map[string]string)

	c.OnHTML("div.flex", func(e *colly.HTMLElement) {
		categoryURL := e.ChildAttr("a", "href")
		categoryName := e.ChildText("h2.ma0")

		if !strings.HasPrefix(categoryURL, "/shop/deals") && !strings.HasPrefix(categoryURL, "/all-departments") {
			if !strings.HasPrefix(categoryURL, "https://www.walmart.com") {
				categoryURL = "https://www.walmart.com" + categoryURL
			}
			proccesCategory(categoryURL)
			retrievedCategories[categoryName] = categoryURL
		}

	})

	c.Visit("https://www.walmart.com/all-departments")

	return retrievedCategories
}
