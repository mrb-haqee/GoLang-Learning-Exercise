# Dasar Pemrograman Backend

## Live Coding - LearnlyApp

### Implementation technique

Siswa akan melaksanakan sesi live code di 15 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 15 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

Kita diminta untuk menyelesaikan tugas pada sistem **Learnly**, yaitu sebuah aplikasi belajar daring. Tugas ini meminta kita untuk menyelesaikan proses login pengguna dan fitur untuk mendapatkan data pelajaran berdasarkan tingkat kesulitannya di Learnly.

Kita diminta untuk menambahkan beberapa fungsi pada code yang sudah diberikan, diantaranya:

- Fungsi `LoginUser`: untuk mengizinkan pengguna untuk melakukan login ke sistem Learnly.

  ```go
  func (l *learnlyApp) LoginUser(email, password string) (model.User, error)
  ```

  - Parameter:
    - `email string`: alamat email pengguna yang ingin masuk.
    - `password string`: kata sandi pengguna yang ingin masuk.
  - Nilai kembalian:
    - `model.User`: objek `User` yang telah masuk ke sistem.
    - `error`: pesan kesalahan jika terjadi error dalam proses masuk.
  - Kemungkinan pesan kesalahan:
    - `"invalid email or password"`: email atau kata sandi salah.

- Fungsi `GetLessonsByDifficulty`: untuk mengambil daftar pelajaran sesuai dengan tingkat kesulitan tertentu dan pengguna yang sedang login.

  ```go
  func (l *learnlyApp) GetLessonsByDifficulty(email string, difficulty int) ([]model.Lesson, error)
  ```

  - Parameter:
    - `email string`: alamat email pengguna yang ingin mengakses pelajaran.
    - `difficulty int`: tingkat kesulitan pelajaran yang diinginkan.
  - Nilai kembalian:
    - `[]model.Lesson`: array yang berisi semua pelajaran dengan tingkat kesulitan yang sesuai.
    - `error`: pesan kesalahan jika terjadi error dalam proses pengambilan pelajaran.
  - Kemungkinan pesan kesalahan:
    - `"you must login first"`: pengguna belum masuk ke sistem.

Anda dapat menambahkan fungsi atau code lainnya jika diperlukan.

### Constraints

Beberapa Constraints yang harus dihadapi yaitu:

- Semua struct `Users` harus memiliki `Name`, `Email`, dan `Password` yang tidak kosong.
- `Age` dari struct `Users` harus antara `0` dan `120` (inklusif).
- `Gender` dari struct `Users` harus berupa "Male" atau "Female".
- `Email` dari struct `Users` harus unik (yaitu, tidak ada dua pengguna yang dapat memiliki alamat email yang sama).
- Saat pengguna register, bidang Session mereka disetel ke `false`.
- Saat pengguna login, `Session` mereka disetel ke `true`.
- Hanya pengguna yang login yang dapat mengambil pelajaran.

### Test Case Examples

#### Test Case 1: Valid Login

**Input**:

```go
//register users
users := []model.User{
    {Name: "John", Email: "john@example.com", Password: "password", Age: 25, Gender: "Male"},
    {Name: "Jane", Email: "jane@example.com", Password: "password", Age: 30, Gender: "Female"},
}

//add lessons
lessons := []model.Lesson{
  {
    Title: "Introduction to Go", 
    Description: "Learn the basics of Go programming language", 
    Category: "Programming", 
    Difficulty: 1
  },
  {
    Title: "Intermediate Go", 
    Description: "Take your Go skills to the next level with advanced topics", 
    Category: "Programming", 
    Difficulty: 2
  },
}

//automatic
app := NewLearnly(users, lessons)

//automatic when we choose "2"
user, err := app.LoginUser("jane@example.com", "password")
```

**Expected Output / Behavior**:

Golang

```go
err == nil
user.Email == "jane@example.com"
user.Session == true
```

Terminal:

```bash
> 2. Login
          - Email: jane@example.com
          - Password: password
> Logged in as Jane # Logged in as {user.Name}
```

**Explanation**:

> Testcase ini memeriksa _user login_ yang berhasil dengan kredensial yang valid. Fungsi `LoginUser` harus mengembalikan `nil` dan seharusnya tidak ada error. Objek `User` yang dikembalikan harus menyetel field `Session` ke `true`.

#### Test Case 2: Invalid Login

**Input**:

```go
//register users
users := []model.User{
    {Name: "John", Email: "john@example.com", Password: "password", Age: 25, Gender: "Male"},
    {Name: "Jane", Email: "jane@example.com", Password: "password", Age: 30, Gender: "Female"},
}

//add lessons
lessons := []model.Lesson{
  {
    Title: "Introduction to Go", 
    Description: "Learn the basics of Go programming language", 
    Category: "Programming", 
    Difficulty: 1
  },
  {
    Title: "Intermediate Go", 
    Description: "Take your Go skills to the next level with advanced topics", 
    Category: "Programming", 
    Difficulty: 2
  },
}

//automatic
l := NewLearnly(users, lessons)

//automatic when we choose "2"
_, err := l.LoginUser("invalid_email@example.com", "invalid_password")
```

**Expected Output / Behavior**:

```go
err == invalid email or password
```

Terminal:

```bash
> 2. Login
          - Email: invalid_email@example.com
          - Password: invalid_password
> invalid email or password
```

**Explanation**:

> Testcase ini memeriksa _user login_ yang berhasil dengan kredensial yang tidak valid. Karena email dan password tidak ditemukan di slice `User`, maka fungsi LoginUser mengembalikan error yang menunjukkan bahwa login tidak valid.

#### Test Case 3: Get Lessons By Difficulty

**Input**:

```go
//register users
users := []model.User{
    {Name: "John", Email: "john@example.com", Password: "password", Age: 25, Gender: "Male"},
    {Name: "Jane", Email: "jane@example.com", Password: "password", Age: 30, Gender: "Female"},
}

//add lessons
lessons := []model.Lesson{
  {
    Title: "Introduction to Go", 
    Description: "Learn the basics of Go programming language", 
    Category: "Programming", 
    Difficulty: 1
  },
  {
    Title: "Intermediate Go", 
    Description: "Take your Go skills to the next level with intermediate topics", 
    Category: "Programming", 
    Difficulty: 2
  },
  {
    Title: "Advanced Go", 
    Description: "Take your Go skills to the next level with advanced topics", 
    Category: "Programming", 
    Difficulty: 3
  },
}

//automatic
l := NewLearnly(users, lessons)

//automatic when we choose "4"
result, err := l.GetLessonsByDifficulty("john@example.com", 2)
```

**Expected Output / Behavior**:

Golang:

```go
[{Intermediate Go Take your Go skills to the next level with intermediate topics Programming 2}], nil
```

Terminal:

```bash
> 4. Get lesson by difficulty
        - Email: john@example.com
        - Difficulty: 2
> Lesson:  [{Intermediate Go Take your Go skills to the next level with intermediate topics Programming 2}]
```

**Explanation**:

> Fungsi `GetLessonsByDifficulty` dipanggil dengan email yang valid dan difficulty 2. Fungsi ini memeriksa apakah user sudah melakukan login. Karena user dengan email "john@example.com" ditemukan dan memiliki session true, maka fungsi ini mengembalikan lesson dengan difficulty 2, yaitu "Intermediate Go".
