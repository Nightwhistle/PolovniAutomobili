package repository

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"stream-api/models"
)

type OffersRepository struct{}

func (or *OffersRepository) GetOffers() []models.Offer {
	records, err := readData("result.csv")
	var offers []models.Offer

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		offer := models.Offer{
			record[0],
			record[1],
			record[2],
			record[3],
			record[4],
			record[5],
			record[6],
		}

		offers = append(offers, offer)
	}

	return offers
}

func (or *OffersRepository) WriteOffers(offers []models.Offer) {
	file, err := os.Create("result.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(models.Headers())
	if err != nil {
		panic(err)
	}

	for id, value := range offers {
		err := writer.Write(value.ToSlice(strconv.Itoa(id)))
		if err != nil {
			panic(err)
		}
	}
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
