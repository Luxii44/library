package system

import (
	"fmt"
	"github.com/Luxii44/library/server/model/common/response"
	"github.com/Luxii44/library/server/model/system"
	systemReq "github.com/Luxii44/library/server/model/system/request"
	"github.com/gin-gonic/gin"
)

// CreateOrg
// @Tags     Base
// @Summary  新增组织架构
// @Produce   application/json
// @Param    data  body      system.SysOrganization                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /org/Create [post]
func (b *BaseApi) CreateOrg(c *gin.Context) {
	var l system.SysOrganization
	err := c.ShouldBindJSON(&l)
	fmt.Println(l)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	result, err := organizationService.CreateOrg(l)
	response.OkWithDetailed(result, "添加成功", c)
}

// GetOrgList
// @Tags     Base
// @Summary  获取组织架构列表
// @Produce   application/json
// @Param    data  body      systemReq.GetOrgList                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /org/getOrgList [get]
func (b *BaseApi) GetOrgList(c *gin.Context) {
	var pageInfo systemReq.GetOrgList
	if err := c.ShouldBind(&pageInfo); err != nil { // ShouldBind()会根据请求的Content-Type自行选择绑定器
		response.FailWithMessage("参数错误", c)
		return
	}
	//page := c.Query("page")
	//pageSize := c.Query("pageSize")
	list, total, err := organizationService.GetOrgList(pageInfo)
	if err == nil {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
