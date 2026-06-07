# 🏠POWERLOG - Sistem Pemantauan Energi(listrik) Rumah Tangga

POWERLOG adalah aplikasi sederhana berbasis **Golang** untuk mempermudah pencatatan energi rumah tangga

# 📝Program ini dibuat untuk memudahkan permasalahan rumah tangga, lebih spesifik pada pencatatan listrik rumah tangga, antara lain:
- Digitalisasi data alat listrik.
- Kalkulasi konsumsi energi harian.
- Deteksi perangkat paling boros.

# 🛠️Program ini berfungsi untuk melatih pemahaman kita di bidang:
- Struct dan Array global.
- Shifting indeks array (Hapus/Ubah).
- Sequential Search pada nama.
- Binary Search pada ruangan.
- Selection Sort pada abjad.
- Insertion Sort pada watt.

# 📊Fitur program ini antara lain adalah:
- CRUD perangkat elektronik.
- Pencarian Nama dan Ruangan.
- Pengurutan Energi Tertinggi dan Abjad.
- Kalkulator statistik total daya.
- Looping menu beranda utama.


# 📁Struktur Data

```go
type elektronik struct {
	nama    string
	ruangan string
	watt    float64
	durasi  float64
}

```

# 📝Menu Program

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

Power Log adalah aplikasi sederhana berbasis **Golang** untuk memantau dan mencatat penggunaan daya listrik pada berbagai perangkat elektronik rumah tangga menggunakan konsep algoritma dan struktur data dasar.

Program ini dibuat untuk melatih pemahaman tentang:

- Array statis
- Struct
- CRUD (Create, Read, Update, Delete)
- Searching (Sequential & Binary Search)
- Sorting (Selection & Insertion Sort)
- Statistik data
- Perulangan & Percabangan

---

# 🚀 Fitur Program

- Tambah data tagihan
- Ubah data tagihan
- Hapus data tagihan
- Menampilkan seluruh data
- Sequential Search (nama / perangkat)
- Binary Search (berdasarkan ruang)
- Selection Sort
- Insertion Sort
- Statistik total penggunaan daya harian dan daftar perangkat yang paling boros energi

---

# 🧱 Struktur Data

```go
type Tagihan struct {
    nama    string
	ruangan string
	watt    float64
	durasi  float64
}

| Atribut    | Tipe Data | Keterangan          |
| ---------- | --------- | ------------------- |
| Nama       | string    | Nama Perangkat      |
| Ruangan    | string    | Nama Ruangan        |
| Watt       | float64   | Besaran Watt        |
| Durasi     | float64   | Lama Penggunaan     |
```

# 📋 Menu Program

```bash
===== SIMTAB =====

1. Tambah Tagihan
2. Ubah Tagihan
3. Hapus Tagihan
4. Tampilkan Data
5. Sequential Search
6. Binary Search
7. Selection Sort
8. Insertion Sort
9. Statistik
10. Keluar
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

## ➕ Tambah Perangkat

| Input       | Contoh  |
| ----------- | ------- |
| Nama        | Kulkas  |
| Ruangan     | Dapur   |
| Watt        | 120     |
| Durasi      | 24      |

---

## 📄 Tampilkan Data

| Menu | Fungsi                 |
| ---- | ---------------------- |
| 4    | Menampilkan semua data |

---

## ✏️ Ubah Data

| Langkah | Keterangan       |
| ------- | ---------------- |
| 1       | Pilih menu 2     |
| 2       | Pilih nomor data |
| 3       | Input data baru  |

---

## ❌ Hapus Data

| Langkah | Keterangan              |
| ------- | ----------------------- |
| 1       | Pilih menu 3            |
| 2       | Pilih data yang dihapus |

---

## 🔎 Sequential Search

| Fitur     | Keterangan         |
| --------- | ------------------ |
| Menu      | 5                  |
| Pencarian | Nama atau Kategori |
| Metode    | Linear Search      |

---

## ⚡ Binary Search

| Fitur  | Keterangan                     |
| ------ | ------------------------------ |
| Menu   | 6                              |
| Data   | Harus terurut berdasarkan Nama |
| Metode | Binary Search                  |

---

## 📊 Sorting Data

| Fitur  | Keterangan                     |
| ------ | ------------------------------ |
| Menu   | 7                              |
| Field  | Nama / Ruangan / Watt / Durasi |
| Urutan | Ascending / Descending         |
| Metode | Selection Sort                 |

---

## Insertion Sort

| Fitur  | Keterangan                     |
| ------ | ------------------------------ |
| Menu   | 8                              |
| Field  | Nama / Ruangan / Watt / Durasi |
| Urutan | Ascending / Descending         |
| Metode | Insertion Sort                 |

---

## 📈 Statistik Konsumsi Listrik

| Statistik                  | Keterangan                       |
| -------------------------- | -------------------------------- |
| Total penggunaan daya      | Jumlah total seluruh penggunaan  |
| Perangkat Terboros         | Daftar perangkat terboros        |


---

# 🧠 Algoritma yang Digunakan

| Algoritma         | Fungsi                           |
| ----------------- | -------------------------------- |
| Sequential Search | Pencarian data secara linear     |
| Binary Search     | Pencarian data pada data terurut |
| Selection Sort    | Mengurutkan data (selection)     |
| Insertion Sort    | Mengurutkan data (insertion)     |
| Array             | Menyimpan data                   |
| Struct            | Struktur data tagihan            |
| Perulangan        | Menampilkan & memproses data     |
| Percabangan       | Logika menu & kondisi program    |

---

# 📦 Contoh Output

```bash
============================================================================
Pilih menu: 4
============================================================================
Nomor Nama Ruangan Daya Durasi: 
Nomor           Nama            Ruangan         Daya(Watt)      Durasi(Jam)
1               Kulkas          Dapur           120             24
2               Kipas Angin     Kamar Tidur     50              8
3               Televisi	    Ruang Keluarga	100         	5
4	            AC	            Kamar Utama	    400         	10
5	            Mesin Cuci	    Area Cuci	    300         	2          
============================================================================
Ketik 0 untuk Exit ke beranda atau Ketik 1 untuk Lanjut
```
