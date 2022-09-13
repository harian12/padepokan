package router

import (
	"padepokan/app/controllers"
	"padepokan/app/repositories"
	"padepokan/app/services"
	"padepokan/config"
)

var db = config.ConnectionDB()

var nasabah_repository repositories.INasabahRepository = repositories.NewNasabahRepository(db)
var nasabah_service services.INasabahService = services.NewNasabahService(nasabah_repository)
var nasabah_controller controllers.INasabahController = controllers.NewNasabahController(nasabah_service)

var transaction_repository repositories.ITransactionRepository = repositories.NewTransactionRepository(db)
var transaction_service services.ITransactionService = services.NewTransactionService(transaction_repository)
var transaction_controller controllers.ITransactionController = controllers.NewTransactionController(transaction_service)
