package main

import "fmt"

type elektronik struct {
	nama    string
	ruangan string
	watt    float64
	durasi  float64
}

var daftarElektronik [1000]elektronik
var perangkat int

func main() {
	var beranda int
	var berandaloop bool = false
	for berandaloop == false {
		fmt.Println("=====================================================")
		fmt.Println("Selamat datang di PowerLog!")
		fmt.Println("1. Tambah Perangkat Elektronik")
		fmt.Println("2. Ubah Perangkat Elektronik")
		fmt.Println("3. Hapus Perangkat Elektronik")
		fmt.Println("4. Tampilkan Perangkat Elektronik")
		fmt.Println("5. Sequential Search Perangkat Elektronik")
		fmt.Println("6. Binary Search Perangkat Elektronik")
		fmt.Println("7. Selection Sort Perangkat Elektronik")
		fmt.Println("8. Insertion Sort Perangkat Elektronik")
		fmt.Println("9. Statistik Perangkat Elektronik")
		fmt.Println("0. Keluar")
		fmt.Println("=====================================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&beranda)

		if beranda == 1 {
			addelektronik()
		} else if beranda == 2 {
			updateelektronik()
		} else if beranda == 3 {
			deleteelektronik()
		} else if beranda == 4 {
		} else if beranda == 5 {
			sequentialSearch()
		} else if beranda == 6 {
			binarySearch()
		} else if beranda == 7 {
		} else if beranda == 8 {

		} else if beranda == 9 {
		} else if beranda == 0 {
			fmt.Println("Program selesai")
			berandaloop = true
		} else {
			fmt.Println("Beranda tidak tersedia")
		}
	}
}

func addelektronik() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		fmt.Println("=====================================================")
		fmt.Println("Masukkan nama perangkat: ")
		fmt.Scan(&daftarElektronik[perangkat].nama)

		fmt.Println("Masukkan ruangan: ")
		fmt.Scan(&daftarElektronik[perangkat].ruangan)

		fmt.Println("Masukkan daya (watt): ")
		fmt.Scan(&daftarElektronik[perangkat].watt)

		fmt.Println("Masukkan durasi penggunaan (Jam): ")
		fmt.Scan(&daftarElektronik[perangkat].durasi)

		perangkat++
		fmt.Println("Perangkat berhasil ditambahkan!")

		fmt.Println("=====================================================")

		fmt.Println("Ketik 0 untuk Exit ke beranda")
		fmt.Println("Ketik 1 untuk Lanjut")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
		if pilihan == "0" {
			addloop = true
		} else {
			addloop = false
		}

	}
}

func updateelektronik() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		var idx int

		fmt.Println("Nomor Nama Ruangan Daya Durasi: ")
		for i := 0; i < perangkat; i++ {
			fmt.Printf("%d %s %s %.1f watt %.1f menit\n", i+1, daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)

		}

		fmt.Print("Masukkan nomor perangkat yang ingin diubah: ")
		fmt.Scan(&idx)
		idx--

		if idx >= 0 && idx < perangkat {
			fmt.Print("Masukkan nama perangkat baru: ")
			fmt.Scan(&daftarElektronik[idx].nama)

			fmt.Print("Masukkan ruangan baru: ")
			fmt.Scan(&daftarElektronik[idx].ruangan)

			fmt.Print("Masukkan daya baru (watt): ")
			fmt.Scan(&daftarElektronik[idx].watt)

			fmt.Print("Masukkan durasi penggunaan baru (Jam): ")
			fmt.Scan(&daftarElektronik[idx].durasi)

			fmt.Println("Perangkat berhasil diubah!")
		} else {
			fmt.Println("Perangkat tidak ditemukan!")
		}

		fmt.Println("Ketik 0 untuk Exit ke beranda")
		fmt.Println("Ketik 1 untuk Lanjut")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
		if pilihan == "0" {
			addloop = true
		} else {
			addloop = false
		}

	}
}

func deleteelektronik() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		var idx int

		fmt.Println("Nomor Nama Ruangan Daya Durasi: ")
		for i := 0; i < perangkat; i++ {
			fmt.Printf("%d %s %s %.1f watt %.1f menit\n", i+1, daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)

		}
		fmt.Print("Masukkan nomor perangkat yang ingin dihapus: ")
		fmt.Scan(&idx)
		idx--

		if idx >= 0 && idx < perangkat {
			for i := idx; i < perangkat; i++ {
				daftarElektronik[i] = daftarElektronik[i+1]
			}
			perangkat--
			fmt.Println("Perangkat berhasil dihapus!")
		} else {
			fmt.Println("Perangkat tidak ditemukan!")
		}
		fmt.Println("Ketik 0 untuk Exit ke beranda")
		fmt.Println("Ketik 1 untuk Lanjut")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
		if pilihan == "0" {
			addloop = true
		} else {
			addloop = false
		}

	}
}

func sequentialSearch() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		var namaCari string
		fmt.Print("Masukkan nama perangkat yang ingin dicari: ")
		fmt.Scan(&namaCari)
		var found bool = false
		for i := 0; i < perangkat; i++ {
			if daftarElektronik[i].nama == namaCari {
				fmt.Printf("Perangkat ditemukan: %s, Ruangan: %s, Daya: %.1f watt, Durasi: %.1f menit\n", daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)
				found = true
			}
		}
		if !found {
			fmt.Println("Perangkat tidak ditemukan!")
		}
		fmt.Println("Ketik 0 untuk Exit ke beranda")
		fmt.Println("Ketik 1 untuk Lanjut")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
		if pilihan == "0" {
			addloop = true
		} else {
			addloop = false
		}

	}
}

func binarySearch() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		var ruanganCari string
		fmt.Print("Masukkan ruangan yang ingin dicari: ")
		fmt.Scan(&ruanganCari)

		var found bool = false
		kiri := 0
		kanan := perangkat - 1

		for kiri <= kanan {
			tengah := (kiri + kanan) / 2
			if daftarElektronik[tengah].ruangan == ruanganCari {
				found = true
				i := tengah
				for i >= 0 && daftarElektronik[i].ruangan == ruanganCari {
					fmt.Printf("Perangkat ditemukan: %s, Ruangan: %s, Daya: %.1f watt, Durasi: %.1f menit\n",
						daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)
					i--
				}
				i = tengah + 1
				for i < perangkat && daftarElektronik[i].ruangan == ruanganCari {
					fmt.Printf("Perangkat ditemukan: %s, Ruangan: %s, Daya: %.1f watt, Durasi: %.1f menit\n",
						daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)
					i++
				}
				break
			} else if daftarElektronik[tengah].ruangan < ruanganCari {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}
		if !found {
			fmt.Println("Perangkat tidak ditemukan!")
		}
		fmt.Println("Ketik 0 untuk Exit ke beranda")
		fmt.Println("Ketik 1 untuk Lanjut")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
		if pilihan == "0" {
			addloop = true
		} else {
			addloop = false
		}
	}
}
