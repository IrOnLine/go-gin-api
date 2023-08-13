package main

import (
	"github.com/ironline/go-gin-api/pkg/books"
	"github.com/ironline/go-gin-api/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// инициализируем Viper для обработки переменных среды;
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	// инициализируем базу данных на основе GORM;
	r := gin.Default()
	h := db.Init(dbUrl)

	// добавляем простой маршрут /;
	books.RegisterRoutes(r, h)
	// register more routes here

	// запускаем приложение
	r.Run(port)
}
