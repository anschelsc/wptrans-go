package main

import (
	"xml"
	"http"
	"strings"
	"fmt"
)

const template = "http://%s.wikipedia.org/w/api.php?action=query&titles=%s&prop=langlinks&redirects&format=xml"

type Dict struct {
	original string
	langs    []string
	trans    map[string]string
}

func emptyDict() *Dict {
	return &Dict{"", make([]string, 0), make(map[string]string)}
}

func (d *Dict) String() string {
	ret := make([]string, 0)
	for i := 0; i != len(d.langs); i++ {
		ret = append(ret, fmt.Sprintf("%s: %s", d.langs[i], d.trans[d.langs[i]]))
	}
	return strings.Join(ret, "\n")
}

func (d *Dict) add(translations []Ll) {
	langs := make([]string, len(translations))
	for i, t := range translations {
		langs[i] = t.Lang
		d.trans[t.Lang] = t.Translation
	}
	d.langs = append(d.langs, langs...)
}

func NewDict(lang, title string) *Dict {
	d := emptyDict()
	parsed := new(Api)
	base := fmt.Sprintf(template, lang, http.URLEscape(title))
	response, _, _ := http.Get(base)
	defer response.Body.Close()
	xml.Unmarshal(response.Body, parsed)
	d.original = parsed.Title
	for {
		d.add(parsed.Ll)
		cont := parsed.Query_continue.Langlinks.Llcontinue
		if cont == "" {
			break
		}
		newResponse, _, _ := http.Get(base + "&llcontinue=" + cont)
		defer newResponse.Body.Close()
		xml.Unmarshal(newResponse.Body, parsed)
	}
	return d
}
