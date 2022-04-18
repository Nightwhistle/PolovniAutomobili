package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type WishlistItem struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Year         string `json:"year"`
	Price        string `json:"price"`
}

type WishlistItems struct {
	WishlistItems []WishlistItem `json:"wishlist"`
}

func GetUrls() []string {
	jsonFile, err := os.Open("wishlist.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var wishlist WishlistItems

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &wishlist)

	var urls []string

	for _, wishlistItem := range wishlist.WishlistItems {
		url := fmt.Sprintf(
			"https://www.polovniautomobili.com/auto-oglasi/pretraga?brand=%s&model%%5B%%5D=%s&price_from=%s&price_to=%s&year_from=%s&year_to=%s&showOldNew=all&submit_1=&without_price=0",
			wishlistItem.Manufacturer,
			wishlistItem.Model,
			strings.Split(wishlistItem.Price, "-")[0],
			strings.Split(wishlistItem.Price, "-")[1],
			strings.Split(wishlistItem.Year, "-")[0],
			strings.Split(wishlistItem.Year, "-")[1],
		)

		urls = append(urls, url)
	}

	return urls
}
