package main

import "fmt"

const NMAX = 100

type calon struct {
	namaCalon  string
	noUrut     int
	thresehold float64
	jumSuara   int
	partai     string
}

type pemilih struct {
	namaPemilih string
	hakSuara    bool
	pilihan     int
	valid       bool
}

type caleg [NMAX - 1]calon
type pemilihSuara [NMAX - 1]pemilih

func namaCalon(tCaleg caleg) {
	i := 0
	for i < len(tCaleg) {
		fmt.Println("Caleg No.", tCaleg[i].noUrut, ": ")
		fmt.Println("Nama caleg \t: ", tCaleg[i].namaCalon)
		fmt.Println("Jumlah tresehold calon \t: ", tCaleg[i].thresehold)
		fmt.Println("Nama partai \t: ", tCaleg[i].partai)
		fmt.Println()
		i++
	}
}

func pemilihan(tPilihan *pemilihSuara, tCaleg caleg) {
	i := 0
	var nama string
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	for i <= len(tPilihan) && tPilihan[i].valid == false {
		if nama == tPilihan[i].namaPemilih {
			tPilihan[i].valid = true
		}
		tPilihan[i].valid = false
		i++
	}
	if tPilihan[j].valid && tPilihan[j].hakSuara == true {
		fmt.Println("Halo", nama, "! Anda berhak untuk menentukan suara anda")
		namaCalon(tCaleg)
		fmt.Print("Anda akan memilih calon no : ")
		fmt.Scan(&tPilihan[j].pilihan)
		fmt.Print("Terima kasih telah melakukan pemilu")
	}
}

func perhitunganSuara(tPemilih *pemilihSuara, tCaleg caleg) {
	i := 0

	for i < len(tCaleg) {
		if tPemilih[i].pilihan == 1 {
			tCaleg[i].jumSuara++
		} else if tPemilih[i].pilihan == 2 {
			tCaleg[i].jumSuara++
		} else {
			tCaleg[i].jumSuara++
		}
		i++
	}
}
func menu() {
	fmt.Println("=====MENU=====")
	fmt.Println("1. Petugas")
	fmt.Println("2. Pemilih")
	fmt.Println("3. Exit")
	fmt.Println("==============")
	fmt.Print("Anda sebagai : ")

}

func main() {
	var pilihan int
	var tPilihan pemilihSuara
	var tCaleg caleg

	tCaleg[0].namaCalon = "Salma Safira"
	tCaleg[0].noUrut = 1
	tCaleg[0].thresehold = 30.0
	tCaleg[0].partai = "PDIP"

	tCaleg[1].namaCalon = "Alya Sabrina"
	tCaleg[1].noUrut = 2
	tCaleg[1].thresehold = 28.0
	tCaleg[1].partai = "Demokrat"

	tCaleg[2].namaCalon = "Salma Sabrina"
	tCaleg[2].noUrut = 3
	tCaleg[2].thresehold = 33.0
	tCaleg[2].partai = "PPP"

	tPilihan[0].namaPemilih = "User 1"
	tPilihan[0].hakSuara = true

	tPilihan[1].namaPemilih = "User 2"
	tPilihan[1].hakSuara = true

	tPilihan[2].namaPemilih = "User 3"
	tPilihan[2].hakSuara = false

	tPilihan[3].namaPemilih = "User 4"
	tPilihan[3].hakSuara = true

	menu()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Println("Coming Soon")
	case 2:
		pemilihan(&tPilihan, tCaleg)
	}
}
