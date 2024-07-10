# coinmarketcap-scraper

Cryptocurrency data scraper for CoinMarketCap in Go.

## About

This is a web scraper written in Go that uses Selenium to collect cryptocurrency data from CoinMarketCap and save it into a CSV file.

## Setting up

Follow these steps to set up the project:

1. **Install Go**: Download and install Go from the [official website](https://golang.org/dl/).
2. **Install Chromedriver**: Download Chromedriver and place the executable in your system's PATH.
3. **Install Selenium**: Use the following command to install Selenium:
      ```sh
      go get -u github.com/tebeka/selenium
      ```

4. **Clone the repository**: Clone this repository to your local machine using command:
      ```sh
      git clone <repository_url>
      cd <repository_name>
      ```

## Usage

To build and run the project, follow these steps:

1. **Build the project**: Navigate to the project directory and build the Go program:
      ```sh
      go build -o scraper main.go coin.go
      ```

2. **Run the scraper**:
    - Run the executable with optional arguments for the first and last page to scrape:
      ```sh
      ./scraper [firstPage] [lastPage]
      ```
    - If no arguments are provided, the scraper will default to scraping pages 1 through 10:
      ```sh
      ./scraper
      ```
    - To scrape specific pages, provide the start and end page numbers as arguments:
      ```sh
      ./scraper 1 5
      ```
    - This command will scrape pages from 1 to 5 and save the data to `coins.csv`.