# Dasar Pemrograman Backend

## Live Coding - Student Management System

### Implementation technique

Siswa akan melaksanakan sesi live code di 15 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 15 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

Proyek ini adalah sebuah program _command-line interface (CLI)_ sederhana yang ditulis dalam bahasa Go, yang memungkinkan kita untuk mengelola data mahasiswa, seperti menambahkan mahasiswa baru, menghapus mahasiswa, menampilkan data mahasiswa, dan mencari mahasiswa berdasarkan nilai.

Program ini memiliki 4 fungsi utama:

- `ViewStudents`: untuk menampilkan data mahasiswa
- `AddStudent`: untuk menambahkan mahasiswa baru ke dalam sistem
- `RemoveStudent`: untuk menghapus mahasiswa dari sistem
- `FindStudent`: untuk mencari mahasiswa berdasarkan nilai dan menampilkan hasil pencarian

Program ini menggunakan `map` untuk menyimpan data mahasiswa, dengan key berupa nama mahasiswa dan value berupa slice yang berisi alamat, nomor telepon, dan nilai mahasiswa.

### Constraints

Program ini dibagi menjadi 4 bagian:

- **Main**: mengontrol keseluruhan aliran program dan memanggil fungsi lainnya
- **ViewStudents**: menampilkan data mahasiswa dalam format tabel
- **AddStudent**: menambahkan mahasiswa baru ke dalam sistem
- **RemoveStudent**: menghapus mahasiswa dari sistem
- **FindStudent**: mencari mahasiswa berdasarkan nilai dan menampilkan hasil pencarian

Program ini menggunakan switch statement untuk memproses input pengguna dan memanggil fungsi yang sesuai untuk mengelola data mahasiswa. Jika pengguna memasukkan input yang tidak valid, program akan menampilkan pesan kesalahan.

Berikut adalah penjelasan dari fungsi-fungsi yang harus diimplementasi:

- Fungsi **RemoveStudent** menerima parameter `*map[string][]interface{}` dan mengembalikan sebuah fungsi yang menerima parameter string, yang akan menghapus data mahasiswa dengan nama yang diberikan dari `map`.

  ```go
  func RemoveStudent(students *map[string][]interface{}) func(string) {
    // TODO: answer here
  }
  ```

- Fungsi **FindStudent** menerima parameter `map[string][]interface{}` dan `int` yang mewakili nilai mahasiswa minimum yang ingin dicari. Fungsi ini akan mencari semua mahasiswa yang memiliki nilai di atas nilai minimum dan mengembalikan sebuah `map[string][]interface{}` yang berisi data mahasiswa yang ditemukan.

  ```go
  func FindStudent(students map[string][]interface{}, scoreThreshold int) map[string][]interface{} {
    // TODO: answer here
  }
  ```

Selain fungsi-fungsi tersebut, program ini juga menggunakan sebuah `map` untuk menyimpan data mahasiswa dan sebuah loop `for` untuk mengontrol aliran program. Loop ini akan meminta input dari pengguna dan memanggil fungsi yang sesuai untuk mengelola data mahasiswa. Jika input pengguna tidak valid, program akan menampilkan pesan kesalahan.

Untuk menyelesaikan live coding ini, siswa harus mengimplementasikan fungsi-fungsi yang diberikan di atas dan memastikan program dapat berjalan dengan baik untuk mengelola data mahasiswa seperti yang dijelaskan.

### Test Case Examples

**Input/Output**:

```bash
> Enter command (add, remove, find, view): add
> Enter name: John
> Enter address: Sudirman
> Enter phone: 081234567890
> Enter score: 90
> Enter command (add, remove, find, view): view
> Student data:
> Name  Address     Phone           Score
> John  Sudirman    081234567890    90

> Enter command (add, remove, find, view): find
> Enter score threshold: 80
> Search result:
> Name  Address     Phone           Score
> John  Sudirman    081234567890    90

> Enter command (add, remove, find, view): find
> Enter score threshold: 95
> Search result:
> Name  Address Phone   Score

> Enter command (add, remove, find, view): remove
> Enter name: John
> Enter command (add, remove, find, view): view
> Student data:
> Name  Address Phone   Score

> Enter command (add, remove, find, view): invalid
> Invalid command. Available commands: add, remove, find, view
```

**Explanation**:

> Test case ini menunjukkan fungsionalitas dasar dari program. Pengguna menambahkan siswa baru bernama John dengan alamat, nomor telepon, dan score. Kita kemudian melihat data siswa, menghapus John dari sistem, dan melihat data siswa yang diperbarui. Terakhir, kita mencari siswa dengan score di atas **80** (termasuk John) dan kemudian mencari siswa dengan skor di atas **95** (tidak termasuk John). Pengguna kemudian memasukkan perintah yang tidak valid dan menerima pesan error.
