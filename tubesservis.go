package main

import "fmt"

const NMAX int = 999

type dataPemilik struct {
	nama   string
	kontak string
}
type dataKendaraan struct {
	jenisKendaraan string
	plat           string
	tahunProduksi  int
	merk           string
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
	fmt.Println("APLIKASI MANAJEMEN DAN RIWAYAT SERVIS KENDARAAN")
	fmt.Println("[1] DATA KENDARAAN & PEMILIK")
	fmt.Println("[2] TAMBAH RIWAYAT SERVIS")
	fmt.Println("[3] SEARCH KENDARAAN")
	fmt.Println("[4] SORTING DAFTAR KENDARAAN")
	fmt.Println("[5] STATISTIK SERVIS")
	fmt.Println("[6] TAMPILKAN SEMUA DATA")
	fmt.Println("[7] KELUAR")
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
		optionData4(kendaraan, *nData)
	case 5:
		optionData5(servis, *nServis)
	case 6:
		optionData6(pemilik, kendaraan, servis, *nData, *nServis)
	case 7:
		return false
	}
	return true
}

func optionData1(T *tabPemilik, A *tabKendaraan, nData *int) {
	var pilih int
	fmt.Println()
	fmt.Println("DATA KENDARAAN & PEMILIK")
	fmt.Println("[1] ADD")
	fmt.Println("[2] EDIT")
	fmt.Println("[3] DELETE")
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
			fmt.Print("NOMOR TELEPON: ")
			fmt.Scan(&T[*nData].kontak)
			fmt.Print("JENIS KENDARAAN (MOBIL/MOTOR): ")
			fmt.Scan(&A[*nData].jenisKendaraan)
			fmt.Print("MERK KENDARAAN: ")
			fmt.Scan(&A[*nData].merk)
			fmt.Print("NOMOR PLAT: ")
			fmt.Scan(&A[*nData].plat)
			fmt.Print("TAHUN PRODUKSI: ")
			fmt.Scan(&A[*nData].tahunProduksi)
			*nData = *nData + 1
			fmt.Println("\nDATA BERHASIL DISIMPAN!")
			fmt.Print("APAKAH ADA DATA SERVIS LAIN YANG INGIN DITAMBAHKAN? (YES/NO): ")
			fmt.Scan(&next)
			fmt.Println()
		}
	case 2:
		fmt.Println("EDIT DATA")
		var targetPlat string
		var idx int = -1
		var pilih int
		var next string = "YES"
		if *nData == 0 {
			fmt.Println("DATA KOSONG. HARAP TAMBAHKAN DATA TERLEBIH DAHULU PADA MENU '[1] ADD'")
		} else {
			fmt.Print("MASUKKAN NOMOR PLAT KENDARAAN YANG INGIN DIEDIT: ")
			fmt.Scan(&targetPlat)
			idx = sequentialSearch(A, *nData, targetPlat)
			if idx != -1 {
				fmt.Printf("DATA KENDARAAN DENGAN NOMOR PLAT %s DITEMUKAN!\n", targetPlat)
				fmt.Printf("%03d.\n", idx+1)
				fmt.Printf("NAMA: %s\n", T[idx].nama)
				fmt.Printf("NOMOR TELEPON: %s\n", T[idx].kontak)
				fmt.Printf("JENIS KENDARAAN: %s\n", A[idx].jenisKendaraan)
				fmt.Printf("MERK KENDARAAN: %s\n", A[idx].merk)
				fmt.Printf("NOMOR PLAT: %s\n", A[idx].plat)
				fmt.Printf("TAHUN PRODUKSI: %d\n", A[idx].tahunProduksi)
				for next == "YES" {
					fmt.Println("PILIH DATA YANG INGIN DIEDIT:")
					fmt.Println("[1] NAMA")
					fmt.Println("[2] NOMOR TELEPON")
					fmt.Println("[3] JENIS KENDARAAN")
					fmt.Println("[4] MERK KENDARAAN")
					fmt.Println("[5] NOMOR PLAT")
					fmt.Println("[6] TAHUN PRODUKSI")
					fmt.Print("PILIH: ")
					fmt.Scan(&pilih)
					fmt.Println()
					switch pilih {
					case 1:
						fmt.Println("DATA NAMA SAAT INI: ", T[idx].nama)
						fmt.Print("MASUKKAN NAMA BARU: ")
						fmt.Scan(&T[idx].nama)
					case 2:
						fmt.Println("DATA NOMOR TELEPON SAAT INI: ", T[idx].kontak)
						fmt.Print("MASUKKAN NOMOR TELEPON BARU: ")
						fmt.Scan(&T[idx].kontak)
					case 3:
						fmt.Println("DATA JENIS KENDARAAN SAAT INI: ", A[idx].jenisKendaraan)
						fmt.Print("MASUKKAN JENIS KENDARAAN BARU (MOBIL/MOTOR): ")
						fmt.Scan(&A[idx].jenisKendaraan)
					case 4:
						fmt.Println("DATA MERK KENDARAAN SAAT INI: ", A[idx].merk)
						fmt.Print("MASUKKAN MERK KENDARAAN BARU: ")
						fmt.Scan(&A[idx].merk)
					case 5:
						fmt.Println("DATA NOMOR PLAT SAAT INI: ", A[idx].plat)
						fmt.Print("MASUKKAN NOMOR PLAT BARU: ")
						fmt.Scan(&A[idx].plat)
					case 6:
						fmt.Println("DATA TAHUN PRODUKSI SAAT INI: ", A[idx].tahunProduksi)
						fmt.Print("MASUKKAN TAHUN PRODUKSI BARU: ")
						fmt.Scan(&A[idx].tahunProduksi)
					}
					fmt.Println("\nDATA BERHASIL DIUPDATE!")
					fmt.Print("APAKAH ADA DATA LAIN YANG INGIN DIEDIT? (YES/NO): ")
					fmt.Scan(&next)
				}
			} else {
				fmt.Printf("DATA DENGAN NOMOR PLAT %s TIDAK DITEMUKAN!, HARAP MASUKKAN TERLEBIH DAHULU PADA MENU '[1] ADD'\n", targetPlat)
			}
		}
	case 3:
		fmt.Println("DELETE")
		var targetPlat string
		var idx int = -1
		var next string = "YES"
		if *nData == 0 {
			fmt.Println("DATA KOSONG. HARAP TAMBAHKAN DATA TERLEBIH DAHULU PADA MENU '[1] ADD'")
		} else {
			fmt.Print("MASUKKAN NOMOR PLAT KENDARAAN YANG INGIN DIHAPUS: ")
			fmt.Scan(&targetPlat)
			idx = sequentialSearch(A, *nData, targetPlat)
			if idx != -1 {
				fmt.Printf("DATA KENDARAAN DENGAN NOMOR PLAT %s DITEMUKAN!\n", targetPlat)
				fmt.Printf("%03d.\n", idx+1)
				fmt.Printf("NAMA: %s\n", T[idx].nama)
				fmt.Printf("NO.KONTAK: %s\n", T[idx].kontak)
				fmt.Printf("JENIS KENDARAAN: %s\n", A[idx].jenisKendaraan)
				fmt.Printf("MERK KENDARAAN: %s\n", A[idx].merk)
				fmt.Printf("NOMOR PLAT: %s\n", A[idx].plat)
				fmt.Printf("TAHUN PRODUKSI: %d\n", A[idx].tahunProduksi)
				fmt.Print("APAKAH ANDA YAKIN INGIN MENGHAPUS DATA INI? (YES/NO): ")
				fmt.Scan(&next)
				fmt.Println()
				if next == "YES" {
					for i := idx; i < *nData-1; i++ {
						T[i] = T[i+1]
						A[i] = A[i+1]
					}
					*nData = *nData - 1
					fmt.Println("DATA BERHASIL DIHAPUS!")
				} else {
					fmt.Println("PENGHAPUSAN DATA DIBATALKAN!")
				}
			} else {
				fmt.Printf("DATA DENGAN NOMOR PLAT %s TIDAK DITEMUKAN!, HARAP MASUKKAN TERLEBIH DAHULU PADA MENU '[1] ADD'\n", targetPlat)
			}
		}

	}
}

func optionData2(T *tabPemilik, A *tabKendaraan, B *tabRiwayat, nData int, nServis *int) {
	var targetPlat string
	var idx int = -1
	var next string = "YES"
	fmt.Println()
	fmt.Println("TAMBAH RIWAYAT SERVIS")
	fmt.Print("MASUKKAN NOMOR PLAT KENDARAAN : ")
	fmt.Scan(&targetPlat)
	idx = sequentialSearch(A, nData, targetPlat)
	if idx != -1 {
		fmt.Printf("\nDATA PELANGGAN DITEMUKAN!\n")
		fmt.Printf("%03d.\n", idx+1)
		fmt.Printf("NAMA: %s\n", T[idx].nama)
		fmt.Printf("NO.KONTAK: %s\n", T[idx].kontak)
		fmt.Printf("JENIS KENDARAAN: %s\n", A[idx].jenisKendaraan)
		fmt.Printf("NOMOR PLAT: %s\n", A[idx].plat)
		fmt.Printf("TAHUN PRODUKSI: %d\n", A[idx].tahunProduksi)
		fmt.Println("\nSILAHKAN MASUKKAN DATA SERVIS KENDARAAN: ")
		fmt.Printf("%03d.\n", *nServis+1)
		for next == "YES" && *nServis < NMAX {
			fmt.Print("TANGGAL SERVIS (TANGGAL BULAN TAHUN)(DD MM YYYY): ")
			fmt.Scan(&B[*nServis].tanggal, &B[*nServis].bulan, &B[*nServis].tahun)
			fmt.Println("OPSI JENIS KERUSAKAN: ")
			fmt.Println("[1] SERVIS_BERKALA")
			fmt.Println("[2] MESIN")
			fmt.Println("[3] KELISTRIKAN")
			fmt.Println("[4] BODY")
			fmt.Println("[5] BAN")
			fmt.Println("[6] KAKI-KAKI")
			fmt.Print("PILIH JENIS KERUSAKAN: ")
			fmt.Scan(&B[*nServis].jenisKerusakan)
			fmt.Print("DETAIL SERVIS: ")
			fmt.Scan(&B[*nServis].detailServis)
			*nServis = *nServis + 1
			fmt.Print("APAKAH ADA DATA SERVIS LAIN YANG INGIN DITAMBAHKAN? (YES/NO): ")
			fmt.Scan(&next)
			fmt.Println()
		}
	} else {
		fmt.Println("DATA PELANGGAN TIDAK DITEMUKAN!")
	}
}

func optionData3(T *tabPemilik, A *tabKendaraan, B *tabRiwayat, nData int, nServis int) {
	var pilih string
	var pilihSorting int
	fmt.Println("SEARCH KENDARAAN (BERDASARKAN NOMOR PLAT)")
	fmt.Println("APAKAH ANDA INGIN MELAKUKAN PENGURUTAN DATA TERLEBIH DAHULU SEBELUM MELAKUKAN PENCARIAN? (YES/NO): ")
	fmt.Scan(&pilih)
	fmt.Println()
	if pilih == "YES" {
		fmt.Println("DATA AKAN DIURUTKAN SECARA:")
		fmt.Println("[1] ASCENDING (TERKECIL -> TERBESAR)")
		fmt.Println("[2] DESCENDING (TERBESAR -> TERKECIL)")
		fmt.Print("PILIH JENIS PENGURUTAN: ")
		fmt.Scan(&pilihSorting)
		switch pilihSorting {
		case 1:
			var plat string
			var hasil int
			fmt.Print("MASUKKAN PLAT YANG INGIN DICARI: ")
			fmt.Scan(&plat)
			fmt.Println()
			SelectionSortAscPlat(A, nData)
			cetakData(*A, nData)
			hasil = binarySearchplat(A, nData, plat)
			if hasil == -1 {
				fmt.Printf("DATA DENGAN NOMOR PLAT %s TIDAK DITEMUKAN!\n", plat)
				fmt.Println()
			} else {
				fmt.Printf("DATA DITEMUKAN!\n")
				fmt.Println("HASIL PENCARIAN: ")
				fmt.Println("\nDATA KENDARAAN:")
				fmt.Printf("JENIS KENDARAAN: %s\n", A[hasil].jenisKendaraan)
				fmt.Printf("NOMOR PLAT: %s\n", A[hasil].plat)
				fmt.Printf("TAHUN PRODUKSI: %d\n", A[hasil].tahunProduksi)
				fmt.Println("\nDATA PEMILIK:")
				fmt.Printf("NAMA: %s\n", T[hasil].nama)
				fmt.Printf("NO.KONTAK: %s\n", T[hasil].kontak)
				fmt.Print("\nRIWAYAT SERVIS: ")
				for i := 0; i < nServis; i++ {
					if A[hasil].plat == A[i].plat {
						fmt.Printf("%03d.\n", i+1)
						fmt.Printf("TANGGAL SERVIS: %02d-%02d-%04d\n", B[i].tanggal, B[i].bulan, B[i].tahun)
						fmt.Printf("JENIS KERUSAKAN: %d\n", B[i].jenisKerusakan)
						fmt.Printf("DETAIL SERVIS: %s\n", B[i].detailServis)
					}
				}
				fmt.Println()
			}

		case 2:
			var plat string
			var hasil int
			fmt.Print("MASUKKAN PLAT YANG INGIN DICARI: ")
			fmt.Scan(&plat)
			fmt.Println()
			SelectionSortDescPlat(A, nData)
			cetakData(*A, nData)
			hasil = binarySearchplat(A, nData, plat)
			if hasil == -1 {
				fmt.Printf("DATA DENGAN NOMOR PLAT %s TIDAK DITEMUKAN!\n", plat)
				fmt.Println()
			} else {
				fmt.Printf("DATA DITEMUKAN!\n")
				fmt.Println("HASIL PENCARIAN: ")
				fmt.Println("\nDATA KENDARAAN:")
				fmt.Printf("JENIS KENDARAAN: %s\n", A[hasil].jenisKendaraan)
				fmt.Printf("NOMOR PLAT: %s\n", A[hasil].plat)
				fmt.Printf("TAHUN PRODUKSI: %d\n", A[hasil].tahunProduksi)
				fmt.Println("\nDATA PEMILIK:")
				fmt.Printf("NAMA: %s\n", T[hasil].nama)
				fmt.Printf("NO.KONTAK: %s\n", T[hasil].kontak)
				fmt.Print("\nRIWAYAT SERVIS: ")
				for i := 0; i < nServis; i++ {
					if A[hasil].plat == A[i].plat {
						fmt.Printf("%03d.\n", i+1)
						fmt.Printf("TANGGAL SERVIS: %02d-%02d-%04d\n", B[i].tanggal, B[i].bulan, B[i].tahun)
						fmt.Printf("JENIS KERUSAKAN: %d\n", B[i].jenisKerusakan)
						fmt.Printf("DETAIL SERVIS: %s\n", B[i].detailServis)
					}
				}
				fmt.Println()
			}
		}
	} else {
		var plat string
		var hasil int
		fmt.Print("MASUKKAN PLAT YANG INGIN DICARI: ")
		fmt.Scan(&plat)
		fmt.Println()
		hasil = sequentialSearch(A, nData, plat)
		if hasil == -1 {
			fmt.Printf("DATA DENGAN NOMOR PLAT %s TIDAK DITEMUKAN!\n", plat)
			fmt.Println()
		} else {
			fmt.Println("HASIL PENCARIAN: ")
			fmt.Println("\nDATA KENDARAAN:")
			fmt.Printf("JENIS KENDARAAN: %s\n", A[hasil].jenisKendaraan)
			fmt.Printf("NOMOR PLAT: %s\n", A[hasil].plat)
			fmt.Printf("TAHUN PRODUKSI: %d\n", A[hasil].tahunProduksi)
			fmt.Println("\nDATA PEMILIK:")
			fmt.Printf("NAMA: %s\n", T[hasil].nama)
			fmt.Printf("NO.KONTAK: %s\n", T[hasil].kontak)
			fmt.Print("\nRIWAYAT SERVIS: ")
			for i := 0; i < nServis; i++ {
				if A[hasil].plat == A[i].plat {
					fmt.Printf("%03d.\n", i+1)
					fmt.Printf("TANGGAL SERVIS: %02d-%02d-%04d\n", B[i].tanggal, B[i].bulan, B[i].tahun)
					fmt.Printf("JENIS KERUSAKAN: %d\n", B[i].jenisKerusakan)
					fmt.Printf("DETAIL SERVIS: %s\n", B[i].detailServis)
				}
			}
			fmt.Println()
		}
	}

}

func optionData4(A *tabKendaraan, nData int) {
	var pilih int
	var pilihSorting int
	var pilihUrutan int
	fmt.Println("SORTING DAFTAR KENDARAAN")
	fmt.Println("SILAHKAN PILIH JENIS SORTING DATA KENDARAAN BERDASARKAN:")
	fmt.Println("[1] PLAT")
	fmt.Println("[2] TAHUN PRODUKSI")
	fmt.Print("PILIH: ")
	fmt.Scan(&pilih)
	fmt.Println()

	fmt.Println("SILAHKAN PILIH JENIS SORTING DATA KENDARAAN SECARA:")
	fmt.Println("[1] SELECTION SORTING")
	fmt.Println("[2] INSERTION SORTING")
	fmt.Print("PILIH JENIS SORTING: ")
	fmt.Scan(&pilihSorting)
	fmt.Println()

	fmt.Println("SILAHKAN PILIH PENGURUTAN DATA SECARA:")
	fmt.Println("[1] SORTING ASCENDING (TERKECIL -> TERBESAR)")
	fmt.Println("[2] SORTING DESCENDING  (TERBESAR -> TERKECIL)")
	fmt.Print("PILIH JENIS SORTING : ")
	fmt.Scan(&pilihUrutan)
	fmt.Println()

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
				InsertionSortAscPlat(A, nData)
				cetakData(*A, nData)
			case 2:
				InsertionSortDescPlat(A, nData)
				cetakData(*A, nData)
			}
		}
	case 2:
		switch pilihSorting {
		case 1:
			switch pilihUrutan {
			case 1:
				SelectionSortAscTahun(A, nData)
				cetakData(*A, nData)
			case 2:
				SelectionSortDescTahun(A, nData)
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

func optionData5(B *tabRiwayat, nServis int) {
	var pilih int
	fmt.Println("STATISTIK SERVIS")
	fmt.Println("[1] STATISTIK JENIS KERUSAKAN")
	fmt.Println("[2] STATISTIK JUMLAH SERVIS PER BULAN")
	fmt.Print("PILIH MENU: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("STATISTIK JENIS KERUSAKAN")
		var jenisKerusakan = [6]string{
			"SERVIS_BERKALA", "MESIN", "KELISTRIKAN", "BODY", "BAN", "KAKI-KAKI"}
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
			fmt.Printf("JENIS KERUSAKAN %s: (%d kasus)\n", jenisKerusakan[i], pilihan[i])
		}
		fmt.Print("JENIS KERUSAKAN YANG PALING SERING TERJADI: ")
		var max = pilihan[0]
		for i = 1; i < 6; i++ {
			if pilihan[i] > max {
				max = pilihan[i]
			}
		}
		fmt.Printf("%s (%d kasus)\n", jenisKerusakan[max], pilihan[max])

		fmt.Print("JENIS KERUSAKAN YANG JARANG TERJADI: ")
		var min = pilihan[0]
		for i = 1; i < 6; i++ {
			if pilihan[i] < min {
				min = pilihan[i]
			}
		}
		fmt.Printf("%s (%d kasus)\n", jenisKerusakan[min], pilihan[min])
	case 2:
		fmt.Println("STATISTIK JUMLAH SERVIS PER BULAN")
		var bulan [12]int
		var namaBulan = [12]string{
			"JANUARI", "FEBRUARI", "MARET", "APRIL", "MEI", "JUNI", "JULI", "AGUSTUS", "SEPTEMBER", "OKTOBER", "NOVEMBER", "DESEMBER"}
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
				fmt.Printf("BULAN %s: (%d kasus)\n", namaBulan[i], bulan[i])
			}
			fmt.Print("BULAN DENGAN JUMLAH SERVIS PALING BANYAK: ")
			var max = bulan[0]
			for i = 1; i < 12; i++ {
				if bulan[i] > max {
					max = bulan[i]
				}
			}
			fmt.Printf("BULAN %s: (%d kasus)\n", namaBulan[max], max)
			fmt.Print("BULAN DENGAN JUMLAH SERVIS PALING SEDIKIT: ")
			var min = bulan[0]
			for i = 1; i < 12; i++ {
				if bulan[i] < min {
					min = bulan[i]
				}
			}
			fmt.Printf("BULAN %s: (%d kasus)\n", namaBulan[min], min)
		}
	}
}

func optionData6(T *tabPemilik, A *tabKendaraan, B *tabRiwayat, nData int, nServis int) {
	var i int
	fmt.Println("\n===============================================================================================================================")

	fmt.Printf("                      				RIWAYAT DATA                   \n")
	fmt.Println("===============================================================================================================================")

	if nData == 0 {
		fmt.Println("                      [!] BELUM ADA DATA MASUKAN [!]                            ")
		fmt.Println("=========================================================================================")
	} else {
		fmt.Printf("%-6s | %-12s | %-16s | %-12s | %-14s || %-12s | %-16s | %-25s\n", "NO ID", "PLAT NOMOR", "JENIS KENDARAAN", "PEMILIK", "NO. TELPON", "TGL SERVIS", "KATEGORI", "DETAIL SERVIS")
		fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------")
		for i = 0; i < nData; i++ {
			if B[i].tanggal != 0 {
				fmt.Printf("%-6s | %-12s | %-16s | %-12s | %-14s || %02d/%02d/%-6d | %-16s | %-25s\n", "", "", "", "", "", B[i].tanggal, B[i].bulan, B[i].tahun, B[i].jenisKerusakan, B[i].detailServis)
			} else {
				fmt.Printf("%-06d | %-12s | %-16s | %-12s | %-14s || %-12s | %-16s | %-25s\n", i+1, A[i].plat, A[i].jenisKendaraan, T[i].nama, T[i].kontak, "-", "-", "BELUM PERNAH SERVIS")
			}
			fmt.Println("-------------------------------------------------------------------------------------------------------------------------")
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

func InsertionSortAscPlat(dK *tabKendaraan, n int) {
	var pass, k int
	var temp dataKendaraan
	for pass = 1; pass < n; pass++ {
		k = pass
		temp = dK[k]
		for k > 0 && temp.plat < dK[k-1].plat {
			dK[k] = dK[k-1]
			k = k - 1
		}
		dK[k] = temp
	}
}

func InsertionSortDescPlat(dK *tabKendaraan, n int) {
	var pass, k int
	var temp dataKendaraan
	for pass = 1; pass < n; pass++ {
		k = pass
		temp = dK[k]
		for k > 0 && temp.plat > dK[k-1].plat {
			dK[k] = dK[k-1]
			k = k - 1
		}
		dK[k] = temp
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

func SelectionSortAscTahun(dK *tabKendaraan, n int) {
	var pass, i, acuan int
	var temp dataKendaraan
	for pass = 1; pass < n; pass++ {
		acuan = pass - 1
		for i = pass; i < n; i++ {
			if dK[acuan].tahunProduksi > dK[i].tahunProduksi {
				acuan = i
			}
		}
		temp = dK[acuan]
		dK[acuan] = dK[pass-1]
		dK[pass-1] = temp
	}
}

func SelectionSortDescTahun(dK *tabKendaraan, n int) {
	var pass, i, acuan int
	var temp dataKendaraan
	for pass = 1; pass < n; pass++ {
		acuan = pass - 1
		for i = pass; i < n; i++ {
			if dK[acuan].tahunProduksi < dK[i].tahunProduksi {
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
