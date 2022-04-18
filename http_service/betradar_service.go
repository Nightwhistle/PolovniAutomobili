package http_service

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"stream-api/models"
)

type BetradarHttpService models.HttpService

// Get request from betradar
func (bs BetradarHttpService) Get(url, clientIp string) (models.Response, error) {
	client := http.Client{}
	var response models.Response

	if envIp := os.Getenv("TEST_IP_ADDRESS"); envIp != "" {
		clientIp = envIp
	}

	req, err := http.NewRequest("GET", os.Getenv("BET_RADAR_API_URL")+url, nil)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("BET_RADAR_ACCESS_TOKEN"))
	req.Header.Add("X-Real-IP", clientIp)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return response, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		panic(err)
	}

	response.Code = resp.StatusCode
	response.Body = body

	return response, err
}
