package controllers

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllCards
// @Summary Get All Cards
// @Security ApiKeyAuth
// @Tags cards
// @Description get list of all card
// @ID get-all-cards
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.Card
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/card [get]
func GetAllCards(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	cards, err := service.GetAllCards(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cards": cards})
}

// GetCardByID
// @Summary Get Card By ID
// @Security ApiKeyAuth
// @Tags cards
// @Description get card by ID
// @ID get-card-by-id
// @Produce json
// @Param id path integer true "id of the card"
// @Success 200 {object} models.Card
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/cards/{id} [get]
func GetCardByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	cardID, err := strconv.Atoi(c.Param("cardID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid card ID"})
		return
	}
	card, err := service.GetCardByID(userID, uint(cardID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, card)
}

// CreateCard
// @Summary Create Card
// @Security ApiKeyAuth
// @Tags cards
// @Description create new card
// @ID create-new-card
// @Accept json
// @Produce json
// @Param input body models.Card true "new card info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/cards [post]
func CreateCard(c *gin.Context) {
	var card models.Card

	if err := c.BindJSON(&card); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	card.UserID = userID // Устанавливаем ID пользователя

	if err := service.CreateCard(card); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, defaultResponse{Message: "Card created successfully"})
}

// UpdateCardBalance
// @Summary Update Card Balance
// @Security ApiKeyAuth
// @Tags cards
// @Description Update the balance of an existing card by its ID
// @ID update-card-balance
// @Accept json
// @Produce json
// @Param id path integer true "ID of the card"
// @Param input body struct { Balance float32 } true "New card balance"
// @Success 200 {object} defaultResponse
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 404 {object} ErrorResponse "Card not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Failure default {object} ErrorResponse

func UpdateCardBalance(c *gin.Context) {
	var updateRequest struct {
		CardID uint    `json:"card_id"`
		Amount float32 `json:"amount"` // Сумма для пополнения
	}

	if err := c.BindJSON(&updateRequest); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.UpdateCardBalance(updateRequest.CardID, updateRequest.Amount); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, defaultResponse{Message: "Card balance updated successfully"})
}

// DeleteCard
// @Summary Delete Card By ID
// @Security ApiKeyAuth
// @Tags cards
// @Description delete card by ID
// @ID delete-card-by-id
// @Param id path integer true "id of the card"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/cards/{id} [delete]
func DeleteCard(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id")) // Получаем ID карты из параметров URL
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx) // Получаем userID из контекста
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.DeleteCard(uint(cardID), userID); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, defaultResponse{Message: "Card deleted successfully"})
}
