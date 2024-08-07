package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/metallust/textbg/pkg/request"
)

func main() {

	fmt.Println("Set Background Image...")

	fmt.Print("Enter Number of lines :")
	var n int
	fmt.Scanf("%d", &n)
	var texts []string
	var fonts []int
	for i := 0; i < n; i++ {
		var s string
        var f int

        fmt.Print("Text : ")
		fmt.Scanln(&s)
        fmt.Print("Fontsize : ")
		fmt.Scanf("%d", &f)
		texts = append(texts, s)
        fonts = append(fonts, f)
	}

	var spacing int
	fmt.Scanf("%d", &spacing)

    url := "<server url>"

    req := request.Setin{
    	Sentence: texts,
    	FontSize: fonts,
    	Spacing:  spacing,
    }

    jsonbytes, _ := json.Marshal(req)
    res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonbytes))
    if err != nil {
        panic(err)
    }
    var result []byte
    res.Body.Read(result)
    fmt.Println("Done ... status:" + res.Status + string(result))

}
