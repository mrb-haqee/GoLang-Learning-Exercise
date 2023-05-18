# Assignment Backend Lanjutan 1

## Student Portal 3

### Description

Kita akan mengimplementasikan http server sederhana untuk membuat aplikasi portal mahasiswa. Aplikasi ini akan digunakan untuk mengelola data mahasiswa, dan mengelola data mata kuliah yang diambil oleh mahasiswa.

Data user pada program ini disimpan dalam variabel Students dengan format string berisi informasi ID, nama, dan jurusan, disimpan ke dalam format string yang digabung dengan _underscore_ (`_`). Formatnya adalah `[id]_[name]_[kode jurusan]`

Data tersebut disimpan ke dalam file `txt` yang ada di `data/users.txt`

Contoh data user:

```txt
A1111_Pratama_SI
B1112_Jaya_TI
```

Program juga menyimpan list jurusan yang ada di dalam file `txt` pada direktori `data/list-study.txt`. Terdapat 15 jurusan yang disedikan dengan format `[kode jurusan]_[nama jurusan]`. Contoh data jurusan:

```txt
TI_Teknik Informatika
TK_Teknik Komputer
SI_Sistem Informasi
...
```

Aplikasi ini akan memiliki fitur sederhana sebagai berikut:

- Get List Study
- Add Student
- Delete Student

### Instruction

Kalian diminta untuk melengkapi handler yang sudah disediakan untuk mengerjakan aplikasi ini. Berikut adalah hal-hal yang harus diperhatikan dalam mengerjakan aplikasi student portal:

- **GetStudyProgram**, adalah handler yang digunakan untuk menampilkan list jurusan yang disimpan di file `data/list-study.txt` dan kemudian mengembalikan datanya dalam bentuk JSON pada end point `/study-program` dengan method **GET**. Aturan yang harus diperhatikan adalah sebagai berikut:
  - Hanya boleh menggunakan method **GET** jika tidak, maka:
    - Berikan response code **405** (Method Not Allowed)
    - Berikan responnse message `{"error":"Method is not allowed!"}`
  - Jika sudah benar, maka:
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

- **AddUser**, adalah handler yang digunakan untuk menambahkan user baru ke dalam file `data/users.txt` pada endpoint `/user/add` menggunakan method **POST**. Aturan yang harus diperhatikan adalah sebagai berikut:
  - Hanya boleh menggunakan method **POST** jika tidak, maka:
    - Berikan response code **405**
    - Berikan responnse message `{"error":"Method is not allowed!"}`
  - Harus memberikan request body dengan format JSON `{"id": "<id user>", "name": <name>, "study": <study>}`. Contoh request body:

    ```json
    {
        "id": "A1234",
        "name": "Pratama",
        "study": "TI"
    }
    ```

  - Jika request body ID, name, atau study kosong(`""`), maka:
    - Berikan response code **400**
    - Berikan responnse message `{"error":"ID, name, or study code is empty"}`
  - Jika ID user sudah ada di file `data/users.txt`, maka:
    - Berikan response code **400**
    - Berikan responnse message `{"error":"user id already exist"}`
  - Jika study code tidak ditemukan di file `data/list-study.txt`, maka:
    - Berikan response code **400**
    - Berikan responnse message `{"error":"study code not found"}`
  - Jika semua oke maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<id user>","message":"add user success"}`
    - Tambahkan data user ke dalam file `data/users.txt` dengan format yang sudah dijelaskan sebelumnya.

- **DeleteUser** adalah handler yang digunakan untuk menghapus user dari file `data/users.txt` pada endpoint `/user/delete` menggunakan method **DELETE**. Aturan yang harus diperhatikan adalah sebagai berikut:
  - Hanya boleh menggunakan method **DELETE** jika tidak, maka:
    - Berikan response code **405**
    - Berikan responnse message `{"error":"Method is not allowed!"}`
  - Harus memberikan request dengan query parameter `id` yang berisi ID user yang akan dihapus. Contoh request:

    ```http
    DELETE localhost:8080/user/delete?id=A1234 => Menghapus user dengan ID: A1234
    ```

    Jika tidak ada query parameter `id` maka:
    - Berikan response code **400**
    - Berikan responnse message `{"error":"user id is empty"}`

  - Jika ID user tidak ada di file `data/users.txt`, maka:
    - Berikan response code **400**
    - Berikan responnse message `{"error":"user id not found"}`

  - Jika semua oke maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<id user>","message":"delete success"}`
    - Hapus data user dari file `data/users.txt` dengan format yang sudah dijelaskan sebelumnya.

> Note: Kalian dapat mengecek Method yang dikirim pada handler dengan menggunakan `r.Method` dan mengecek apakah method yang dikirim sama dengan method yang diharapkan atau tidak.

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

#### Test Case 2

**Input**:

```bash
> curl -i -X POST http://localhost:8080/user/add -H "Content-Type:application/json" -d "{ \"id\":\"A1112\", \"name\":\"Pratama\",  \"study_code\":\"TI\"}" 
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 200 OK
Date: Wed, 28 Sep 2022 06:31:50 GMT
Content-Length: 52
Content-Type: text/plain; charset=utf-8

{"username":"A1112","message":"add user success"}
```

#### Test Case 3

**Input**:

```bash
> curl -i -X DELETE http://localhost:8080/user/delete?id=A1112 
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 200 OK
Date: Wed, 19 Apr 2023 02:23:16 GMT
Content-Length: 48
Content-Type: text/plain; charset=utf-8

{"username":"A1112","message":"delete success"}
```

Kalian dapat mencoba semua fungsi yang ada di atas dengan menggunakan `curl` atau `postman`.

### Template

Project Structure:

```txt
- 📁 data
  - 📄 list-study.txt
  - 📄 users.txt
- 📁 model
  - 📄 model.go
- 📄 main.go
```
