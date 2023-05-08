# Dasar Pemrograman Backend

## Student Portal 1

### Description

Student Portal 1 adalah sebuah program sederhana berbasis terminal yang dibangun dengan menggunakan bahasa pemrograman Go. Program ini menyediakan beberapa menu seperti Login, Register, dan Get Program Study yang dapat digunakan oleh pengguna untuk memanipulasi data mahasiswa.

Data mahasiswa pada program ini disimpan dalam variabel `Students` dengan format string berisi informasi ID, nama, dan jurusan setiap mahasiswa, dipisahkan dengan tanda koma dan spasi antar mahasiswa. Contoh format data mahasiswa dalam variabel `Students` adalah sebagai berikut:

```go
"A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
```

Program juga menyimpan informasi mengenai kode dan nama program studi dalam variabel `studentStudyPrograms` dengan format string berisi informasi kode dan nama setiap program studi, dipisahkan dengan tanda koma dan spasi antar program studi. Contoh format data program studi dalam variabel `studentStudyPrograms` adalah sebagai berikut:

```go
"TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"
```

Program Student Portal 1 memiliki **tiga** fungsi utama, yaitu:

- **Login**: Memeriksa apakah data mahasiswa yang diinput oleh pengguna sudah terdaftar di dalam variabel `Students`.
  - Jika input `id` dan `name` kosong, program menampilkan pesan `"ID or Name is undefined!"`
  - Jika input `id` kurang atau lebih dari 5 karakter, program akan menampilkan pesan `"ID must be 5 characters long!"`
  - Jika data tidak terdaftar, program akan menampilkan pesan `"Login gagal: data mahasiswa tidak ditemukan"`.
  - Jika data terdaftar, maka program akan menampilkan pesan login berhasil beserta nama dan jurusan mahasiswa. Format `"Login berhasil: <nama> (<jurusan mahasiswa>)"` contoh: `"Login berhasil: Aditira (TI)"`

- **Register**: Memeriksa apakah ID yang diinput oleh pengguna sudah terdaftar di dalam variabel `Students`.
  - Jika input `id`, `name` dan `major` kosong, program menampilkan pesan `"ID, Name or Major is undefined!"`
  - Jika input `id` kurang atau lebih dari 5 karakter, program akan menampilkan pesan `"ID must be 5 characters long!"`
  - Jika ID sudah terdaftar, maka program akan menampilkan pesan `"Registrasi gagal: id sudah digunakan"`.
  - Jika ID belum terdaftar, maka program akan menambahkan data mahasiswa baru ke dalam variabel `Students` dan menampilkan pesan registrasi berhasil. Format `"Registrasi berhasil: <nama> (<jurusan mahasiswa>)"` contoh: `"Registrasi berhasil: Aditira (TI)"`

- **Get Study Program**: Mencari nama program studi di dalam variabel `StudentStudyPrograms` berdasarkan kode program studi yang diinput oleh pengguna.
  - Jika input `code` kosong, program menampilkan pesan `"Code is undefined!"`
  - Jika kode dari program studi ditemukan, maka program akan menampilkan nama program studi. Contoh input `"TI"` maka keluar output `"Teknik Informatika"`

Program Student Portal 1 dijalankan dengan menggunakan perulangan `for` yang tidak terbatas, sehingga program akan berjalan terus menerus sampai pengguna memilih menu "4. Keluar".

### Constraints

- ID mahasiswa hanya berupa kombinasi dari huruf dan angka, dengan panjang 5 karakter.
- Nama mahasiswa hanya berupa huruf, dengan panjang antara 1 hingga 50 karakter.
- Jurusan mahasiswa hanya berupa kombinasi dari huruf, dengan panjang antara 1 hingga 50 karakter.
- Kode program studi hanya berupa kombinasi dari huruf, dengan panjang 2 karakter.
- Nama program studi hanya berupa huruf, dengan panjang antara 1 hingga 50 karakter.

### Test Case Examples

#### Test Case 1

**Input**:

```go
Login("A1234", "Aditira")
```

**Expected Output / Behavior**:

```bash
> Login berhasil: Aditira (TI)
```

**Explanation**:

- Fungsi ini dipanggil dengan ID "A1234" dan name "Aditira".
- Fungsi memeriksa setiap data dalam variabel `Students` untuk menemukan ID dan nama yang cocok.
- Fungsi mencari data `"A1234_Aditira_TI"` dan mengekstraksi jurusan `"TI"`.
- Fungsi mengembalikan pesan `"Login berhasil: Aditira (TI)"`.

#### Test Case 2

**Input**:

```go
Register("C1234", "Citra", "TS")
```

**Expected Output / Behavior**:

```txt
> Registrasi berhasil: Citra (TS)
```

**Explanation**:

- Fungsi dipanggil dengan ID "C1234", name "Citra", dan major "TS". (Teknik Sipil)
- Fungsi memeriksa apakah ID belum digunakan dalam variabel `Students`.
- Fungsi menambahkan data baru `"C1234_Citra_TS"` ke variabel `Students`.
- Fungsi mengembalikan pesan `"Registrasi berhasil: Citra (TS)"`.

#### Test Case 3

**Input**:

```go
GetStudyProgram("SI")
```

**Expected Output / Behavior**:

```bash
> Sistem Informasi
```

**Explanation**:

- Fungsi ini dipanggil dengan program code study "SI".
- Fungsi iterasi melalui variabel `studentStudyPrograms` untuk menemukan kode yang cocok.
- Fungsi mencari data `"SI_Sistem Informasi"` dan mengekstrak nama program studi - `"Sistem Informasi"`.
- Fungsi mengembalikan `"Sistem Informasi"`.
