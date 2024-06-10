package main

import (
	controller "Kak-Meyda-Ku/Controller"
	connection "Kak-Meyda-Ku/Models/Connections"

	"github.com/gin-gonic/gin"
)

func main() {

	port := ":1213"
	r := gin.Default()
	connection.ConnectDatabase()

	 r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            return
        }
        c.Next()
    })



	//###BEGIN WEB API Keterangan
	// Get data
	r.GET("/api/DataMahasiswa/GetJurusan", controller.GetJurusanByID)

	//Insert data
	r.POST("/api/DataMahasiswa/InsertJurusan", controller.InsertJurusan)

	// Update data
	r.PUT("/api/DataMahasiswa/UpdateJurusan", controller.UpdateJurusan)

	//Delete data
	r.DELETE("/api/DataMahasiswa/DeleteJurusan", controller.DeleteJurusan)
	//###END WEB API Keterangan


	r.Run(port)
}