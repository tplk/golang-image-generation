package main

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"fmt"
)

func InitServer() {
	var generator = GenerateImage
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc(
		"/VericalGradient",
		generateImageResponse(
			generator,
			ConvertToGray16AlgoFunc(VericalGradient),
		),
	)
	http.HandleFunc(
		"/HorizontalGradient",
		generateImageResponse(
			generator,
			ConvertToGray16AlgoFunc(HorizontalGradient),
		),
	)
	http.HandleFunc(
		"/CornerGradient",
		generateImageResponse(
			generator,
			ConvertToGray16AlgoFunc(CornerGradient),
		),
	)
	http.HandleFunc(
		"/CryptoRandom",
		generateImageResponse(
			generator,
			ConvertToGray16AlgoFunc(CryptoRandom),
		),
	)
	http.HandleFunc(
		"/CryptoRandomThreshold",
		generateImageResponse(
			generator,
			ConvertToGray16AlgoFunc(CryptoRandomThreshold(0.5)),
		),
	)
	http.HandleFunc(
		"/SimplexNoise",
		generateImageResponse(
			generator,
			ConvertToGray16AlgoFunc(SimplexNoise()),
		),
	)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func generateImageResponse(
	generator func(algo algoFunc) string,
	algoFunc algoFunc,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("serve-image.go: %v\n", r.RequestURI)
		w.Header().Set("Content-Type", "image/png")
		content, err := base64.StdEncoding.DecodeString(generator(algoFunc))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(content)))
		w.Write(content)
		return
	}
}
