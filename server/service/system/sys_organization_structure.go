package system

import (
	"errors"
	"github.com/Luxii44/library/server/global"
	"github.com/Luxii44/library/server/model/system"
	"github.com/Luxii44/library/server/model/system/request"
	"github.com/Luxii44/library/server/model/system/response"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysOrganization
//@return: userInter system.SysOrganization, err error

type OrganizationService struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysOrganization

func (OrganizationService *OrganizationService) FindUserById(id int) (user *system.SysOrganization, err error) {
	var u system.SysOrganization
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysOrganization

func (OrganizationService *OrganizationService) FindUserByUuid(uuid string) (user *system.SysOrganization, err error) {
	var u system.SysOrganization
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateOrg
//@description: 新增组织架构
//@param: o system.SysOrganization
//@return: *system.SysOrganization,error
func (OrganizationService *OrganizationService) CreateOrg(o system.SysOrganization) (*system.SysOrganization, error) {
	o.UUID = uuid.NewV4()
	a := global.GVA_DB.Model(&system.SysOrganization{})
	var org system.SysOrganization
	if strings.Contains(o.OrganizationCode, "gro") {
		o.OrganizationFullName = o.OrganizationName
		o.TopOrganizationUUID = o.UUID
		o.CompanyUUID = nil
	} else {
		for {
			if org.UUID == o.TopOrganizationUUID {
				break //跳出结束循环
			} else {
				if errors.Is(a.Where("uuid = ?", &o.ParentUUID).First(&org).Error, gorm.ErrRecordNotFound) {
					break
					return nil, gorm.ErrRecordNotFound
				}
				o.OrganizationFullName = org.OrganizationName + "-" + o.OrganizationName
			}
		}
	}
	err := a.Create(&o).Error
	return &o, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetOrgList
//@description: 获取组织架构列表
//@param: info request.GetOrgList
//@return: list interface{}, total int64, err error
func (OrganizationService *OrganizationService) GetOrgList(info request.GetOrgList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysOrganization{})
	var orgList []response.Organization
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Table("sys_organization AS A").
		Select("A.*,C.organization_name AS company_name, D.organization_name AS top_organization_name").
		Joins("LEFT JOIN sys_organization AS B ON A.parent_uuid = B.uuid").
		Joins("LEFT JOIN sys_organization AS C ON A.company_uuid = C.uuid").
		Joins("LEFT JOIN sys_organization AS D ON A.top_organization_uuid = D.uuid").
		Limit(limit).Offset(offset).Find(&orgList).Error
	//err = db.Limit(limit).Offset(offset).Find(&orgList).Error
	return orgList, total, err
}
