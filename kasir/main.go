package main

import "fmt"

func main() {
	items := [3]string{"Mie Ayam", "Nasi Goreng", "Es Teh"}
	prices := [3]int{10000, 15000, 5000}
	total := 0

	var meja int
	var kondisi bool
	var pesanan int

	fmt.Println("Menu:")
	fmt.Println("1:", items[0], "Rp", prices[0])
	fmt.Println("2:", items[1], "Rp", prices[1])
	fmt.Println("3:", items[2], "Rp", prices[2])
	fmt.Println("Ketik nomor menu untuk memesan, atau ketik 4 untuk selesai:")

	for {
		fmt.Scan(&pesanan)

		if pesanan == 1 {
			total += prices[0]
			fmt.Println("Anda memesan", items[0], "- Total sementara:", total)
		} else if pesanan == 2 {
			total += prices[1]
			fmt.Println("Anda memesan", items[1], "- Total sementara:", total)
		} else if pesanan == 3 {
			total += prices[2]
			fmt.Println("Anda memesan", items[2], "- Total sementara:", total)
		} else if pesanan == 4 {
			break
		} else {
			fmt.Println("Pesanan tidak valid. Silakan pilih menu yang benar.")
		}

		fmt.Println("Ingin memesan lagi? Ketik nomor menu (1-3), atau ketik 4 untuk selesai:")
	}

	fmt.Println("Total pesanan Anda:", total)
	fmt.Println("Apakah Anda sudah yakin dengan pesanan Anda? (ketik 1 untuk ya, 0 untuk tidak)")
	fmt.Scan(&kondisi)

	if kondisi {
		fmt.Println("Anda berada di meja nomor berapa?")
		fmt.Scan(&meja)
		fmt.Printf("Pesanan Anda sebesar Rp %d akan dikirim ke meja %d.\n", total, meja)
	} else {
		fmt.Println("Anda tidak yakin? Pesanan dibatalkan.")
	}

	fmt.Println("Terima kasih telah memesan.")
}
