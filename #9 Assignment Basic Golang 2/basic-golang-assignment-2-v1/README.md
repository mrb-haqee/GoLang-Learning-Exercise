# Dasar Pemrograman Backend

## Student Portal 2

### Description

Program ini adalah sebuah student portal sederhana yang digunakan untuk melakukan registrasi mahasiswa baru, login, mengambil program studi dari suatu mahasiswa, dan mengubah program studi dari suatu mahasiswa.

### Cara Penggunaan

1. Jalankan program dengan menjalankan `go run main.go`
2. Pilih menu yang diinginkan:
   - Pilihan 1: Login
     - Masukkan `id` dan `name` yang sesuai dengan data mahasiswa yang telah terdaftar.
       - Jika `id` atau `name` tidak di isi, program akan menampilkan pesan `"ID or Name is undefined!"`.
       - Jika gagal, program akan menampilkan pesan `"Login gagal: data mahasiswa tidak ditemukan"`.
       - Jika berhasil, program akan menampilkan pesan `"Login berhasil: <nama>"`.
   - Pilihan 2: Register
     - Masukkan `id`, `name`, dan `major` dari mahasiswa baru yang ingin didaftarkan.
       - Jika `id`, `name`, atau `major` tidak di isi, program akan menampilkan pesan `"ID, Name or Major is undefined!"`.
       - Jika gagal, program akan menampilkan pesan `"Registrasi gagal: id sudah digunakan"`.
       - Jika berhasil, program akan menampilkan pesan `Registrasi berhasil: <nama> (<major>)`.
   - Pilihan 3: Get Program Study
     - Masukkan kode program studi yang ingin dicari.
       - Jika kode program studi ditemukan, program akan menampilkan nama program studi yang sesuai.
       - Jika tidak ditemukan, program akan menampilkan pesan `"Kode program studi tidak ditemukan"`.
   - Pilihan 4: Change student study program
     - Masukkan nama mahasiswa yang ingin diubah program studinya dan masukkan program studi baru yang diinginkan.
       - Jika nama mahasiswa ditemukan, program studi mahasiswa akan diubah dan program akan menampilkan pesan `"Program studi mahasiswa berhasil diubah."`.
       - Jika nama mahasiswa tidak ditemukan, program akan menampilkan pesan `"Mahasiswa tidak ditemukan."`
   - Pilihan 5: Keluar
     - Keluar dari program.

Dalam program tersebut terdapat sebuah slice bernama `Students` yang berisi data mahasiswa dengan format string `"id_nama_jurusan"`. Terdapat juga sebuah map `StudentStudyPrograms` yang memetakan kode program studi ke nama program studi. Di sini, kamu diminta untuk melengkapi fungsi di bawah ini dengan ketentuan sebagai berikut:

- Fungsi `Login` menerima input `id` dan `name` yang merupakan id dan nama mahasiswa. Fungsi ini akan memeriksa apakah data mahasiswa terdapat dalam slice `Students`.
  - Jika input `id` atau `name` kosong, program menampilkan pesan `"ID or Name is undefined!"`
  - Jika ditemukan, fungsi akan mengembalikan pesan `"Login berhasil: name"`.
  - Jika tidak ditemukan, fungsi akan mengembalikan pesan `"Login gagal: data mahasiswa tidak ditemukan"`.
- Fungsi `Register` menerima input `id`, `name`, dan `major` yang merupakan id, nama, dan program studi mahasiswa yang baru ingin didaftarkan. Fungsi ini akan memeriksa apakah id mahasiswa telah digunakan sebelumnya.
  - Jika input `id`, `name`, atau `major` kosong, program menampilkan pesan `"ID, Name or Major is undefined!"`
  - Jika id telah digunakan, fungsi akan mengembalikan pesan `"Registrasi gagal: id sudah digunakan"`.
  - Jika id belum pernah digunakan sebelumnya, fungsi akan menambahkan data mahasiswa baru ke dalam slice Students dan mengembalikan pesan `"Registrasi berhasil: name (major)"`.
- Fungsi `GetStudyProgram` menerima input `code` yang merupakan kode program studi. Fungsi ini akan mencari nama program studi yang sesuai berdasarkan kode program studi dari map `StudentStudyPrograms`.
  - Jika ditemukan, fungsi akan mengembalikan nama program studi.
  - Jika tidak ditemukan, fungsi akan mengembalikan pesan `"Kode program studi tidak ditemukan"`.
- Fungsi `ModifyStudent` menerima input `programStudi`, `nama`, dan `fn` yang merupakan program studi baru, nama mahasiswa, dan sebuah fungsi dengan _type_ `studentModifier` yang akan memodifikasi data mahasiswa pada slice `Students`:

  ```go
  type studentModifier func(string, *string)
  ```

  > Pada program ini, parameter fungsi `fn` akan di isi dengan fungsi `UpdateStudyProgram`

  Contoh implementasinya:

  ```go
  ModifyStudent("TI", "Aditira", UpdateStudyProgram)
  ```
  
  Fungsi ini akan mencari data mahasiswa dengan nama yang sesuai pada slice `Students`.
  - Jika ditemukan, fungsi `fn` akan dipanggil dengan input `programStudi` dan data mahasiswa pada slice `Students` akan dimodifikasi menggunakan fungsi `fn` yaitu `UpdateStudyProgram` dan mengembalikan pesan `"Program studi mahasiswa berhasil diubah."`
  - Jika modifikasi gagal, maka program ini akan mengembalikan pesan `"Mahasiswa tidak ditemukan."`.
- Untuk fungsi `UpdateStudyProgram` ini menerima input `programStudi` dengan type `string` dan `students` dengan type `*string` yang merupakan program studi baru dan data mahasiswa dalam format string `"<id>_<nama>_<programstudy>"`. Fungsi ini akan mengubah data mahasiswa pada `students` sehingga program studi mahasiswa tersebut menjadi `programStudi`. Contoh:

  ```go
  var Students = []string{
    "A1234_Aditira_TI",
    "B2131_Dito_TK",
    "A3455_Afis_MI",
  }

  fmt.Println(Students) // Output: "A1234_Aditira_TI", "B2131_Dito_TK", "A3455_Afis_MI",

  UpdateStudyProgram("GA", *Students[1])

  fmt.Println(Students) // Output: "A1234_Aditira_TI", "B2131_Dito_GA", "A3455_Afis_MI",
  ```

  > Terlihat bahwa data `Students` dengan index 1 yaitu `"B2131_Dito_TK"` berubah menjadi `"B2131_Dito_GA"` sesuai dengan parameter `programStudi`

### Test Case Examples

#### Test Case 1

**Input**:

```go
Login("", "") // empty ID and name
```

**Expected Output / Behavior**:

```bash
> ID or Name is undefined!
```

**Explanation**: Fungsi `Login` memeriksa apakah `id` dan `name` kosong. Jika salah satunya kosong, fungsi mengembalikan `"ID or Name is undefined!"`.

#### Test Case 2

**Input**:

```go
Login("A1234", "Aditira")
```

**Expected Output / Behavior**:

```bash
> Login gagal: data mahasiswa tidak ditemukan
```

**Explanation**: Fungsi `Login` memeriksa apakah ada siswa dengan `id` dan `name` yang diberikan dalam `Students` slice. Karena tidak ada siswa seperti itu di dalam slice, fungsi mengembalikan `"Login gagal: data mahasiswa tidak ditemukan"`.

#### Test Case 3

**Input**:

```go
Register("A1234", "Aditira", "TI") // register a new student
Login("A1234", "Aditira") // try to login with the registered student
```

**Expected Output / Behavior**:

```bash
> Registrasi berhasil: Aditira (TI)
> Login berhasil: Aditira
```

**Explanation**: Fungsi `Register` menambahkan siswa baru ke `Students` slice dengan `id`, `name`, dan `major` yang diberikan. Fungsi mengembalikan `"Registrasi berhasil: Aditira (TI)"`. Kemudian, fungsi `Login` mencoba login dengan `id` dan nama siswa yang terdaftar. Karena siswa sudah terdaftar, fungsi Login mengembalikan `"Login berhasil: Aditira"`.

#### Test Case 4

**Input**:

```go
GetStudyProgram("TI") // get study program name for code "TI"
```

**Expected Output / Behavior**:

```bash
> Teknik Informatika
```

**Explanation**: Fungsi `GetStudyProgram` mengembalikan nama program studi yang terkait dengan kode (`"TI"`) yang diberikan, yaitu `"Teknik Informatika"`, dari map `StudentStudyPrograms`.

#### Test Case 5

**Input**:

```go
ModifyStudent("SI", "Afis", UpdateStudyProgram) // change study program for student named "Afis"
GetStudyProgram("MI") // get study program name for code "MI" after modification
```

**Expected Output / Behavior**:

```bash
> Program studi mahasiswa berhasil diubah.
> Manajemen Informasi
```

**Explanation**: Fungsi `ModifyStudent` menemukan siswa bernama `"Afis"` dalam `Students` slice dan mengubah program studinya menjadi `"SI"`. Fungsi mengembalikan `"Program studi mahasiswa berhasil diubah."`. Kemudian, fungsi `GetStudyProgram` mengembalikan nama program studi yang dikaitkan dengan kode `"MI"` yaitu `"Manajemen Informasi"`. Hal ini menegaskan bahwa program studi untuk mahasiswa "Afis" memang telah dimodifikasi.
