package db

var connection string

// init function akan dipanggil otomatis saat package diimpor dan digunakan untuk inisialisasi package
func init() {
	connection = "MySQL Connect"
}

func GetDatabase() string {
	return connection
}
