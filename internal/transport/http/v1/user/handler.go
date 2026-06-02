package user

import (
	"net/http"
	"strconv"

	userv1 "github.com/buqiuwenda/gin-template/api/gen/go/v1/user"
	appuser "github.com/buqiuwenda/gin-template/internal/application/user"
	"github.com/buqiuwenda/gin-template/internal/meta"
	"github.com/gin-gonic/gin"
)

// Handler 传输层：Gin 路由 + proto 契约 DTO
type Handler struct {
	svc *appuser.Service
}

func NewHandler(svc *appuser.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(r gin.IRoutes) {
	r.POST("/users", h.CreateUser)
	r.GET("/users/:id", h.GetUser)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req userv1.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, meta.Fail(400, err.Error()))
		return
	}
	u, err := h.svc.CreateUser(c.Request.Context(), req.Username, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, meta.Fail(500, err.Error()))
		return
	}
	c.JSON(http.StatusOK, meta.OK(&userv1.CreateUserReply{
		User: &userv1.User{Id: u.ID, Username: u.Username, Email: u.Email},
	}))
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, meta.Fail(400, "invalid id"))
		return
	}
	u, err := h.svc.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, meta.Fail(500, err.Error()))
		return
	}
	c.JSON(http.StatusOK, meta.OK(&userv1.GetUserReply{
		User: &userv1.User{Id: u.ID, Username: u.Username, Email: u.Email},
	}))
}
