package models

type Offer struct {
	Id      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Price   string `json:"price" bson:"price"`
	Volume  string `json:"volume" bson:"volume"`
	Year    string `json:"year" bson:"year"`
	Link    string `json:"link" bson:"link"`
	Mileage string `json:"mileage" bson:"mileage"`
}

func (o *Offer) ToSlice(id string) []string {
	return []string{id, o.Name, o.Price, o.Volume, o.Year, o.Link, o.Mileage}
}

func Headers() []string {
	return []string{"id", "name", "price", "volume", "year", "link", "mileage"}
}
