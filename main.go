package main

import (
	"Project/common"
	"Project/modules/item/model"
	ginitem "Project/modules/item/transport/gin"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Mở kết nối cơ sở dữ liệu
	dsn := os.Getenv("Dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection failed:", err)
		log.Fatalln(err)
	}

	log.Println("Connected:", db)

	// Tạo một đối tượng Gin router
	r := gin.Default()

	// Nhóm các route liên quan đến API
	api := r.Group("/api/")
	{
		api.POST("/create", ginitem.CreateItem(db))
		//api.GET("/getById/:id", GetItemById(db))
		//api.GET("/getAll", GetAll(db))
		//api.PATCH("/items/:id", editItemById(db))
		//api.PATCH("/delete/:id", deleteItemById(db))
	}

	// Chạy server trên cổng 8000
	err = r.Run(":8000")
	if err != nil {
		log.Fatalln("Server failed to start:", err)
	}
}

func GetItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ToDoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.Id = id
		if err := db.First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}

func deleteItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
		if err := db.Table(model.ToDoItem{}.TableName()).Where(`id = ?`, id).Updates(map[string]interface{}{
			"status": "Deleted",
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
func editItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ToDoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)

			return
		}
		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
func GetAll(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.DataPaging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)

			return
		}
		paging.Process()
		var data []model.ToDoItem

		db.Where("status <> ?", "Deleted")

		if err := db.Table(model.ToDoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		if err := db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, common.NewSuccessRes(data, paging, nil))

	}
}
