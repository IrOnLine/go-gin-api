https://medium.com/nuances-of-programming/%D0%BA%D0%B0%D0%BA-%D0%BF%D0%BE%D1%81%D1%82%D1%80%D0%BE%D0%B8%D1%82%D1%8C-%D0%BC%D0%B0%D1%81%D1%88%D1%82%D0%B0%D0%B1%D0%B8%D1%80%D1%83%D0%B5%D0%BC%D1%8B%D0%B9-api-%D0%BD%D0%B0-go-%D1%81-%D0%BF%D0%BE%D0%BC%D0%BE%D1%89%D1%8C%D1%8E-gin-b20be1c9ae58

Создание базы данных
Для начала нужно создать базу данных. Я знаю, что все делают это по-своему. Некоторые используют GUI, но мы будем работать из терминала.
Должна быть установлена PostgreSQL. В этом случае нижеприведенные команды будут работать в системах Linux, Mac и Windows:

        // открываем командную строку psql под пользователем postgres
        $ psql postgres
        //cоздаем базу данных
        $ CREATE DATABASE go_medium_api;
        //выводим список всех баз данных
        $ \l
        //закрываем командную строку
        $ \q

Настройка проекта
Далее мы инициируем проект и устанавливаем все необходимые модули.
! Внимание: замените YOUR_USERNAME на свое имя пользователя Github.

        $ mkdir go-gin-api-medium
        $ cd go-gin-api-medium 
        $ code . 
        $ go mod init github.com/YOUR_USERNAME/go-gin-api-medium

Теперь установим Gin, GORM и Viper.
С помощью Viper мы будем управлять переменными среды.

        $ go get github.com/spf13/viper 
        $ go get github.com/gin-gonic/gin 
        $ go get gorm.io/gorm 
        $ go get gorm.io/driver/postgres 

Переменные среды
        pkg/common/envs/.env

Конфигурация
        pkg/common/config/config.go

Модели/сущности
        pkg/common/models/book.go

Инициализация базы данных (GORM):
        pkg/common/db/db.go
