package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"sigitprd/golang-restful-api/app"
	"sigitprd/golang-restful-api/controller"
	"sigitprd/golang-restful-api/helper"
	"sigitprd/golang-restful-api/middleware"
	"sigitprd/golang-restful-api/repository"
	"sigitprd/golang-restful-api/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
