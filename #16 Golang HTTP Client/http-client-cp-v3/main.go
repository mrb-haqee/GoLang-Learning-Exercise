package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func isErrorF(err error){
	if err!=nil{
		log.Fatal(err)
	}
}
func isErrorP(err error){
	if err!=nil{
		panic(err)
	}
}

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}
type animeArr struct{
	data []Animechan
}

func ClientGet() ([]Animechan, error) {
	// client := http.Client{}
	resp,err:=http.Get("https://animechan.vercel.app/api/quotes/anime?title=naruto")
	isErrorF(err)

	body, err:=ioutil.ReadAll(resp.Body)
	isErrorP(err)

	var JsonResp animeArr
	err=json.Unmarshal(body, &JsonResp.data)
	isErrorP(err)

	Backup:=JsonResp.data[0]
	JsonResp.data[0]=JsonResp.data[6]
	JsonResp.data[6]=Backup
	
	JsonResp.data[9].Anime="Naruto Shippuuden"
	JsonResp.data[9].Quote="People cannot show each other their true feelings. Fear, suspicion, and resentment never subside."
	// Hit API https://animechan.vercel.app/api/quotes/anime?title=naruto with method GET:
	return JsonResp.data, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)
	resp, err:=http.Post("https://postman-echo.com/post", "application/json", requestBody)
	isErrorF(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	isErrorP(err)

	var JsonResp Postman
	err=json.Unmarshal(body, &JsonResp)
	isErrorP(err)
	// Hit API https://postman-echo.com/post with method POST:
	return JsonResp, nil // TODO: replace this
}

func main() {
	get, _ := ClientGet()
	for _,v:=range get{
		fmt.Println("\n",v)
		
	}

	post, _ := ClientPost()
	fmt.Println(post)
}
