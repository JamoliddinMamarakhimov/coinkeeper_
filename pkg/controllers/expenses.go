package controllers

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllExpenses
// @Summary Get All Expenses
// @Security ApiKeyAuth
// @Tags expenses
// @Description get list of all expense
// @ID get-all-expenses
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.Expense
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/expense [get]
func GetAllExpenses(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	expenses, err := service.GetAllExpenses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}

// GetExpenseByID
// @Summary Get Expense By ID
// @Security ApiKeyAuth
// @Tags expenses
// @Description get expense by ID
// @ID get-expense-by-id
// @Produce json
// @Param id path integer true "id of the expense"
// @Success 200 {object} models.Expense
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/expenses/{id} [get]
func GetExpenseByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	expenseID, err := strconv.Atoi(c.Param("expenseID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expense ID"})
		return
	}

	expense, err := service.GetExpenseByID(userID, uint(expenseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expense)
}

// CreateExpense
// @Summary Create Expense
// @Security ApiKeyAuth
// @Tags expenses
// @Description create new expense
// @ID create-new-expense
// @Accept json
// @Produce json
// @Param input body models.Expense true "new expense info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/expenses [post]
func CreateExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.BindJSON(&expense); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	expense.UserID = userID
	if err := service.CreateExpense(expense); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "expense created successfully"})
}

// UpdateExpense
// @Summary Update Expense
// @Security ApiKeyAuth
// @Tags expenses
// @Description update existed expense
// @ID update-expense
// @Accept json
// @Produce json
// @Param id path integer true "id of the expense"
// @Param input body models.Expense true "expense update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/expenses/{id} [put]
func UpdateExpense(c *gin.Context) {
	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var expense models.Expense
	if err = c.BindJSON(&expense); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	expense.ID = uint(expenseID)
	expense.UserID = userID
	if err = service.UpdateExpense(expense); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "expense updated successfully"})
}

// DeleteExpense
// @Summary Delete Expense By ID
// @Security ApiKeyAuth
// @Tags expenses
// @Description delete expense by ID
// @ID delete-expense-by-id
// @Param id path integer true "id of the expense"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/expenses/{id} [delete]
func DeleteExpense(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err = service.DeleteExpense(uint(expenseID), userID); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "expense deleted successfully"})
}
