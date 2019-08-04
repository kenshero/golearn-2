package newsmodel

import (
	"fmt"
)

var news = []string{}

func AddNews() {
	var data string = "Test"
	news = append(news, data)
	fmt.Println(news)
}

func GetNews() []string {
	return news
}

func DeleteNews() []string {

	return news
}
