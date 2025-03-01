package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	db "gitlab.com/alexandrudeac/minibank/db/sqlc"
)

type createAccountParams struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=RON EUR USD"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}
	acc, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	})
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}
	ctx.JSON(201, acc)
}

type getAccountParams struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}
	acc, err := server.store.GetAccount(ctx, req.Id)
	if errors.Is(err, db.ErrRecordNotFound) {
		ctx.JSON(404, errorResponse(err))
		return
	}
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}
	ctx.JSON(200, acc)
}

type listAccountsParams struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (server *Server) listAccounts(ctx *gin.Context) {

	var req listAccountsParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	accs, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
		Owner:  "admin",
	})
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}
	ctx.JSON(200, gin.H{"accounts": accs})
}
