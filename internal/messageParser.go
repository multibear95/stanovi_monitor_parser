package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseMessages(url string) []FlatInfo {
	var parsedMessages []FlatInfo

	var messages []Message
	unparsedMessages := getJSON(url)
	err := json.Unmarshal(unparsedMessages, &messages)
	if err != nil {
		return nil
	}

	for _, message := range messages {
		convertedFlatInfo := ConvertFlatInfo(message)
		if convertedFlatInfo.ID != 0 {
			parsedMessages = append(parsedMessages, convertedFlatInfo)
		}
	}

	return parsedMessages
}

func getJSON(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	responseAsByteArr, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return responseAsByteArr
}
