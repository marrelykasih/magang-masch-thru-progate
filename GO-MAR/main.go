package main

import (
	"encoding/json"
	"fmt"
)

type Tugas struct {
	ID      int
	Judul   string
	Selesai bool
}


func tampilkanTugas(daftar []Tugas) {
	fmt.Println("=== Daftar Tugas Hari Ini ===")
	for _, tugas := range daftar {
		status := "Belum"
		if tugas.Selesai {
			status = "Selesai"
		}
		fmt.Printf("[%d] %s - Status: %s\n", tugas.ID, tugas.Judul, status)
	}
}


func tandaiSelesaiDanCetakJSON(tugas *Tugas) {
	
	tugas.Selesai = true


	jsonData, err := json.MarshalIndent(tugas, "", "  ")
	if err != nil {
		fmt.Println("Terjadi error saat membuat JSON:", err)
		return
	}

	fmt.Println("\n=== Data JSON Setelah Update ===")
	fmt.Println(string(jsonData))
}

func main() {

	daftarTugas := []Tugas{
		{ID: 1, Judul: "Belajar Golang Dasar", Selesai: false},
		{ID: 2, Judul: "Membuat Program Kecil Golang", Selesai: false},
	}

	tandaiSelesaiDanCetakJSON(&daftarTugas[0])

}