package main
import "fmt"

const NMAX int = 100000
type dataPemilik struct{
    nama string
    kontak string
}
type dataKendaraan struct{
    jenisKendaraan string
    noKendaraan string
    tahunProduksi int
}

type riwayatServis struct{
    tanggalServis string
    jenisKerusakan string
    detailServis string
}

type tabPemilik[NMAX]dataPemilik
type tabKendaraan[NMAX]dataKendaraan
type tabRiwayat[NMAX]riwayatServis
func main(){
    var
}

func mainMenu(){
    var angka int
    fmt.Println("AUTOCARE")
    fmt.Println("Aplikasi Manajemen dan Riwayat Servis Kendaraan")
    fmt.Println("1. Data Kendaraan & Pemilik")
    fmt.Println("2. Tambah Riwayat Servis Baru")
    fmt.Println("3. Search Kendaraan")
    fmt.Println("4. Urutan Daftar Kendaraan")
    fmt.Println("5. Statistik Servis")
    fmt.Println("6. Keluar")
    fmt.Print("PILIH MENU: ")
    fmt.Scan(&angka)
    
    
}

func optionData1(T tabPemilik, A tabKendaraan){
    fmt.Println("DATA KENDARAAN & PEMILIK")
    fmt.Println("1. ADD")
    fmt.Println("2. EDIT")
    fmt.Println("3. DELETE")
    fmt.Print.("PILIH MENU: ")
    switch pilihan {
    case 1:
        fmt.Println("ADD")
        fmt.Print("NAMA: ")
        fmt.Scan(&T.nama)
        fmt.Print("KONTAK: ")
        fmt.Scan(&T.kontak)
        fmt.Print("JENIS KENDARAAN: ")
        fmt.Scan(&A.jenisKendaraan)
        fmt.Print("NOMOR PLAT: ")
        fmt.Scan(&A.noKendaraan)
        fmt.Print("TAHUN PRODUKSI: ")
        fmt.Scan(&A.tahunProduksi)
    
    case 2:
    case 3:
    }

    func optionData2(){

    }

    funcoptionData3(){
        
    }
}

func BinarySearch(n int)int {
	var left, right, mid, found int
	
	left = 0
	right = n-1
	found = -1
	
	for left <= right && found == -1 {
		mid = (right+left)/2
		if {
			found = mid
		}else if { //ketika data yg dicari > data array mid//
			left = mid + 1
		}else{ //ketika data yg dicari < data array mid//
			right = mid - 1
		}
	}
	return found
}

func Sequential()int {
	var found, r int
	found = -1
	r = 0
	for r <= n-1 && found == -1{
		if{ //jika yg dicari == array r
			found = r
		}
		r++
	}
}

//selection itu procedure dengan array sebagai in/out
func SelectionSortAsc(){
	var pass, i, acuan, temp int
	for pass = 1; pass < ; pass++{
		acuan = pass - 1
		for i = pass; i < ; i++{
			if // array acuan > array i// {
				acuan = i
			}
		}
		temp = // array acuan
		//array acuan = array pass - 1
		//array pass - 1 = temp
	}
}

func SelectionSortDesc(){
	var pass, i, acuan, temp int
	for pass = 1; pass < ; pass++{
		acuan = pass - 1
		for i = pass; i < ; i++{
			if // array acuan > array i// {
				acuan = i
			}
		}
		temp = // array acuan
		//array acuan = array pass - 1
		//array pass - 1 = temp
	}
}

func tambahData(dP *tabPemilik, n *int){
	var namaPemilik string
	fmt.Scan(&namaPemilik)
	*n = 0
	for namaPemilik != "SELESAI" {
		dP[*n].nama = namaPemilik
		fmt.Scan(&dP[*n].kontak)
		*n++
		fmt.Scan(&namaPemilik)
	}
}