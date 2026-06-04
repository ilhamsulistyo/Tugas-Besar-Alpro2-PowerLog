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
			InsertionSortWatt()
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
		fmt.Print("Masukkan nama perangkat: ")
		fmt.Scan(&daftarElektronik[perangkat].nama)

		fmt.Print("Masukkan ruangan: ")
		fmt.Scan(&daftarElektronik[perangkat].ruangan)

		fmt.Print("Masukkan daya (watt): ")
		fmt.Scan(&daftarElektronik[perangkat].watt)

		fmt.Print("Masukkan durasi penggunaan (menit): ")
		fmt.Scan(&daftarElektronik[perangkat].durasi)

		perangkat++
		fmt.Println("Perangkat berhasil ditambahkan!")

		fmt.Println("=====================================================")
		
		fmt.Println("Ketik Exit untuk kembali ke beranda")
		fmt.Scan(&pilihan)
		if pilihan == "Exit" {
			addloop = true
		}

	}
}

func updateelektronik() {
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

		fmt.Print("Masukkan durasi penggunaan baru (menit): ")
		fmt.Scan(&daftarElektronik[idx].durasi)

		fmt.Println("Perangkat berhasil diubah!")
	} else {
		fmt.Println("Perangkat tidak ditemukan!")
	}
}

func deleteelektronik() {
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
}


func SelectionSortPerangkat(perangkat string) {


	var idx_min int
	var T elektronik

	for i:= 0; i < length(T[].perangkat); i++ {
		idx_min = i - 1
		j := i
		for j < n {
			if [idx_min].perangkat < T[j].perangkat {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}

}

func InsertionSortWatt(watt int) {
	


	for i:= 0; i < length(T[].watt); i++ {
		idx_min = i - 1
		j := i
		for j < n {
			if [idx_min].watt < T[j].watt {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}

}