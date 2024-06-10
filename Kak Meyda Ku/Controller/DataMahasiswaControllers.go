package controler

import (
	"Kak-Meyda-Ku/Models"
	JurusanModels "Kak-Meyda-Ku/Models/JurusanModels"
	DataMahasiswaRepository "Kak-Meyda-Ku/Repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJurusanByID(c *gin.Context) {
	var request JurusanModels.DataMahasiswa
	var response Models.BaseResponseModels

	response = DataMahasiswaRepository.GetJurusanByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertJurusan(c *gin.Context) {
	var request JurusanModels.DataMahasiswa
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = DataMahasiswaRepository.InsertJurusan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateJurusan(c *gin.Context) {
	var request JurusanModels.DataMahasiswa
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = DataMahasiswaRepository.UpdateJurusan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteJurusan(c *gin.Context) {
	var request JurusanModels.DataMahasiswa
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = DataMahasiswaRepository.DeleteJurusan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}