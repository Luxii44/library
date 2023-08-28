package example

import (
	"fmt"
	exaRequest "github.com/Luxii44/library/server/model/example/request"
)

type MarkdownService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddRecruitAccount
//@description: 创建账号
//@param: e model.ExaRecruitAccount
//@return: err error

func (exa *MarkdownService) Save(e exaRequest.MarkdownResData) (err error) {
	fmt.Println(e)
	//err = global.GVA_DB.Create(&e).Error
	return err
}
