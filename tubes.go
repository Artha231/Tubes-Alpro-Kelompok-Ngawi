package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxPasien = 100

type Pasien struct {
	ID       int    `json:"ID"`
	Nama     string `json:"Nama"`
	Umur     int    `json:"Umur"`
	Diagnosa string `json:"Diagnosa"`
}

var dataPasien [maxPasien]Pasien
var jumlahPasien int = 0

func tambahPasien(id int, nama string, umur int, diagnosa string) {
	if jumlahPasien < maxPasien {
		dataPasien[jumlahPasien] = Pasien{id, nama, umur, diagnosa}
		jumlahPasien++
		simpanKeJSON()
	} else {
		fmt.Println("Data penuh!")
	}
}

func tampilkanPasien() {
	fmt.Println("\nDaftar Pasien:")
	for i := 0; i < jumlahPasien; i++ {
		tampilkanData(dataPasien[i])
	}
}

func tampilkanData(p Pasien) {
	fmt.Printf("ID: %d | Nama: %s | Umur: %d | Diagnosa: %s\n", p.ID, p.Nama, p.Umur, p.Diagnosa)
}

func sequentialSearch(nama string) int {
	for i := 0; i < jumlahPasien; i++ {
		if strings.EqualFold(dataPasien[i].Nama, nama) {
			return i
		}
	}
	return -1
}

func binarySearch(id int) int {
	low := 0
	high := jumlahPasien - 1
	for low <= high {
		mid := (low + high) / 2
		if dataPasien[mid].ID == id {
			return mid
		} else if dataPasien[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func selectionSortUmur(asc bool) {
	for i := 0; i < jumlahPasien-1; i++ {
		idx := i
		for j := i + 1; j < jumlahPasien; j++ {
			if (asc && dataPasien[j].Umur < dataPasien[idx].Umur) || (!asc && dataPasien[j].Umur > dataPasien[idx].Umur) {
				idx = j
			}
		}
		swap(i, idx)
	}
}

func insertionSortNama(asc bool) {
	for i := 1; i < jumlahPasien; i++ {
		key := dataPasien[i]
		j := i - 1
		for (j >= 0) && ((asc && strings.ToLower(dataPasien[j].Nama) > strings.ToLower(key.Nama)) || (!asc && strings.ToLower(dataPasien[j].Nama) < strings.ToLower(key.Nama))) {
			dataPasien[j+1] = dataPasien[j]
			j--
		}
		dataPasien[j+1] = key
	}
}

func insertionSortID(asc bool) {
	for i := 1; i < jumlahPasien; i++ {
		key := dataPasien[i]
		j := i - 1
		for (j >= 0) && ((asc && dataPasien[j].ID > key.ID) || (!asc && dataPasien[j].ID < key.ID)) {
			dataPasien[j+1] = dataPasien[j]
			j--
		}
		dataPasien[j+1] = key
	}
}

func swap(i, j int) {
	temp := dataPasien[i]
	dataPasien[i] = dataPasien[j]
	dataPasien[j] = temp
}

func inputPasien() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pasien: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	fmt.Print("Masukkan Nama Pasien: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Umur Pasien: ")
	umurStr, _ := reader.ReadString('\n')
	umurStr = strings.TrimSpace(umurStr)
	umur, _ := strconv.Atoi(umurStr)

	fmt.Print("Masukkan Diagnosa: ")
	diagnosa, _ := reader.ReadString('\n')
	diagnosa = strings.TrimSpace(diagnosa)

	tambahPasien(id, nama, umur, diagnosa)
}

func simpanKeJSON() {
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
		return
	}
	defer file.Close()

	list := make([]Pasien, jumlahPasien)
	for i := 0; i < jumlahPasien; i++ {
		list[i] = dataPasien[i]
	}

	jsonData, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		fmt.Println("Gagal mengubah data menjadi JSON:", err)
		return
	}

	file.Write(jsonData)
}

func muatDariJSON() {
	file, err := os.Open("data.json")
	if err != nil {
		// file tidak ditemukan, biarkan kosong
		return
	}
	defer file.Close()

	var list []Pasien
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&list)
	if err != nil {
		fmt.Println("Gagal memuat data JSON:", err)
		return
	}

	jumlahPasien = 0
	for i := 0; i < len(list) && i < maxPasien; i++ {
		dataPasien[i] = list[i]
		jumlahPasien++
	}
}

func menu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Tampilkan Semua Pasien")
		fmt.Println("3. Cari Pasien berdasarkan Nama (Sequential Search)")
		fmt.Println("4. Cari Pasien berdasarkan ID (Binary Search)")
		fmt.Println("5. Urutkan berdasarkan Umur (Selection Sort)")
		fmt.Println("6. Urutkan berdasarkan Nama (Insertion Sort)")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "1" {
			inputPasien()
		} else if input == "2" {
			tampilkanPasien()
		} else if input == "3" {
			fmt.Print("Masukkan nama yang dicari: ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			idx := sequentialSearch(nama)
			if idx != -1 {
				tampilkanData(dataPasien[idx])
			} else {
				fmt.Println("Pasien tidak ditemukan.")
			}
		} else if input == "4" {
			fmt.Print("Masukkan ID yang dicari: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)
			insertionSortID(true)
			idx := binarySearch(id)
			if idx != -1 {
				tampilkanData(dataPasien[idx])
			} else {
				fmt.Println("Pasien tidak ditemukan.")
			}
		} else if input == "5" {
			fmt.Print("Urutkan berdasarkan Umur (asc/desc): ")
			opt, _ := reader.ReadString('\n')
			opt = strings.TrimSpace(opt)
			selectionSortUmur(strings.ToLower(opt) == "asc")
			fmt.Println("Data telah diurutkan berdasarkan umur.")
		} else if input == "6" {
			fmt.Print("Urutkan berdasarkan Nama (asc/desc): ")
			opt, _ := reader.ReadString('\n')
			opt = strings.TrimSpace(opt)
			insertionSortNama(strings.ToLower(opt) == "asc")
			fmt.Println("Data telah diurutkan berdasarkan nama.")
		} else if input == "7" {
			fmt.Println("Terima kasih!")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	muatDariJSON()
	menu()
}
