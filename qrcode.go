package test

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
)

func GenQrcode() {
	qrcode.WriteFile("http://www.le.com", qrcode.Medium, 256, "./le_qrcode.png")
}

func GenQrcodeBytes() {
	b, err := qrcode.Encode("http://www.le.com", qrcode.Medium, 256)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(b)
	fmt.Println(string(b))
}

func NewQrcode() {
	qr, err := qrcode.New("http://www.le.com", qrcode.Medium)
	if err != nil {
		fmt.Println(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./le2_qrcode.png")
	}
}
