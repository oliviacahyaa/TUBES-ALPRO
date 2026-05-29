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
	tanggal, bulan, tahun int
	jenisKerusakan        int
	detailServis          string
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
		optionData3(pemilik, kendaraan, servis, *nData, *nServis)
	case 4:
		optionData4(pemilik, kendaraan, *nData)
	case 5:
		optionData5(servis, *nServis)
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
	fmt.Println()

	switch pilih {
	case 1:
		fmt.Println("ADD")
		var next string = "YES"
		for next == "YES" && *nData < NMAX {
			fmt.Printf("%03d.\n", *nData+1)
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
		fmt.Println("EDIT DATA")
		var targetPlat string
		var idx int = -1
		var pilih int
		var next string = "YES"
		if *nData == 0 {
			fmt.Println("Data kosong. Harap tambahkan data terlebih dahulu.")
		} else {
			fmt.Print("Masukkan nomor plat kendaraan yang ingin diedit: ")
			fmt.Scan(&targetPlat)
			idx = sequentialSearch(A, *nData, targetPlat)
			if idx != -1 {
				fmt.Printf("Data kendaraan dengan nomor plat %s ditemukan!\n", targetPlat)
				fmt.Printf("%03d.\n", idx+1)
				fmt.Printf("NAMA: %s\n", T[idx].nama)
				fmt.Printf("NO.KONTAK: %s\n", T[idx].kontak)
				fmt.Printf("JENIS KENDARAAN: %s\n", A[idx].jenisKendaraan)
				fmt.Printf("NOMOR PLAT: %s\n", A[idx].plat)
				fmt.Printf("TAHUN PRODUKSI: %d\n", A[idx].tahunProduksi)
				for next == "YES" {
					fmt.Println("Pilih data yang ingin diedit:")
					fmt.Println("1. NAMA")
					fmt.Println("2. NO.KONTAK")
					fmt.Println("3. JENIS KENDARAAN")
					fmt.Println("4. NOMOR PLAT")
					fmt.Println("5. TAHUN PRODUKSI")
					fmt.Print("PILIH: ")
					fmt.Scan(&pilih)
					switch pilih {
					case 1:
						fmt.Println("Data nama saat ini: ", T[idx].nama)
						fmt.Print("Masukkan nama baru: ")
						fmt.Scan(&T[idx].nama)
					case 2:
						fmt.Println("Data nomor kontak saat ini: ", T[idx].kontak)
						fmt.Print("Masukkan nomor kontak baru: ")
						fmt.Scan(&T[idx].kontak)
					case 3:
						fmt.Println("Data jenis kendaraan saat ini: ", A[idx].jenisKendaraan)
						fmt.Print("Masukkan jenis kendaraan baru: ")
						fmt.Scan(&A[idx].jenisKendaraan)
					case 4:
						fmt.Println("Data nomor plat saat ini: ", A[idx].plat)
						fmt.Print("Masukkan nomor plat baru: ")
						fmt.Scan(&A[idx].plat)
					case 5:
						fmt.Println("Data tahun produksi saat ini: ", A[idx].tahunProduksi)
						fmt.Print("Masukkan tahun produksi baru: ")
						fmt.Scan(&A[idx].tahunProduksi)
					}
				}
			}
		}
	case 3:
		fmt.Println("DELETE")
		var targetPlat string
		var idx int = -1
		var pilih int
		var next string = "YES"
		if *nData == 0 {
			fmt.Println("Data kosong. Harap tambahkan data terlebih dahulu.")
		} else {
			fmt.Print("Masukkan nomor plat kendaraan yang ingin dihapus: ")
			fmt.Scan(&targetPlat)
			idx = sequentialSearch(A, *nData, targetPlat)
			if idx != -1 {
				fmt.Printf("Data kendaraan dengan nomor plat %s ditemukan!\n", targetPlat)
				fmt.Printf("%03d.\n", idx+1)
				fmt.Printf("NAMA: %s\n", T[idx].nama)
				fmt.Printf("NO.KONTAK: %s\n", T[idx].kontak)
				fmt.Printf("JENIS KENDARAAN: %s\n", A[idx].jenisKendaraan)
				fmt.Printf("NOMOR PLAT: %s\n", A[idx].plat)
				fmt.Printf("TAHUN PRODUKSI: %d\n", A[idx].tahunProduksi)
				fmt.Print("Apakah Anda yakin ingin menghapus data ini? (YES/NO): ")
				fmt.Scan(&next)
				if next == "YES" {
					for i := idx; i < *nData-1; i++ {
						T[i] = T[i+1]
						A[i] = A[i+1]
					}
					*nData = *nData - 1
					fmt.Println("Data berhasil dihapus!")
				} else {
					fmt.Println("Penghapusan data dibatalkan.")
				}
			}
		}

	}
}

func optionData2(T *tabPemilik, A *tabKendaraan, B *tabRiwayat, nData int, nServis *int) {
	var targetPlat string
	var idx int = -1
	var next string = "YES"
	fmt.Print("MASUKKAN NOMOR PLAT KENDARAAN : ")
	fmt.Scan(&targetPlat)
	idx = sequentialSearch(A, nData, targetPlat)
	if idx != -1 {
		fmt.Println("DATA PELANGGAN DITEMUKAN!")
		fmt.Printf("%03d.\n", idx+1)
		fmt.Printf("NAMA: %s\n", T[idx].nama)
		fmt.Printf("NO.KONTAK: %s\n", T[idx].kontak)
		fmt.Printf("JENIS KENDARAAN: %s\n", A[idx].jenisKendaraan)
		fmt.Printf("NOMOR PLAT: %s\n", A[idx].plat)
		fmt.Printf("TAHUN PRODUKSI: %d\n", A[idx].tahunProduksi)
		fmt.Printf("%03d.\n", *nServis+1)
		for next == "YES" && *nServis < NMAX {
			fmt.Print("TANGGAL SERVIS (TANGGAL BULAN TAHUN)(DD MM YYYY): ")
			fmt.Scan(&B[*nServis].tanggal, &B[*nServis].bulan, &B[*nServis].tahun)
			fmt.Print("OPSI JENIS KERUSAKAN: ")
			fmt.Println("1. Servis_Berkala")
			fmt.Println("2. Mesin")
			fmt.Println("3. Kelistrikan")
			fmt.Println("4. Body")
			fmt.Println("5. Ban")
			fmt.Println("6. Kaki-kaki")
			fmt.Print("PILIH JENIS KERUSAKAN: ")
			fmt.Scan(&B[*nServis].jenisKerusakan)
			fmt.Print("DETAIL SERVIS: ")
			fmt.Scan(&B[*nServis].detailServis)
			*nServis = *nServis + 1
			fmt.Print("Apakah ingin tambah data lagi? (YES/NO): ")
			fmt.Scan(&next)
		}
	}

}

func optionData3(T *tabPemilik, A *tabKendaraan, B *tabRiwayat, nData int, nServis int) {
	var pilih string
	var pilihSorting int
	fmt.Println("SEARCH KENDARAAN")
	fmt.Println("Apakah data ingin disorting terlebih dahulu? (YES/NO): ")
	fmt.Scan(&pilih)
	if pilih == "YES" {
		fmt.Print("Data ingin disorting secara:")
		fmt.Println("1. Ascending (terkecil -> terbesar)")
		fmt.Println("2. Sorting Descending  (terbesar -> terkecil)")
		fmt.Print("PILIH JENIS SORTING : ")
		fmt.Scan(&pilihSorting)
		switch pilihSorting {
		case 1:
			var plat string
			fmt.Print("Masukkan plat yang ingin dicari: ")
			fmt.Scan(&plat)
			SelectionSortAscPlat(A, nData)
			binarySearchplat(A, nData, plat)
			cetakData(*A, nData)
		case 2:
			var plat string
			fmt.Print("Masukkan plat yang ingin dicari: ")
			fmt.Scan(&plat)
			SelectionSortDescPlat(A, nData)
			binarySearchplat(A, nData, plat)
			cetakData(*A, nData)
		}
	} else {
		var plat string
		fmt.Print("Masukkan plat yang ingin dicari: ")
		fmt.Scan(&plat)
		sequentialSearch(A, nData, plat)
		cetakData(*A, nData)

	}
}

func optionData4(T *tabPemilik, A *tabKendaraan, nData int) {
	var pilih int
	var pilihSorting int
	var pilihUrutan int
	fmt.Println("SORTING DAFTAR KENDARAAN")
	fmt.Println("Silahkan pilih pengurutan data berdasarkan:")
	fmt.Println("1. PLAT")
	fmt.Println("2. TAHUN PRODUKSI")
	fmt.Print("PILIH:")
	fmt.Scan(&pilih)

	fmt.Println("Silahkan pilih jenis pengurutan data secara:")
	fmt.Println("1. Selection")
	fmt.Println("2. Insertion")
	fmt.Print("PILIH JENIS SORTING : ")
	fmt.Scan(&pilihSorting)

	fmt.Println("Silahkan pilih pengurutan data secara:")
	fmt.Println("1. Ascending (terkecil -> terbesar)")
	fmt.Println("2. Sorting Descending  (terbesar -> terkecil)")
	fmt.Print("PILIH JENIS SORTING : ")
	fmt.Scan(&pilihUrutan)

	switch pilih {
	case 1:
		switch pilihSorting {
		case 1:
			switch pilihUrutan {
			case 1:
				SelectionSortAscPlat(A, nData)
				cetakData(*A, nData)
			case 2:
				SelectionSortDescPlat(A, nData)
				cetakData(*A, nData)
			}
		case 2:
			switch pilihUrutan {
			case 1:
				InsertionSortAscTahun(A, nData)
				cetakData(*A, nData)
			case 2:
				InsertionSortDescTahun(A, nData)
				cetakData(*A, nData)
			}
		}
	}

}

func optionsData5(B *tabRiwayat, nServis int) {
	var pilih int
	fmt.Println("STATISTIK SERVIS")
	fmt.Println("1. Statistik Jenis Kerusakan")
	fmt.Println("2. Statistik Jumlah Servis per Bulan")
	fmt.Print("PILIH MENU: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("Statistik Jenis Kerusakan")
		var pilihan [6]int
		var i int
		for i = 0; i < nServis; i++ {
			switch B[i].jenisKerusakan {
			case 1:
				pilihan[0]++
			case 2:
				pilihan[1]++
			case 3:
				pilihan[2]++
			case 4:
				pilihan[3]++
			case 5:
				pilihan[4]++
			case 6:
				pilihan[5]++
			}
		}
		fmt.Println("HASIL STATISTIK JENIS KERUSAKAN:")
		for i = 0; i < 6; i++ {
			fmt.Printf("Jenis kerusakan %d: %d\n", i+1, pilihan[i])
		}
		fmt.Print("Jenis kerusakan yang sering terjadi: ")
		var max = pilihan[0]
		for i = 1; i < 6; i++ {
			if pilihan[i] > max {
				max = pilihan[i]
			}
		}
		fmt.Println(max)

		fmt.Print("Jenis kerusajan yang jarang terjadi: ")
		var min = pilihan[0]
		for i = 1; i < 6; i++ {
			if pilihan[i] < min {
				min = pilihan[i]
			}
		}
		fmt.Println(min)
	case 2:
		fmt.Println("Statistik Jumlah Servis per Bulan")
		var bulan [12]int
		var i int
		for i = 0; i < nServis; i++ {
			switch B[i].bulan {
			case 1:
				bulan[0]++
			case 2:
				bulan[1]++
			case 3:
				bulan[2]++
			case 4:
				bulan[3]++
			case 5:
				bulan[4]++
			case 6:
				bulan[5]++
			case 7:
				bulan[6]++
			case 8:
				bulan[7]++
			case 9:
				bulan[8]++
			case 10:
				bulan[9]++
			case 11:
				bulan[10]++
			case 12:
				bulan[11]++
			}
			fmt.Println("HASIL STATISTIK JUMLAH SERVIS PER BULAN:")
			for i = 0; i < 12; i++ {
				fmt.Printf("Bulan %d: %d\n", i+1, bulan[i])
			}
			fmt.Print("Bulan dengan jumlah servis terbanyak: ")
			var max = bulan[0]
			for i = 1; i < 12; i++ {
				if bulan[i] > max {
					max = bulan[i]
				}
			}
			fmt.Println(max)
			fmt.Print("Bulan dengan jumlah servis paling sedikit: ")
			var min = bulan[0]
			for i = 1; i < 12; i++ {
				if bulan[i] < min {
					min = bulan[i]
				}
			}
			fmt.Println(min)
		}
	}
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

func InsertionSortDescTahun(dK *tabKendaraan, n int) {
	var pass, k int
	var temp dataKendaraan

	for pass = 1; pass < n; pass++ {
		k = pass
		temp = dK[k]
		for k > 0 && temp.tahunProduksi > dK[k-1].tahunProduksi {
			dK[k] = dK[k-1]
			k = k - 1
		}
		dK[k] = temp
	}
}

func InsertionSortAscTahun(dK *tabKendaraan, n int) {
	var pass, k int
	var temp dataKendaraan

	for pass = 1; pass < n; pass++ {
		k = pass
		temp = dK[k]
		for k > 0 && temp.tahunProduksi < dK[k-1].tahunProduksi {
			dK[k] = dK[k-1]
			k = k - 1
		}
		dK[k] = temp
	}
}

func cetakData(dK tabKendaraan, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %d\n", dK[i].jenisKendaraan, dK[i].plat, dK[i].tahunProduksi)
	}
}
