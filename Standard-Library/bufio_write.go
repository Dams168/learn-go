package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	// Menulis beberapa string ke output dengan buffer
	// Penggunaan bufio.NewWriter membantu mengurangi jumlah operasi tulis langsung ke os.Stdout
	_, _ = writer.WriteString("Halo, dunia!\n")
	_, _ = writer.WriteString("Ini adalah contoh penggunaan bufio untuk menulis data dengan buffer.\n")
	writer.Flush()
}
