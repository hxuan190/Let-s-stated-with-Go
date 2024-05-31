package ginitem

import (
	"Project/common"
	"Project/modules/item/biz"
	"Project/modules/item/model"
	"Project/modules/item/storage/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateToDoItemsPresenter
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)

			return
		}
		store := mysql.NewStorage(db)
		business := biz.NewCreateItemBiz(store)
		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessRes(data.Id))
	}
}
