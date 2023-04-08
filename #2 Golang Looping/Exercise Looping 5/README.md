# Exercise Looping

## Latihan Ke-5 - Reverse Word in String

### Description

Buatlah sebuah _code_ yang dapat membalik sekumpulan kata yang ada di dalam _string_ yang diberikan. Jika _string_ yang diberikan tidak memiliki kata, maka outputnya adalah _string_ kosong. Terdapat fungsi `ReverseWord` dengan parameter `str`bertipe `string` yang akan digunakan sebagai _input_ data. Terakhir kembalikan _output_ berupa `string` yang berisi kata-kata yang sudah dibalik.

Contoh sebagai berikut, jika diberikan _string_ `"Aku Sayang Ibu"` maka _output_ yang diharapkan adalah `"Uka Gnayans Ubi"`. Atau `"A bird fly to the Sky"`, maka _otput_ yang diharapkan adalah `"A drib ylf ot eht Yks"`.

Perlu diperhatikan, jika huruf awal dari kata tersebut adalah huruf kapital, maka ketika kata tersebut dibalik harus tetap berawalan huruf kapital.

> hint : dapat menggunakan `strings.ToUpper()` untuk mengubah dari huruf kecil ke huruf kapital. Dan `strings.ToLower()` untuk mengubah dari huruf kapital ke huruf kecil. Dan gunakan `unicode.IsUpper()` untuk mengecek apakah merupakan huruf besar.

### Constraints

- Input `str` akan selalu berupa _string_ yang terdiri dari karakter huruf dan spasi (**tidak ada** karakter simbol seperti `!@#$%^&*()-+=.,`), dengan panjang karakter 1 <= `str` <= 100.000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
str = "Aku Sayang Ibu"
```

**Expected Output / Behavior**:

```txt
"Uka Gnayas Ubi"
```

**Explanation**:

Kata pertama yang perlu dibalik adalah `"Aku"` menjadi `"ukA"`. Namun, jika tedapat huruf kapital di awal kata, maka kata tersebut harus tetap berawalan huruf kapital. Jadi, kata `"Aku"` menjadi `"Uka"` (balikan kata `ukA` akan mengubah `u` menjadi huruf besar dan `a` menjadi huruf kecil ). Kemudian, kata kedua yang perlu dibalik adalah `"Sayang"` menjadi `"Gnayas"`. Dan kata terakhir yang perlu dibalik adalah `"Ibu"` menjadi `"Ubi"`. Hasilnya adalah `"Uka Gnayas Ubi"`

#### Test Case 2

**Input**:

```txt
str = "ini terlalu mudah"
```

**Expected Output / Behavior**:

```txt
"ini ulalret hadum"
```

**Explanation**:

Kata pertama yang perlu dibalik adalah `"ini"` menjadi `"ini"` (karena semua huruf kecil, maka tidak perlu mengubah huruf awal menjadi huruf besar). Kemudian, kata kedua yang perlu dibalik adalah `"terlalu"` menjadi `"ulalret"`. Dan kata terakhir yang perlu dibalik adalah `"mudah"` menjadi `"hadum"`. Hasilnya adalah `"ini ulalret hadum"`

#### Test Case 3

**Input**:

```txt
str = "KITA SELALU BERSAMA"
```

**Expected Output / Behavior**:

```txt
"ATIK ULALES AMASREB"
```

**Explanation**:

Kata pertama yang perlu dibalik adalah `"KITA"` menjadi `"ATIK"` (karena semua huruf besar, maka tidak ada yang perlu diubah awalan atau akhiran kata menjadi huruf kecil). Kemudian, kata kedua yang perlu dibalik adalah `"SELALU"` menjadi `"ULALES"`. Dan kata terakhir yang perlu dibalik adalah `"BERSAMA"` menjadi `"AMASREB"`. Hasilnya adalah `"ATIK ULALES AMASREB"`
