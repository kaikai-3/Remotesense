package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)
func main(){
	//create a router
	router := gin.Default()
	//config static file service
	err := os.MkdirAll("./images",os.ModePerm)
	if err != nil {
		fmt.Println("Failed ot create images directory:",err)
		return
	}

	router.POST("/upload",func (c *gin.Context)  {
		file , err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"fail to uploda file"})
			return
		}
		
		//save image to localhost
		uploadpath := "./uploads/"
		err = os.MkdirAll(uploadpath,os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"Failed to create upload directory"})
			return
		}
		filePath := filepath.Join(uploadpath,file.Filename)
		if err := c.SaveUploadedFile(file,file.Filename);err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"Failed to save file"})
			return
		}

		//simulate image processing
		time.Sleep(2*time.Second)

		processedPath := "./images/"+file.Filename
		err = os.MkdirAll("./images",os.ModePerm)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"Failed to create images directory"})
			return
		}
		if err := os.Rename(filePath,processedPath);err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"Failed to process file"})
			return 
		}

		//return processed image
		c.JSON(http.StatusOK, gin.H{
			"url":"/images/"+file.Filename,
		})
	})

	router.Run(":8081")
}