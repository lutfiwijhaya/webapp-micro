package main

import (
	"fmt"
	"reimbursement-service-go/config"
	"reimbursement-service-go/models"
	"reimbursement-service-go/routes"
)

func main() {
	// Inisialisasi database
	config.InitDB()

	// AutoMigrate tabel
	config.DB.AutoMigrate(&models.Reimbursement{}, &models.Category{}, &models.Log{})

	fmt.Println("Reimbursement service running on :8085")

	// Jalankan server di port 8085
	r := routes.SetupRouter()
	r.Run(":8085")
}
