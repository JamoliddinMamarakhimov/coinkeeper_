package controllers

import (
	"coinkeeper/configs"
	_ "coinkeeper/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingPong)

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	apiG := r.Group("/api", checkUserAuthentication)

	incomeG := apiG.Group("/income")
	{
		incomeG.GET("", GetAllIncome)
		incomeG.POST("", CreateIncome)
		incomeG.GET("/:id", GetIncomeByID)
		incomeG.PUT("/:id", UpdateIncome)
		incomeG.DELETE("/:id", DeleteIncome)
	}

	outcomeG := apiG.Group("/outcome")
	{
		outcomeG.GET("", GetAllOutcome)
		outcomeG.POST("", CreateOutcome)
		outcomeG.GET("/:id", GetOutcomeByID)
		outcomeG.PUT("/:id", UpdateOutcome)
		outcomeG.DELETE("/:id", DeleteOutcome)
	}

	expenseG := apiG.Group("/expenses")
	{
		expenseG.GET("", GetAllExpenses)
		expenseG.POST("", CreateExpense)
		expenseG.GET("/:id", GetExpenseByID)
		expenseG.PUT("/:id", UpdateExpense)
		expenseG.DELETE("/:id", DeleteExpense)
	}

	cardG := apiG.Group("/cards")
	{
		cardG.GET("", GetAllCards)
		cardG.POST("", CreateCard)
		cardG.GET("/:id", GetCardByID)
		cardG.PUT("/:id", UpdateCardBalance)
		cardG.DELETE("/:id", DeleteCard)
	}

	return r
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
