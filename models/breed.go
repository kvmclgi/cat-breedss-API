package models

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// define cache
var Cache = cache.New(10*time.Minute, 10*time.Minute)

// for breed setcache
func SetCache(key string, breed interface{}) bool {
	Cache.Set(key, breed, cache.NoExpiration)
	return true
}

// for breed getcache
func GetCache(key string) (Breeds, bool) {
	var breeds Breeds
	var found bool
	data, found := Cache.Get(key)
	if found {
		breeds = data.(Breeds)
	}
	return breeds, found
}

type Breeds []struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Breeds_images []struct {
	Breeds []struct {
		Weight struct {
			Imperial string `json:"imperial"`
			Metric   string `json:"metric"`
		} `json:"weight"`
		ID               string `json:"id"`
		Name             string `json:"name"`
		CfaURL           string `json:"cfa_url"`
		VetstreetURL     string `json:"vetstreet_url"`
		VcahospitalsURL  string `json:"vcahospitals_url"`
		Temperament      string `json:"temperament"`
		Origin           string `json:"origin"`
		CountryCodes     string `json:"country_codes"`
		CountryCode      string `json:"country_code"`
		Description      string `json:"description"`
		LifeSpan         string `json:"life_span"`
		Indoor           int    `json:"indoor"`
		Lap              int    `json:"lap"`
		Adaptability     int    `json:"adaptability"`
		AffectionLevel   int    `json:"affection_level"`
		ChildFriendly    int    `json:"child_friendly"`
		CatFriendly      int    `json:"cat_friendly"`
		DogFriendly      int    `json:"dog_friendly"`
		EnergyLevel      int    `json:"energy_level"`
		Grooming         int    `json:"grooming"`
		HealthIssues     int    `json:"health_issues"`
		Intelligence     int    `json:"intelligence"`
		SheddingLevel    int    `json:"shedding_level"`
		SocialNeeds      int    `json:"social_needs"`
		StrangerFriendly int    `json:"stranger_friendly"`
		Vocalisation     int    `json:"vocalisation"`
		Bidability       int    `json:"bidability"`
		Experimental     int    `json:"experimental"`
		Hairless         int    `json:"hairless"`
		Natural          int    `json:"natural"`
		Rare             int    `json:"rare"`
		Rex              int    `json:"rex"`
		SuppressedTail   int    `json:"suppressed_tail"`
		ShortLegs        int    `json:"short_legs"`
		WikipediaURL     string `json:"wikipedia_url"`
		Hypoallergenic   int    `json:"hypoallergenic"`
		ReferenceImageID string `json:"reference_image_id"`
	} `json:"breeds"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
