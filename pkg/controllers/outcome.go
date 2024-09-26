package controllers

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllOutcome
// @Summary Get All outcome
// @Security ApiKeyAuth
// @Tags outcomes
// @Description get list of all outcome
// @ID get-all-outcome
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.Outcome
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/outcome [get]
func GetAllOutcome(c *gin.Context) {
	query := c.Query("q")

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	outcome, err := service.GetAllOutcome(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"outcome": outcome})

	//outcome, err := service.GetAllOutcome()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"outcome": outcome,
	//})
}

// GetOutcomeByID
// @Summary Get Outcome By ID
// @Security ApiKeyAuth
// @Tags outcomes
// @Description get outcome by ID
// @ID get-outcome-by-id
// @Produce json
// @Param id path integer true "id of the outcome"
// @Success 200 {object} models.Outcome
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/outcomes/{id} [get]
func GetOutcomeByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	outcomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	outcome, err := service.GetOutcomeByID(userID, uint(outcomeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, outcome)
}

// CreateOutcome
// @Summary Create Outcome
// @Security ApiKeyAuth
// @Tags outcomes
// @Description create new outcome
// @ID create-new-outcome
// @Accept json
// @Produce json
// @Param input body models.Outcome true "new outcome info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/outcomes [post]
func CreateOutcome(c *gin.Context) {
	var outcome models.Outcome
	if err := c.BindJSON(&outcome); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	outcome.UserID = userID
	if err := service.CreateOutcome(outcome); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, defaultResponse{Message: "outcome created successfully"})
}

// UpdateOutcome
// @Summary Update Outcome
// @Security ApiKeyAuth
// @Tags outcomes
// @Description update existed outcome
// @ID update-outcome
// @Accept json
// @Produce json
// @Param id path integer true "id of the outcome"
// @Param input body models.Outcome true "outcome update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/outcomes/{id} [put]
func UpdateOutcome(c *gin.Context) {
	outcomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var outcome models.Outcome
	if err = c.BindJSON(&outcome); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	outcome.ID = uint(outcomeID)

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	outcome.ID = uint(outcomeID)
	outcome.UserID = userID
	if err = service.UpdateOutcome(outcome); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, defaultResponse{Message: "outcome updated successfully"})
}

// DeleteOutcome
// @Summary Delete Outcome By ID
// @Security ApiKeyAuth
// @Tags outcomes
// @Description delete outcome by ID
// @ID delete-outcome-by-id
// @Param id path integer true "id of the outcome"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/outcomes/{id} [delete]
func DeleteOutcome(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	outcomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err = service.DeleteOutcome(outcomeID, uint(userID)); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, defaultResponse{Message: "outcome deleted successfully"})
}
