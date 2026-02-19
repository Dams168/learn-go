package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second, //fungsi untuk mengatur waktu tunggu idle, yaitu waktu yang diizinkan untuk menjaga koneksi tetap terbuka tanpa aktivitas sebelum dianggap tidak aktif dan ditutup.
		ReadTimeout:  5 * time.Second, //fungsi untuk mengatur waktu tunggu baca, yaitu waktu maksimum yang diizinkan untuk membaca permintaan dari klien sebelum dianggap gagal.
		WriteTimeout: 5 * time.Second, //fungsi untuk mengatur waktu tunggu tulis, yaitu waktu maksimum yang diizinkan untuk menulis respons ke klien sebelum dianggap gagal.
	})

	app.Use("/api", func(ctx fiber.Ctx) error {
		fmt.Println("I'm middleware before processing request")
		err := ctx.Next()
		fmt.Println("I'm middleware after processing request")
		return err
	})

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	if fiber.IsChild() {
		fmt.Println("I'm child process")
	} else {
		fmt.Println("I'm parent process")
	}

	err := app.Listen(":3000", fiber.ListenConfig{EnablePrefork: true})
	if err != nil {
		panic(err)
	}
}
