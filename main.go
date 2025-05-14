package main

import (
	"Gin_PlayGround/post"
	"Gin_PlayGround/user"
	"fmt"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, Gin!")
	err := godotenv.Load()
	if err != nil {
		log.Println(".env 파일을 찾을 수 없거나 로드할 수 없습니다.")
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	var userRepoImpl = user.UserRepositoryImpl{}
	var userRepo user.UserRepository = &userRepoImpl

	if err := userRepo.AutoMigrate(db); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
	if err := post.AutoMigrate(db); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	r := gin.Default()
	_ = r.SetTrustedProxies(nil)
	

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	user.RegisterRoutes(r, db)
	post.RegisterRoutes(r, db)

	r.Run() // 기본적으로 :8080 포트에서 실행
} 
