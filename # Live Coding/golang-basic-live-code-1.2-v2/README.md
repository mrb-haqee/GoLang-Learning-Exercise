# Dasar Pemrograman Backend

## Live Coding - Kalkulator Konversi Satuan Panjang

### Implementation technique

Siswa akan melaksanakan sesi live code di 15 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 15 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

Proyek ini adalah sebuah _command-line interface (CLI)_ program yang ditulis dalam bahasa Go yang memungkinkan kita untuk mengkonversi satuan panjang. Program ini dapat mengkonversi meter (m), sentimeter (cm), kaki (ft), dan inci (in) satu sama lain. User akan diminta memasukkan nilai panjang yang akan dikonversi, satuan awal, dan satuan tujuan, kemudian program akan memproses input tersebut menggunakan fungsi ConvertLength dan menampilkan hasilnya dalam bentuk: `nilai awal` `satuan awal` = `nilai hasil` `satuan hasil`. Program juga memberikan opsi untuk mengkonversi lagi atau keluar dari aplikasi setelah konversi selesai.

### Constraints

Program ini dibagi menjadi 2 bagian:

- **Main** (mengontrol keseluruhan aliran program dan memanggil fungsi ConvertLength)
- **ConvertLength** (mengkonversi nilai panjang dari satuan awal ke satuan tujuan)

Program ini menggunakan statement `loop` untuk memproses konversi dan statement `switch` untuk memilih operasi konversi yang tepat berdasarkan satuan awal dan satuan tujuan. Jika pengguna memasukkan satuan yang tidak valid, program akan menampilkan pesan kesalahan.

Berikut adalah penjelasan dari fungsi yang harus dilengkapi:

- Fungsi `ConvertLength` menerima 3 parameter yaitu `float64` yang mewakili nilai panjang yang akan dikonversi, dan 2 parameter string yang mewakili satuan awal dan satuan tujuan, dan mengembalikan `float64` yang mewakili nilai hasil konversi.

  ```go
  func ConvertLength(panjang float64, dari, ke string) float64 {
    // TODO: answer here
  }
  ```

  Aturan dan Formula:

  - Nilai `panjang` kurang dari atau sama dengan 0, maka kembalikan 0
  - Nilai `dari` dan `ke` tidak valid, maka kembalikan nilai `panjang`
  - Nilai `dari` dan `ke` sama, maka kembalikan nilai `panjang`
  - m (meter):
    - ke satuan cm dengan cara mengalikan nilai `panjang` dengan `100`
    - ke satuan ft dengan cara mengalikan nilai `panjang` dengan `3.281`
    - ke satuan in dengan cara mengalikan nilai `panjang` dengan `39.37`
  - cm (centimeter):
    - ke satuan m dengan cara membagi nilai `panjang` dengan `100`
    - ke satuan ft dengan cara membagi nilai `panjang` dengan `30.48`
    - ke satuan in dengan cara membagi nilai `panjang` dengan `2.54`
  - ft (feet)
    - ke satuan m dengan cara membagi nilai `panjang` dengan `3.281`
    - ke satuan cm dengan cara mengalikan nilai `panjang` dengan `30.48`
    - ke satuan in dengan cara mengalikan nilai `panjang` dengan `12`
  - in (inch)
    - ke satuan m dengan cara membagi nilai `panjang` dengan `39.37`
    - ke satuan cm dengan cara mengalikan nilai `panjang` dengan `2.54`
    - ke satuan ft dengan cara membagi nilai `panjang` dengan `12`

### Test Case Examples

#### Test Case 1

**Input/Output:**:

```bash
> === Kalkulator Konversi Satuan Panjang ===
> Masukkan panjang: 10
> Masukkan satuan dari: m
> Masukkan satuan ke: cm
> 10.00 m = 1000.00 cm
> Apakah Anda ingin mengkonversi kembali? (y/n): n
```

**Explanation**:

> Pada test case ini, pengguna memasukkan panjang 10 dalam satuan meter (m) dan ingin dikonversi ke satuan centimeter (cm). Hasil konversi dari 10.00 m ke cm adalah 1000.00 cm. Program akan menampilkan hasil konversi dalam format "%.2f %s = %.2f %s\n". Kemudian, program akan menanyakan apakah pengguna ingin mengkonversi kembali. Karena pengguna memilih "n", maka program akan berhenti.

#### Test Case 2

**Input/Output**:

```bash
> === Kalkulator Konversi Satuan Panjang ===
> Masukkan panjang: 15
> Masukkan satuan dari: in
> Masukkan satuan ke: m
> 15.00 in = 0.38 m
> Apakah Anda ingin mengkonversi kembali? (y/n): y
> === Kalkulator Konversi Satuan Panjang ===
> Masukkan panjang: 50
> Masukkan satuan dari: ft
> Masukkan satuan ke: cm
> 50.00 ft = 1524.00 cm
> Apakah Anda ingin mengkonversi kembali? (y/n): n
```

**Explanation**:

> Pada test case ini, pengguna memasukkan dua konversi, yaitu dari satuan inchi (in) ke meter (m) dan dari satuan kaki (ft) ke centimeter (cm). Untuk konversi pertama, hasil konversi dari 15.00 in ke m adalah 0.38 m. Kemudian program akan menanyakan apakah pengguna ingin mengkonversi kembali. Karena pengguna memilih "y", maka program akan meminta input konversi yang baru. Untuk konversi kedua, hasil konversi dari 50.00 ft ke cm adalah 1524.00 cm. Program akan menampilkan hasil konversi dalam format "%.2f %s = %.2f %s\n". Kemudian, program akan menanyakan apakah pengguna ingin mengkonversi kembali. Karena pengguna memilih "n", maka program akan berhenti.

#### Test Case 3

**Input/Output**:

```bash
> === Kalkulator Konversi Satuan Panjang ===
> Masukkan panjang: 7.5
> Masukkan satuan dari: m
> Masukkan satuan ke: ft
> 7.50 m = 24.61 ft
> Apakah Anda ingin mengkonversi kembali? (y/n): n
```

**Explanation**:

> Pada test case ini, pengguna memasukkan panjang 7.5 dalam satuan meter (m) dan ingin dikonversi ke satuan kaki (ft). Hasil konversi dari 7.50 m ke ft adalah 24.61 ft. Program akan menampilkan hasil konversi dalam format "%.2f %s = %.2f %s\n". Kemudian, program akan menanyakan apakah pengguna ingin mengkonversi kembali. Karena pengguna memilih "n", maka program akan berhenti.
