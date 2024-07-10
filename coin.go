package main

import (
	"fmt"
	"github.com/tebeka/selenium"
)

type Coin struct {
	Name      string
	Price     string
	LastHour  string
	LastDay   string
	LastWeek  string
	MarketCap string
	Volume    string
}

func (coin *Coin) parseFromHTML(node selenium.WebElement) {
	cells, err := node.FindElements(selenium.ByTagName, "td")
	if err != nil {
		return
	}

	var nameParts []string
	ps, _ := cells[2].FindElements(selenium.ByTagName, "p")
	for _, p := range ps {
		text, _ := p.Text()
		nameParts = append(nameParts, text)
	}

	name := ""

	if len(nameParts) >= 2 {
		name = fmt.Sprintf("%s (%s)", nameParts[0], nameParts[1])
	} else if len(nameParts) == 1 {
		name = nameParts[0]
	}
	coin.Name = name

	priceEl, err := cells[3].FindElement(selenium.ByTagName, "span")
	if err == nil {
		price, _ := priceEl.Text()
		coin.Price = price
	}

	lastHourEl, err := cells[4].FindElement(selenium.ByTagName, "span")
	if err == nil {
		lastHour, _ := lastHourEl.Text()
		coin.LastHour = lastHour
	}

	lastDayEl, err := cells[5].FindElement(selenium.ByTagName, "span")
	if err == nil {
		lastDay, _ := lastDayEl.Text()
		coin.LastDay = lastDay
	}

	lastWeekEl, err := cells[6].FindElement(selenium.ByTagName, "span")
	if err == nil {
		lastWeek, _ := lastWeekEl.Text()
		coin.LastWeek = lastWeek
	}

	marketCapEls, _ := cells[7].FindElements(selenium.ByTagName, "span")
	if len(marketCapEls) >= 2 {
		marketCap, _ := marketCapEls[1].Text()
		coin.MarketCap = marketCap
	}

	volumeEl, err := cells[8].FindElement(selenium.ByTagName, "p")
	if err == nil {
		volume, _ := volumeEl.Text()
		coin.Volume = volume
	}
}
