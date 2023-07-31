package main

import (
	"fmt"
	nomorteleponvalidatorindonesia "github.com/adityarizkyramadhan/nomor-telepon-validator-indonesia"
)

func main() {
	nope := nomorteleponvalidatorindonesia.CekNomor("085701234567")
	fmt.Println(nope.Validasi())
	fmt.Println(nope.NomorNasional())
	fmt.Println(nope.NomorInternasional())
	fmt.Println(nope.PenyediaLayanan())
	fmt.Println(nope.LinkWhatsApp())
	fmt.Println(nope.LinkWhatsAppDenganPesan("Hallo"))
}
