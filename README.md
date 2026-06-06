# POWERLOG - Sistem Pemantauan Enegrgi(listrik) Rumah Tangga

POWERLOG adalah aplikasi sederhana berbasis **Golang** untuk mempermudah pencatatan energi rumah tangga

Program ini dibuat untuk memudahkan permasalahan rumah tangga, lebih spesifik pada pencatatan listrik rumah tangga, antara lain:
-Digitalisasi data alat listrik.
-Kalkulasi konsumsi energi harian.
-Deteksi perangkat paling boros.

Program ini berfungsi untuk melatih pemahaman kita di bidang:
-Struct dan Array global.
-Shifting indeks array (Hapus/Ubah).
-Sequential Search pada nama.
-Binary Search pada ruangan.
-Selection Sort pada abjad.
-Insertion Sort pada watt.

Fitur program ini antara lain adalah:
-CRUD perangkat elektronik.
-Pencarian Nama dan Ruangan.
-Pengurutan Energi Tertinggi dan Abjad.
-Kalkulator statistik total daya.
-Looping menu beranda utama.


# Struktur Data

```go
type elektronik struct {
	nama    string
	ruangan string
	watt    float64
	durasi  float64
}

```

# Menu Program

```bash
=====================================================
Selamat datang di PowerLog!
1. Tambah Perangkat Elektronik
2. Ubah Perangkat Elektronik
3. Hapus Perangkat Elektronik
4. Tampilkan Perangkat Elektronik
5. Sequential Search Perangkat Elektronik
6. Binary Search Perangkat Elektronik
7. Selection Sort Perangkat Elektronik
8. Insertion Sort Perangkat Elektronik
9. Statistik Perangkat Elektronik
0. Keluar
=====================================================
```


---

# ⚙️ Cara Menjalankan Program

## 1. Install Golang

Cek apakah Golang sudah terinstall:

```bash
go version
```

---

## 2. Simpan File

Simpan source code dengan nama:

```text
main.go
```

---

## 3. Jalankan Program

```bash
go run main.go
```

---

# 📌 Cara Penggunaan Program
