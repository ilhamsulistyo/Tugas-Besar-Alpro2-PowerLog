package main

import "fmt"

type elektronik struct {
	perangkat string
	watt      int
	durasi    int
}

var daftarElektronik [1000]elektronik

func main() {
	var input bool = true
	i := 0
	for input {
		fmt.Print("Masukan Barang Elektronik: ")
		fmt.Scanln(&daftarElektronik[i].perangkat)
		if daftarElektronik[i].perangkat == "Selesai" {
			input = false
			break
		}
		fmt.Print("Masukan watt Elektronik: ")
		fmt.Scanln(&daftarElektronik[i].watt)
		fmt.Print("Masukan Durasi Elektronik: ")
		fmt.Scanln(&daftarElektronik[i].durasi)
		i++
	}
}
