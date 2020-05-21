package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	scrapper "stacew/gostudy/nomadcoders/4.JobScrapper"
)

//github.com/labstack/echo

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(scrapper.FileName) //삭제가 좀... 잘 안되넴..
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	//서버에 있는 이름 파일을, 사용자에게 리턴
	return c.Attachment(scrapper.FileName, scrapper.FileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":3000"))
}
