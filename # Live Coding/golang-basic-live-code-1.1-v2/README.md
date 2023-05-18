# Dasar Pemrograman Backend

## Live Coding - Geometry Calculator

### Implementation technique

Siswa akan melaksanakan sesi live code di 15 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 15 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

Proyek ini adalah sebuah _command-line interface (CLI)_ program yang ditulis dalam bahasa Go yang memungkinkan kita untuk melakukan kalkulasi geometri sederhana. Program ini dapat menghitung luas dan keliling dari persegi dan persegi panjang. User akan diminta memasukkan bentuk geometri yang akan dihitung, kemudian program akan memproses input tersebut menggunakan switch statement dan memanggil fungsi-fungsi `SelectForm`, `CalculateSquare`, dan `CalculateRectangle` untuk menghitung hasilnya. Program juga memberikan opsi untuk menghitung lagi atau keluar dari aplikasi setelah perhitungan selesai.

### Constraints

Program ini dibagi menjadi 4 bagian:

1. **Main** (mengontrol keseluruhan aliran program dan memanggil fungsi lainnya)
2. **SelectForm** (menentukan jenis bentuk berdasarkan input pengguna)
3. **CalculateSquare** (menghitung luas dan keliling persegi)
4. **CalculateRectangle** (menghitung luas dan keliling persegi panjang)

Program ini menggunakan statement _`loop`_ dan _`switch`_ untuk memproses pilihan bentuk dan memanggil fungsi yang sesuai untuk menghitung luas dan keliling. Jika pengguna memasukkan formulir yang tidak valid, program akan menampilkan pesan kesalahan.

Berikut adalah penjelasan dari fungsi-fungsi yang harus di lengkapi:

- Fungsi **SelectForm** menerima parameter `string` yang mewakili bentuk dan mengembalikan nama bentuk atau pesan "Bentuk geometri tidak valid!" jika input tidak valid dalam tipe `string`.

    ```go
    func SelectForm(bentuk string) string {
    // TODO: answer here
    }
    ```

- Fungsi **CalculateSquare** menerima parameter `float64` yang mewakili panjang sisi persegi dan mengembalikan 3 nilai yaitu 2 nilai `float64` yang mewakili luas dan keliling persegi juga pesan "sisi harus lebih besar dari 0" jika input tidak valid dalam tipe `string`.

    ```go
    func CalculateSquare(sisi float64) (float64, float64, string) {
    // TODO: answer here
    }
    ```

- Fungsi `CalculateRectangle` menerima dua parameter `float64` yang mewakili panjang dan lebar persegi panjang dan mengembalikan 3 nilai yaitu 2 nilai `float64` yang mewakili luas dan keliling persegi panjang juga pesan "panjang dan lebar harus lebih besar dari 0" jika input tidak valid dalam tipe `string`.

    ```go
    func CalculateRectangle(panjang, lebar float64) (float64, float64, string) {
    // TODO: answer here
    }
    ```

### Test Case Examples

#### Test Case 1

**Input/Output**:

```bash
> Masukkan bentuk geometri (persegi/persegi-panjang): segitiga
> Bentuk geometri tidak valid!
> Apakah Anda ingin menghitung lagi? (y/n): n
```

**Explanation**:

> Pada test case ini, input yang diberikan adalah "segitiga" sebagai bentuk geometri. Namun bentuk geometri yang tersedia di program adalah "persegi" dan "persegi-panjang". Maka, program akan menampilkan pesan "Bentuk geometri tidak valid!" dan menanyakan apakah pengguna ingin menghitung lagi. Karena pengguna memilih "n", maka program akan berhenti.

#### Test Case 2

**InputOutput**:

```bash
> Masukkan bentuk geometri (persegi/persegi-panjang): persegi
> Masukkan sisi: 0
> sisi harus lebih besar dari 0
> Luas persegi: 0.00
> Keliling persegi: 0.00
> Apakah Anda ingin menghitung lagi? (y/n): y
> Masukkan bentuk geometri (persegi/persegi-panjang): persegi-panjang
> Masukkan panjang: 10
> Masukkan lebar: -5
> panjang dan lebar harus lebih besar dari 0
> Apakah Anda ingin menghitung lagi? (y/n): n
```

**Explanation**:

> Pada test case ini, input yang diberikan adalah "persegi" sebagai bentuk geometri. Kemudian, pengguna diminta untuk memasukkan nilai sisi. Namun, pengguna memasukkan nilai `0`. Karena sisi harus lebih besar dari `0`, maka program akan menampilkan hasil luas dan keliling dengan nilai `0` dan pesan "sisi harus lebih besar dari `0`". Kemudian, program akan menanyakan apakah pengguna ingin menghitung lagi. Karena pengguna memilih "y", maka program akan meminta pengguna untuk memasukkan bentuk geometri yang ingin dihitung, yaitu "persegi-panjang". Kemudian program meminta pengguna untuk memasukkan nilai panjang dan lebar. Namun, pengguna memasukkan nilai lebar dengan nilai negatif (`-5`). Karena nilai panjang dan lebar harus lebih besar dari `0`, maka program akan menampilkan pesan "panjang dan lebar harus lebih besar dari `0`". Karena pengguna memilih "n", maka program akan berhenti.

#### Test Case 3

**InputOutput**:

```bash
> Masukkan bentuk geometri (persegi/persegi-panjang): persegi
> Masukkan sisi: 5
> Luas persegi: 25.00
> Keliling persegi: 20.00
> Apakah Anda ingin menghitung lagi? (y/n): y
> Masukkan bentuk geometri (persegi/persegi-panjang): persegi-panjang
> Masukkan panjang: 6
> Masukkan lebar: 3
> Luas persegi panjang: 18.00
> Keliling persegi panjang: 18.00
> Apakah Anda ingin menghitung lagi? (y/n): n

```

**Explanation**:

> Pada test case ini, input yang diberikan adalah "persegi" sebagai bentuk geometri. Kemudian, pengguna diminta untuk memasukkan sisi persegi dengan nilai `5`. Setelah itu, program menghitung luas dan keliling persegi dengan rumus yang sesuai, yaitu `(luas = sisi x sisi)` dan `(keliling = 4 x sisi)`. Output yang dihasilkan adalah luas persegi sebesar `25.00` dan keliling persegi sebesar `20.00`.
>
> Kemudian, program kembali menanyakan apakah pengguna ingin menghitung lagi atau tidak. Karena pengguna memilih untuk menghitung lagi (dengan memasukkan "y"), program menanyakan bentuk geometri yang akan dihitung, dan pada kali ini pengguna memilih "persegi-panjang". Selanjutnya, program meminta pengguna untuk memasukkan nilai panjang dan lebar persegi-panjang dengan masing-masing nilai `6` dan `3`. Setelah itu, program menghitung luas dan keliling persegi-panjang dengan rumus yang sesuai, yaitu `(luas = panjang x lebar)` dan `(keliling = 2 x (panjang + lebar))`. Output yang dihasilkan adalah luas persegi-panjang sebesar `18.00` dan keliling persegi-panjang sebesar `18.00`.
>
> Terakhir, program kembali menanyakan apakah pengguna ingin menghitung lagi atau tidak. Karena pengguna memilih untuk tidak menghitung lagi (dengan memasukkan "n"), program berakhir.
