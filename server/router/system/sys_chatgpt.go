package system

import (
	v1 "github.com/Luxii44/library/server/api/v1"
	"github.com/Luxii44/library/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChatGptRouter struct{}

func (s *ChatGptRouter) InitChatGptRouter(Router *gin.RouterGroup) {
	chatGptRouter := Router.Group("chatGpt").Use(middleware.OperationRecord())
	chatGptApi := v1.ApiGroupApp.SystemApiGroup.ChatGptApi
	{
		chatGptRouter.POST("createSK", chatGptApi.CreateSK)
		chatGptRouter.GET("getSK", chatGptApi.GetSK)
		chatGptRouter.DELETE("deleteSK", chatGptApi.DeleteSK)
		chatGptRouter.POST("getTable", chatGptApi.GetTable)
	}
}
