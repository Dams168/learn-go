package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// GetConnection function untuk menghubungkan aplikasi ke database MySQL
// dan mengembalikan objek *sql.DB yang mewakili koneksi database.
func GetConnection() *sql.DB {
	// dsn := "username:password@protocol(address:port)/dbname?param=value"
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database?parseTime=true")

	if err != nil {
		panic(err)
	}
	// Fungsi setmaxidleconns adalanya untuk mengatur jumlah koneksi idle maksimal yang dapat disimpan dalam pool koneksi database.
	// Dengan mengatur batas ini, kita dapat mengontrol berapa banyak koneksi yang tetap terbuka dan siap digunakan,
	// sehingga mengurangi overhead pembuatan koneksi baru setiap kali ada permintaan ke database.
	db.SetMaxIdleConns(10)

	// Fungsi setmaxopenconns adalah untuk mengatur jumlah koneksi maksimal yang dapat dibuka ke database pada satu waktu.
	// Dengan mengatur batas ini, kita dapat mengontrol berapa banyak koneksi yang dapat digunakan secara bersamaan,
	// sehingga mencegah kelebihan beban pada database dan memastikan kinerja yang optimal.
	db.SetMaxOpenConns(100)

	// Fungsi setconnmaxidletime adalah untuk mengatur durasi maksimal koneksi dapat tetap idle sebelum ditutup.
	// Jika koneksi tetap idle lebih lama dari durasi ini, maka koneksi tersebut akan ditutup secara otomatis oleh driver database.
	// Hal ini berguna untuk mengelola sumber daya dan mencegah koneksi yang tidak terpakai tetap terbuka terlalu lama.
	// Dengan mengatur waktu maksimal koneksi idle, kita dapat memastikan bahwa koneksi yang tidak aktif akan dibersihkan secara berkala,
	// sehingga mengurangi beban pada database dan meningkatkan efisiensi penggunaan koneksi.
	db.SetConnMaxIdleTime(5 * time.Minute)

	// Fungsi setconnmaxlifetime adalah untuk mengatur durasi maksimal koneksi dapat digunakan sebelum ditutup dan dibuat ulang.
	// Setelah koneksi mencapai batas waktu ini, koneksi akan ditutup secara otomatis oleh driver database,
	// dan ketika ada permintaan koneksi baru, koneksi baru akan dibuat untuk menggantikan koneksi yang sudah ditutup.
	db.SetConnMaxLifetime(60 * time.Minute)

	fmt.Println("Connected Cihuyyy")

	return db
}
