package repository

import (
	"Kak-Meyda-Ku/Models"
	connections "Kak-Meyda-Ku/Models/Connections"
	jurusanmodels "Kak-Meyda-Ku/Models/JurusanModels"
)

func InsertJurusan(DataMahasiswa jurusanmodels.DataMahasiswa) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "INSERT INTO DataMahasiswa (Id, Nim, Nama_Mahasiswa, Gender, Jurusan, Tahun_Ajaran, Jenis_Reg, Tanggal_Masuk_Perkuliahan, Semester, Jumlah_Sks, Ipk, Jumlah_MataKuliah_Lulus, Jumlah_MataKuliah_Td_Lulus, Jumlah_Sks_Lulus, Jumlah_Sks_Td_Lulus, Keterangan) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	tempResult := db.Exec(query, DataMahasiswa.Id, DataMahasiswa.Nim, DataMahasiswa.Nama_Mahasiswa, DataMahasiswa.Gender, DataMahasiswa.Jurusan, DataMahasiswa.Tahun_Ajaran, DataMahasiswa.Jenis_Reg, DataMahasiswa.Tangal_Masuk_Perkuliahan, DataMahasiswa.Semester, DataMahasiswa.Jumlah_Sks, DataMahasiswa.Ipk, DataMahasiswa.Jumlah_MataKuliah_Lulus, DataMahasiswa.Jumlah_MataKuliah_Td_Lulus, DataMahasiswa.Jumlah_Sks_Lulus, DataMahasiswa.Jumlah_Sks_Td_Lulus, DataMahasiswa.Keterangan)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Models.BaseResponseModels{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data berhasil ditambahkan.",
			Data:          nil,
		}
	}

	return result
}


func UpdateJurusan(DataMahasiswa jurusanmodels.DataMahasiswa) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE DataMahasiswa SET Nim = ?, Nama_Mahasiswa = ?, Gender = ?, Jurusan = ?, Tahun_Ajaran = ?, Jenis_Reg = ?, Tanggal_Masuk_Perkuliahan = ?, Semester = ?, Jumlah_Sks = ?, Ipk = ?, Jumlah_MataKuliah_Lulus = ?, Jumlah_MataKuliah_Td_Lulus = ?, Jumlah_Sks_Lulus = ?, Jumlah_Sks_Td_Lulus = ?, Keterangan = ? WHERE id = ?"

	tempResult := db.Exec(query,DataMahasiswa.Nim, DataMahasiswa.Nama_Mahasiswa, DataMahasiswa.Gender, DataMahasiswa.Jurusan, DataMahasiswa.Tahun_Ajaran, DataMahasiswa.Jenis_Reg, DataMahasiswa.Tangal_Masuk_Perkuliahan, DataMahasiswa.Semester, DataMahasiswa.Jumlah_Sks, DataMahasiswa.Ipk, DataMahasiswa.Jumlah_MataKuliah_Lulus, DataMahasiswa.Jumlah_MataKuliah_Td_Lulus, DataMahasiswa.Jumlah_Sks_Lulus, DataMahasiswa.Jumlah_Sks_Td_Lulus, DataMahasiswa.Keterangan, DataMahasiswa.Id)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Models.BaseResponseModels{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil diubah.",
				Data:          nil,
			}
		}
	}

	return result
}


func DeleteJurusan(request jurusanmodels.DataMahasiswa) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM DataMahasiswa WHERE Id = ?"

	tempResult := db.Exec(query, request.Id)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Periksa apakah ada baris yang terpengaruh oleh perintah DELETE
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Models.BaseResponseModels{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil dihapus.",
				Data:          nil,
			}
		}
	}

	return result
}

func GetJurusanByID(request jurusanmodels.DataMahasiswa) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var DataMahasiswa []jurusanmodels.DataMahasiswa
	db := connections.DB

	if request.Id != 0 {
		query = "SELECT * FROM DataMahasiswa WHERE Id = ?"
	} else {
		query = "SELECT * FROM DataMahasiswa"
	}

	tempResult := db.Raw(query).Find(& DataMahasiswa)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {	
		result = Models.BaseResponseModels{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data retrieved successfully",
			Data:         	DataMahasiswa,
		}
	}

	return result
}
