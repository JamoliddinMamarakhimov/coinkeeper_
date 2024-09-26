package controllers

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllIncomes
// @Summary Get All Incomes
// @Security ApiKeyAuth
// @Tags incomes
// @Description get list of all income
// @ID get-all-incomes
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.Income
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/income [get]
func GetAllIncome(c *gin.Context) {
	query := c.Query("q")

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	income, err := service.GetAllIncome(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income": income})

	//income, err := service.GetAllIncome()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"income": income,
	//})
}

// GetIncomeByID
// @Summary Get Income By ID
// @Security ApiKeyAuth
// @Tags incomes
// @Description get Income by ID
// @ID get-income-by-id
// @Produce json
// @Param id path integer true "id of the income"
// @Success 200 {object} models.Income
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/incomes/{id} [get]
func GetIncomeByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	incomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	income, err := service.GetIncomeByID(userID, uint(incomeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, income)
}

// CreateIncome
// @Summary Create Income
// @Security ApiKeyAuth
// @Tags incomes
// @Description create new income
// @ID create-new-income
// @Accept json
// @Produce json
// @Param input body models.Income true "new income info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/incomes [post]
func CreateIncome(c *gin.Context) {
	var income models.Income
	if err := c.BindJSON(&income); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	income.UserID = userID
	if err := service.CreateIncome(income); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, defaultResponse{Message: "income created successfully"})
}

// UpdateIncome
// @Summary Update Income
// @Security ApiKeyAuth
// @Tags incomes
// @Description update existed income
// @ID update-income
// @Accept json
// @Produce json
// @Param id path integer true "id of the income"
// @Param input body models.Income true "income update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/incomes/{id} [put]
func UpdateIncome(c *gin.Context) {
	incomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var income models.Income
	if err = c.BindJSON(&income); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	income.ID = uint(incomeID)
	income.UserID = userID
	if err = service.UpdateIncome(income); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, defaultResponse{Message: "income updated successfully"})
}

// DeleteIncome
// @Summary Delete Income By ID
// @Security ApiKeyAuth
// @Tags incomes
// @Description delete income by ID
// @ID delete-income-by-id
// @Param id path integer true "id of the income"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/incomes/{id} [delete]
func DeleteIncome(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	incomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err = service.DeleteIncome(incomeID, uint(userID)); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, defaultResponse{Message: "income deleted successfully"})
}
