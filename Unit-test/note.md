testing di Go

Unit testing adalah proses pengujian bagian terkecil dari kode program, yaitu unit, untuk memastikan bahwa unit tersebut berfungsi sesuai yang diharapkan. Di Go, unit testing dilakukan dengan menggunakan paket "testing" yang sudah disediakan secara bawaan.

Berikut adalah langkah-langkah dasar untuk membuat unit test di Go:

1. Buat file test dengan akhiran `_test.go`. Misalnya, jika Anda memiliki file `hello_world.go`, buat file `hello_world_test.go`.
2. Impor paket "testing" di dalam file test.
3. Buat fungsi test dengan nama yang diawali dengan `Test` diikuti dengan nama fungsi yang akan diuji. Fungsi ini harus menerima parameter bertipe `*testing.T`.
4. Gunakan metode dari `*testing.T` untuk melaporkan kegagalan atau keberhasilan pengujian, seperti `t.Error`, `t.Fail`, atau `t.Log`.

Command untuk menjalankan unit test:

1. `go test` - Menjalankan semua unit test di dalam paket saat ini.
2. `go test -v` - Menjalankan unit test dengan output yang lebih rinci (verbose).
3. `go test ./...` - Menjalankan unit test di semua subdirektori dari paket saat ini.
4. `go test -run TestFunctionName` - Menjalankan hanya unit test tertentu yang sesuai dengan pola nama fungsi yang diberikan.
5. `go test -cover` - Menampilkan informasi tentang cakupan kode (code coverage) dari unit test yang dijalankan.

Benchmarking di Go
Benchmarking adalah proses mengukur kinerja suatu fungsi atau bagian kode program untuk mengetahui seberapa cepat atau efisien kode tersebut berjalan. Di Go, benchmarking dilakukan dengan menggunakan paket "testing" yang sama dengan unit testing.

Command untuk Benchmarking:

1. `go test -bench=.` - Menjalankan semua benchmark di dalam paket saat ini.
2. `go test -bench=BenchmarkFunctionName` - Menjalankan hanya benchmark tertentu yang sesuai dengan pola nama fungsi yang diberikan.
3. `go test -bench=. -benchtime=5s` - Menjalankan semua benchmark dengan waktu eksekusi 5 detik per benchmark.

Mock di Go
Mocking adalah teknik dalam pengujian perangkat lunak di mana objek tiruan (mock objects) digunakan untuk meniru perilaku objek nyata. Ini berguna untuk mengisolasi unit yang diuji dan mengontrol lingkungan pengujian. Di Go, mocking dapat dilakukan dengan menggunakan pustaka pihak ketiga seperti "testify/mock" atau "gomock".
