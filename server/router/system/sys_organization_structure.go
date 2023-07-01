package system

import (
	v1 "github.com/Luxii44/library/server/api/v1"
	"github.com/Luxii44/library/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrganizationRouter struct{}

func (s *OrganizationRouter) InitOrganizationRouter(Router *gin.RouterGroup) {
	OrganizationRouter := Router.Group("org").Use(middleware.OperationRecord())
	//OrganizationRouterWithoutRecord := Router.Group("org")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		OrganizationRouter.POST("createOrg", baseApi.CreateOrg)  // 新增组织架构
		OrganizationRouter.GET("getOrgList", baseApi.GetOrgList) // 获取组织架构列表
	}
	{
		//OrganizationRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		//OrganizationRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
	}
}
