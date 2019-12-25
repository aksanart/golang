package models

type Buku struct {
	Idbuku int
	Judul  string
	Harga  string
	Path   string
}
type Bukus []Buku

func SetBuku() *Buku {
	return &Buku{}
}
