package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EditProfileRequestBody struct {
	Name       string `json:"name,omitempty"` //test
	Email      string `json:"email,omitempty"`
	PanNo      string `json:"pan,omitempty" validate:"len=10,number"`
	Phone      string `json:"phone,omitempty"`
	Gender     string `json:"gender,omitempty" validate:"oneof='male' 'female' 'others'"`
	Address    string `json:"address,omitempty" validate:"max=50,ascii"`
	ProfilePic string `json:"profilepic,omitempty"`
	Dob        string `json:"dob,omitempty"`
}

// Edit Profile  godoc
//
//	@Summary		Edit profile
//	@Description	User can edit profile
//	@Description	Choose Gender : male,female and others
//	@Description	Date of Birth Format : 2006-01-02 00:00:00
//	@Tags			User Profile
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			body	body		EditProfileRequestBody	true	"Edit profile Data"
//	@Success		200		{object}	pb.EditProfileResponse
//	@Failure		400		{object}	pb.EditProfileResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.EditProfileResponse
//	@Router			/user/profile/edit  [patch]
func EditProfile(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Edit profile started...",int32(ctx.GetInt64("userid")))
	body := EditProfileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.EditProfileResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			User:     &pb.UserProfile{},
		})
		return
	}
	//validator
	validator := validator.New()
	if err := validator.Struct(body); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.EditProfileResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			User:     &pb.UserProfile{},
		})
		return
	}
	res, err := c.EditProfile(context.Background(), &pb.UserProfile{
		Id:             int32(ctx.GetInt64("userid")),
		Name:           body.Name,
		Email:          body.Email,
		Phone:          body.Phone,
		Gender:         body.Gender,
		DoB:            body.Dob,
		Address:        body.Address,
		PAN:            body.PanNo,
		ProfilePicture: body.ProfilePic,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.EditProfileResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			User:     &pb.UserProfile{},
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
