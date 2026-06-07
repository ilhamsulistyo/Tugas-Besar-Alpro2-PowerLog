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
		} else if beranda == 6 {
		} else if beranda == 7 {
			SelectionSortPerangkat()
		} else if beranda == 8 {
			InsertionSortWatt()
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
		fmt.Println("Ketik 0 untuk Exit ke beranda atau Ketik 1 untuk Lanjut")
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

		fmt.Println("=====================================================")
		fmt.Println("Daftar Perangkat Saat Ini:")
		if perangkat == 0 {
			fmt.Println("[ Data Kosong ]")
		} else {
			for i := 0; i < perangkat; i++ {
				fmt.Printf("%d. Nama: %s | Ruangan: %s | Daya: %.1f watt | Durasi: %.1f jam\n",
					i+1, daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt, daftarElektronik[i].durasi)
			}
		}
		fmt.Println("=====================================================")

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

		fmt.Println("=====================================================")
		fmt.Println("Ketik 0 untuk Exit ke beranda atau Ketik 1 untuk Lanjut")
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

		fmt.Println("=====================================================")
		fmt.Println("Daftar Perangkat Saat Ini:")
		if perangkat == 0 {
			fmt.Println("[ Data Kosong ]")
		} else {
			for i := 0; i < perangkat; i++ {
				fmt.Printf("%d. %s (Ruangan: %s, %.1f watt)\n", i+1, daftarElektronik[i].nama, daftarElektronik[i].ruangan, daftarElektronik[i].watt)
			}
		}
		fmt.Println("=====================================================")

		fmt.Print("Masukkan nomor perangkat yang ingin dihapus: ")
		fmt.Scan(&idx)
		idx--

		if idx >= 0 && idx < perangkat {

			for i := idx; i < perangkat-1; i++ {
				daftarElektronik[i] = daftarElektronik[i+1]
			}

			daftarElektronik[perangkat-1] = elektronik{}

			perangkat--
			fmt.Println("Perangkat berhasil dihapus!")
		} else {
			fmt.Println("Perangkat tidak ditemukan atau nomor salah!")
		}

		fmt.Println("=====================================================")
		fmt.Println("Ketik 0 untuk Exit ke beranda atau Ketik 1 untuk Lanjut")
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
			fmt.Println("=====================================================")
		}
		var temp [1000]elektronik
		for i := 0; i < perangkat; i++ {
			temp[i] = daftarElektronik[i]
		}
		var idx_min int

		for i := 0; i < perangkat-1; i++ {
			idx_min = i
			for j := i + 1; j < perangkat; j++ {
				if temp[j].nama < temp[idx_min].nama { // Urut abjad nama
					idx_min = j
				}
			}
			temp[i], temp[idx_min] = temp[idx_min], temp[i]
		}
		fmt.Println("=====================================================")
		fmt.Println("Daftar perangkat diurutkan berdasarkan Abjad Nama:")
		for i := 0; i < perangkat; i++ {
			fmt.Printf("%d. Nama: %s, Ruangan: %s\n", i+1, temp[i].nama, temp[i].ruangan)
		}
		fmt.Println("=====================================================")
		fmt.Println("Ketik 0 untuk Exit ke beranda atau Ketik 1 untuk Lanjut")
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
		var temp [1000]elektronik
		for i := 0; i < perangkat; i++ {
			temp[i] = daftarElektronik[i]
		}

		var i int
		for i = 1; i < perangkat; i++ {
			key := daftarElektronik[i]
			j := i - 1

			for j >= 0 && daftarElektronik[j].watt < key.watt {
				temp[j+1] = temp[j]
				j = j - 1
			}
			temp[j+1] = key
		}
		fmt.Println("=====================================================")
		fmt.Println("Konsumsi daya tertinggi sampai terendah")
		for i := 0; i < perangkat; i++ {
			fmt.Printf("%d. Nama: %s, Ruangan: %s, Daya: %.1f watt, Durasi: %.1f menit\n", i+1, temp[i].nama, temp[i].ruangan, temp[i].watt, temp[i].durasi)
		}
		fmt.Println("=====================================================")
		fmt.Println("Ketik 0 untuk Exit ke beranda atau Ketik 1 untuk Lanjut")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
		if pilihan == "0" {
			addloop = true
		} else {
			addloop = false
		}
	}
}
