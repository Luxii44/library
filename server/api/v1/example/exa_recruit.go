package example

import (
	"encoding/json"
	"fmt"
	"github.com/Luxii44/library/server/global"
	"github.com/Luxii44/library/server/model/common/request"
	"github.com/Luxii44/library/server/model/common/response"
	"github.com/Luxii44/library/server/model/example"
	exaRequest "github.com/Luxii44/library/server/model/example/request"
	"github.com/Luxii44/library/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type RecruitApi struct{}

// AddRecruitAccount
// @Tags      ExaRecruit
// @Summary   添加招聘账号
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaRecruitAccount  true  "账号, 密码, 平台, 类型"
// @Success   200   {object}  response.Response{msg=string}  "添加招聘账号"
// @Router    /recruit/account [post]
func (e *RecruitApi) AddRecruitAccount(c *gin.Context) {
	var recruitAccount example.ExaRecruitAccount
	err := c.ShouldBindJSON(&recruitAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(recruitAccount, utils.RecruitAccountVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	recruitAccount.Owner = utils.GetUserAuthorityId(c)
	err = recruitService.AddRecruitAccount(recruitAccount)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// UpdateRecruitAccount
// @Tags      ExaRecruit
// @Summary   更新招聘账号
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaRecruitAccount  true  "账号, 密码, 平台, 类型"
// @Success   200   {object}  response.Response{msg=string}  "更新招聘账号"
// @Router    /recruit/account [put]
func (e *RecruitApi) UpdateRecruitAccount(c *gin.Context) {
	var recruitAccount example.ExaRecruitAccount
	err := c.ShouldBindJSON(&recruitAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//不知道下面这一部分校验什么
	err = utils.Verify(recruitAccount.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(recruitAccount, utils.RecruitAccountVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = recruitService.UpdateRecruitAccount(&recruitAccount)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteRecruitAccount
// @Tags      ExaRecruit
// @Summary   删除招聘账号
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaRecruitAccount  true  "账号, 密码, 平台, 类型"
// @Success   200   {object}  response.Response{msg=string}  "删除招聘账号"
// @Router    /recruit/account [delete]
func (e *RecruitApi) DeleteRecruitAccount(c *gin.Context) {
	var recruitAccount example.ExaRecruitAccount
	err := c.ShouldBindJSON(&recruitAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//不知道下面这一部分校验什么
	err = utils.Verify(recruitAccount.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(recruitAccount, utils.RecruitAccountVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = recruitService.DeleteRecruitAccount(&recruitAccount)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetRecruitAccountList
// @Tags      ExaRecruit
// @Summary   分页获取权限招聘账号列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限招聘账号列表,返回包括列表,总数"
// @Router    /recruit/recruitAccountList [get]
func (e *RecruitApi) GetRecruitAccountList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	recruitAccountList, total, err := recruitService.GetRecruitAccountList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:  recruitAccountList,
		Total: total,
	}, "获取成功", c)
}

// CreateRecruitPlan
// @Tags      ExaRecruit
// @Summary   创建招聘计划
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaRecruitPlan  true  "关键字, 城市, 区域, 行业, 是否投递简历, 开始时间, 间隔时间"
// @Success   200   {object}  response.Response{msg=string}  "创建招聘计划"
// @Router    /recruit/plan [post]
func (e *RecruitApi) CreateRecruitPlan(c *gin.Context) {
	var recruitPlan example.ExaRecruitPlan
	err := c.ShouldBindJSON(&recruitPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(recruitPlan, utils.RecruitAccountVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = recruitService.CreateRecruitPlan(recruitPlan)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// UpdateRecruitPlan
// @Tags      ExaRecruit
// @Summary   更新招聘计划
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaRecruitPlan  true  "关键字, 城市, 区域, 行业, 是否投递简历, 开始时间, 间隔时间"
// @Success   200   {object}  response.Response{msg=string}  "更新招聘计划"
// @Router    /recruit/plan [put]
func (e *RecruitApi) UpdateRecruitPlan(c *gin.Context) {
	var recruitPlan example.ExaRecruitPlan
	err := c.ShouldBindJSON(&recruitPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//不知道下面这一部分校验什么
	err = utils.Verify(recruitPlan.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(recruitPlan, utils.RecruitAccountVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = recruitService.UpdateRecruitPlan(&recruitPlan)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteRecruitPlan
// @Tags      ExaRecruit
// @Summary   删除招聘计划
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaRecruitPlan  true  "关键字, 城市, 区域, 行业, 是否投递简历, 开始时间, 间隔时间"
// @Success   200   {object}  response.Response{msg=string}  "删除招聘计划"
// @Router    /recruit/plan [delete]
func (e *RecruitApi) DeleteRecruitPlan(c *gin.Context) {
	var recruitPlan example.ExaRecruitPlan
	err := c.ShouldBindJSON(&recruitPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//不知道下面这一部分校验什么
	err = utils.Verify(recruitPlan.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(recruitPlan, utils.RecruitAccountVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = recruitService.DeleteRecruitPlan(&recruitPlan)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetRecruitPlanList
// @Tags      ExaRecruit
// @Summary   分页获取权限招聘账号列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限招聘计划列表,返回包括列表,总数"
// @Router    /recruit/recruitPlanList [get]
func (e *RecruitApi) GetRecruitPlanList(c *gin.Context) {
	var planPageInfo exaRequest.GetRecruitPlanList
	err := c.ShouldBindQuery(&planPageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//err = utils.Verify(planPageInfo, utils.PageInfoVerify)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	recruitPlanList, total, err := recruitService.GetRecruitPlanList(utils.GetUserAuthorityId(c), planPageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:  recruitPlanList,
		Total: total,
	}, "获取成功", c)
}

// GetPlanBaseData
// @Tags      ExaRecruit
// @Summary   获取基本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限招聘计划列表,返回包括列表,总数"
// @Router    /recruit/recruitPlanList [get]
func (e *RecruitApi) GetPlanBaseData(c *gin.Context) {
	url := "https://www.zhipin.com/wapi/zpgeek/businessDistrict.json?cityCode=" + c.Query("cityCode")
	// 创建 HTTP 客户端
	client := &http.Client{}
	// 创建 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		response.FailWithMessage("创建请求失败:"+err.Error(), c)
		fmt.Println("创建请求失败:", err)
		return
	}
	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		response.FailWithMessage("发送请求失败:"+err.Error(), c)
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.FailWithMessage("读取响应失败:"+err.Error(), c)
		fmt.Println("读取响应失败:", err)
		return
	}
	// 输出响应内容
	response.OkWithDetailed(body, "获取成功", c)
}

// GetBusinessDistrict
// @Tags      ExaRecruit
// @Summary   获取基本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限招聘计划列表,返回包括列表,总数"
// @Router    /recruit/businessDistrict [get]
func (e *RecruitApi) GetBusinessDistrict(c *gin.Context) {
	url := "https://www.zhipin.com/wapi/zpgeek/businessDistrict.json?cityCode=" + c.Query("cityCode")
	// 创建 HTTP 客户端
	client := &http.Client{}
	// 创建 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		response.FailWithMessage("创建请求失败:"+err.Error(), c)
		fmt.Println("创建请求失败:", err)
		return
	}
	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		response.FailWithMessage("发送请求失败:"+err.Error(), c)
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.FailWithMessage("读取响应失败:"+err.Error(), c)
		fmt.Println("读取响应失败:", err)
		return
	}
	var v interface{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return
	}
	data := v.(map[string]interface{})
	m1 := data["zpData"].(map[string]interface{})
	//fmt.Println(m1["businessDistrict"])
	// 输出响应内容
	response.OkWithDetailed(m1["businessDistrict"], "获取成功", c)
}
