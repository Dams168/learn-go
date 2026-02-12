# Tipe data Column

- Tipe data column adalah tipe data yang digunakan untuk menyimpan nilai dalam sebuah kolom pada tabel database. Tipe data ini menentukan jenis nilai yang dapat disimpan dalam kolom tersebut, seperti angka, teks, tanggal, dan lain-lain. Pemilihan tipe data yang tepat sangat penting untuk memastikan integritas data dan kinerja database yang optimal. Beberapa contoh tipe data column yang umum digunakan adalah INT, VARCHAR, DATE, dan BOOLEAN.

# Mapping Tipe Data Column ke Tipe Data Go

- Mapping tipe data column ke tipe data Go adalah proses menghubungkan tipe data yang digunakan dalam database dengan tipe data yang digunakan dalam bahasa pemrograman Go. Hal ini penting untuk memastikan bahwa data yang diambil dari database dapat diproses dengan benar dalam aplikasi Go. Misalnya, tipe data INT dalam database dapat dipetakan ke tipe data int di Go, sementara tipe data VARCHAR dapat dipetakan ke tipe data string di Go. Mapping yang tepat membantu mencegah kesalahan dan memastikan bahwa aplikasi dapat berfungsi dengan baik saat berinteraksi dengan database.
  Berikut adalah contoh mapping tipe data column ke tipe data Go:
  | Tipe Data Column | Tipe Data Go |
  |------------------|--------------|
  | INT | int |
  | VARCHAR | string |
  | DATE | time.Time |
  | BOOLEAN | bool |
  | FLOAT | float64 |
  | DECIMAL | float64 |
  | TEXT | string |
  | BOOLEAN | bool |
  | BLOB | []byte |
  | TIMESTAMP | time.Time |
  | DATETIME | time.Time |
  | TIME | time.Time |

# Nullable Type Di GO

- Nullable type di Go adalah tipe data yang dapat menyimpan nilai null atau tidak ada nilai. Dalam konteks database, nullable type digunakan untuk merepresentasikan kolom yang dapat memiliki nilai null. Go tidak memiliki tipe data built-in untuk nullable type, tetapi kita dapat menggunakan pointer atau package seperti "sql.NullString", "sql.NullInt64", dan lain-lain untuk menangani nullable type. Dengan menggunakan nullable type, kita dapat dengan mudah memeriksa apakah sebuah nilai ada atau tidak sebelum menggunakannya dalam aplikasi Go.

# Prepared Statement

- Prepared statement adalah fitur dalam database yang memungkinkan kita untuk menyiapkan sebuah pernyataan SQL sekali dan kemudian mengeksekusinya berkali-kali dengan parameter yang berbeda. Hal ini meningkatkan kinerja dan keamanan aplikasi karena pernyataan SQL hanya perlu di-parse dan di-compile satu kali, serta membantu mencegah serangan SQL injection dengan memisahkan data dari kode SQL. Dalam Go, kita dapat menggunakan metode "Prepare" dari package "database/sql" untuk membuat prepared statement, dan kemudian menggunakan metode "Exec" atau "Query" untuk mengeksekusinya dengan parameter yang diinginkan.

Perbedaan antara prepared statement dan regular statement adalah bahwa prepared statement memungkinkan kita untuk menyiapkan pernyataan SQL sekali dan mengeksekusinya berkali-kali dengan parameter yang berbeda, sedangkan regular statement harus di-parse dan di-compile setiap kali dieksekusi. Prepared statement juga lebih aman karena memisahkan data dari kode SQL, sehingga membantu mencegah serangan SQL injection. Regular statement, di sisi lain, rentan terhadap serangan SQL injection jika tidak digunakan dengan hati-hati.

# Database Transaction

- Database transaction adalah serangkaian operasi database yang dilakukan sebagai satu unit kerja yang tidak dapat dibagi. Transaksi memastikan bahwa semua operasi dalam transaksi berhasil atau tidak ada yang berhasil sama sekali, sehingga menjaga integritas data. Dalam Go, kita dapat menggunakan metode "Begin" dari package "database/sql" untuk memulai sebuah transaksi, dan kemudian menggunakan metode "Commit" untuk menyimpan perubahan atau "Rollback" untuk membatalkan perubahan jika terjadi kesalahan. Transaksi sangat berguna dalam situasi di mana beberapa operasi harus dilakukan bersama-sama, seperti saat memindahkan dana antara dua akun bank.

- Perbedaan antara transaction dan non-transaction adalah bahwa transaction melibatkan serangkaian operasi yang dilakukan sebagai satu unit kerja yang tidak dapat dibagi, sedangkan non-transaction melibatkan operasi yang dilakukan secara individual tanpa jaminan konsistensi data. Dalam transaction, jika salah satu operasi gagal, semua perubahan yang dilakukan dalam transaksi akan dibatalkan, sedangkan dalam non-transaction, perubahan yang sudah dilakukan tidak dapat dibatalkan jika terjadi kesalahan pada operasi berikutnya. Transaction memberikan jaminan konsistensi data, sementara non-transaction tidak memiliki jaminan tersebut.

# Repository Pattern

- Repository pattern adalah pola desain perangkat lunak yang digunakan untuk memisahkan logika akses data dari logika bisnis dalam aplikasi. Pola ini menyediakan antarmuka yang konsisten untuk berinteraksi dengan sumber data, seperti database, sehingga memungkinkan perubahan pada sumber data tanpa mempengaruhi logika bisnis. Dalam Go, kita dapat mengimplementasikan repository pattern dengan membuat interface repository yang mendefinisikan metode untuk operasi CRUD (Create, Read, Update, Delete), dan kemudian membuat implementasi konkret dari interface tersebut yang berinteraksi dengan database.

- Flow repository pattern di go
  1. Definisikan interface repository yang mendefinisikan metode untuk operasi CRUD.
  2. Buat struktur data (struct) yang mewakili entitas bisnis yang akan disimpan dalam database.
  3. Buat implementasi konkret dari interface repository yang berinteraksi dengan database menggunakan package "database/sql" atau ORM (Object-Relational Mapping).
  4. Gunakan dependency injection untuk menyuntikkan instance repository ke dalam layanan atau handler yang membutuhkan akses data.
  5. Gunakan metode repository dalam logika bisnis untuk melakukan operasi CRUD tanpa perlu mengetahui detail implementasi akses data.

- Berikut adalah bentuk gamabaran flow repository pattern di Go:
  Business Logic <-Call-> Repository Interface <-Implement-> Repository Implementation <-Interact-> Database
  Repository Interface<-Use-> Entity Struct

- Entity / model adalah struktur data yang mewakili entitas bisnis dalam aplikasi. Entity biasanya berisi atribut atau properti yang sesuai dengan kolom dalam tabel database, serta metode yang berkaitan dengan entitas tersebut. Entity digunakan untuk memodelkan data yang akan disimpan, diambil, atau dimanipulasi dalam aplikasi. ketika kita query ke repository, dibanding kita mengembalikan array of map[string]interface{} yang kurang type safety, kita bisa mengembalikan entity struct yang sudah didefinisikan sebelumnya sehingga lebih aman dan mudah digunakan dalam logika bisnis.
