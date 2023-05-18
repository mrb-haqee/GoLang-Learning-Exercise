# Final Project Backend Lanjutan

## Student Portal Final

### Description

Kita akan menyempurnakan project kita sebelumnya yaitu portal mahasiswa dengan menggunakan http server mux untuk membuat aplikasi portal mahasiswa. Aplikasi ini akan digunakan untuk mengelola data mahasiswa, dan mengelola data mata kuliah yang diambil oleh mahasiswa.

Data user pada program ini disimpan dalam variabel Students dengan format string berisi informasi ID, nama, jurusan dan role, disimpan ke dalam format string yang digabung dengan _underscore_ (`_`). Formatnya adalah `[id]_[name]_[kode jurusan]_[role]`

Terdapat 2 role yang disediakan yaitu `admin` dan `user`. Admin dapat mengakses semua fitur yang ada, sedangkan user hanya dapat mengakses fitur yang diperbolehkan saja.

Data tersebut disimpan ke dalam file `txt` yang ada di `data/users.txt`

Contoh data user:

```txt
A1111_Pratama_SI_admin
B1112_Jaya_TI_user
```

Program juga menyimpan list jurusan yang ada di dalam file `txt` pada direktori `data/list-study.txt`. Terdapat 15 jurusan yang disedikan dengan format `[kode jurusan]_[nama jurusan]`. Contoh data jurusan:

```txt
TI_Teknik Informatika
TK_Teknik Komputer
SI_Sistem Informasi
...
```

Aplikasi ini akan memiliki fitur sederhana seperti:

- Register
- Login
- Logout

- Get List Study
- Add Student
- Delete Student

- Get Weather

Beberapa fitur memiliki akses yang berbeda-beda berdasarkan role user yang login. Berikut adalah tabel akses yang diberikan:

| Fitur          | Admin | User |
| -------------- | ----- | ---- |
| Register       | ✅    | ✅   |
| Login          | ✅    | ✅   |
| Logout         | ✅    | ✅   |
| Get List Study | ✅    | ✅   |
| Add Student    | ✅    | ❌   |
| Delete Student | ✅    | ❌   |
| Get Weather    | ✅    | ✅   |

### Instruction

**Handler**

Kalian diminta untuk melengkapi handler yang sudah disediakan di `/handler/handler.go`. Berikut adalah hal-hal yang harus diperhatikan dalam mengerjakan aplikasi student portal:

> NOTE: tugas handler di bawah ini adalah pengembangan dari handler di assignment sebelumnya, jadi kalian bisa menggunakan handler yang sudah dikerjakan di assignment sebelumnya jika diperlukan.

- **Register** dengan method **POST** dan end point `/register`
  - Hanya boleh menggunakan method **POST** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Harus memberikan request body dengan format JSON `{"id": "<id user>", "name": <name>, "role" : <admin / user>, "study_code" : <code>}`. Contoh request body:

    ```json
    {
        "id": "A1234",
        "name": "Pratama",
        "role": "admin",
        "study_code": "TI"
    }
    ```

    - Jika request body ID, name, role atau study Code kosong(`""`), maka:
      - Berikan response code **400**
      - Berikan response message `{"error":"ID, name, study code or role is empty"}`
    - Pastikan bahwa role yang diberikan hanya `admin` atau `user`, jika tidak, maka:
      - Berikan response code **400**
      - Berikan response message `{"error":"role must be admin or user"}`
    - Pastikan juga study code sesuai dengan yang ada di file `data/list-study.txt`, jika tidak, maka:
      - Berikan response code **400**
      - Berikan response message `{"error":"study code not found"}`
    - Jika semua ID user sudah ada di penyimpanan `data/users.txt`, maka:
      - Berikan response **400**
      - Berikan response message `{"error":"user id already exist"}`
    - Jika semua Oke, maka:
      - Berikan response **200**
      - Berikan response message `{"username":"<id user>","message":"register success"}`

  - Simpan data user terbaru ke dalam file `data/users.txt` dengan format yang sudah dijelaskan sebelumnya.

- **Login** dengan method **POST** dan end point `/login`
  - Hanya boleh menggunakan method **POST** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Harus memberikan request body dengan format JSON `{"id": "<id user>", "name": <name>}`. Contoh request body:

    ```json
    {
        "id": "A1234",
        "name": "Pratama"
    }
    ```

  - Jika request body ID atau name kosong(`""`), maka:
    - Berikan response code **400**
    - Berikan response message `{"error":"ID or name is empty"}`
  - Jika ID user dan name tidak ditemukan di file `data/users.txt`, maka:
    - Berikan response code **400**
    - Berikan response message `{"error":"user id not found"}`
  - Jika semua oke maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<id user>","message":"login success"}`
    - Berikan cookie dengan key `user_login_id` dan value `<id user>`
    - Berikan cookie dengan key `user_login_role` dan value `<role user>`
    - Simpan data user yang login ke dalam variable map `UserLogin` dengan key `<id user>` dan value berupa data user.

  > hint: untuk membuat cookie, kalian bisa menggunakan function `http.SetCookie()` dengan parameter `w http.ResponseWriter` dan `cookie *http.Cookie` yang berisi key dan value cookie. Contoh: `http.SetCookie(w, &http.Cookie{Name: "user_login_id", Value: "A1234"})`

- **Logout** dengan method **POST** dan end point `/logout`
  - Hanya boleh menggunakan method **POST** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`

  - Pastikan user sudah login dengan mengecek juga cookie yang dikirimkan oleh user, jika tidak ada cookie dengan key `user_login_id` maka:
    - Berikan response code **401**
    - Berikan response message error `{"error":"user login id not found"}`

  - Jika semua oke maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<id user>","message":"logout success"}`
    - Hapus cookie dengan key `user_login_id`
    - Hapus cookie dengan key `user_login_role`
    - Hapus data user yang login dari variable map `UserLogin`

- **GetStudyProgram** dengan method **GET** dan end point `/study-program`
  - Hanya boleh menggunakan method **GET** jika tidak, maka:
    - Berikan response code **405** (Method Not Allowed)
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Pastikan user sudah login dengan mengecek juga cookie yang dikirimkan oleh user, jika tidak ada cookie dengan key `user_login_id` maka:
    - Berikan response code **401**
    - Berikan response message error `{"error":"user login id not found"}`

  - Jika semua oke maka:
    - Berikan response code **200**
    - Berikan response message berupa list study program yang ada di file `data/list-study.txt` dengan format sebagai berikut:

    ```json
    [
        {
          "code": "TI",
          "name": "Teknik Informatika"
        },
        {
          "code": "SI",
          "name": "Sistem Informasi"
        },
        {
          ...
        }
    ]
    ```

- **AddUser** dengan method **POST** dan end point `/user/add`, fungsi ini hanya dapat diakses oleh user yang sudah login dengan role `"admin"`.
  - Hanya boleh menggunakan method **POST** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Pastikan user sudah login dengan mengecek juga cookie yang dikirimkan oleh user, jika tidak ada cookie dengan key `user_login_id` maka:
    - Berikan response code **401**
    - Berikan response message error `{"error":"user login id not found"}`
  - Pastikan user yang login memiliki role `"admin"` dengan mengecek cookie yang dikirim oleh user pada `user_login_role`, jika role bukan `"admin"` maka:
    - Berikan response code **401**
    - Berikan response message error `{"error":"user login role not Admin"}`

  - Harus memberikan request body dengan format JSON `{"id": "<id user>", "name": <name>, "study": <study>}`. Contoh request body:

    ```json
    {
        "id": "A1234",
        "name": "Pratama",
        "study": "TI"
    }
    ```

    Data yang dikirim tidak perlu mengirim role, karena role selalu di set menjadi `"user"`.
  - Jika request body ID, name, atau study kosong(`""`), maka:
    - Berikan response code **400**
    - Berikan response message `{"error":"ID, name, or study code is empty"}`
  - Jika ID user sudah ada di file `data/users.txt`, maka:
    - Berikan response code **400**
    - Berikan response message `{"error":"user id already exist"}`
  - Jika study code tidak ditemukan di file `data/list-study.txt`, maka:
    - Berikan response code **400**
    - Berikan response message `{"error":"study code not found"}`
  - Jika semua oke maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<id user>","message":"add user success"}`
    - Tambahkan data user ke dalam file `data/users.txt` dengan format yang sudah dijelaskan sebelumnya.

- **DeleteUser** dengan method **DELETE** dan end point `/user/add`, fungsi ini hanya dapat diakses oleh user yang sudah login dengan role `"admin"`.
  - Hanya boleh menggunakan method **DELETE** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Pastikan user sudah login dengan mengecek juga cookie yang dikirimkan oleh user, jika tidak ada cookie dengan key `user_login_id` maka:
    - Berikan response code **401**
    - Berikan response message error `{"error":"user login id not found"}`
  - Pastikan user yang login memiliki role `"admin"` dengan mengecek cookie yang dikirim oleh user pada `user_login_role`, jika role bukan `"admin"` maka:
    - Berikan response code **401**
    - Berikan response message error `{"error":"user login role not Admin"}`

  - Harus memberikan request dengan query parameter `id` yang berisi ID user yang akan dihapus. Contoh request:

    ```http
    DELETE /user/delete?id=A1234
    ```

  - Jika ID user tidak ada di file `data/users.txt`, maka:
    - Berikan response code **400**
    - Berikan response message `{"error":"user id not found"}`

  - Jika semua oke maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<id user>","message":"delete success"}`
    - Hapus data user dari file `data/users.txt` dengan format yang sudah dijelaskan sebelumnya.

- **GetWeather** dengan method **GET** dan end point `/get-weather`, fungsi ini akan menampilkan informasi cuaca berdasarkan 10 wilayah
  - Informasi cuaca dari 10 wilayah tersebut adalah:
    - Jakarta
    - Bandung
    - Surabaya
    - Yogyakarta
    - Medan
    - Makassar
    - Manado
    - Palembang
    - Semarang
    - Bali

  - Kalian diminta untuk mengimplementasikan goroutine untuk mengambil informasi cuaca dari 10 wilayah tersebut secara bersamaan. Untuk mendapatkan informasi cuaca, sudah disediakan fungsi `GetWeatherByRegion` di package `client` yang akan mengembalikan data cuaca berupa `Weather` struct. Contoh penggunaan fungsi tersebut adalah sebagai berikut:

    ```go
    weather, err := client.GetWeatherByRegion("jakarta")
    if err != nil {
        // handle error
    }
    ```

  - Jika terdapat error maka:
    - Berikan response code **500**
    - Berikan response message error
  - Jika semua oke maka:
    - Berikan response code **200**
    - Berikan response message berupa list informasi cuaca dari 10 wilayah tersebut dengan data yang sama dari fungsi `GetWeatherByRegion`

**Auth**

Lengkapi middleware `AuthAdmin` yang ada di `handler/handler.go` untuk melakukan pengecekan apakah user yang login adalah role `admin` atau tidak.
Kalian dapat mengambil cookie dengan key `user_login_role` untuk mendapatkan role user yang login. Jika bukan role `admin` maka :

- Berikan response code **401**
- Besponse message `{"error":"user login role not Admin"}`.

> hint: dapatkan cookie dari *http.Request menggunakan method `r.Cookie`

### Constraints

- ID mahasiswa hanya berupa kombinasi dari huruf dan angka, dengan panjang 5 karakter.
- Nama mahasiswa hanya berupa huruf, dengan panjang antara 1 hingga 50 karakter.
- Jurusan mahasiswa hanya berupa kombinasi dari huruf, dengan panjang antara 1 hingga 50 karakter.
- Kode program studi hanya berupa kombinasi dari huruf, dengan panjang 2 atau 3 karakter.
- Nama program studi hanya berupa huruf, dengan panjang antara 1 hingga 50 karakter.

### Test Case Examples

#### Test Case 1

**Input**:

```bash
> curl -i -X POST http://localhost:8080/user/register -H "Content-Type:application/json" -d "{ \"id\":\"A1112\", \"name\":\"Pratama\",  \"study\":\"TI\", \"role\":\"user\"}" 
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 200 OK
Date: Wed, 28 Mar 2023 06:31:50 GMT
Content-Length: 52
Content-Type: text/plain; charset=utf-8

{"username":"A1112","message":"register success"}
```

**Explanation**:

```txt
Test ini menjalankan end point /user/register dengan method POST dan request body {"id":"A1112", "name":"Pratama", "study":"TI", "role":"user"}.

Karena request body sudah valid, maka server akan mengembalikan response code 200 dan response message {"username":"A1112","message":"register success"}.
```

#### Test Case 2

```bash
> curl -i -X POST http://localhost:8080/user/login -H "Content-Type:application/json" -d "{ \"id\":\"A1112\", \"name\":\"Pratama\"}" 
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 200 OK
Date: Wed, 28 Mar 2023 06:31:50 GMT
Content-Length: 52
Content-Type: text/plain; charset=utf-8

{"username":"A1112","message":"login success"}
```

**Explanation**:

```txt
Test ini menjalankan end point /user/login dengan method POST dan request body {"id":"A1112", "name":"Pratama"}.

Karena request body sudah valid dan ID A1112 sudah didaftarkan, maka server akan mengembalikan response code 200 dan response message {"username":"A1112","message":"login success"}.
```

#### Test Case 3

**Input**:

```bash
> curl -i -X GET http://localhost:8080/study-program -H "Content-Type:application/json"
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 200 OK
Date: Wed, 28 Mar 2023 06:31:50 GMT
Content-Length: 52
Content-Type: text/plain; charset=utf-8

[{"code": "TI","name": "Teknik Informatika"}, { "code": "SI", "name": "Sistem Informasi" }, {...
```

**Explanation**:

```txt
Test ini menjalankan end point /study-program dengan method GET. 

Karena kita sudah login sebelumnya, maka server akan mengembalikan response code 200 dan response message berupa list program studi yang sudah didaftarkan.
```

### Template

Project Structure:

```txt
- 📁 data
  - 📄 list-study.txt
  - 📄 users.txt
- 📁 client
  - 📄 weather.go
- 📁 handler
  - 📄 handler.go
- 📁 middleware
  - 📄 middleware.go
- 📁 model
  - 📄 model.go
- 📄 main.go
```
