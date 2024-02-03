package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MonthlyGoalBody struct {
	Category int `json:"category" validate:"required,min=1,max=5,number"` //note add
	Amount   int    `json:"amount" validate:"required,min=100,max=10000,number"`
	Day      int    `json:"day" validate:"required,min=1,max=28,number"`
}

// MonthlyGoal godoc
//
//	@Summary		User MonthlyGoal
//	@Description	User can see MonthlyGoal
//	@Tags			User Profile
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Success		200	{object}	pb.GetMonthlyGoalResponse
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{object}	pb.GetMonthlyGoalResponse
//	@Router			/user/profile/monthly-goal  [get]
func GetMonthlyGoal(ctx *gin.Context, c pb.UserServiceClient) {
	userid:=ctx.GetInt64("userid")
	res, err := c.GetMonthlyGoal(context.Background(), &pb.GetMonthlyGoalRequest{
		Userid: int32(userid),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.GetMonthlyGoalResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server:" + err.Error(),
			Amount:   0,
			Day:      0,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)
}

// Add MonthlyGoal godoc
//
//	@Summary		Add MonthlyGoal
//	@Description	User can add MonthlyGoal
//	@Description	category: 1. medical,2. child care,3. animal care, 4. Education,5. Memorial
//	@Tags			User Profile
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			monthlyGoalBody	body		MonthlyGoalBody	true	"Monthly Goal Data"
//	@Success		200				{object}	pb.AddMonthlyGoalResponse
//	@Failure		400				{object}	pb.AddMonthlyGoalResponse
//	@Failure		403				{string}	string	"You have not logged in"
//	@Failure		502				{object}	pb.AddMonthlyGoalResponse
//	@Router			/user/profile/monthly-goal  [post]
func AddMonthlyGoal(ctx *gin.Context, c pb.UserServiceClient) {
	monthlyGoalBody := MonthlyGoalBody{}
	if err := ctx.BindJSON(&monthlyGoalBody); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.AddMonthlyGoalResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Category: 0,
			Amount:   0,
			Day:      0,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(monthlyGoalBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.AddMonthlyGoalResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Category: 0,
			Amount:   0,
			Day:      0,
		})
		return
	}
	userid:=ctx.GetInt64("userid")
	res, err := c.AddMonthlyGoal(context.Background(), &pb.AddMonthlyGoalRequest{
		Userid:   int32(userid),
		Category: int32(monthlyGoalBody.Category),
		Amount:   int64(monthlyGoalBody.Amount),
		Day:      int32(monthlyGoalBody.Day),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.AddMonthlyGoalResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server:" + err.Error(),
			Category: 0,
			Amount:   0,
			Day:      0,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)

}

// Edit MonthlyGoal godoc
//
//	@Summary		Edit MonthlyGoal
//	@Description	User can edit an MonthlyGoal
//	@Description	category: 1. medical,2. child care,3. animal care, 4. Education,5. Memorial
//	@Tags			User Profile
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			monthlyGoalBody	body		MonthlyGoalBody	true	"MonthlyGoal Data"
//	@Success		200				{object}	pb.EditMonthlyGoalResponse
//	@Failure		400				{object}	pb.EditMonthlyGoalResponse
//	@Failure		403				{string}	string	"You have not logged in"
//	@Failure		502				{object}	pb.EditMonthlyGoalResponse
//	@Router			/user/profile/monthly-goal  [put]
func EditMonthlyGoal(ctx *gin.Context, c pb.UserServiceClient) {
	monthlyGoalBody := MonthlyGoalBody{}
	if err := ctx.BindJSON(&monthlyGoalBody); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.EditMonthlyGoalResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Category: 0,
			Amount:   0,
			Day:      0,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(monthlyGoalBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.EditMonthlyGoalResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data",
			Category: 0,
			Amount:   0,
			Day:      0,
		})
		return
	}
	userid:=ctx.GetInt64("userid")
	res, err := c.EditMonthlyGoal(context.Background(), &pb.EditMonthlyGoalRequest{
		Userid:   int32(userid),
		Category: int32(monthlyGoalBody.Category),
		Amount:   int64(monthlyGoalBody.Amount),
		Day:      int32(monthlyGoalBody.Day),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.EditMonthlyGoalResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server:" + err.Error(),
			Category: 0,
			Amount:   0,
			Day:      0,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)

}
