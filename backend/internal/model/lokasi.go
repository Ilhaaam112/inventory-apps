package model

type Lokasi struct {
	ID         int    `json:"id"`
	NamaLokasi string `json:"nama_lokasi"`
	Keterangan string `json:"keterangan"`
}
