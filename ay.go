package main

import (
	"fmt"
	"sort"
	"time"
)

const jumlahSlotMaks = 10 // Kapasitas maksimum slot parkir

type Kendaraan struct {
	PlatNomor  string
	Jenis      string
	WaktuMasuk time.Time
}

var slotParkir [jumlahSlotMaks]Kendaraan

func main() {
	for {
		fmt.Println("Selamat Datang di Aplikasi Parkir")
		fmt.Println("====================================")
		fmt.Println("1. Menu User")
		fmt.Println("2. Menu Admin")
		fmt.Println("3. Keluar")
		fmt.Println("====================================")

		var pilihanMenuUtama int
		fmt.Print("Masukkan pilihan menu: ")
		fmt.Scan(&pilihanMenuUtama)

		switch pilihanMenuUtama {
		case 1:
			menuUser()
		case 2:
			menuAdmin()
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi parkir.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func menuUser() {
	for {
		fmt.Println("====================================")
		fmt.Println("Menu User:")
		fmt.Println("1. Parkir Kendaraan")
		fmt.Println("2. Keluarkan Kendaraan")
		fmt.Println("3. Mencari Lahan Parkir Yang Kosong")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Println("====================================")

		var pilihanUser int
		fmt.Print("Masukkan pilihan menu user: ")
		fmt.Scan(&pilihanUser)

		switch pilihanUser {
		case 1:
			parkirKendaraan()
		case 2:
			keluarkanKendaraan()
		case 3:
			cariSlotKosong()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func menuAdmin() {
	for {
		fmt.Println("====================================")
		fmt.Println("Menu Admin:")
		fmt.Println("1. Tampilkan Status Parkir")
		fmt.Println("2. Urutkan Slot Parkir (Descending)")
		fmt.Println("3. Urutkan Slot Parkir (Ascending)")
		fmt.Println("4. Tambah Data Parkir")
		fmt.Println("5. Hapus Data Parkir")
		fmt.Println("6. Ubah Data Parkir")
		fmt.Println("7. Binary Search Nomor Plat")
		fmt.Println("8. Total Pendapatan")
		fmt.Println("9. Kembali ke Menu Utama")
		fmt.Println("====================================")

		var pilihanAdmin int
		fmt.Print("Masukkan pilihan menu admin: ")
		fmt.Scan(&pilihanAdmin)

		switch pilihanAdmin {
		case 1:
			tampilkanStatusParkir()
		case 2:
			urutkanSlotParkirDescending()
		case 3:
			urutkanSlotParkirAscending()
		case 4:
			tambahDataParkir()
		case 5:
			hapusDataParkir()
		case 6:
			ubahDataParkir()
		case 7:
			binarySearchPlatNomor()
		case 8:
			totalPendapatan()
		case 9:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func totalPendapatan() {
	totalMobil := 0
	totalMotor := 0

	for i := 0; i < jumlahSlotMaks; i++ {
		if slotParkir[i].PlatNomor != "" {
			biaya := hitungBiaya(slotParkir[i].WaktuMasuk, time.Now(), slotParkir[i].Jenis)
			if slotParkir[i].Jenis == "mobil" {
				totalMobil += biaya
			} else if slotParkir[i].Jenis == "motor" {
				totalMotor += biaya
			}
		}
	}

	fmt.Printf("Total Pendapatan Parkir Mobil: Rp%d\n", totalMobil)
	fmt.Printf("Total Pendapatan Parkir Motor: Rp%d\n", totalMotor)
	fmt.Printf("Total Pendapatan: Rp%d\n", totalMobil+totalMotor)
}

func parkirKendaraan() {
	slotKosong := cariSlotKosong()
	if slotKosong == -1 {
		fmt.Println("Parkir penuh! Pilih slot lain.")
		return
	}

	var platNomor string
	var jenis string
	fmt.Print("Masukkan nomor plat kendaraan: ")
	fmt.Scan(&platNomor)
	fmt.Print("Jenis Kendaraan (mobil/motor): ")
	fmt.Scan(&jenis)
	kendaraan := Kendaraan{
		PlatNomor:  platNomor,
		Jenis:      jenis,
		WaktuMasuk: time.Now(),
	}
	slotParkir[slotKosong] = kendaraan
	fmt.Printf("Kendaraan dengan nomor plat %s diparkir di slot %d\n", platNomor, slotKosong+1)
}

func keluarkanKendaraan() {
	var platNomor string
	fmt.Print("Masukkan nomor plat kendaraan yang ingin dikeluarkan: ")
	fmt.Scan(&platNomor)

	// Mengurutkan slot parkir sebelum melakukan pencarian
	urutkanSlotParkirAscending()

	index := binarySearch(platNomor)
	if index == -1 {
		fmt.Println("Kendaraan dengan nomor plat tersebut tidak ditemukan.")
		return
	}

	waktuMasuk := slotParkir[index].WaktuMasuk
	waktuKeluar := time.Now()
	biaya := hitungBiaya(waktuMasuk, waktuKeluar, slotParkir[index].Jenis)
	durasiParkir := waktuKeluar.Sub(waktuMasuk)
	fmt.Printf("Kendaraan dengan nomor plat %s dikeluarkan dari slot %d.\n", slotParkir[index].PlatNomor, index+1)
	fmt.Printf("Waktu Masuk: %s\n", waktuMasuk.Format("15:04:05"))
	fmt.Printf("Waktu Keluar: %s\n", waktuKeluar.Format("15:04:05"))
	fmt.Printf("Durasi Parkir: %s\n", durasiParkir)
	fmt.Printf("Biaya parkir: Rp%d\n", biaya)
	slotParkir[index] = Kendaraan{}
}

func tampilkanStatusParkir() {
	fmt.Println("Status Parkir:")
	for i := 0; i < jumlahSlotMaks; i++ {
		if slotParkir[i].PlatNomor == "" {
			fmt.Printf("Slot %d: Kosong\n", i+1)
		} else {
			fmt.Printf("Slot %d: %s (%s)\n", i+1, slotParkir[i].PlatNomor, slotParkir[i].Jenis)
		}
	}
}

func hitungBiaya(waktuMasuk, waktuKeluar time.Time, jenis string) int {
	const tarifAwalMotor = 2000
	const tarifAwalMobil = 5000
	const tarifPerJam = 1000

	durasi := waktuKeluar.Sub(waktuMasuk)
	jam := int(durasi.Hours())
	if durasi.Minutes() > float64(jam*60) {
		jam++ // Menambah satu jam jika ada sisa menit
	}

	var biaya int
	if jenis == "mobil" {
		biaya = tarifAwalMobil
		if jam > 1 {
			biaya += (jam - 1) * tarifPerJam
		}
	} else if jenis == "motor" {
		biaya = tarifAwalMotor
		if jam > 1 {
			biaya += (jam - 1) * tarifPerJam
		}
	}
	return biaya
}

func urutkanSlotParkirDescending() {
	sort.Slice(slotParkir[:], func(i, j int) bool {
		return slotParkir[i].PlatNomor > slotParkir[j].PlatNomor
	})
	fmt.Println("Slot Parkir berhasil diurutkan secara Descending.")
	tampilkanStatusParkir()
}

func urutkanSlotParkirAscending() {
	sort.Slice(slotParkir[:], func(i, j int) bool {
		return slotParkir[i].PlatNomor < slotParkir[j].PlatNomor
	})
	fmt.Println("Slot Parkir berhasil diurutkan secara Ascending.")
	tampilkanStatusParkir()
}

func cariSlotKosong() int {
	for i := 0; i < jumlahSlotMaks; i++ {
		if slotParkir[i].PlatNomor == "" {
			return i
		}
	}
	return -1
}

func tambahDataParkir() {
	slotKosong := cariSlotKosong()
	if slotKosong == -1 {
		fmt.Println("Parkir penuh! Tidak ada slot kosong.")
		return
	}

	var platNomor string
	var jenis string
	fmt.Print("Masukkan nomor plat kendaraan: ")
	fmt.Scan(&platNomor)
	fmt.Print("Jenis Kendaraan (mobil/motor): ")
	fmt.Scan(&jenis)
	kendaraan := Kendaraan{
		PlatNomor:  platNomor,
		Jenis:      jenis,
		WaktuMasuk: time.Now(),
	}
	slotParkir[slotKosong] = kendaraan
	fmt.Printf("Kendaraan dengan nomor plat %s diparkir di slot %d\n", platNomor, slotKosong+1)
}

func hapusDataParkir() {
	var platNomor string
	fmt.Print("Masukkan nomor plat kendaraan yang ingin dihapus: ")
	fmt.Scan(&platNomor)

	found := false
	for i := 0; i < jumlahSlotMaks; i++ {
		if slotParkir[i].PlatNomor == platNomor {
			found = true
			fmt.Printf("Kendaraan dengan nomor plat %s berhasil dihapus dari slot %d.\n", slotParkir[i].PlatNomor, i+1)
			slotParkir[i] = Kendaraan{}
			break
		}
	}
	if !found {
		fmt.Println("Kendaraan dengan nomor plat tersebut tidak ditemukan.")
	}
}

func ubahDataParkir() {
	var platNomor string
	fmt.Print("Masukkan nomor plat kendaraan yang ingin diubah: ")
	fmt.Scan(&platNomor)

	found := false
	for i := 0; i < jumlahSlotMaks; i++ {
		if slotParkir[i].PlatNomor == platNomor {
			found = true
			fmt.Printf("Kendaraan dengan nomor plat %s ditemukan di slot %d.\n", slotParkir[i].PlatNomor, i+1)
			fmt.Println("Masukkan data baru:")
			var platNomorBaru string
			var jenisBaru string
			fmt.Print("Nomor Plat Kendaraan baru: ")
			fmt.Scan(&platNomorBaru)
			fmt.Print("Jenis Kendaraan baru (mobil/motor): ")
			fmt.Scan(&jenisBaru)
			slotParkir[i].PlatNomor = platNomorBaru
			slotParkir[i].Jenis = jenisBaru
			fmt.Printf("Data kendaraan dengan nomor plat %s berhasil diubah menjadi nomor plat %s dan jenis %s.\n", platNomor, platNomorBaru, jenisBaru)
			break
		}
	}
	if !found {
		fmt.Println("Kendaraan dengan nomor plat tersebut tidak ditemukan.")
	}
}

func binarySearchPlatNomor() {
	var platNomor string
	fmt.Print("Masukkan nomor plat kendaraan yang ingin dicari: ")
	fmt.Scan(&platNomor)

	urutkanSlotParkirAscending()

	index := binarySearch(platNomor)
	if index == -1 {
		fmt.Println("Kendaraan dengan nomor plat tersebut tidak ditemukan.")
	} else {
		fmt.Printf("Kendaraan dengan nomor plat %s ditemukan di slot %d.\n", slotParkir[index].PlatNomor, index+1)
	}
}

func binarySearch(platNomor string) int {
	low, high := 0, jumlahSlotMaks-1
	for low <= high {
		mid := (low + high) / 2
		if slotParkir[mid].PlatNomor == platNomor {
			return mid
		} else if slotParkir[mid].PlatNomor < platNomor {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
