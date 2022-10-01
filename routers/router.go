package routers

import (
	"fmt"
	"golang-be/config"
	"golang-be/controllers"
	"golang-be/utils/helper"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GoRouter struct {
	GinFunc gin.HandlerFunc
	Router  *gin.Engine
}

// Server ..
func Server(listenAddress string) (err error) {
	debugMode := helper.GetEnv("APPS_DEBUG", "debug")
	gin.SetMode(debugMode)
	hrisRouter := GoRouter{}
	hrisRouter.Routers()

	err = http.ListenAndServe(listenAddress, hrisRouter.Router)

	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}
	fmt.Println("Routing successfully: ", listenAddress)

	return err
}

func (goRouter *GoRouter) Routers() {
	db, _ := config.DatabaseOpen()

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	AuthController := controllers.NewAuthController(db)
	UserController := controllers.NewUserController(db)
	ItemController := controllers.NewItemController(db)
	POHeaderController := controllers.NewPoHeaderController(db)
	PODetailController := controllers.NewPoDetailController(db)

	api := router.Group("/api/v1")
	{
		login := api.Group("/login")
		{
			login.POST("", AuthController.Login)
		}

		user := api.Group("/user")
		{
			user.POST("/", UserController.SaveUser)
			user.PUT("/", UserController.UpdateUser)
			user.POST("/id", UserController.GeUserByID)
			user.POST("/delete", UserController.Delete)

		}

		item := api.Group("/item")
		{
			item.POST("/", ItemController.SaveItem)
			item.PUT("/", ItemController.UpdateItem)
			item.POST("/id", ItemController.GeItemByID)
			item.POST("/delete", ItemController.Delete)

		}

		poHeader := api.Group("/po-header")
		{
			poHeader.POST("/", POHeaderController.SavePoHeader)
			poHeader.PUT("/", POHeaderController.UpdatePoHeader)
			poHeader.POST("/id", POHeaderController.GePoHeaderByID)
			poHeader.POST("/delete", POHeaderController.Delete)

		}

		poDetail := api.Group("/po-detail")
		{
			poDetail.POST("/", PODetailController.SavePoDetail)
			poDetail.PUT("/", PODetailController.UpdatePoDetail)
			poDetail.POST("/id", PODetailController.GePoDetailByID)
			poDetail.POST("/delete", PODetailController.Delete)

		}

	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	goRouter.Router = router
}
