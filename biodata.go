package main

import (
	"fmt"
	"os"
)

// Teman struct represents a friend data - student Hacktiv8

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Database friend - class Golang Hacktiv8
var database = map[int]Teman{

	1: {"Danvers", "jalan kebenaran no 33", "mahasiswa", "Mudah Dipelajari dan relatif mudah dipahami"},
	2: {"herdiyan adam putra", "jalan jati kelapa no 31", "mahasiswa", "performa tinggi dan dapat di kompilasi ke machine code"},
	3: {"fitriana", "jalan kayu putih selatan IV B", "ibu rumah tangga", "golang memiliki built in support untuk microservices"},
	4: {"Kanikayudha", "jalan rawa buntu selamanya", "ojol", "golang menjadi bahasa pemograman yang dibutuhkan di industri"},
	5: {"Nuraeni", "jalan kacang panjang raya", "pedagang", "golang temen nya bolang"},
	6: {"Komang Adyanata", "Jalan kebenaran utama no 99", "lawyer", "golang bahasa yang powerful"},
	7: {"Zikrurridho Afwani", "jalan waru selatan no 13", "ustad", "golang bahasa yang populer"},
	8: {"Muhammad Fariz Al-Pasha", "jalan proklamasi timur no 18", "politisi", "bahasa golang memiliki prospek yang baik untuk pekerjaan saya dan karir saya"},
	9: {"Akbar Alhazir", "jalan merdeka selatan no 86", "buruh", "golang menjadi bahasa yang istimewa saat ini"},
}

// showData untuk menampilkan data friend base line number absen
func showData(absen int) {
	teman, found := database[absen]
	if !found {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan!")

		return

	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih golang:", teman.Alasan)

	return
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run biodata.go [nomor_absen]")
		return
	}

	absen := args[0]
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor Absen harus berupa bilangan bulat!")
		return
	}
	showData(absenInt) // menampilkan data teman berdasarkan no absen
}
