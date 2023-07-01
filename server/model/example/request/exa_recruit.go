package request

import (
	"github.com/Luxii44/library/server/model/common/request"
	"time"
)

// Get Recruit Plan List structure
type GetRecruitPlanList struct {
	request.PageInfo
	Account *string `json:"account"`
}

// Get Recruit Plan List structure
type RecruitPlan struct {
	Query                     string    `json:"query" gorm:"comment:关键字"`                       // 关键字
	City                      string    `json:"city" gorm:"comment:城市"`                         // 城市
	CityName                  string    `json:"cityName" gorm:"comment:城市名"`                    // 城市名
	MultiBusinessDistrict     string    `json:"multiBusinessDistrict" gorm:"comment:区域"`        //区域
	MultiBusinessDistrictName string    `json:"multiBusinessDistrictName" gorm:"comment:区域名"`   //区域名
	Industry                  string    `json:"industry" gorm:"comment:行业"`                     // 公司行业
	Experience                string    `json:"experience" gorm:"comment:经验"`                   // 经验
	PayType                   string    `json:"payType" gorm:"comment:经验"`                      // 经验
	PartTime                  string    `json:"partTime" gorm:"comment:经验"`                     // 经验
	Degree                    string    `json:"degree" gorm:"comment:经验"`                       // 经验
	Scale                     string    `json:"scale" gorm:"comment:经验"`                        // 经验
	Stage                     string    `json:"stage" gorm:"comment:经验"`                        // 经验
	Position                  string    `json:"position" gorm:"comment:经验"`                     // 经验
	JobType                   string    `json:"jobType" gorm:"comment:经验"`                      // 经验
	Salary                    string    `json:"salary" gorm:"comment:经验"`                       // 经验
	MultiSubway               string    `json:"multiSubway" gorm:"comment:经验"`                  // 经验
	Delivery                  int       `json:"delivery" gorm:"default:1;comment:是否投递简历 1是 2否"` //是否投递简历 1是 2否
	BeginTime                 time.Time `json:"beginTime" gorm:"comment:开始时间"`                  //开始时间
	IntervalTime              int       `json:"intervalTime" gorm:"comment:间隔时间"`               //间隔时间
	//Account      string    `json:"account" gorm:"comment:关联账号"`                     //关联账号
	Enable int `json:"enable" gorm:"default:1;comment:账号是否被冻结 1正常 2冻结"` //账号是否被冻结 1正常 2冻结
	//SysUser        system.SysUser    `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`              // 管理详情
}
