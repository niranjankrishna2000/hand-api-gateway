package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

// Upload image godoc
//
//	@Summary		Upload image
//	@Description	User can upload image
//	@Tags			User Post
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		api_key
//	@Param			image	formData	file	true	"image"
//	@Success		200		{string}	fileLink
//	@Failure		400		{string}	string "failed to upload image"
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{string}	string "Error in internal server"
//	@Router			/user/post/upload-image [post]
func UploadImage(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Upload image started")
	userId := ctx.GetInt64("userId")
	log.Println("User ID:", userId)
	fileLink := ""
	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data" + err.Error())
		return
	}

	log.Println("Image header", image.Filename, image.Header)

	log.Println("Image Found")
	filename := fmt.Sprintf("user-00%d-%s.jpg", userId, time.Now().Format("20060102-150405"))
	err = ctx.SaveUploadedFile(image, "src/assets/"+filename)
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, "Error in internal server")
		return
	}

	fileLink = "/src/assets/" + filename
	log.Println("File saved at:", fileLink)

	ctx.JSON(http.StatusCreated, fileLink)
}
