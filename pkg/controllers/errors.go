package controllers

import (
	"coinkeeper/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

func handleError(c *gin.Context, err error) {
	if errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	} else if errors.Is(err, errs.ErrRecordNotFound) ||
		errors.Is(err, errs.ErrOperationNotFound) {
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	} else {
		c.JSON(http.StatusInternalServerError, newErrorResponse(errs.ErrSomethingWentWrong.Error()))
	}
}
