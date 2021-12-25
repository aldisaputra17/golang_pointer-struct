package main

import "fmt"

type Mahasiswa struct {
	Id      int
	Nama    string
	Jurusan string
	Nim     int
}

func mahasiswa(m *Mahasiswa) {
	if m.Id > 2 {
		m.Nama = "Dono"
		fmt.Println("Berhasil di ubah")
	} else {
		fmt.Println("Gagal di ubah")
	}
}

func main() {
	hasil := &Mahasiswa{Id: 1, Nama: "Aldi", Jurusan: "Ti", Nim: 141515}
	mahasiswa(hasil)
	fmt.Println(hasil)
}
