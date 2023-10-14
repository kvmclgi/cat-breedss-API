package controllers

import (
	breed "cat-breeds/models"
	utils "cat-breeds/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/beego/beego/v2/server/web"

	// for env
	"github.com/joho/godotenv"
)

type MainController struct {
	web.Controller
}

// for environment setup
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func (c *MainController) Get() {

	channel := make(chan string)

	// load .env file
	api_key := goDotEnvVariable("API_KEY")

	// here get all api breeds
	go utils.Get_api_request("https://api.thecatapi.com/v1/breeds", channel)
	get_breed_data := <-channel

	// here define structure for breeds
	breeds := breed.Breeds{}
	// unmarshall here
	err_1 := json.Unmarshal([]byte(get_breed_data), &breeds)
	if err_1 != nil {
		fmt.Println("Here some error get")
	}

	// set data in cache
	breed.SetCache("breeds_data", breeds)

	// get data from cache
	breeds_data, found := breed.GetCache("breeds_data")
	if found {
		fmt.Println("Data get from cache:")
	} else {
		fmt.Println("Error! Not found key into cache")
		return
	}

	// get request params
	var id string
	c.Ctx.Input.Bind(&id, "breed")

	// check params or not
	if id == "" {
		id = breeds[0].Id
	}
	c.Data["reid"] = id
	fmt.Println("URL: " + id)

	// here api call for images
	go utils.Get_api_request("https://api.thecatapi.com/v1/images/search?limit=25&breed_ids="+id+"&api_key="+api_key, channel)
	get_breeds_image_data := <-channel

	// here define structure for breedsimage
	breeds_image := breed.Breeds_images{}
	// unmarshall here
	err_img := json.Unmarshal([]byte(get_breeds_image_data), &breeds_image)
	if err_img != nil {
		fmt.Println("Here some error get")
	}

	c.Data["name"] = breeds_image[0].Breeds[0].Name
	c.Data["cat_id"] = breeds_image[0].Breeds[0].ID
	c.Data["description"] = breeds_image[0].Breeds[0].Description
	c.Data["temperament"] = breeds_image[0].Breeds[0].Temperament
	c.Data["origin"] = breeds_image[0].Breeds[0].Origin
	c.Data["weight"] = breeds_image[0].Breeds[0].Weight.Metric
	c.Data["life_span"] = breeds_image[0].Breeds[0].LifeSpan
	c.Data["wikipedia_url"] = breeds_image[0].Breeds[0].WikipediaURL

	fmt.Println(breeds_image[0].Breeds[0].Name)

	c.Data["breeds"] = breeds_data
	c.Data["breeds_images"] = breeds_image
	c.TplName = "index.html"
}
