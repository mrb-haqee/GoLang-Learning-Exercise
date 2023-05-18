# Dasar Pemrograman Backend

## Student Portal 3

### Description

Program ini adalah sebuah student portal sederhana yang digunakan untuk melakukan registrasi mahasiswa baru, login, mengambil program studi dari suatu mahasiswa, dan mengubah program studi dari suatu mahasiswa. Implementasi mencakup struct `InMemoryStudentManager`, yang memiliki daftar internal mahasiswa dan program studinya, dan interface `StudentManager` yang mendefinisikan fungsionalitas student portal.

### Constraints

Ada beberapa constraints yang harus Anda pertimbangkan saat menerapkan student portal dalam memori:

- `ID` mahasiswa harus unik, tidak boleh ada dua mahasiswa dengan `ID` yang sama.
- Saat melakukan login, program hanya memeriksa apakah `ID` dan `name` yang dimasukkan sudah ada di slice `students`. Program tidak memeriksa apakah `ID` dan `name` tersebut merupakan pasangan yang valid.
- Kode program studi (`studyProgram`) harus valid, artinya harus terdaftar di map `studentStudyPrograms`.
- Saat melakukan registrasi, program memeriksa apakah `ID` yang dimasukkan sudah ada di slice `students`. Jika sudah ada, registrasi akan gagal.
- Saat melakukan modifikasi data mahasiswa, program hanya memeriksa apakah nama mahasiswa sudah terdaftar di slice `students`. Jika nama mahasiswa ditemukan, program akan memanggil fungsi dengan type `model.StudentModifier` yang diberikan sebagai argumen dan mengirimkan data mahasiswa yang bersesuaian.

📝 Di sini, kamu diminta untuk melengkapi fungsi di file `main.go` dengan ketentuan sebagai berikut:

- Fungsi `GetStudents` tidak menerima input dan akan mengembalikan seluruh data mahasiswa dalam bentuk slice `[]Student`.
  - Jika tidak ada mahasiswa yang terdaftar, fungsi akan mengembalikan slice kosong.

- Fungsi `Login` menerima input `id` dan `name` ber-_type_ `string` yang merupakan id dan nama mahasiswa. Fungsi ini akan memeriksa apakah data mahasiswa yang terdapat dalam slice `students`.
  - Jika input `id` atau `name` kosong, fungsi akan mengembalikan pesan `""` dan error `"ID or Name is undefined!"`.
  - Jika ditemukan, fungsi akan mengembalikan pesan `"Login berhasil: <name>"` beserta error `nil`.
  - Jika tidak ditemukan, fungsi akan mengembalikan pesan `""` dan error `"Login gagal: data mahasiswa tidak ditemukan"`.

  Contoh penggunaan:
  
    ```go
    sm := NewInMemoryStudentManager()
    message, err := sm.Login("C1234", "Cindy") // Output: "Login berhasil: Cindy", nil
    ```

- Fungsi `Register` menerima input `id`, `name`, dan `studyProgram` ber-_type_ `string` yang merupakan id, nama, dan program studi mahasiswa yang akan diregistrasikan. Fungsi ini akan menambahkan data mahasiswa baru ke dalam slice `Students`.
  - Jika input `id`, `name`, atau `studyProgram` kosong, program mengembalikan pesan `""` dan error `"ID, Name or StudyProgram is undefined!"`
  - Jika kode program studi yang dimasukkan tidak terdapat dalam map `studentStudyPrograms`, program mengembalikan pesan `""` dan error `"Study program <studyProgram> is not found"`.
  - Jika id mahasiswa yang dimasukkan sudah terdapat di dalam slice `Students`, program mengembalikan pesan `""` dan error `"Registrasi gagal: id sudah digunakan"`.
  - Jika berhasil, fungsi akan mengembalikan pesan `"Registrasi berhasil: name (studyProgram)"` dan error `nil`.

  Contoh penggunaan:

    ```go
    sm := NewInMemoryStudentManager()
    message, err := sm.Register("C1234", "Cindy", "SI") // Output: "Registrasi berhasil: Cindy (Sistem Informasi)", nil
    ```

- Fungsi `GetStudyProgram` menerima input code yang merupakan kode program studi. Fungsi ini akan mencari nama program studi dari map `studentStudyPrograms` berdasarkan kode program studi yang dimasukkan.
  - Jika input code kosong, program mengembalikan pesan `""` dan error `"Code is undefined!"`
  - Jika kode program studi yang dimasukkan tidak ditemukan, program mengembalikan pesan `""` dan error `"Kode program studi tidak ditemukan"`.
  - Jika berhasil, fungsi akan mengembalikan nama program studi yang dicari dan error `nil`.

  Contoh penggunaan:

  ```go
  sm := NewInMemoryStudentManager()
  message, err := sm.GetStudyProgram("TK") // Output: "Teknik Komputer", nil
  ```

- Fungsi `ModifyStudent` menerima input `name` yang merupakan nama mahasiswa dan `fn` yang merupakan sebuah fungsi `studentModifier`. Fungsi ini akan mencari data mahasiswa berdasarkan nama yang dimasukkan, dan akan memodifikasinya menggunakan fungsi `fn`.
  - Jika mahasiswa dengan nama yang dimasukkan tidak ditemukan, program mengembalikan pesan `""` dan error `"Mahasiswa tidak ditemukan."`.
  - Jika terdapat error pada fungsi `fn`, program akan mengembalikan pesan `""` dan error tersebut.
  - Jika berhasil, fungsi akan mengembalikan pesan `"Program studi mahasiswa berhasil diubah."` dan error `nil`.

  Contoh penggunaan:

  ```go
  sm := NewInMemoryStudentManager()
  message, err := sm.ModifyStudent("Dito", sm.ChangeStudyProgram("SI")) // Output: "Program studi mahasiswa berhasil diubah.", error
  ```

- Fungsi `ChangeStudyProgram` menerima input `programStudi` yang merupakan kode program studi yang akan diganti. Fungsi ini akan mengembalikan sebuah fungsi dengan type `model.tudentModifier` yang akan memodifikasi program studi dari mahasiswa yang diberikan.
  - Jika kode program studi yang dimasukkan tidak ditemukan, program akan mengembalikan error `"Kode program studi tidak ditemukan"`.
  - Jika berhasil, fungsi dengan type `model.StudentModifier` yang dihasilkan akan mengubah program studi dari mahasiswa yang diberikan.

  Contoh penggunaan:

  ```go
  sm := NewInMemoryStudentManager()
  modifier := sm.ChangeStudyProgram("SI")
  sm.ModifyStudent("Dito", modifier) // "Program studi mahasiswa berhasil
  ```

### Test Case Examples

#### Test Case 1

**Input**:

```go
sm := NewInMemoryStudentManager()
sp, err := sm.GetStudyProgram("TI")
```

**Expected Output / Behavior**:

```bash
"Teknik Informatika", nil
```

**Explanation**:

Method `GetStudyProgram` harus mengembalikan `"Teknik Informatika"` untuk kode masukan `"TI"`.

#### Test Case 2

**Input**:

```go
sm := NewInMemoryStudentManager()
res, err := sm.Register("C12345", "Charlie", "SI")
```

**Expected Output / Behavior**:

```bash
"Registrasi berhasil: Charlie (SI)", nil
```

**Explanation**:

Method `Register` harus mendaftarkan mahasiswa baru dengan ID `C12345`, name `Charlie`, dan studyProgram `SI`. Sehingga output yang diharapkan adalah `"Registrasi Berhasil: Charlie (SI)", nil`.

#### Test Case 3

**Input**:

```go
sm := NewInMemoryStudentManager()
res, err := sm.ModifyStudent("Aditira", sm.ChangeStudyProgram("SI"))
```

**Expected Output / Behavior**:

```bash
"Program studi mahasiswa berhasil diubah.", nil
```

**Explanation**:

Metode `ModifyStudent` harus mengubah program studi siswa dengan nama "Aditira" menjadi "SI". Oleh karena itu output yang diharapkan adalah `"Program studi mahasiswa berhasil diubah.", nil`.
