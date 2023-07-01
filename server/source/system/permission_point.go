package system

import (
	"context"
	sysModel "github.com/Luxii44/library/server/model/system"
	"github.com/Luxii44/library/server/service/system"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const initOrderPermission = initOrderPosition + 1

type initPermission struct{}

// auto run
func init() {
	system.RegisterInit(initOrderPermission, &initPermission{})
}

func (i *initPermission) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysPermissionPoint{})
}

func (i *initPermission) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysPermissionPoint{})
}

func (i initPermission) InitializerName() string {
	return sysModel.SysPermissionPoint{}.TableName()
}

func (i *initPermission) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	Nulluuid, err := uuid.FromString("00000000-0000-0000-0000-000000000000")

	entities := []sysModel.SysPermissionPoint{
		{
			UUID:           Nulluuid,
			PermissionCode: "per-001",
			PermissionName: "新增",
		},
		{
			UUID:           Nulluuid,
			PermissionCode: "per-002",
			PermissionName: "修改",
		},
		{
			UUID:           Nulluuid,
			PermissionCode: "per-003",
			PermissionName: "查看",
		},
		{
			UUID:           Nulluuid,
			PermissionCode: "per-004",
			PermissionName: "删除",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysPermissionPoint{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	apiEntities, ok := ctx.Value(initApi{}.InitializerName()).([]sysModel.SysApi)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [权限点-API] 关联失败, 未找到API表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("APIs").Replace(apiEntities); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("APIs").Replace(apiEntities[:1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initPermission) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysPermissionPoint
	if errors.Is(db.Where("permissionName = ?", "新增").
		Preload("APIs").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
