package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func Get(w http.ResponseWriter) {

	data := []byte(`{
		"max_results": 500000,
		"_source":["message_Id","to", "from","date","subject","message"]
	}`)
	h := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:4080/api/enron_mail/_search", bytes.NewBuffer(data))
	req.SetBasicAuth("admin", "Complexpass#123")

	rs, err := h.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if rs.Body != nil {
		defer rs.Body.Close()
	}
	body, readErr := ioutil.ReadAll(rs.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var result Response
	json.Unmarshal(body, &result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

type Response struct {
	Took      int  `json:"took"`
	Timed_Out bool `json:"timed_out"`
	Shards    struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Max_Score int `json:"max_score"`
		Hits      []struct {
			Index     string `json:"_index"`
			Type      string `json:"_type"`
			Id        string `json:"_id"`
			Score     int    `json:"_score"`
			TimeStamp string `json:"@timestamp"`
			Source    struct {
				Subject    string `json:"Subject"`
				Date       string `json:"Date"`
				From       string `json:"From"`
				Message    string `json:"Message"`
				Message_ID string `json:"Message-ID"`
				To         string `json:"To"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
