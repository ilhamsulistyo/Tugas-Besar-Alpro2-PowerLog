package main

import "fmt"

func addelektronik(nama, ruangan string, watt, durasi float64) {
	elektronik := elektronik{
		ID:      nextID,
		nama:    nama,
		ruangan: ruangan,
		watt:    watt,
		durasi:  durasi,
	}
	elektroniks = append(elektroniks, elektronik)
	nextID++
	fmt.Printf("Perangkat %s berhasil ditambahkan.\n", nama)
}

func updateelektronik(id int, nama, ruangan string, watt, durasi float64) {
	for i, d := range elektroniks {
		if d.ID == id {
			elektroniks[i].nama = nama
			elektroniks[i].ruangan = ruangan
			elektroniks[i].watt = watt
			elektroniks[i].durasi = durasi
			fmt.Printf("Perangkat dengan ID %d berhasil diperbarui.\n", id)
			return
		}
	}
	fmt.Printf("Perangkat dengan ID %d tidak ditemukan.\n", id)
}

func deleteelektronik(id int) {
	for i, d := range elektroniks {
		if d.ID == id {
			elektroniks = append(elektroniks[:i], elektroniks[i+1:]...)
			fmt.Printf("Perangkat dengan ID %d berhasil dihapus.\n", id)
			return
		}
	}
	fmt.Printf("Perangkat dengan ID %d tidak ditemukan.\n", id)
}

