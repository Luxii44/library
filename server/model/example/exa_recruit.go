package example

import (
	"github.com/Luxii44/library/server/global"
	"github.com/Luxii44/library/server/model/system"
	"time"
)

type ExaRecruitPlatform struct {
	global.GVA_MODEL
	PlatformName   string `json:"platformName" gorm:"index;comment:平台名"`         // 平台名称
	PlatformJobUrl string `json:"platformJobUrl" gorm:"index;comment:平台工作列表url"` // 平台工作列表url
	//RecruitAccount ExaRecruitAccount `json:"recruitAccount" gorm:"foreignKey:Platform;references:PlatformName"`
}

func (ExaRecruitPlatform) TableName() string {
	return "exa_recruit_platform"
}

type ExaRecruitAccount struct {
	global.GVA_MODEL
	Account         string           `json:"account" gorm:"comment:账号"`                                   // 账号
	Platform        string           `json:"platform"`                                                    // 平台
	Type            int              `json:"type" gorm:"default:1;comment:1招聘者 2企业"`                      //1招聘者 2企业
	Owner           uint             `json:"owner" gorm:"comment:账号所属人"`                                  //账号所属人
	ExaRecruitPlans []ExaRecruitPlan `json:"exaRecruitPlan" gorm:"foreignKey:Account;references:Account"` //关联计划
	Enable          int              `json:"enable" gorm:"default:1;comment:账号是否被冻结 1正常 2冻结"`
	//Password string             `json:"password" gorm:"comment:密码"`                      // 密码              //账号是否被冻结 1正常 2冻结
}

func (ExaRecruitAccount) TableName() string {
	return "exa_recruit_account"
}

type ExaRecruitPlan struct {
	global.GVA_MODEL
	Query                     string `json:"query" gorm:"comment:关键字"`                     // 关键字
	City                      int64  `json:"city" gorm:"comment:城市"`                       // 城市
	CityName                  string `json:"cityName" gorm:"comment:城市名"`                  // 城市名
	MultiBusinessDistrict     string `json:"multiBusinessDistrict" gorm:"comment:区域"`      //区域
	MultiBusinessDistrictName string `json:"multiBusinessDistrictName" gorm:"comment:区域名"` //区域名
	Industry                  string `json:"industry" gorm:"comment:行业"`                   // 公司行业
	Experience                string `json:"experience" gorm:"comment:经验"`                 // 经验
	PayType                   string `json:"payType" gorm:"comment:经验"`                    // 经验
	PartTime                  string `json:"partTime" gorm:"comment:经验"`                   // 经验
	Degree                    string `json:"degree" gorm:"comment:经验"`                     // 经验
	Scale                     string `json:"scale" gorm:"comment:经验"`                      // 经验
	Stage                     string `json:"stage" gorm:"comment:经验"`                      // 经验
	Position                  int64  `json:"position" gorm:"comment:经验"`                   // 经验
	JobType                   string `json:"jobType" gorm:"comment:经验"`                    // 经验
	Salary                    string `json:"salary" gorm:"comment:经验"`                     // 经验
	MultiSubway               string `json:"multiSubway" gorm:"comment:经验"`                // 经验

	Delivery     int    `json:"delivery" gorm:"default:1;comment:是否投递简历 1是 2否"`  //是否投递简历 1是 2否
	DeliverCount int    `json:"deliverCount" gorm:"comment:投递数量"`                //投递数量
	Account      string `json:"account" gorm:"comment:关联账号"`                     //关联账号
	UrlPath      string `json:"urlPath" gorm:"type:text(500);comment:经验"`        // 经验
	Enable       int    `json:"enable" gorm:"default:1;comment:账号是否被冻结 1正常 2冻结"` //账号是否被冻结 1正常 2冻结
	//SysUser        system.SysUser    `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`              // 管理详情
}

func (ExaRecruitPlan) TableName() string {
	return "exa_recruit_plan"
}

type ExaJob struct {
	Position       string `json:"position" gorm:"comment:职位名称"`        // 职位名称
	Synopsis       string `json:"synopsis" gorm:"comment:职位简介"`        // 职位简介
	Address        string `json:"address" gorm:"comment:地点"`           // 地点
	CompanyName    string `json:"companyName" gorm:"comment:公司名称"`     //公司名称
	Industry       string `json:"industry" gorm:"comment:行业"`          // 公司行业
	CompanyAddress string `json:"companyAddress" gorm:"comment:公司地址"`  //公司地址
	Salary         string `json:"salary" gorm:"comment:薪酬"`            //薪酬
	Plan           string `json:"plan" gorm:"comment:关联计划"`            //关联计划
	UrlPath        string `json:"urlPath" gorm:"comment:url"`          // url主键
	Delivered      uint   `json:"delivered" gorm:"comment:是否投递：1是，2否"` // 是否投递：1是，2否
}

func (ExaJob) TableName() string {
	return "exa_job"
}

type ExaComponyBlack struct {
	global.GVA_MODEL
	ComponyName     string            `json:"componyName" gorm:"comment:公司名"`                          // 公司名
	City            string            `json:"city" gorm:"comment:城市"`                                  // 城市
	Area            string            `json:"area" gorm:"comment:区域"`                                  //区域
	Industry        string            `json:"industry" gorm:"comment:行业"`                              // 平台
	Hiredate        time.Time         `json:"hiredate" gorm:"comment:入职日期"`                            //入职日期
	ResignationDate time.Time         `json:"resignationDate" gorm:"comment:离职日期"`                     //离职日期
	LastSalary      int               `json:"intervalTime" gorm:"comment:间隔时间"`                        //间隔时间
	RecruitAccount  ExaRecruitAccount `json:"recruitAccount" gorm:"default:1;comment:账号是否被冻结 1正常 2冻结"` //账号是否被冻结 1正常 2冻结
	SysUser         system.SysUser    `json:"sysUser" form:"sysUser" gorm:"comment:关联用户"`              //关联用户
}

func (ExaComponyBlack) TableName() string {
	return "exa_compony_black"
}
