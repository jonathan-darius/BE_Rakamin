package routes

import (
	"BE_Rakamin/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/AddPerson", controllers.CreatePerson())
	incomingRoutes.GET("/GetUsers", controllers.GetAllUser())
	incomingRoutes.GET("/GetUser/:user_id", controllers.GetUser())
	incomingRoutes.DELETE("DeleteUser/:user_id", controllers.DeleteUser())
	incomingRoutes.PATCH("UpdateUser/:user_id", controllers.UpdateUser())

}
