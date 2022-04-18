package models

type Content struct {
	IsMain      bool       `json:"is_main" bson:"is_main"`
	Id          string     `json:"id" bson:"id"`
	Name        string     `json:"name" bson:"name"`
	ContentType IdNamePair `json:"content_type" bson:"content_type"`
	Streams     []Stream   `json:"stream" bson:"stream"`
}

type Stream struct {
	Id              string           `json:"id" bson:"id"`
	StartTime       string           `json:"start_time" bson:"start_time"`
	EndTime         string           `json:"end_time" bson:"end_time"`
	Product         IdNamePair       `json:"product" bson:"product"`
	Distribution    IdNamePair       `json:"distribution" bson:"distribution"`
	StreamStatus    IdNamePair       `json:"stream_status" bson:"stream_status"`
	GeoRestrictions []GeoRestriction `json:"geo_restrictions" bson:"geo_restrictions"`
}

type GeoRestriction struct {
	CountryIsoAlpha2Codes string     `json:"country_iso_alpha2_codes" bson:"country_iso_alpha2_codes"`
	DeviceCategory        IdNamePair `json:"device_category" bson:"device_category"`
}

type Venue struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Competitor struct {
	Id        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	Qualifier string `json:"qualifier" bson:"qualifier"`
}

type IdNamePair struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Event struct {
	Id                  string       `json:"id" bson:"id"`
	SportEventId        string       `json:"sport_event_id" bson:"sport_event_id"`
	StartTime           string       `json:"start_time" bson:"start_time"`
	EndTime             string       `json:"end_time" bson:"end_time"`
	Venue               IdNamePair   `json:"venue" bson:"venue"`
	EventStatus         IdNamePair   `json:"event_status" bson:"event_status"`
	FirstLevelCategory  IdNamePair   `json:"first_level_category" bson:"first_level_category"`
	SecondLevelCategory IdNamePair   `json:"second_level_category" bson:"second_level_category"`
	ThirdLevelCategory  IdNamePair   `json:"third_level_category" bson:"third_level_category"`
	Competitors         []Competitor `json:"competitors" bson:"competitors"`
	Contents            []Content    `json:"contents" bson:"contents"`
}
