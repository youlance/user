package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/youlance/user/db/sqlc"
)

type createUserFollowerRequest struct {
	FollowerID string `json:"follower_id" binding:"required,alphanum"`
	FolloweeID string `json:"followee_id" binding:"required,alphanum"`
}

func (server *Server) CreateUserFollower(ctx *gin.Context) {
	var req createUserFollowerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFollowerParams{
		FollowerID: req.FollowerID,
		FolloweeID: req.FolloweeID,
	}

	userFollower, err := server.db.CreateFollower(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userFollower)
}

type listFollowersRequest struct {
	FolloweeID string `json:"followee_id" binding:"required"`
	PageID     int32  `json:"page_id" binding:"required,min=1"`
	PageSize   int32  `json:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListFollowers(ctx *gin.Context) {
	var req listFollowersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListFollowersParams{
		FolloweeID: req.FolloweeID,
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
	}

	followers, err := server.db.ListFollowers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, followers)
}

type listFolloweesRequest struct {
	FollowerID string `json:"follower_id" binding:"required"`
	PageID     int32  `json:"page_id" binding:"required,min=1"`
	PageSize   int32  `json:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListFollowees(ctx *gin.Context) {
	var req listFolloweesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListFolloweesParams{
		FollowerID: req.FollowerID,
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
	}

	followers, err := server.db.ListFollowees(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, followers)
}

type deleteFollowerRequest struct {
	FollowerID string `json:"follower_id" binding:"required"`
	FolloweeID string `json:"followee_id" binding:"required"`
}

func (server *Server) DeleteFollower(ctx *gin.Context) {
	var req deleteFollowerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteFollowerParams{
		FollowerID: req.FollowerID,
		FolloweeID: req.FolloweeID,
	}

	err := server.db.DeleteFollower(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
