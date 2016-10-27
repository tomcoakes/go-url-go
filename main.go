package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url1String := r.URL.Query().Get("url1")
		url2String := r.URL.Query().Get("url2")
		fmt.Printf("urls are: %s, %s ", url1String, url2String)
		bodyToPrint, _ := Concatenator(url1String, url2String)
		fmt.Fprint(w, bodyToPrint)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

type urlAndRespBody struct {
	url      string
	respBody string
}

func fetchUrls(urls []string) ([]urlAndRespBody, error) {
	uAndRs := make([]urlAndRespBody, 0)
	responseChannel := make(chan urlAndRespBody, len(urls))
	for _, u := range urls {
		go fetchUrl(u, responseChannel)
		//if err != nil {
		//	continue  // todo handle
		//}
	}
	for i := 0; i < len(urls); i++ {
		uAndRs = append(uAndRs, <-responseChannel)
	}
	return uAndRs, nil
}

func fetchUrl(url string, responseChannel chan urlAndRespBody) {
	resp, err := http.Get(url)
	if err != nil {
		responseChannel <- urlAndRespBody{"oh", "no"}
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }

	responseChannel <- urlAndRespBody{
		url:      url,
		respBody: string(bodyBytes),
	}
}

func Concatenator(url ...string) (megabody string, err error) {
	urlsAndBodies, _ := fetchUrls(url)
	for _, u := range urlsAndBodies {
		megabody += u.respBody
	}
	return
}
