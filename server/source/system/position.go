package system

import (
	"context"
	sysModel "github.com/Luxii44/library/server/model/system"
	"github.com/Luxii44/library/server/service/system"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const initOrderPosition = initOrderOrg + 1

type initPosition struct{}

// auto run
func init() {
	system.RegisterInit(initOrderPosition, &initPosition{})
}

func (i *initPosition) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysPosition{}, &sysModel.SysUser{})
}

func (i *initPosition) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysPosition{})
}

func (i initPosition) InitializerName() string {
	return sysModel.SysPosition{}.TableName()
}

func (i *initPosition) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	Nulluuid, err := uuid.FromString("00000000-0000-0000-0000-000000000000")
	Topuuid := uuid.NewV4()

	entities := []sysModel.SysPosition{
		{
			UUID:             Topuuid,
			PositionName:     "董事长",
			OrganizationUUID: Nulluuid,
			ParentUUID:       Nulluuid,
			CompanyUUID:      Nulluuid,
		},
		{
			UUID:             Topuuid,
			PositionName:     "财务总监",
			OrganizationUUID: Nulluuid,
			ParentUUID:       Nulluuid,
			CompanyUUID:      Nulluuid,
		},
		{
			UUID:             Topuuid,
			PositionName:     "财务经理",
			OrganizationUUID: Nulluuid,
			ParentUUID:       Nulluuid,
			CompanyUUID:      Nulluuid,
		},
		{
			UUID:             Topuuid,
			PositionName:     "人事总监",
			OrganizationUUID: Nulluuid,
			ParentUUID:       Nulluuid,
			CompanyUUID:      Nulluuid,
		},
		{
			UUID:             Topuuid,
			PositionName:     "技术总监",
			OrganizationUUID: Nulluuid,
			ParentUUID:       Nulluuid,
			CompanyUUID:      Nulluuid,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysPosition{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	userEntities, ok := ctx.Value(initUser{}.InitializerName()).([]sysModel.SysUser)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [岗位-用户] 关联失败, 未找到用户表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("Multiuser").Replace(userEntities); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("Multiuser").Replace(userEntities[:1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initPosition) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysPosition
	if errors.Is(db.Where("positionName = ?", "董事长").
		Preload("Multiuser").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Multiuser) > 0 && record.Multiuser[0].AuthorityId == 1
}
