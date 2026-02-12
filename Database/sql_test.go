package database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customers(id, name) VALUES('james', 'james bond')"

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customers")
}

func TestQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customers"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Fungsi rows.Next() digunakan untuk iterasi melalui hasil query yang dikembalikan oleh database.
	// Fungsi ini akan mengembalikan true jika masih ada baris hasil yang dapat diproses, dan false jika tidak ada lagi baris yang tersedia.
	// Dengan menggunakan rows.Next(), kita dapat melakukan iterasi melalui setiap baris hasil query dan mengambil data yang diperlukan untuk diproses lebih lanjut.
	for rows.Next() {
		var id string
		var name string
		// Scan digunakan untuk memindai nilai dari kolom hasil query ke dalam variabel yang sesuai.
		// disesuaikan dengan urutan kolom pada query SELECT.
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id, "Name:", name)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func TestSqlComplext(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customers"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var email sql.NullString // untuk menangani nilai NULL pada kolom email, karna go tidak mendukung null secara langsung
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var married bool
		var createdAt time.Time
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("===========================================")
		fmt.Println("id :", id)
		fmt.Println("name :", name)
		if email.Valid {
			fmt.Println("email :", email.String)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		if birthDate.Valid {
			fmt.Println("birthDate :", birthDate.Time)
		}
		fmt.Println("married :", married)
		fmt.Println("createdAt :", createdAt)
	}

}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin'; #" // ini adalah contoh sql injection,
	password := "passwordsalah"

	// query di bawah ini rentan terhadap sql injection
	// karena input dari user langsung dimasukkan ke dalam query tanpa validasi atau sanitasi terlebih dahulu
	// sehingga user dapat memanipulasi query sesuai keinginannya
	// paling parah bisa menghapus seluruh data pada tabel users
	// contoh input user di atas akan menghasilkan query seperti ini:
	// SELECT username FROM users WHERE username = 'admin'; #' AND password = 'passwordsalah' LIMIT 1
	// yang mana query tersebut akan mengabaikan bagian AND password = 'passwordsalah' LIMIT 1
	// sehingga query tersebut akan selalu mengembalikan data user dengan username 'admin'
	// terlepas dari password yang dimasukkan user
	query := "SELECT username FROM users WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success", username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin"
	password := "admin"

	// menggunakan parameterized query untuk mencegah sql injection
	// dengan cara mengganti input user dengan tanda tanya (?) pada query dan memasukkan input user sebagai argumen pada method
	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"

	// ini adalah cara yang aman karena input user tidak langsung dimasukkan ke dalam query
	// sehingga tidak dapat memanipulasi query sesuai keinginannya, karena input user akan dianggap sebagai data biasa
	//urutan parameter harus sesuai dengan urutan tanda tanya pada query
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success", username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestPreparedStatement(t *testing.T) {

	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	// ini digunakan ketika kita ingin mengeksekusi query yang sama berkali-kali dengan parameter yang berbeda, sehingga kita tidak perlu menulis ulang query setiap kali ingin mengeksekusinya
	// dengan menggunakan prepared statement, kita dapat meningkatkan performa aplikasi karena query hanya perlu di-parse dan di-compile sekali, serta membantu mencegah serangan SQL injection dengan memisahkan data dari kode SQL
	// misal data dari csv atau data dari input user yang banyak
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "Dam" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		// mendapatkan last insert id, jika menggunakan auto increment
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "Dam" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}
	// err = tx.Commit() // untuk menyimpan perubahan ke database
	err = tx.Rollback() // untuk membatalkan semua perubahan yang dilakukan dalam transaksi
	if err != nil {
		panic(err)
	}
}
