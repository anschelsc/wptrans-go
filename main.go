package main

import (
	"fmt"
	"http"
	"xml"
	"os"
)

const (
	pageName = "Zamenhof"
	lang = "en"
	template = "http://%s.wikipedia.org/w/api.php?action=query&titles=%s&prop=langlinks&redirects&format=xml"
)

type Redirect struct {
	From string "attr"
	To   string "attr"
}

type Redirects struct {
	R Redirect
}

type Query struct {
	Redirects
}

type Api struct {
	Query
}

func main() {
	response, _, err := http.Get(fmt.Sprintf(template, lang, http.URLEscape(pageName)))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer response.Body.Close()
	result := new(Api)
	xml.Unmarshal(response.Body, result)
	fmt.Println(result.Query.Redirects.R.To)
}
