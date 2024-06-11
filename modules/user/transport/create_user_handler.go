package transport

import (
	"Project/common"
	"Project/modules/user/biz"
	"Project/modules/user/model"
	"Project/modules/user/storage/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.Register
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)

			return
		}
		store := mysql.NewStorage(db)
		business := biz.NewCreateUserBiz(store)
		if err := business.CreateNewUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessRes(data))
	}
}
