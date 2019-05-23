package test

import (
	"fmt"
	"regexp"
)

func UrlGroup() {
	url := "http://www.flysnow.org/2018/02/09/go-regexp-extract-text.html"
	regexpString := `^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`
	regexpCompile := regexp.MustCompile(regexpString)
	params := regexpCompile.FindStringSubmatch(url)

	for _, param := range params {
		fmt.Println(param)
	}
}
