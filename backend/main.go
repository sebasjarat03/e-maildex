package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/indexer/{word}", SearchDocumentsByfilter)

	fmt.Println("Serving on port 3000")
	http.ListenAndServe(":3000", r)
}

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			ID        string                 `json:"_id"`
			Timestamp string                 `json:"@timestamp"`
			Score     float64                `json:"_score"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}

func getEmails(word string) (*SearchResponse, error) {

	bodyResponse := &SearchResponse{}
	query := `{
        "search_type": "matchphrase",
	"query": {
		"term": "` + word + `",
		"start_time": "2022-01-02T14:28:31.894Z",
		"end_time": "2022-12-02T15:28:31.894Z"
	},
	"from": 0,
	"max_results": 1000,
	"_source": []
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Add("Content-Type", "application/json")
	req.Close = true

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	json.NewDecoder(resp.Body).Decode(bodyResponse)

	return bodyResponse, nil
}

type Email struct {
	Id   string   `json:"id"`
	From string   `json:"sender"`
	To   []string `json:"recipient"`

	Subject string `json:"subject"`
	Message string `json:"message"`
}

type searchDocumentsByfilter struct {
	Emails []Email `json:"emails"`
}

func SearchDocumentsByfilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	word := chi.URLParam(r, "word")
	dResponse, err := getEmails(word)

	if err != nil {
		return
	}

	emails, err := mapMails(*dResponse)

	if err != nil {
		fmt.Println("estoy aqui si")
		return

	}

	response := &searchDocumentsByfilter{
		Emails: emails,
	}

	render.JSON(w, r, response)

}
func mapMails(response SearchResponse) ([]Email, error) {
	var emails []Email

	for _, hit := range response.Hits.Hits {
		var email Email

		bytes, err := json.Marshal(hit.Source)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bytes, &email)
		if err != nil {

		}
		emails = append(emails, email)

	}
	print("before -> ", emails)
	return emails, nil
}
