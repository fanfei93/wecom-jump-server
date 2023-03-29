package group

import (
	"github.com/gin-gonic/gin"
	"wecom-jump-server/service/group"
)

type CreateGroupRequest struct {
	GroupName string   `json:"groupName" binding:"required"`
	OwnerID   string   `json:"ownerID" binding:"required"`
	UserList  []string `json:"userList" binding:"required"`
}

func HandleCreateGroup(c *gin.Context) {
	req := new(CreateGroupRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	chatID, err := group.CreateGroup(req.GroupName, req.OwnerID, req.UserList)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    chatID,
	})
	return
}
