# Parallel Programming

Parallel programming adalah suatu paradigma pemrograman yang memungkinkan eksekusi beberapa proses atau thread secara bersamaan untuk meningkatkan efisiensi dan kinerja aplikasi. Dalam konteks bahasa pemrograman Go, parallel programming sering diimplementasikan menggunakan goroutine.

Contoh Parallel adalah ketika sebuah aplikasi web menerima banyak permintaan dari pengguna secara bersamaan. Misal, ketika banyak pengguna mengakses halaman web yang sama, server web dapat memproses permintaan tersebut secara paralel menggunakan goroutine untuk menangani setiap permintaan secara independen. Hal ini memungkinkan server untuk merespons lebih cepat dan meningkatkan pengalaman pengguna.

# Process vs Thread vs Goroutine

- **Process** adalah program yang sedang berjalan yang memiliki ruang alamat memori sendiri. Setiap proses berjalan secara independen dan tidak berbagi memori dengan proses lain. Komunikasi antar proses biasanya dilakukan melalui mekanisme seperti inter-process communication (IPC). Process menkonsumsi sumber daya sistem yang lebih besar dibandingkan dengan thread atau goroutine.
- **Thread** adalah unit eksekusi terkecil dalam sebuah proses. Beberapa thread dalam satu proses berbagi ruang alamat memori yang sama, sehingga mereka dapat berkomunikasi lebih mudah dan efisien. Namun, ini juga berarti bahwa kesalahan pada satu thread dapat mempengaruhi thread lain dalam proses yang sama. Thread menggunakan memori dan sumber daya sistem yang lebih sedikit dibandingkan dengan proses, tetapi lebih banyak dibandingkan dengan goroutine.
- **Goroutine** adalah konsep unik dalam bahasa Go yang merupakan lightweight thread. Goroutine dikelola oleh runtime Go dan memiliki overhead yang sangat kecil dibandingkan dengan thread tradisional. Ribuan goroutine dapat dijalankan secara bersamaan dalam satu proses, dan mereka berkomunikasi menggunakan channel, yang menyediakan cara yang aman untuk berbagi data antar goroutine tanpa perlu khawatir tentang kondisi balapan (race conditions). Goroutine sangat efisien dalam hal penggunaan memori dan sumber daya sistem, membuatnya ideal untuk aplikasi yang membutuhkan concurrency tinggi.

# Parralel vs Concurrent

- **Paralel** mengacu pada eksekusi beberapa tugas secara bersamaan pada waktu yang sama, biasanya pada multiple CPU cores. Dalam konteks paralelisme, tugas-tugas tersebut benar-benar berjalan bersamaan, sehingga dapat mempercepat penyelesaian pekerjaan. Contohnya adalah ketika sebuah aplikasi melakukan komputasi berat yang dibagi menjadi beberapa bagian dan setiap bagian dijalankan pada core CPU yang berbeda secara bersamaan.
- **Concurrent** mengacu pada kemampuan untuk menangani beberapa tugas pada waktu yang sama, tetapi tidak harus secara bersamaan. Dalam konteks concurrency, tugas-tugas tersebut dapat dimulai, dijeda, dan dilanjutkan pada waktu yang berbeda, tergantung pada ketersediaan sumber daya. Contohnya adalah ketika sebuah aplikasi web menangani banyak permintaan dari pengguna dengan cara mengalihkan perhatian antara permintaan-permintaan tersebut, sehingga terlihat seperti semuanya berjalan bersamaan, meskipun sebenarnya hanya satu permintaan yang diproses pada satu waktu.

Perbedaan utama antara paralel dan concurrent adalah bahwa paralelisme melibatkan eksekusi simultan dari tugas-tugas pada multiple CPU cores, sedangkan concurrency melibatkan pengelolaan beberapa tugas pada waktu yang sama tanpa harus menjalankannya secara bersamaan.

# CPU Bound vs I/O Bound

- **CPU Bound** adalah kondisi di mana kinerja suatu program atau proses dibatasi oleh kecepatan prosesor (CPU). Dalam situasi ini, program menghabiskan sebagian besar waktunya untuk melakukan perhitungan atau operasi yang intensif pada CPU, sehingga peningkatan kecepatan CPU akan secara langsung meningkatkan kinerja program. Contoh dari CPU bound adalah aplikasi yang melakukan komputasi matematis yang kompleks, seperti rendering grafis atau simulasi ilmiah.

- **I/O Bound** adalah kondisi di mana kinerja suatu program atau proses dibatasi oleh kecepatan operasi input/output (I/O). Dalam situasi ini, program menghabiskan sebagian besar waktunya menunggu operasi I/O selesai, seperti membaca atau menulis data ke disk, jaringan, atau perangkat lain. Peningkatan kecepatan CPU tidak akan banyak mempengaruhi kinerja program karena bottleneck terletak pada operasi I/O. Contoh dari I/O bound adalah aplikasi yang sering mengakses database atau berkomunikasi melalui jaringan.

Perbedaan utama antara CPU bound dan I/O bound adalah sumber utama pembatas kinerja program: CPU bound dibatasi oleh kecepatan prosesor, sedangkan I/O bound dibatasi oleh kecepatan operasi input/output.

# Goroutine

Goroutine adalah sebuah thread ringan yang dikelola oleh runtime Go. Goroutine memungkinkan eksekusi fungsi secara konkuren (concurrent) dengan overhead yang sangat kecil dibandingkan dengan thread tradisional. Goroutine dibuat menggunakan kata kunci `go` diikuti oleh pemanggilan fungsi. Ukuran goroutine sangat kecil, biasanya hanya beberapa kilobyte, sehingga memungkinkan ribuan goroutine berjalan secara bersamaan dalam satu proses.

Cara Kerja Goroutine:

1. **Pembuatan Goroutine**: Ketika sebuah goroutine dibuat menggunakan kata kunci `go`, runtime Go mengalokasikan memori yang diperlukan untuk goroutine tersebut dan menambahkannya ke dalam scheduler goroutine.
2. **Scheduler Goroutine**: Runtime Go memiliki scheduler yang mengelola eksekusi goroutine. Scheduler ini bertanggung jawab untuk menjadwalkan goroutine pada thread OS yang tersedia, sehingga goroutine dapat berjalan secara konkuren.
3. **Stack Dinamis**: Goroutine menggunakan stack yang dapat tumbuh dan menyusut secara dinamis sesuai kebutuhan. Ini memungkinkan goroutine untuk menggunakan memori secara efisien, karena stack awalnya sangat kecil dan hanya akan tumbuh jika diperlukan.
4. **Komunikasi melalui Channel**: Goroutine dapat berkomunikasi satu sama lain menggunakan channel, yang menyediakan cara yang aman untuk berbagi data antar goroutine tanpa perlu khawatir tentang kondisi balapan (race conditions).
