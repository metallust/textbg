package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/metallust/textbg/pkg/request"
)

func main() {
	CreateImage([]string{"Grid 6.0"}, []int{75}, 0)
	http.HandleFunc("/", handler)
	http.HandleFunc("/set", set)
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	var in request.Setin
    json.NewDecoder(r.Body).Decode(&in)
	fmt.Println(in)
    CreateImage(in.Sentence, in.FontSize, in.Spacing)
    w.Write([]byte("Done"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "output.png")
}

func CreateImage(sentence []string, fontSize []int, spacing int) {
	// Create a context
	const W = 1920
	const H = 1200
	context := gg.NewContext(W, H)

	// Set background color
	context.SetRGB(0, 0, 0) // Black color
	context.Clear()

	// read font file
	file, err := os.Open("font.ttf")
	if err != nil {
		panic(err)
	}
	fontBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	// parse font
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, v := range fontSize {
		sum += v + spacing
	}

	context.SetRGB(1, 1, 1) // White color
	for i := 0; i < len(sentence); i++ {
		y := H/2 - sum/2 + i*(fontSize[i]+spacing)
		context.SetFontFace(truetype.NewFace(f, &truetype.Options{Size: float64(fontSize[i])}))
		context.DrawStringAnchored(sentence[i], W/2, float64(y), 0.5, 0.5)
	}
	context.SavePNG("output.png")
}
