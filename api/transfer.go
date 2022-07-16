package api

import (
	"fmt"
	db "go-simple-bank/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) transfer(ctx *gin.Context) {
	var transferRequest TransferRequest
	if err := ctx.ShouldBindJSON(&transferRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.isCurrencyValid(transferRequest.FromAccountID, transferRequest.ToAccountID, transferRequest.Currency, ctx) {
		return
	}

	result, err := server.store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: transferRequest.FromAccountID,
		ToAccountID:   transferRequest.ToAccountID,
		Amount:        transferRequest.Amount,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) isCurrencyValid(FromAccountID int64, toAccountID int64, currency string, ctx *gin.Context) bool {
	fromAccount, err := server.store.GetAccount(ctx, FromAccountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return false
	}
	toAccount, err := server.store.GetAccount(ctx, FromAccountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return false
	}

	if fromAccount.Currency != currency || toAccount.Currency != currency {

		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid currency. FromAccount: %s - ToAccount: %s - request currency: %s", fromAccount.Currency, toAccount.Currency, currency)))
		return false

	}
	return true
}
