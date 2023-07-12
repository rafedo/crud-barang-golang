// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package Route

import (
	"crud-barang/Controller"
	"crud-barang/Repository"
	"crud-barang/Services"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func barangDI(db *gorm.DB) *Controller.BarangControllerImpl {
	barangRepositoryImpl := Repository.BarangRepositoryControllerProvider(db)
	barangServiceImpl := Services.BarangServiceControllerProvider(barangRepositoryImpl)
	barangControllerImpl := Controller.BarangControllerControllerProvider(barangServiceImpl)
	return barangControllerImpl
}
