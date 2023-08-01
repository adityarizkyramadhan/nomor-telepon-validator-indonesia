package test

import (
	nomorteleponvalidatorindonesia "github.com/adityarizkyramadhan/nomor-telepon-validator-indonesia"
	"testing"
)

func TestNomorHpRandom(t *testing.T) {
	testCase := []struct {
		Nama         string
		NomorTelepon string
		expected     error
	}{
		{
			Nama:         "Nomor Telepon Kosong",
			NomorTelepon: "",
			expected:     nomorteleponvalidatorindonesia.ErrPanjangTidakSesuai,
		}, {
			Nama:         "Nomor Telepon Mengandung Huruf",
			NomorTelepon: "08123456789a",
			expected:     nomorteleponvalidatorindonesia.ErrMengandungHuruf,
		},
		{
			Nama:         "Nomor Telepon Kurang Dari 10 Digit",
			NomorTelepon: "081234567",
			expected:     nomorteleponvalidatorindonesia.ErrPanjangTidakSesuai,
		},
		{
			Nama:         "Nomor Telepon Lebih Dari 13 Digit",
			NomorTelepon: "08123456789123",
			expected:     nomorteleponvalidatorindonesia.ErrPanjangTidakSesuai,
		},
		{
			Nama:         "Nomor Telepon Tidak Menggunakan Prefix 0 atau 62",
			NomorTelepon: "81234567891",
			expected:     nomorteleponvalidatorindonesia.ErrPrefix,
		},
		{
			Nama:         "Nomor Telepon Menggunakan Prefix 0",
			NomorTelepon: "081234567891",
			expected:     nil,
		},
		{
			Nama:         "Nomor Telepon Menggunakan Prefix 62",
			NomorTelepon: "6281234567891",
			expected:     nil,
		},
		{
			Nama:         "Nomor Telepon Tidak Sesuai Prefix Penyedia Layanan",
			NomorTelepon: "099234567890",
			expected:     nomorteleponvalidatorindonesia.ErrLayanan,
		},
		{
			Nama:         "Nomor Telepon Prefix Sesuai oleh Penyedia Layanan",
			NomorTelepon: "081234567891",
			expected:     nil,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Nama, func(t *testing.T) {
			err := nomorteleponvalidatorindonesia.CekNomor(tc.NomorTelepon).Validasi()
			if err != tc.expected {
				t.Errorf("Expected: %v, Got: %v", tc.expected, err)
			}
		})
	}
}
