package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"parseAndSendData_service/RetrieveDataFromDepartments"
	"parseAndSendData_service/ScrapingDataFromWeb"
	"sync"
	"time"
)

type PageData struct {
	Title            string
	Status           []string
	IsScraping       bool
	IsSendingToKafka bool
	ScrapingTimeLeft int64
	KafkaTimeLeft    int64
}

type Controller struct {
	ctx              context.Context
	cancel           context.CancelFunc
	logChan          chan string
	progressMux      sync.Mutex
	progress         int
	isScraping       bool
	isSendingToKafka bool
	scrapingTimer    *time.Timer
	kafkaTimer       *time.Timer
	scrapingEndTime  time.Time
	kafkaEndTime     time.Time
}

func NewController() *Controller {
	ctx, cancel := context.WithCancel(context.Background())
	return &Controller{
		ctx:     ctx,
		cancel:  cancel,
		logChan: make(chan string, 100),
	}
}

func (c *Controller) StartScraping(w http.ResponseWriter, r *http.Request) {
	if c.isScraping || c.isSendingToKafka {
		http.Redirect(w, r, "/go_client/admin", http.StatusSeeOther)
		return
	}
	c.isScraping = true
	c.progress = 0
	c.updateProgress(0)
	c.scrapingEndTime = time.Now().Add(45 * time.Minute)

	c.scrapingTimer = time.AfterFunc(45*time.Minute, func() {
		c.isScraping = false
		c.updateProgress(100)
		time.Sleep(2 * time.Second)
		c.updateProgress(0)
	})

	go func() {
		defer c.scrapingTimer.Stop()
		ctx, cancel := context.WithTimeout(c.ctx, 45*time.Minute)
		defer cancel()

		arrURL := ScrapingDataFromWeb.ScrapeDataFromDepartment(c.logChan)

		for i, url := range arrURL {
			c.updateProgress((i + 1) * 100 / len(arrURL))
			time.Sleep(5 * time.Second)
			ScrapingDataFromWeb.ScrapeDataFromCategory(ctx, url, c.logChan)
			time.Sleep(1 * time.Minute)
			log.Println("\n\n\n")
		}

		http.Redirect(w, r, "/go_client/admin", http.StatusSeeOther)
	}()

	http.Redirect(w, r, "/go_client/admin", http.StatusSeeOther)
}

func (c *Controller) SendDataToKafka(w http.ResponseWriter, r *http.Request) {
	if c.isScraping || c.isSendingToKafka {
		http.Redirect(w, r, "/go_client/admin", http.StatusSeeOther)
		return
	}

	c.isSendingToKafka = true
	c.kafkaEndTime = time.Now().Add(2 * time.Minute)

	c.kafkaTimer = time.AfterFunc(2*time.Minute, func() {
		c.isSendingToKafka = false
		log.Println("Kafka sending completed.")
	})

	go func() {
		defer c.kafkaTimer.Stop()
		ctx, cancel := context.WithTimeout(c.ctx, 1*time.Minute)
		defer cancel()

		err := RetrieveDataFromDepartments.SendAllDataToKafka(ctx)
		if err != nil {
			log.Println("Kafka sending error:", err)
		} else {
			log.Println("Kafka sending completed.")
		}

		http.Redirect(w, r, "/go_client/admin", http.StatusSeeOther)
	}()

	http.Redirect(w, r, "/go_client/admin", http.StatusSeeOther)
}

func (c *Controller) AdminPage(w http.ResponseWriter, r *http.Request) {
	var logMessages []string

	for len(c.logChan) > 0 {
		logMessages = append(logMessages, <-c.logChan)
	}

	scrapingTimeLeft := int64(0)
	if c.isScraping {
		scrapingTimeLeft = int64(c.scrapingEndTime.Sub(time.Now()).Seconds())
	}

	kafkaTimeLeft := int64(0)
	if c.isSendingToKafka {
		kafkaTimeLeft = int64(c.kafkaEndTime.Sub(time.Now()).Seconds())
	}

	data := PageData{
		Title:            "Admin Panel",
		Status:           logMessages,
		IsScraping:       c.isScraping,
		IsSendingToKafka: c.isSendingToKafka,
		ScrapingTimeLeft: scrapingTimeLeft,
		KafkaTimeLeft:    kafkaTimeLeft,
	}

	tmpl, err := template.ParseFiles("./templates/admin.html")
	if err != nil {
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	errAdminTemplate := tmpl.Execute(w, data)
	if errAdminTemplate != nil {
		http.Error(w, "Failed to execute template: "+errAdminTemplate.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) updateProgress(progress int) {
	c.progressMux.Lock()
	defer c.progressMux.Unlock()
	c.progress = progress
}

func (c *Controller) GetProgress() int {
	c.progressMux.Lock()
	defer c.progressMux.Unlock()
	return c.progress
}

func GetController() {
	controller := NewController()

	http.HandleFunc("/go_client/admin", controller.AdminPage)
	http.HandleFunc("/go_client/start-scraping", controller.StartScraping)
	http.HandleFunc("/go_client/send-to-kafka", controller.SendDataToKafka)

	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
