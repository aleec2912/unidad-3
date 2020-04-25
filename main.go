package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type Libro struct {
	ID        			uint   `json:"id"`
	Titulo 				string `json:"titulo"`
	Descripcion  		string `json:"descripcion"`
	Autor      			string `json:"autor"`
	Editorial			string `json:"editorial"`
	FechaPublicacion	string `json:"fechapublicacion"` 
}


func main() {
	db, _ = gorm.Open("mysql", "root:@/biblioteca?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	r := gin.Default()
	r.GET("/libro/:id", GetLibro)
	r.GET("/libro/", GettLibro)
	r.POST("/libro", CreateLibro)
	r.PUT("/libro/:id", UpdateLibro)
	r.DELETE("/libro/:id", DeleteLibro)

	r.Run(":8080")
}

func DeleteLibro(c *gin.Context) {
	id := c.Params.ByName("id")
	var libro Libro
	d := db.Where("id = ?", id).Delete(&libro)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateLibro(c *gin.Context) {

	var libro Libro
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&libro)

	db.Save(&libro)
	c.JSON(200, libro)

}

func CreateLibro(c *gin.Context) {

	var libro Libro
	c.BindJSON(&libro)

	db.Create(&libro)
	c.JSON(200, libro)
}

func GetLibro(c *gin.Context) {
	id := c.Params.ByName("id")
	var libro Libro
	if err := db.Where("id = ?", id).First(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, libro)
	}
}
func GettLibro(c *gin.Context) {
	var libro []Libro
	if err := db.Find(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, libro)
	}

}
