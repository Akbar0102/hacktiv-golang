package main

import (
	"assignment1/helpers"
	"fmt"
	"os"
)

func main() {
	var people = []helpers.Person{
		{Id: "0", Nama: "Akbar", Alamat: "Bandung", Pekerjaan: "Programmer", Alasan: "Nambah skill baru"},
		{Id: "1", Nama: "Penelope", Alamat: "1103 Quilmes Boulevard", Pekerjaan: "Data Science", Alasan: "Update skill"},
		{Id: "2", Nama: "Davis", Alamat: "844 Bucuresti Place", Pekerjaan: "Manager", Alasan: "Cari pengalaman"},
		{Id: "3", Nama: "Grace", Alamat: "1325 Fukuyama Street", Pekerjaan: "Consultan", Alasan: "Coba hal baru"},
		{Id: "4", Nama: "Matthew", Alamat: "1074 Sanaa Parkway", Pekerjaan: "Staff", Alasan: "Tugas dari kantor"},
	}

	var inputNama = os.Args[1]                               //ini inputan dari cmd
	var biodata, exists = helpers.FindBio(inputNama, people) //panggil fungsi dengan param inputan user dan slice of struct
	if exists {                                              //kalau ditemukan (true)
		helpers.ShowBio(biodata)
	} else {
		fmt.Println("Biodata user tidak ditemukan")
	}

}
