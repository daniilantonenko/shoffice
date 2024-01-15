package main

import (
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func generateQrByte(url string) []byte {
	qrCodeImageData, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Println(err)
	}
	return qrCodeImageData
}
