package example

import (
	"github.com/Luxii44/library/server/global"
	"github.com/Luxii44/library/server/model/common/response"
	"github.com/Luxii44/library/server/model/example"
	"github.com/Luxii44/library/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MarkdownApi struct{}

// CreateWS
// @Tags      ExaChat
// @Summary   创建聊天室
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaRecruitAccount  true  "消息"
// @Success   200   {object}  response.Response  "创建聊天室"
// @Router    /chat/ping [get]
func (e *MarkdownApi) SaveMarkdown(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = utils.GetUserID(c)
	customer.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = customerService.CreateExaCustomer(customer)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
