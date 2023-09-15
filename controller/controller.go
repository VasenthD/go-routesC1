package controller

import (
	"mongodb1/models"
	services "mongodb1/transactionservice"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileService services.Transactionservice
}

func InitProfileController(profileService services.Transactionservice) ProfileController {
	return ProfileController{profileService}
}

func (pc *ProfileController) CreateProfile(ctx *gin.Context) {
	var profile *models.Transaction
	//context .bindJson convert the json to go lang structure
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return

	}

	// the go lang struture is passed as parameter to Createprofile
	newProfile, err := pc.ProfileService.Createtransaction(profile)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}
func (pc *ProfileController) EditProfile(context *gin.Context)        {}
func (pc *ProfileController) GetProfileById(context *gin.Context)     {}
func (pc *ProfileController) GetProfileBySearch(context *gin.Context) {}
func (pc *ProfileController) DeleteProfile(context *gin.Context)      {}

// dependent on only services: just give the service which will insert the data
