package main

import (
	"encoding/base64"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func generateQr(url string) string {
	qrCodeImageData, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Println(err)
	}

	encodedData := base64.StdEncoding.EncodeToString(qrCodeImageData)

	return encodedData
}
