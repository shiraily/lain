package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `<html>
こんにちは。梅雨の時期の洗濯に悩んでいませんか。<br>
<a href="https://www.irisplaza.co.jp/media/A13916021914">梅雨時期の洗濯のコツとカビ対策記事</a>を紹介しますのでご覧ください。
</html>`); err != nil {
			return
		}
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("failed to listen/serve")
	}
	return
}
