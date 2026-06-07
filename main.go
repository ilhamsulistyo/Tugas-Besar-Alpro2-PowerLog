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
			mencatatPerangkat()
		} else if beranda == 5 {
		} else if beranda == 6 {
			sequentialSearch()
		} else if beranda == 7 {
			SelectionSortPerangkat()
		} else if beranda == 8 {
			InsertionSortWatt()
		} else if beranda == 9 {
			tampilStatistik()
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

		mencatatPerangkat()

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

		mencatatPerangkat()
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

func mencatatPerangkat() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		fmt.Println("Nomor Nama Ruangan Daya Durasi: ")
		for i := 0; i < perangkat; i++ {
			fmt.Printf("%d %s %s %.1f watt %.1f menit\n", i+1, daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)

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
				break
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
				fmt.Printf("Perangkat ditemukan: %s, Ruangan: %s, Daya: %.1f watt, Durasi: %.1f menit\n", daftarElektronik[tengah].nama, daftarElektronik[tengah].ruangan, daftarElektronik[tengah].watt, daftarElektronik[tengah].durasi)
				found = true
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

func SelectionSortPerangkat() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		if perangkat == 0 {
			fmt.Println("=====================================================")
			fmt.Println("Belum ada perangkat yang terdaftar untuk diurutkan!")
			for addloop == false {
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
			return
		}
		var idx_min int

		for i := 0; i < perangkat-1; i++ {
			idx_min = i
			for j := i + 1; j < perangkat; j++ {
				if daftarElektronik[idx_min].watt < daftarElektronik[j].watt {
					idx_min = j
				}
			}
			temp := daftarElektronik[idx_min]
			daftarElektronik[idx_min] = daftarElektronik[i]
			daftarElektronik[i] = temp
		}
		fmt.Println("=====================================================")
		fmt.Println("Konsumsi daya tertinggi sampai terendah")
		for i := 0; i < perangkat; i++ {
			fmt.Printf("%d. Nama: %s, Ruangan: %s, Daya: %.1f watt, Durasi: %.1f menit\n", i+1, daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)
		}
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

func InsertionSortWatt() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		if perangkat == 0 {
			fmt.Println("=====================================================")
			fmt.Println("Belum ada watt yang terdaftar untuk diurutkan!")
			for addloop == false {
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
			return
		}

		var i int
		for i = 1; i < perangkat; i++ {
			key := daftarElektronik[i]
			j := i - 1

			for j >= 0 && daftarElektronik[j].watt < key.watt {
				daftarElektronik[j+1] = daftarElektronik[j]
				j = j - 1
			}
			daftarElektronik[j+1] = key
		}
		fmt.Println("=====================================================")
		fmt.Println("Konsumsi daya tertinggi sampai terendah")
		for i := 0; i < perangkat; i++ {
			fmt.Printf("%d. Nama: %s, Ruangan: %s, Daya: %.1f watt, Durasi: %.1f menit\n", i+1, daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)
		}
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

func konsumsiHarian(e elektronik) float64 {
	return e.watt * (e.durasi / 60)
}

func tampilStatistik() {
	var pilihan string
	var addloop bool = false
	for addloop == false {
		if perangkat == 0 {
			fmt.Println("Belum ada perangkat yang tercatat.")
			return
		}

		var totalDayaHarian float64 = 0

		for i := 0; i < perangkat; i++ {
			totalDayaHarian += daftarElektronik[i].watt
		}

		var temp [1000]elektronik
		for i := 0; i < perangkat; i++ {
			temp[i] = daftarElektronik[i]
		}

		for i := 0; i < perangkat-1; i++ {
			for j := 0; j < perangkat-i-1; j++ {
				if temp[j].watt < temp[j+1].watt {
					temp[j], temp[j+1] = temp[j+1], temp[j]
				}
			}
		}
		konsumsi := konsumsiHarian(temp[0])

		fmt.Println("=====================================================")
		fmt.Println("=== Statistik Konsumsi Listrik Harian ===")
		fmt.Printf("Total Penggunaan Daya : %.2f Watt-menit (%.2f Wh)\n", totalDayaHarian, totalDayaHarian/60)
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Daftar Perangkat Paling Boros Energi:")
		fmt.Printf("- Nama     : %s\n", temp[0].nama)
		fmt.Printf("- Ruangan  : %s\n", temp[0].ruangan)
		fmt.Printf("- Konsumsi : %.2f Wh\n", konsumsi)
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
