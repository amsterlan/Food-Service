package routes

import "github.com/gin-gonic/gin"

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetMenus())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetMenu())
	incomingRoutes.POST("/invoices", controller.CreateMenu())
	incomingRoutes.PATCH("/foods/:invoice_id", controller.UpdateMenu())
}
