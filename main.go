package main

import (
	"encoding/csv"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var rng = rand.New(source)

const (
	requestURL   = "https://coinmarketcap.com/"
	defaultFirst = 1
	defaultLast  = 10
	outputFile   = "coins.csv"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:85.0) Gecko/20100101 Firefox/85.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.705.50 Safari/537.36 Edg/88.0.705.50",
}

func parseURL(requestURL string) []Coin {
	fmt.Println("Scraping: ", requestURL)

	var err error

	var coins = make([]Coin, 0)

	driver := getDriver("./chromedriver")

	err = driver.Get(requestURL)
	if err != nil {
		log.Fatalf("Error navigating to URL %s: %v", requestURL, err)
	}

	time.Sleep(2 * time.Second)

	scrollPage(driver)

	table, _ := driver.FindElement(selenium.ByCSSSelector, "table.cmc-table tbody")
	rows, _ := table.FindElements(selenium.ByTagName, "tr")

	for _, row := range rows {
		var coin Coin
		coin.parseFromHTML(row)
		coins = append(coins, coin)
	}

	err = driver.Quit()
	if err != nil {
		log.Fatalf("Error quitting the driver: %v", err)
	}

	return coins
}

func getDriver(path string) selenium.WebDriver {
	_, err := selenium.NewChromeDriverService(path, 4444)

	if err != nil {
		log.Fatalf("Error starting the ChromeDriver service: %v", err)
	}

	userAgent := userAgents[rng.Intn(len(userAgents))]

	caps := selenium.Capabilities{
		"pageLoadStrategy": "eager",
	}
	caps.AddChrome(chrome.Capabilities{
		Args: []string{
			"--headless",
			"--user-agent=" + userAgent,
		},
	})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatalf("Error creating new remote session: %v", err)
	}

	return driver
}

func scrollPage(driver selenium.WebDriver) {
	for {
		currentHeight, err := driver.ExecuteScript("return document.body.scrollHeight", nil)
		if err != nil {
			log.Fatalf("Error getting current height: %v", err)
		}

		_, _ = driver.ExecuteScript("window.scrollBy(0, 1000);", nil)

		time.Sleep(200 * time.Millisecond)

		newHeight, _ := driver.ExecuteScript("return document.body.scrollHeight", nil)

		if currentHeight == newHeight {
			break
		}
	}
}

func exportToCsv(coins []Coin, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Failed to create output CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"Name",
		"Price",
		"1h %",
		"24h %",
		"7d %",
		"Market Cap",
		"Volume(24h)",
	}
	writer.Write(headers)

	for _, coin := range coins {
		record := []string{
			coin.Name,
			coin.Price,
			coin.LastHour,
			coin.LastDay,
			coin.LastWeek,
			coin.MarketCap,
			coin.Volume,
		}

		writer.Write(record)
	}
	defer writer.Flush()
}

func main() {
	var err error

	var firstPage, lastPage int

	switch len(os.Args) {
	case 1:
		firstPage = defaultFirst
		lastPage = defaultLast
	case 3:
		firstPage, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Invalid first page number: %v", err)
		}
		lastPage, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid last page number: %v", err)
		}
	default:
		log.Fatalf("Invalid number of arguments. Usage: go run . <firstPage> <lastPage>")
	}

	var coins = make([]Coin, 0)
	for i := firstPage; i <= lastPage; i++ {
		coins = append(
			coins,
			parseURL(fmt.Sprintf("%s?page=%d", requestURL, i))...,
		)
		time.Sleep(time.Duration(1000+rng.Intn(1000)) * time.Millisecond)
	}

	exportToCsv(coins, outputFile)
	fmt.Println("Scraped!")
}
