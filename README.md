# Nomor Telepon Validator Indonesia

Library ini merupakan open source yang ditulis dalam bahasa pemrograman Golang yang bertujuan untuk membantu pengembang dalam melakukan validasi nomor telepon di Indonesia. 
Library ini dirancang untuk membantu melakukan validasi apakah sebuah nomor telepon sesuai dengan format yang berlaku di Indonesia.

## Versi
Kode versi 1.0

## Fitur Utama

- Validasi Nomor Telepon
- Rubah Menjadi Nomor Internasional
- Rubah Menjadi Nomor Nasional
- Mengetahui Layanan yang Dipakai
- Merubah Menjadi Link WhatsApp
- Merubah Menjadi Link WhatsApp dan Ada Pesan


## Lisensi

Project ini dirilis di bawah lisensi MIT, yang memungkinkan penggunaan, distribusi, dan modifikasi kode dengan syarat mencantumkan hak cipta dan lisensi aslinya.

## Contoh Penggunaan

```go
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
```

Output : <br>

nil
085701234567 <br>
6285701234567 <br>
Indosat Ooredoo IM3 <br>
https://api.whatsapp.com/send/?phone=6285701234567 <br>
https://api.whatsapp.com/send/?phone=6285701234567&text=Hallo <br>

## Cara Berkontribusi dan Penambahan Feature

Langsung saja pull request atau buat issues.




