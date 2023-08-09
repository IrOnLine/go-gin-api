package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-gin-api-medium/pkg/books"
	"github.com/hellokvn/go-gin-api-medium/pkg/common/db"
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
