package controllers

import (
	"regexp"
	"stream-api/models"
	"stream-api/repository"
	"stream-api/services"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type OfferController struct{}

func (oc OfferController) Fetch(ctx *gin.Context) {
	var offers []models.Offer

	er := repository.OffersRepository{}

	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("article.classified", func(i int, a *colly.HTMLElement) {
			offer := models.Offer{}
			space := regexp.MustCompile(`\s+`)

			offer.Name = strings.TrimSpace(a.ChildText("a.ga-title"))

			price := strings.TrimSpace(a.ChildText(".price"))
			offer.Price = space.ReplaceAllString(price, " ")
			offer.Link = a.ChildAttr("a.ga-title", "href")
			offer.Mileage = a.ChildText(".setInfo:nth-child(2) .top")

			offers = append(offers, offer)
		})

		er.WriteOffers(offers)
	})

	c.OnHTML("ul.uk-pagination a[rel='next']", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnError(func(r *colly.Response, e error) {
		// fmt.Println("Got this error:", e)
	})

	urls := services.GetUrls()
	for _, url := range urls {
		go visitUrls(c, url)
	}
	// c.JSON(200, mongoDocument.Events)
}

func (oc OfferController) Index(c *gin.Context) {
}

func visitUrls(c *colly.Collector, url string) {
	println("Visiting", url)
	c.Visit(url)
}
