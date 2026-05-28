package main

import "fmt"

const NMAX int = 1000

type dataPemilik struct {
	nama   string
	kontak string
}

type dataKendaraan struct {
	jenisKendaraan string
	plat           string
	tahunProduksi  int
	tanggal_Servis int
	bulan_Servis   int
	tahun_Servis   int
	jenisKerusakan string
	detailServis   string
}

type tabPemilik [NMAX]dataPemilik
type tabKendaraan [NMAX]dataKendaraan

func tambahData(pemilik *tabPemilik, kendaraan *tabKendaraan, n *int) {
	fmt.Print("Nama Pemilik: ")
	fmt.Scan(&pemilik[*n].nama)
	fmt.Print("Kontak Pemilik: ")
	fmt.Scan(&pemilik[*n].kontak)
	fmt.Print("Jenis Kendaraan: ")
	fmt.Scan(&kendaraan[*n].jenisKendaraan)
	fmt.Print("Plat Kendaraan: ")
	fmt.Scan(&kendaraan[*n].plat)
	fmt.Print("Tahun Produksi: ")
	fmt.Scan(&kendaraan[*n].tahunProduksi)
	fmt.Print("Tanggal Servis (DD): ")
	fmt.Scan(&kendaraan[*n].tanggal_Servis)
	fmt.Print("Bulan Servis (MM): ")
	fmt.Scan(&kendaraan[*n].bulan_Servis)
	fmt.Print("Tahun Servis (YYYY): ")
	fmt.Scan(&kendaraan[*n].tahun_Servis)
	*n = *n + 1
}

func sequentialSearch(kendaraan *tabKendaraan, n int, plat string) int {
	var found int = -1
	var i int = 0
	for i <= n-1 && found == -1 {
		if plat == kendaraan[i].plat {
			found = i
		}
		i = i + 1
	}
	return found
}

func binarySearchplat(kendaraan *tabKendaraan, n int, plat string) int {
	var found int = -1
	var left int = 0
	var right int = n - 1
	var mid int
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if plat == kendaraan[mid].plat {
			found = mid
		} else if plat < kendaraan[mid].plat {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return found
}

func SelectionSortAscPlat(dK *tabKendaraan, n int) {
	var pass, i, acuan int
	var temp dataKendaraan
	for pass = 1; pass < n; pass++ {
		acuan = pass - 1
		for i = pass; i < n; i++ {
			if dK[acuan].plat > dK[i].plat {
				acuan = i
			}
		}
		temp = dK[acuan]
		dK[acuan] = dK[pass-1]
		dK[pass-1] = temp
	}
}

func SelectionSortDescPlat(dK *tabKendaraan, n int) {
	var pass, i, acuan int
	var temp dataKendaraan
	for pass = 1; pass < n; pass++ {
		acuan = pass - 1
		for i = pass; i < n; i++ {
			if dK[acuan].plat < dK[i].plat {
				acuan = i
			}
		}
		temp = dK[acuan]
		dK[acuan] = dK[pass-1]
		dK[pass-1] = temp
	}
}

func cetakData(dK tabKendaraan, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %d\n", dK[i].jenisKendaraan, dK[i].plat, dK[i].tahunProduksi)
	}
}

func main() {
	var pemilik tabPemilik
	var kendaraan tabKendaraan
	var n int = 0
	var pilihan int
	switch pilihan {
	case 1:
		tambahData(&pemilik, &kendaraan, &n)
	case 2:
		optionData3(&kendaraan, n)
	}
}

func optionData3(kendaraan *tabKendaraan, n int) {
	var angka int
	fmt.Println("Apakah data plat ingin diurutkan?")
	fmt.Println("1. YA")
	fmt.Println("2. TIDAK")
	fmt.Scan(&angka)
	switch angka {
	case 1:
		var urutan int
		fmt.Println("Urutan data plat:")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Scan(&urutan)
		switch urutan {
		case 1:
			var plat string
			fmt.Print("Masukkan plat yang ingin dicari: ")
			fmt.Scan(&plat)
			SelectionSortAscPlat(&kendaraan, n)
			binarySearchplat(&kendaraan, n, plat)
			cetakData(kendaraan, n)
		case 2:
			var plat string
			fmt.Print("Masukkan plat yang ingin dicari: ")
			fmt.Scan(&plat)
			SelectionSortDescPlat(&kendaraan, n)
			binarySearchplat(&kendaraan, n, plat)
			cetakData(kendaraan, n)
		}
	case 2:
		var plat string
		fmt.Print("Masukkan plat yang ingin dicari: ")
		fmt.Scan(&plat)
		sequentialSearch(&kendaraan, n, plat)
		cetakData(kendaraan, n)
	}
}
