package jurusanmodels

import "github.com/shopspring/decimal"

type DataMahasiswa struct {
	Id                         int64   `json:"Id" from:"Id"`
	Nim                        float64 `JSON:"Nim" from:"Nim"`
	Nama_Mahasiswa             string  `json:"Nama Mahasiswa" from:"Nama Mahasiswa"`
	Gender                     string  `json:"Gender" from:"Gender"`
	Jurusan                    string  `json:"Jurusan" from:"Jurusan"`
	Tahun_Ajaran               float64 `json:"Tahun Ajaran" from:"Tahun Ajaran"`
	Jenis_Reg                  string  `json:"Jenis Reg" from:"Jenis Reg"`
	Tangal_Masuk_Perkuliahan   float64 `json:"Tanggal Masuk Perkuliahan" from:"Tanggal Masuk Perkuliahan"`
	Semester                   int64   `json:"Semester" from:"Semester"`
	Jumlah_Sks                 int64   `json:"Jumlah Sks" from:"Jumlah Sks"`
	Ipk                        decimal.Decimal `json:"Ipk" from:"Ipk"`
	Jumlah_MataKuliah_Lulus    float64 `json:"Jumlah Matakuliah Lulus" from:"Jumlah MataKuluah Lulus"`
	Jumlah_MataKuliah_Td_Lulus float64 `json:"Jumlah MataKuliah Td Lulus" from:"Jumlah MataKuliah Td Lulus"`
	Jumlah_Sks_Lulus           float64 `json:"Jumlah Sks Lulus" from:"Jumlah Sks Lulus"`
	Jumlah_Sks_Td_Lulus        float64 `json:"Jumlah Sks Td Lulus" from:"Jumlah Sks Td Lulus"`
	Keterangan                 string  `Json:"Keterangan" from:"Keterangan"`
}