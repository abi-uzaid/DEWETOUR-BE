package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Upload file
		file, err := c.FormFile("image")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Error Retrieving the File")
			return
		}

		// Open the uploaded file
		src, err := file.Open()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, "Error Opening the File")
			return
		}
		defer src.Close()

		// Create a temporary file
		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, "Error Creating the Temporary File")
			return
		}
		defer tempFile.Close()

		// Copy the file content to the temporary file
		if _, err = io.Copy(tempFile, src); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, "Error Writing the File")
			return
		}

		// Get the path of the temporary file
		data := tempFile.Name()

		filename := data[8:]

		// Add the file path to the context
		c.Set("dataFile", filename)
		next(c)
	}
}
