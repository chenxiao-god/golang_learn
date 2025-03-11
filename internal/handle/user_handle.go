package handle

import (
	"github.com/gin-gonic/gin"
	"gomod/internal/module"
	"gomod/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}
func (u *UserHandler) Creat(c *gin.Context) {
	var user module.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if error := u.userService.Create(&user); error != nil {
		c.JSON(400, gin.H{"code": 400, "error": error.Error()})
	} else {
		c.JSON(200, gin.H{"code": 200, "message": "创建成功", "data": &user})
	}

}
