package nomorteleponvalidatorindonesia

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrMengandungHuruf    = errors.New("nomor ponsel mengandung huruf")
	ErrPanjangTidakSesuai = errors.New("panjang nomor ponsel tidak sesuai standar")
	ErrPrefix             = errors.New("nomor telepon tidak sesuai standar prefix")
	ErrLayanan            = errors.New("layanan tidak tersedia")
)

var penyediaLayanan = map[string]string{
	"852": "Telkomsel Kartu AS",
	"853": "Telkomsel Kartu AS",
	"811": "Telkomsel Kartu Halo",
	"812": "Telkomsel Kartu Simpati atau Halo",
	"813": "Telkomsel Kartu Simpati",
	"821": "Telkomsel Kartu Simpati",
	"822": "Telkomsel Kartu Loop",
	"823": "Telkomsel Kartu AS",
	"851": "Telkomsel Kartu AS atau By.u",
	"855": "Indosat Ooredoo Matrix",
	"856": "Indosat Ooredoo IM3",
	"857": "Indosat Ooredoo IM3",
	"858": "Indosat Ooredoo Mentari",
	"814": "Indosat Ooredoo Indosat M2",
	"815": "Indosat Ooredoo Matrix atau Mentari",
	"816": "Indosat Ooredoo Matrix atau Mentari",
	"817": "XL Axiata",
	"818": "XL Axiata",
	"819": "XL Axiata",
	"859": "XL Axiata",
	"877": "XL Axiata",
	"878": "XL Axiata",
	"832": "Axis",
	"833": "Axis",
	"838": "Axis",
	"895": "Three (3)",
	"896": "Three (3)",
	"897": "Three (3)",
	"898": "Three (3)",
	"899": "Three (3)",
	"881": "Smartfren",
	"882": "Smartfren",
	"883": "Smartfren",
	"884": "Smartfren",
	"885": "Smartfren",
	"886": "Smartfren",
	"887": "Smartfren",
	"888": "Smartfren",
	"889": "Smartfren",
}

const (
	minPanjang  = 10
	maksPanjang = 13
)

type NomorTelepon struct {
	Nomor            string
	nomorTanpaPrefix string
	layanan          string
	validasi         bool
}

type NomorTeleponValidator interface {
	Validasi() error
	NomorNasional() string
	NomorInternasional() string
	PenyediaLayanan() string
	LinkWhatsApp() string
	LinkWhatsAppDenganPesan(pesan string) string
}

func CekNomor(nomor string) NomorTeleponValidator {
	return &NomorTelepon{Nomor: nomor, validasi: false}
}

func (nt *NomorTelepon) Validasi() error {
	rgx := regexp.MustCompile(`^[0-9]*$`)
	if !rgx.MatchString(nt.Nomor) {
		return ErrMengandungHuruf
	}

	if len(nt.Nomor) > maksPanjang || len(nt.Nomor) < minPanjang {
		return ErrPanjangTidakSesuai

	}

	if !(nt.Nomor[0] == '0' || nt.Nomor[:2] == "62") {
		return ErrPrefix
	}

	if nt.Nomor[0] == '0' {
		nt.nomorTanpaPrefix = nt.Nomor[1:]
	} else if nt.Nomor[:2] == "62" {
		nt.nomorTanpaPrefix = nt.Nomor[2:]
	} else {
		return ErrPrefix
	}
	layanan := penyediaLayanan[nt.nomorTanpaPrefix[:3]]

	if layanan == "" {
		return ErrLayanan
	}

	nt.validasi = true
	nt.layanan = layanan

	return nil
}

func (nt *NomorTelepon) NomorNasional() string {
	if !nt.validasi {
		return "validasi terlebih dahulu"
	}
	return "0" + nt.nomorTanpaPrefix
}

func (nt *NomorTelepon) NomorInternasional() string {
	if !nt.validasi {
		return "validasi terlebih dahulu"
	}
	return "62" + nt.nomorTanpaPrefix
}

func (nt *NomorTelepon) PenyediaLayanan() string {
	if !nt.validasi {
		return "validasi terlebih dahulu"
	}
	return nt.layanan
}

func (nt *NomorTelepon) LinkWhatsApp() string {
	return "https://api.whatsapp.com/send/?phone=" + nt.NomorInternasional()
}

func (nt *NomorTelepon) LinkWhatsAppDenganPesan(pesan string) string {
	kalimat := strings.ReplaceAll(pesan, " ", "+")
	return nt.LinkWhatsApp() + "&text=" + kalimat
}
