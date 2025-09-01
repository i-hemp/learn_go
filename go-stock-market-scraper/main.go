package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type Stock struct {
	company, price, change string
}

func main() {
	var ticker = []string{
		"NVDA",
		"AAPL",
		"MSFT",
		"GOOGL",
		"TSLA",
		"AMZN",
		"META",
		"BBY",
		"AMD",
		"INTC",
	}

	stocks := []Stock{}
	c := colly.NewCollector()

	// c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

	c.Limit(&colly.LimitRule{
		// DomainGlob: "*",
		Parallelism: 1,
		Delay:       2 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		stock := Stock{}
		h1Text := strings.TrimSpace(e.ChildText("h1"))
		stock.company = strings.TrimPrefix(h1Text, "Yahoo Finance")

		priceElements := e.DOM.Find("fin-streamer[data-field='regularMarketPrice']")
		if priceElements.Length() > 0 {
			stock.price = strings.TrimSpace(priceElements.First().Text())
		}

		changeElements := e.DOM.Find("fin-streamer[data-field='regularMarketChangePercent']")
		if changeElements.Length() > 0 {
			stock.change = strings.TrimSpace(changeElements.First().Text())
		}

		if stock.price != "" {
			re := regexp.MustCompile(`^([0-9,]+\.?[0-9]*)`) //for floats
			matches := re.FindStringSubmatch(stock.price)
			if len(matches) > 1 {
				stock.price = matches[1]
			}
		}

		if stock.change != "" {
			re := regexp.MustCompile(`^([-+]?[0-9,]*\.?[0-9]*%?)`) //for percentage values
			matches := re.FindStringSubmatch(stock.change)
			if len(matches) > 1 {
				stock.change = matches[1]
			}
		}

		fmt.Println(stock.company, stock.price, stock.change)

		if stock.company != "" && (stock.price != "" || stock.change != "") {
			stocks = append(stocks, stock)
		} else {
			fmt.Println("Insufficient data for this ticker")
		}
	})

	for i, t := range ticker {
		fmt.Println("Current ticker: ", i+1, len(ticker), t)
		err := c.Visit("https://finance.yahoo.com/quote/" + t + "/")
		if err != nil {
			log.Printf("Failed to visit %s: %v", t, err)
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("\nTotal found: %d\n", len(stocks))

	if len(stocks) == 0 {
		fmt.Println("No daata found")
		return
	}

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("Failed to create a csv file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"company", "price", "change"}
	err = writer.Write(headers)
	if err != nil {
		log.Fatalln("Failed to write headers: ", err)
	}

	for _, stock := range stocks {
		record := []string{stock.company, stock.price, stock.change}
		err := writer.Write(record)
		if err != nil {
			log.Printf("Failed to write record %v: %v", record, err)
		}
	}
	if len(stocks) <= 1 {

		fmt.Printf("finished writing %v stocks\n", len(stocks))
	}

	fmt.Printf("finished writing %v stocks\n", len(stocks))
}
