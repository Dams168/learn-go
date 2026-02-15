# Flow dalam pembuatan RESTful API dengan Go

1. Pertama-tama, kita akan membuat sebuah struct yang akan menjadi model dari data yang akan kita simpan di database. Misalnya, kita akan membuat sebuah struct bernama "Category" yang memiliki field "Id" dan "Name".

2. Setelah itu, kita akan membuat sebuah interface yang akan menjadi contract dari service yang akan kita buat. Misalnya, kita akan membuat sebuah interface bernama "CategoryService" yang memiliki method "Create", "Update", "Delete", dan "FindById".

3. Kemudian, kita akan membuat sebuah struct yang akan menjadi implementasi dari interface yang sudah kita buat. Misalnya, kita akan membuat sebuah struct bernama "CategoryServiceImpl" yang memiliki field "CategoryRepository", "DB", dan "Validate".

4. Setelah itu, kita akan mengimplementasikan method-method yang sudah didefinisikan pada interface "CategoryService" di dalam struct "CategoryServiceImpl". Misalnya, kita akan mengimplementasikan method "Create" yang akan membuat sebuah category baru di database, method "Update" yang akan mengupdate sebuah category di database, method "Delete" yang akan menghapus sebuah category di database, dan method "FindById" yang akan mencari sebuah category di database berdasarkan id.

5. Setelah itu, kita akan membuat sebuah handler yang akan menjadi endpoint dari RESTful API yang akan kita buat. Misalnya, kita akan membuat sebuah handler bernama "CategoryHandler" yang memiliki field "CategoryService".

6. Kemudian, kita akan mengimplementasikan method-method yang akan menjadi endpoint dari RESTful API di dalam struct "CategoryHandler". Misalnya, kita akan mengimplementasikan method "Create" yang akan menjadi endpoint untuk membuat sebuah category baru, method "Update" yang akan menjadi endpoint untuk mengupdate sebuah category, method "Delete" yang akan menjadi endpoint untuk menghapus sebuah category, dan method "FindById" yang akan menjadi endpoint untuk mencari sebuah category berdasarkan id.

7. Setelah itu, kita akan membuat sebuah router yang akan mengatur endpoint-endpoint dari RESTful API yang akan kita buat. Misalnya, kita akan membuat sebuah router bernama "Router" yang memiliki field "CategoryHandler".

8. Kemudian, kita akan mengatur endpoint-endpoint dari RESTful API di dalam struct "Router". Misalnya, kita akan mengatur endpoint "/categories" untuk method "Create", endpoint "/categories/{id}" untuk method "Update", endpoint "/categories/{id}" untuk method "Delete", dan endpoint "/categories/{id}" untuk method "FindById".

9. Setelah itu, kita akan menjalankan server yang akan menjadi host dari RESTful API yang akan kita buat. Misalnya, kita akan menjalankan server di port 8080 dengan menggunakan router yang sudah kita buat.

10. Setelah server berjalan, kita bisa mengakses endpoint-endpoint dari RESTful API yang sudah kita buat dengan menggunakan tools seperti Postman atau curl untuk melakukan operasi-operasi seperti membuat category baru, mengupdate category, menghapus category, dan mencari category berdasarkan id.

# Catatan Penting yang harus diperhatikan di GO

- Di Go, kita tidak bisa menggunakan try-catch seperti di bahasa pemrograman lain. Sebagai gantinya, kita harus menggunakan panic dan recover untuk menangani error. Namun, sebaiknya kita tidak menggunakan panic dan recover secara langsung, melainkan menggunakan helper yang sudah kita buat untuk menangani error dengan lebih baik.

- Di Go, kita tidak bisa menggunakan exception seperti di bahasa pemrograman lain. Sebagai gantinya, kita harus menggunakan error untuk menangani error. Namun, sebaiknya kita tidak menggunakan error secara langsung, melainkan menggunakan helper yang sudah kita buat untuk menangani error dengan lebih baik.

- Di Go, kita tidak bisa menggunakan class seperti di bahasa pemrograman lain. Sebagai gantinya, kita harus menggunakan struct untuk membuat sebuah tipe data baru. Namun, sebaiknya kita tidak menggunakan struct secara langsung, melainkan menggunakan interface untuk membuat sebuah contract yang harus diimplementasikan oleh struct yang akan kita buat.

- Di Go, kita tidak bisa menggunakan inheritance seperti di bahasa pemrograman lain. Sebagai gantinya, kita harus menggunakan composition untuk membuat sebuah tipe data baru yang memiliki field dari tipe data lain. Namun, sebaiknya kita tidak menggunakan composition secara langsung, melainkan menggunakan interface untuk membuat sebuah contract yang harus diimplementasikan oleh struct yang akan kita buat.

- Di Go, kita tidak bisa menggunakan package seperti di bahasa pemrograman lain. Sebagai gantinya, kita harus menggunakan module untuk mengatur dependensi dari project yang kita buat. Namun, sebaiknya kita tidak menggunakan module secara langsung, melainkan menggunakan go mod untuk mengatur dependensi dari project yang kita buat dengan lebih baik.

- Di Go, kita tidak bisa menggunakan class method seperti di bahasa pemrograman lain. Sebagai gantinya, kita harus menggunakan function untuk membuat sebuah method yang bisa dipanggil oleh struct yang kita buat. Namun, sebaiknya kita tidak menggunakan function secara langsung, melainkan menggunakan method yang sudah kita buat untuk membuat sebuah method yang bisa dipanggil oleh struct yang kita buat dengan lebih baik.

- Di Go, kita tidak bisa menggunakan constructor seperti di bahasa pemrograman lain. Sebagai gantinya, kita harus menggunakan function untuk membuat sebuah instance dari struct yang kita buat. Namun, sebaiknya kita tidak menggunakan function secara langsung, melainkan menggunakan function yang sudah kita buat untuk membuat sebuah instance dari struct yang kita buat dengan lebih baik.

- Di Go, kita tidak bisa menggunakan destructor seperti di bahasa pemrograman lain. Sebagai gantinya, kita harus menggunakan defer untuk menjalankan sebuah function setelah function yang sedang berjalan selesai. Namun, sebaiknya kita tidak menggunakan defer secara langsung, melainkan menggunakan helper yang sudah kita buat untuk menjalankan sebuah function setelah function yang sedang berjalan selesai dengan lebih baik.
