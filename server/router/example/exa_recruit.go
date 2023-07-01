package example

import (
	v1 "github.com/Luxii44/library/server/api/v1"
	"github.com/Luxii44/library/server/middleware"
	"github.com/gin-gonic/gin"
)

type RecruitRouter struct{}

func (e *RecruitRouter) InitRecruitRouter(Router *gin.RouterGroup) {
	recruitRouter := Router.Group("recruit").Use(middleware.OperationRecord())
	//recruitRouterWithoutRecord := Router.Group("recruit")
	exaRecruitApi := v1.ApiGroupApp.ExampleApiGroup.RecruitApi
	{
		recruitRouter.POST("account", exaRecruitApi.AddRecruitAccount)      // 创建招聘账号
		recruitRouter.PUT("account", exaRecruitApi.UpdateRecruitAccount)    // 更新招聘账号
		recruitRouter.DELETE("account", exaRecruitApi.DeleteRecruitAccount) // 删除招聘账号
		//recruitRouter.GET("account", exaRecruitApi.GetRecruitAccount)         // 获取单一招聘账号信息
		recruitRouter.GET("accountList", exaRecruitApi.GetRecruitAccountList) // 获取招聘账号列表

		recruitRouter.POST("plan", exaRecruitApi.CreateRecruitPlan)              // 创建招聘计划
		recruitRouter.PUT("plan", exaRecruitApi.UpdateRecruitPlan)               // 更新招聘计划
		recruitRouter.DELETE("plan", exaRecruitApi.DeleteRecruitPlan)            // 删除招聘计划
		recruitRouter.GET("planList", exaRecruitApi.GetRecruitPlanList)          // 获取招聘账号列表
		recruitRouter.GET("basedata", exaRecruitApi.GetPlanBaseData)             // 创建招聘计划
		recruitRouter.GET("businessDistrict", exaRecruitApi.GetBusinessDistrict) // 获取对应城市的地区
	}
	{
		//recruitRouterWithoutRecord.GET("account", exaRecruitApi.GetExaCustomer)          // 获取单一招聘账号信息
		//recruitRouterWithoutRecord.GET("customerList", exaRecruitApi.GetExaCustomerList) // 获取招聘账号列表
	}
}
