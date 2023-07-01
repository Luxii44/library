package system

import (
	"context"
	sysModel "github.com/Luxii44/library/server/model/system"
	"github.com/Luxii44/library/server/service/system"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const initOrderOrg = initOrderUser + 1

type initOrganization struct{}

// auto run
func init() {
	system.RegisterInit(initOrderOrg, &initOrganization{})
}

//待优化
func (i *initOrganization) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysOrganization{})
}

func (i *initOrganization) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysOrganization{})
}

func (i initOrganization) InitializerName() string {
	return sysModel.SysOrganization{}.TableName()
}

func (i *initOrganization) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	Nulluuid, err := uuid.FromString("00000000-0000-0000-0000-000000000000")
	Topuuid := uuid.NewV4()
	Componyuuid := uuid.NewV4()
	Departmentuuid := uuid.NewV4()

	entities := []sysModel.SysOrganization{
		{
			UUID:                 Topuuid,
			OrganizationCode:     "",
			OrganizationName:     "强盛集团",
			OrganizationFullName: "强盛集团",
			ParentUUID:           Nulluuid,
			CompanyUUID:          nil,
			TopOrganizationUUID:  Nulluuid,
		},
		{
			UUID:                 Componyuuid,
			OrganizationName:     "XXX公司",
			OrganizationFullName: "强盛集团-XXX公司",
			ParentUUID:           Topuuid,
			CompanyUUID:          &Componyuuid,
			TopOrganizationUUID:  Topuuid,
		},
		{
			UUID:                 Departmentuuid,
			OrganizationName:     "管理部门",
			OrganizationFullName: "强盛集团-XXX公司-管理部门",
			ParentUUID:           Componyuuid,
			CompanyUUID:          &Componyuuid,
			TopOrganizationUUID:  Topuuid,
		},
		{
			UUID:                 Departmentuuid,
			OrganizationName:     "财务部门",
			OrganizationFullName: "强盛集团-XXX公司-财务部门",
			ParentUUID:           Componyuuid,
			CompanyUUID:          &Componyuuid,
			TopOrganizationUUID:  Topuuid,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysOrganization{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, err
}

func (i *initOrganization) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysOrganization
	if errors.Is(db.Where("organizationName = ?", "XXX公司").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
