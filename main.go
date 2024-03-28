package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"ugmo/handler"
	"ugmo/repository"
	"ugmo/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	r := setupRoute()
	r.Run(viper.GetString("app.port"))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func setupRoute() *gin.Engine {
	r := gin.Default()

	// test := repository.Video{
	// 	Uni:        "KMITL",
	// 	Year:       "2023",
	// 	Faculty:    "Engineering",
	// 	Department: "School of International and Interdisciplinary Engineering Programs (SIIE)",
	// 	Curriculum: "Computer Innovation Engineering (International program)",
	// }

	db := initDB()
	videoRepo := repository.NewRepositoryDB(db)
	videoService := service.NewVideoService(videoRepo)
	videoHandler := handler.NewVideoHandler(videoService)

	r.Use(corsMiddleware())
	r.POST("/search", videoHandler.GetVideosId)
	r.GET("/image/:id", videoHandler.GetImage)
	r.GET("/video/:id", videoHandler.GetVideo)
	// a, err := videoRepo.GetVideosId(test)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(a)
	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}

func initDB() *sql.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.hostname"),
		viper.GetInt("db.port"),
		viper.GetString("db.dbname"))

	db, err := sql.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db
}
