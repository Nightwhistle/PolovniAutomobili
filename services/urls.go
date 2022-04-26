package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
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

	// Novi sad: region%%5B%%5D=2550&
	// Beograd:  region%%5B%%5D=Beograd
	for _, wishlistItem := range wishlist.WishlistItems {
		url := fmt.Sprintf(
			"https://www.polovniautomobili.com/auto-oglasi/pretraga?"+
				"brand=%s"+
				"&model%%5B%%5D=%s"+
				"&price_from=%s"+
				"&price_to=%s"+
				"&year_from=%s"+
				"&year_to=%s"+
				"&showOldNew=all"+
				"&region%%5B%%5D=Beograd"+
				"&door_num=3013"+
				"&mileage_to=200000"+
				"&chassis%%5B%%5D=277"+
				"&chassis%%5B%%5D=2631"+
				"&submit_1="+
				"&without_price=0",
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

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
