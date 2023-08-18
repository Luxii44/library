package example

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Luxii44/library/server/global"
	"github.com/Luxii44/library/server/model/common/request"
	"github.com/Luxii44/library/server/model/example"
	exaRequest "github.com/Luxii44/library/server/model/example/request"
	"github.com/Luxii44/library/server/utils/timer"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ChatService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddRecruitAccount
//@description: 创建账号
//@param: e model.ExaRecruitAccount
//@return: err error

func (exa *ChatService) ReceiveMessages(e example.ChatMessage) (err error) {
	fmt.Println(e)
	//如果数据中有room的字段，则是群聊，如果没有，在redis和数据库里面寻找
	//db := global.GVA_DB.Model(&example.ChatMessage{})
	switch e.MessageType {
	case 0:
		//db.Create(&e)
		fmt.Println("文本为：" + e.MessageContent)
	case 1:
		fmt.Println("文本为：" + e.MessageContent)
	default:
		fmt.Println(e)
	}
	return err
}

func (exa *ChatService) StartTimer(f func()) {
	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-t1.C:
			println("5s timer")
			t1.Reset(time.Second * 5)

		case <-t2.C:
			//go f()
			println("10s timer")
			t2.Reset(time.Second * 10)
		}
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRecruitAccount
//@description: 删除账号
//@param: e model.ExaRecruitAccount
//@return: err error

func (exa *ChatService) DeleteRecruitAccount(e *example.ExaRecruitAccount) (err error) {
	e.Enable = 2
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRecruitAccount
//@description: 更新账号
//@param: e *model.ExaRecruitAccount
//@return: err error

func (exa *ChatService) UpdateRecruitAccount(e *example.ExaRecruitAccount) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: customer model.ExaRecruitAccount, err error

func (exa *ChatService) GetExaCustomer(id uint) (customer example.ExaRecruitAccount, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRecruitAccountList
//@description: 分页获取客户列表,根据权限判断是否可以修改禁用和查看密码
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *ChatService) GetRecruitAccountList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//判断是否有权限
	var RecruitAccount []example.ExaRecruitAccount
	db := global.GVA_DB.Model(&example.ExaRecruitAccount{})
	err = db.Limit(limit).Offset(offset).Where("owner = ?", sysUserAuthorityID).Find(&RecruitAccount).Error
	total = int64(len(RecruitAccount))
	return RecruitAccount, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRecruitPlan
//@description: 创建招聘计划
//@param: e model.ExaRecruitPlan
//@return: err error

func (exa *ChatService) CreateRecruitPlan(e example.ExaRecruitPlan) (err error) {
	//根据账号表关联，从平台表获取工作列表url
	u, err := url.Parse("https://www.zhipin.com/wapi/zpgeek/search/joblist.json")
	if err != nil {
		return
	}
	//拼接url
	v := url.Values{} //拼接query参数
	v.Add("scene", "1")
	v.Add("query", e.Query)
	v.Add("city", utils.ToString(e.City))
	v.Add("experience", e.Experience)
	v.Add("payType", e.PayType)
	v.Add("partTime", e.PartTime)
	v.Add("degree", e.Degree)
	v.Add("industry", e.Industry)
	v.Add("scale", e.Scale)
	v.Add("stage", e.Stage)
	v.Add("position", utils.ToString(e.Position))
	v.Add("jobType", e.JobType)
	v.Add("salary", e.Salary)
	v.Add("multiBusinessDistrict", e.MultiBusinessDistrict)
	v.Add("multiSubway", e.MultiSubway)
	v.Add("page", "1")
	v.Add("pageSize", utils.ToString(e.DeliverCount))
	u.RawQuery = v.Encode()
	e.UrlPath = u.String()
	err = global.GVA_DB.Create(&e).Error
	if err == nil {
		for i := 1; i < 8; i++ {
			v.Add("page", strconv.Itoa(i))
			u.RawQuery = v.Encode()
			// 创建 HTTP 客户端
			client := &http.Client{}
			var req *http.Request
			var resp *http.Response
			var body []byte
			// 创建 GET 请求
			req, err = http.NewRequest("GET", e.UrlPath, nil)
			if err != nil {
				fmt.Println("创建请求失败:", err)
				return errors.New("创建请求失败:" + err.Error())
			}
			// 发送请求并获取响应(需要登陆后的cookie)
			resp, err = client.Do(req)
			if err != nil {
				fmt.Println("发送请求失败:", err)
				return errors.New("发送请求失败:" + err.Error())
			}
			defer resp.Body.Close()

			// 读取响应内容
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("读取响应失败:", err)
				return errors.New("读取响应失败:" + err.Error())
			}
			var v interface{}
			err = json.Unmarshal(body, &v)
			if err != nil {
				return
			}
			data := v.(map[string]interface{})
			m1 := data["zpData"].(map[string]interface{})
			fmt.Println(m1, 123)
			time.Sleep(5)
			//并将列表写入jobs列表
		}
	}
	if global.GVA_Timer == nil {
		global.GVA_Timer = timer.NewTimerTask()
	}
	spec := "0 */30 8-" + strconv.Itoa(e.DeliverCount/2) + " * *"
	_, err = global.GVA_Timer.AddTaskByFunc("DeliverResume", spec, ExecutePlan)
	if err != nil {
		fmt.Println("添加失败:", err)
		return
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRecruitPlan
//@description: 删除招聘计划
//@param: e model.ExaRecruitPlan
//@return: err error

func (exa *ChatService) DeleteRecruitPlan(e *example.ExaRecruitPlan) (err error) {
	e.Enable = 2
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRecruitPlan
//@description: 更新账号
//@param: e *model.ExaRecruitPlan
//@return: err error

func (exa *ChatService) UpdateRecruitPlan(e *example.ExaRecruitPlan) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRecruitAccountList
//@description: 分页获取客户列表,根据权限判断是否可以修改禁用和查看密码
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *ChatService) GetRecruitPlanList(sysUserAuthorityID uint, info exaRequest.GetRecruitPlanList) (list interface{}, total int64, err error) {
	var RecruitPlan []example.ExaRecruitPlan
	db := global.GVA_DB.Model(&example.ExaRecruitPlan{})
	if info.Account != nil {
		err = db.Where("owner = ?", sysUserAuthorityID).Find(&RecruitPlan).Error
		total = int64(len(RecruitPlan))
		return RecruitPlan, total, err
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = db.Limit(limit).Offset(offset).Where("account = ?", sysUserAuthorityID).Find(&RecruitPlan).Error
	total = int64(len(RecruitPlan))
	return RecruitPlan, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetReply
//@description: 从接口根据职位url获取回复的数据并写入数据库
//@param:
//@return: isSuccess string

func (exa *ChatService) GetReply() (isSuccess string) {
	//从接口获取回复列表,再从数据库查出来,再更新
	var exaJobs []example.ExaJob
	db := global.GVA_DB.Model(&example.ExaJob{})
	db.Where("").Find(&exaJobs)
	//db.Update("")
	return "完成"
	//做参考、还需考虑多进程并发的问题、还有执行挂了的问题
	//timer := time.NewTimer(durationUntilSix)
	//
	//for {
	//	<-timer.C
	//	// 执行你的任务
	//	fmt.Println("定时任务执行")
	//
	//	// 重置定时器为30分钟后
	//	timer.Reset(30 * time.Minute)
	//}
}

//数据校验
//https://apm-fe.zhipin.com/wapi/zpApm/actionLog/fe/common.json

//https://www.zhipin.com/wapi/zppassport/get/zpToken
//V1RNsjFu302ldpVtRvxxUYIC-17TrXxi0~
//刷新token-url 先后
//https: //www.zhipin.com/wapi/zppassport/get/zpToken?v2=1688093969514
//V1RNsjFu302ldpVtRvxxUYIC-17TrWxio~
//https://www.zhipin.com/wapi/zppassport/get/zpToken?v=1688094450203
//V1RNsjFu302ldpVtRvxxUYIC-55TrRxyo~

//获取工作列表
//https://www.zhipin.com/wapi/zpgeek/search/joblist.json?scene=1&query=golang&city=101280600&experience=&payType=&partTime=&degree=&industry=&scale=&stage=&position=&jobType=&salary=&multiBusinessDistrict=&multiSubway=&page=1&pageSize=30
//根据工作列表中的lid、securityid、sessionid获取html（内含职位id）
//https://www.zhipin.com/job_detail/447ab0adc848c8fa1XZz2921F1RS.html?lid=8wgma4JiI9Q.search.1&securityId=6nb2uuBPNDB6i-M1bQgTV6HZcKvSCi0Q6-6brayMA7NAstHxV-L6rNG57jcujgvgBfGkLVxVimmfMcf7CcSgLcsaRu155Ry0UOl6mYrEFFoyUqS48L_l-TxmatLo6m_3fcSmDmIg0po2pyK_TWXvEf1PSfv-maMFUAADEHBjMpNFy1U~&sessionId=
//发起投递
//https://www.zhipin.com/wapi/zpgeek/friend/add.json?securityId=oE3sl8eO7k8MZ-e1TGLQiJ5pS5xGVRVFF61LKO5luwHEvl8wvkBNWRXtxskbPoGe26KCnwuaTlQ8xJCMLaESNf3BSmqTMh6yUh-fJdDw_E7GoExmOwICHb4_XVjQ6wpIXMThd2VgNGjpVzLgcYaRUsbG2MjFvreWYKTAfPg3T6itPdk~&jobId=447ab0adc848c8fa1XZz2921F1RS&lid=8wgma4JiI9Q.search.1

//30天内联系人
//https://www.zhipin.com/wapi/zprelation/friend/geekFilterByLabel?labelId=0
//根据联系人聊天列表
//https://www.zhipin.com/wapi/zprelation/friend/getGeekFriendList.json?friendIds=64341212,21111561
