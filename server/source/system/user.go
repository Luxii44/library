package system

import (
	"context"
	sysModel "github.com/Luxii44/library/server/model/system"
	"github.com/Luxii44/library/server/service/system"
	"github.com/Luxii44/library/server/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const initOrderUser = initOrderAuthority + 1

type initUser struct{}

// auto run
func init() {
	system.RegisterInit(initOrderUser, &initUser{})
}

func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysUser{}, &sysModel.SysChatGptOption{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	password := utils.BcryptHash("6447985")
	adminPassword := utils.BcryptHash("123456")

	entities := []sysModel.SysUser{
		{
			UUID:        uuid.NewV4(),
			Account:     "admin",
			Username:    "admin",
			Password:    adminPassword,
			NickName:    "Luxii",
			HeaderImg:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			AuthorityId: 1,
			Phone:       "19875606386",
			Email:       "Luxii0826@163.com",
		},
		{
			UUID:        uuid.NewV4(),
			Account:     "user01",
			Username:    "Luxii",
			Password:    utils.BcryptHash("Lx345501"),
			NickName:    "Luxii",
			HeaderImg:   "https:///qmplusimg.henrongyi.top/1572075907logo.png",
			AuthorityId: 8880,
			Phone:       "19875606386",
			Email:       "Luxii0826@163.com",
		},
		{
			UUID:        uuid.NewV4(),
			Account:     "finance01",
			Username:    "小明",
			Password:    utils.BcryptHash("123456"),
			NickName:    "小明",
			HeaderImg:   "https:///qmplusimg.henrongyi.top/1572075907logo.png",
			AuthorityId: 1100,
			Phone:       "19875606386",
			Email:       "Luxii0826@163.com",
		},
		{
			UUID:        uuid.NewV4(),
			Account:     "finance02",
			Username:    "小红",
			Password:    utils.BcryptHash("123456"),
			NickName:    "小红",
			HeaderImg:   "https:///qmplusimg.henrongyi.top/1572075907logo.png",
			AuthorityId: 1101,
			Phone:       "19875606386",
			Email:       "Luxii0826@163.com",
		},
		{
			UUID:        uuid.NewV4(),
			Account:     "test01",
			Username:    "test",
			Password:    password,
			NickName:    "测试用户01",
			HeaderImg:   "https:///qmplusimg.henrongyi.top/1572075907logo.png",
			AuthorityId: 9528,
			Phone:       "19875606386",
			Email:       "Luxii0826@163.com",
		},
		{
			UUID:        uuid.NewV4(),
			Account:     "Reserve01",
			Username:    "Reserve",
			Password:    password,
			NickName:    "备用账号01",
			HeaderImg:   "https:///qmplusimg.henrongyi.top/1572075907logo.png",
			AuthorityId: 1,
			Phone:       "19875606386",
			Email:       "Luxii0826@163.com",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	authorityEntities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("Authorities").Replace(authorityEntities); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("Authorities").Replace(authorityEntities[:1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysUser
	if errors.Is(db.Where("username = ?", "Luxii").
		Preload("Authorities").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Authorities) > 0 && record.Authorities[0].AuthorityId == 1
}
