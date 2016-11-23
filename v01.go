package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Link01 struct {
	Version string `json:"version"`
	Feed    Feed   `json:"Feed"`
}
type Feed struct {
	Xmlns string  `json:"xmlns"`
	Title T1      `json:"title"`
	Entry []Entry `json:"entry"` //!!!!!!!!!!!!!!!!  I believe that this is the error.
}

type Entry struct {
	TitleB  T1       `json:"title"` // !!! or this is the error.  I'm not sure how to import a slice
	Content ContentB `json:"content"`
	GSXID   GSXID    `json:"gsx$id"`
	GSXCOLB GSXCOLB  `json:"gsx$colb"`
	GSXCOLC GSXCOLC  `json:"gsx$colc"`
	GSXCOLD GSXCOLD  `json:"gsx$cold"`
}

type GSXID struct {
	SVal string `json:"$t"`
}

type GSXCOLB struct {
	SVal string `json:"$t"`
}
type GSXCOLC struct {
	SVal string `json:"$t"`
}
type GSXCOLD struct {
	SVal string `json:"$t"`
}

type ContentB struct {
	Tpp string `json:"type"`
	Val string `json:"$t"`
}

// type TitleB struct {
// 	TitleB string `json:"$t"`
// }

type ID struct {
	t string `json:"$t"`
}

type Content struct {
	tpee string `json:"$t"`
}

type T1 struct {
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
	log.Println("============  logging the result end =================")
	log.Println(v.Feed.Title.Title)
	log.Println(v.Feed.Entry[1].Content.Val)
	log.Println("====   doing entry ID ===")
	log.Println(v.Feed.Entry[1].GSXID.SVal)

	log.Println("====   doing entry title ===")
	log.Println(v.Feed.Entry[1].TitleB.Title)

	log.Println("====   doing Content Val ===")
	log.Println(v.Feed.Entry[1].Content.Val)

	log.Println("====   doing entry COLB===")
	log.Println(v.Feed.Entry[1].GSXCOLB.SVal)

	log.Println("====   doing entry COLC===")
	log.Println(v.Feed.Entry[1].GSXCOLC.SVal)

	log.Println("====   doing entry COLD===")
	log.Println(v.Feed.Entry[1].GSXCOLD.SVal)

	log.Println("====   the number of entries ===")
	log.Println(len(v.Feed.Entry))

	num := len(v.Feed.Entry)
	log.Println(num)
	for i := 0; i <= num-1; i++ {
		log.Println(v.Feed.Entry[i].GSXID.SVal + "  =  " + v.Feed.Entry[i].GSXCOLB.SVal + "  =  " + v.Feed.Entry[i].GSXCOLC.SVal + "  =  " + v.Feed.Entry[i].GSXCOLD.SVal)
		// log.Println(v.Feed.Entry[i].GSXID.SVal )
	}

}
