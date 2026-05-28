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
}

type riwayatServis struct {
	tanggalServis  string
	jenisKerusakan string
	detailServis   string
}

type tabPemilik [NMAX]dataPemilik
type tabKendaraan [NMAX]dataKendaraan
type tabRiwayat [NMAX]riwayatServis

func main() {
	var daftarPemilik tabPemilik
	var daftarKendaraan tabKendaraan
	var daftarRiwayat tabRiwayat
	var nData int = 0
	var nServis int = 0
	var keluar bool = true
	for keluar == true {
		keluar = mainMenu(&daftarPemilik, &daftarKendaraan, &daftarRiwayat, &nData, &nServis)
	}
}
func mainMenuUI() {
	fmt.Println("AUTOCARE")
	fmt.Println("Aplikasi Manajemen dan Riwayat Servis Kendaraan")
	fmt.Println("1. Data Kendaraan & Pemilik")
	fmt.Println("2. Tambah Riwayat Servis Baru")
	fmt.Println("3. Search Kendaraan")
	fmt.Println("4. Sorting Daftar Kendaraan")
	fmt.Println("5. Statistik Servis")
	fmt.Println("6. Keluar")
	fmt.Print("PILIH MENU: ")
}

func mainMenu(pemilik *tabPemilik, kendaraan *tabKendaraan, servis *tabRiwayat, nData *int, nServis *int) bool {
	var angka int
	mainMenuUI()
	fmt.Scan(&angka)
	switch angka {
	case 1:
		optionData1(pemilik, kendaraan, nData)
	case 2:
		optionData2(pemilik, kendaraan, servis, *nData, nServis)
	case 3:
		optionData3(kendaraan, *nData)
	case 4:
	case 5:
	case 6:
		return false
	}
	return true
}

func optionData1(T *tabPemilik, A *tabKendaraan, nData *int) {
	var pilih int
	fmt.Println("DATA KENDARAAN & PEMILIK")
	fmt.Println("1. ADD")
	fmt.Println("2. EDIT")
	fmt.Println("3. DELETE")
	fmt.Print("PILIH MENU: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("ADD")
		var next string = "YES"
		fmt.Printf("%03d.\n", *nData+1)
		for next == "YES" && *nData < NMAX {
			fmt.Print("NAMA: ")
			fmt.Scan(&T[*nData].nama)
			fmt.Print("NO.KONTAK: ")
			fmt.Scan(&T[*nData].kontak)
			fmt.Print("JENIS KENDARAAN: ")
			fmt.Scan(&A[*nData].jenisKendaraan)
			fmt.Print("NOMOR PLAT: ")
			fmt.Scan(&A[*nData].plat)
			fmt.Print("TAHUN PRODUKSI: ")
			fmt.Scan(&A[*nData].tahunProduksi)
			*nData = *nData + 1
			fmt.Println("Data berhasil disimpan!")
			fmt.Print("Apakah ingin tambah data lagi? (YES/NO): ")
			fmt.Scan(&next)
		}
	case 2:
		fmt.Println("EDIT")
	case 3:
		fmt.Println("DELETE")

	}
}

func optionData2(T *tabPemilik, A *tabKendaraan, B *tabRiwayat, nData int, nServis *int) {
	var targetPlat string
	var idx int = -1
	fmt.Print("MASUKKAN NOMOR PLAT KENDARAAN : ")
	fmt.Scan(&targetPlat)
	cariPlat(A, nData, targetPlat, &idx)
	if idx != -1 {
		fmt.Println("DATA PELANGGAN DITEMUKAN!")
		fmt.Printf("%03d.\n", idx+1)
		fmt.Printf("NAMA: %s\n", T[idx].nama)
		fmt.Printf("NO.KONTAK: %s\n", T[idx].kontak)
		fmt.Printf("JENIS KENDARAAN: %s\n", A[idx].jenisKendaraan)
		fmt.Printf("NOMOR PLAT: %s\n", A[idx].plat)
		fmt.Printf("TAHUN PRODUKSI: %d\n", A[idx].tahunProduksi)
		fmt.Print("TANGGAL SERVIS: ")
		fmt.Scan(&B[*nServis].tanggalServis)
		fmt.Printf("%03d.\n", *nServis+1)
		for B[*nServis].tanggalServis != "SELESAI" && *nServis < NMAX {
			fmt.Print("JENIS KERUSAKAN: ")
			fmt.Scan(&B[*nServis].jenisKerusakan)
			fmt.Print("DETAIL SERVIS: ")
			fmt.Scan(&B[*nServis].detailServis)
			*nServis = *nServis + 1
			fmt.Scan(&B[*nServis].tanggalServis)
		}
	}

}

func cariPlat(A *tabKendaraan, n int, targetPlat string, idx *int) int {
	var i int
	for i = 0; i < n; i++ {
		if A[i].plat == targetPlat {
			*idx = i
			return i
		}
	}
	*idx = -1
	return -1
}
