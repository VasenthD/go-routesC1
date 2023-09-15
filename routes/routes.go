package routes

import (
	"mongodb1/controller"

	"github.com/gin-gonic/gin"
)

func ProfileRoute(router *gin.Engine, controller controller.ProfileController) {
	router.POST("/api/profile/create", controller.CreateProfile)
	router.POST("/api/profile/edit", controller.EditProfile)
	router.DELETE("/api/profile/delete/:id", controller.DeleteProfile)
	router.GET("/api/profile/:id", controller.GetProfileById)
	router.POST("/api/profile/search", controller.GetProfileBySearch)
}

//routes is expecting the controller
