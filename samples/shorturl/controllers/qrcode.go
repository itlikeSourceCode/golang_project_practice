package controllers

import (
	"github.com/astaxie/beego"
	"github.com/skip2/go-qrcode"
)

type QRCodeController struct {
	beego.Controller
}

func (this *QRCodeController) Get() {
	longurl := this.Input().Get("longurl")

	var png []byte
	png, err := qrcode.Encode(longurl, qrcode.Medium, 256)
	if err != nil {
		this.Abort("500")
	}

	this.Ctx.WriteString(string(png))
}
