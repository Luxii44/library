package example

import (
	"fmt"
	"github.com/Luxii44/library/server/global"
	"github.com/Luxii44/library/server/model/common/response"
	exaRequest "github.com/Luxii44/library/server/model/example/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MarkdownApi struct{}

// SaveMarkdown
// @Tags      ExaMarkdown
// @Summary   保存markdown文件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      exaRequest.MarkdownResData  true  "消息"
// @Success   200   {object}  response.Response  "保存markdown文件"
// @Router    /markdown/saveMarkdown [post]
func (e *MarkdownApi) SaveMarkdown(c *gin.Context) {
	var mdData exaRequest.MarkdownResData
	err := c.ShouldBindJSON(&mdData)
	fmt.Println(mdData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = markdownService.Save(mdData)
	if err != nil {
		global.GVA_LOG.Error("保存失败!", zap.Error(err))
		response.FailWithMessage("保存失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}
