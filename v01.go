package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Link01 struct {
	Version string `json:"version"`
	Title   Feed   `json:"Feed"`
}
type Feed struct {
	Xmlns string `json:"xmlns"`
	Title Title  `json:"title"`
	Entry Entry  `json:"entry"` //!!!!!!!!!!!!!!!!  I believe that this is the error.
}

type Entry struct {
	Title2 Title2 `json:"title"` // !!! or this is the error.  I'm not sure how to import a slice
}

type Title2 struct {
	Title3 string `json:"$t"`
}

type ID struct {
	t string `json:"$t"`
}

type Content struct {
	tpee string `json:"$t"`
}

type Title struct {
	Tpe   string `json:"type"`
	Title string `json:"$t"`
}

type Fo struct {
	Message  string
	Subtitle string
}

func main() {
	url01 := "https://spreadsheets.google.com/feeds/list/1brWkVApKRJ3aRL6OfIQ4jlyx-wXmDhfpGjHIK6ACdQE/1/public/values?alt=json"

	resp, err := http.Get(url01)
	if err != nil {
		// handle error
		println("=============  error in first get =============")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		println("=============  error in 2ND get =============")
	}
	println("===============   the body =====================")
	println(body)
	println("=========   this does not work very well  ==========")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url01, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/json")
	resp02, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp02.Body.Close()
	decoder := json.NewDecoder(resp02.Body)
	//	v := Fo{}
	v := Link01{}
	err = decoder.Decode(&v)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("============  logging the result =================")
	log.Println(v)
	log.Println(v.Title.Title)
	log.Println(v.Title.Title.Title)
	log.Println(v)

}
