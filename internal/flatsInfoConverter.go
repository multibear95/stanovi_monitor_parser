package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func ConvertFlatInfo(message Message) FlatInfo {
	var flatInfo FlatInfo
	switch chatId := message.ChatId; chatId {
	case -1001825948322:
		flatInfo = parseHtmlContent(retrieveUrl(message), message.Id)
	default:
		fmt.Println("Not implemented")
	}

	return flatInfo
}

func retrieveUrl(message Message) string {
	textContent := message.Content.Text.Text
	isPriceReduceMessage := strings.Index(textContent, "Понравились варианты, выходившие на канале, но цена была кусачей?") != -1
	if isPriceReduceMessage {
		return ""
	}

	urlIdx := strings.Index(textContent, "http")
	if !(urlIdx > -1) {
		return ""
	}
	url := textContent[urlIdx:]
	return url
}

func parseHtmlContent(url string, messageId int) FlatInfo {
	if url == "" {
		return FlatInfo{}
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	contentInBytes, err := ioutil.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return FlatInfo{}
	}
	if err != nil {
		log.Fatal(err)
	}
	htmlContent := string(contentInBytes)

	reInfo := regexp.MustCompile(`<p>(▪️.*)</p>`)
	reNumbers := regexp.MustCompile("[0-9]+")

	tagContents := reInfo.FindAllStringSubmatch(htmlContent, -1)[0][0]
	splitFlatInfo := strings.Split(tagContents, "</p>")

	area := tagContents[strings.Index(tagContents, "https"):strings.Index(tagContents, `" target`)]

	flatInfo := FlatInfo{
		ID:    messageId,
		Price: reNumbers.FindAllString(splitFlatInfo[0], -1)[0],
		Area:  area,
		Size:  reNumbers.FindAllString(splitFlatInfo[1], -1)[0],
	}

	return flatInfo
}
