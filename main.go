// Copyright (c) 2015-2018, Christopher Hall
// see: LICENSE
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func main() {

	server := &http.Server{
		Addr:           "127.0.0.1:8016",
		Handler:        pageHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}

type aHandler struct{}

func pageHandler() http.Handler {

	return &aHandler{}
}

const htmlPage = `<!DOCTYPE html>
<html>
<head><meta charset="UTF-8">
<title>{{ .Title }}</title>
<link rel="shortcut icon" href="/icon/{{ .Icon }}.png" />
<script type="text/javascript">
function new_tab(url) {
    window.open(url, '_blank');
    return false;
}
function do_link() {
    var url = document.getElementById('the-url').value;
    if ('' === url) {
        return false;
    }
    if ('.' == url[0]) {
        new_tab(url.substring(1));
    } else if (/^https?:/.test(url)) {
        new_tab(url);
    } else {
        new_tab('http://' + url);
    }
    return false;
}
</script>
</head>
<body>
  <h1>{{ .Title }}</h1>
  <form onsubmit="return do_link();">
    <input name="url" id="the-url" size="60">
    <br />
    <br />
    <a href="/New%20Tab" onclick="return do_link();" target="_blank">New Tab</a>
  </form>
</body>
</html>
`

type data struct {
	Title string
	Icon  string
}

var tp = template.Must(template.New("page").Parse(htmlPage))

const defaultIcon = "tick"

var iconRegexp = regexp.MustCompilePOSIX("^/icon/(.+)[.]png$")

func (f aHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//fmt.Printf("req: %v\n", r)
	//fmt.Printf("path: %q\n", r.URL.Path)

	if "/favicon.png" == r.URL.Path || "/favicon.ico" == r.URL.Path {
		r.Header.Add("Content-Type", "image/png")
		w.Write(icons[defaultIcon])
		return
	} else if strings.HasPrefix(r.URL.Path, "/icon/") {
		r.Header.Add("Content-Type", "image/png")
		m := iconRegexp.FindStringSubmatch(r.URL.Path)
		if len(m) > 1 {
			if icon, ok := icons[m[1]]; ok {
				w.Write(icon)
				return
			}
		}
		w.Write(icons[defaultIcon])
		return
	}

	str := r.URL.Path
	if '/' == str[0] {
		str = str[1:] // strip leading '/'
	}
	runes := []rune(str)
	first := string(unicode.ToLower(runes[0]))
	//fmt.Printf("first: %q\n", runes[0])
	//fmt.Printf("first: %q\n", string(unicode.ToLower(runes[0])))

	d := data{
		Title: str,
		Icon:  first,
	}
	err := tp.Execute(w, d)
	if nil != err {
		fmt.Printf("error: %s\n", err)
	}
}
