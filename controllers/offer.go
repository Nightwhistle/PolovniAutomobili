package controllers

import (
	"net/http"
	"regexp"
	"stream-api/models"
	"stream-api/repository"
	"stream-api/services"
	"sync"

	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type OfferController struct{}

func (oc OfferController) Fetch(ctx *gin.Context) {
	var wg sync.WaitGroup
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
			offer.Link = "https://www.polovniautomobili.com" + a.ChildAttr("a.ga-title", "href")

			offer.Year = a.ChildText(".setInfo:nth-child(1) .top")
			offer.Volume = a.ChildText(".setInfo:nth-child(1) .bottom")

			mileage := strings.ReplaceAll(a.ChildText(".setInfo:nth-child(2) .top"), "km", "")
			mileage = strings.ReplaceAll(mileage, ".", "")
			mileage = strings.TrimSpace(mileage)
			mileageInt, err := strconv.Atoi(mileage)
			offer.Mileage = mileage

			if err != nil {
				panic(err)
			}

			if mileageInt > 180000 {
				return
			}

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
		wg.Add(1)

		go visitUrls(c, url, &wg)
	}

	wg.Wait()

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Offers": offers,
	})
}

func (oc OfferController) Index(c *gin.Context) {
}

func visitUrls(c *colly.Collector, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Visiting", url)
	c.Visit(url)
}
