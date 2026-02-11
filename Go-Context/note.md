# Golang Context

Context di golang digunakan untuk mengatur batas waktu (timeout), pembatalan (cancellation), dan untuk meneruskan nilai-nilai antar goroutine. Context sangat berguna dalam aplikasi yang memerlukan kontrol terhadap eksekusi goroutine, terutama dalam aplikasi web atau layanan mikro.

## Membuat Context

Untuk membuat context, kita bisa menggunakan fungsi `context.Background()` atau `context.TODO()`.

Perbedaan utama antara keduanya adalah bahwa `context.Background()` digunakan sebagai konteks dasar, sedangkan `context.TODO()` digunakan ketika kita belum memutuskan konteks apa yang akan digunakan.
