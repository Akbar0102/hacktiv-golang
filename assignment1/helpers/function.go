package helpers

import (
	"fmt"
	"strings"
)

func greet() {
	fmt.Println("Hallo dari function greet")
}

func Greet() { //kapital agar bisa dipanggil di package lain
	fmt.Println("Hallo dari function Greet")
}

func FindBio(nama string, p []Person) (Person, bool) {
	var bio = Person{}
	var found bool
	for _, user := range p {
		//cek jika inputan user (di lowerkan) sama dengan properti Nama di slice of struct-nya
		if strings.ToLower(nama) == strings.ToLower(user.Nama) {
			bio = user   //set data user ke bio
			found = true //ditemukan, jadi true
		}
	}

	return bio, found //return data user, dan penanda ditemukan
}

func ShowBio(bio Person) {
	fmt.Println("ID\t: ", bio.Id)
	fmt.Println("Nama\t: ", bio.Nama)
	fmt.Println("Alamat\t: ", bio.Alamat)
	fmt.Println("Alasan\t: ", bio.Alasan)
}
