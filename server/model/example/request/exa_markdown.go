package request

// Get Recruit Plan List structure
type MarkdownResData struct {
	Description     string `json:"description" gorm:"comment:描述"`               // 关键字
	MarkdownContent string `json:"markdownContent" gorm:"comment:markdown文本内容"` // markdown文本内容
	Title           string `json:"title" gorm:"comment:标题"`                     // 标题
	Id              int64  `json:"id" gorm:"comment:保存Id"`                      // 保存Id
	Status          int64  `json:"status" gorm:"comment:状态"`                    // 状态
	//authorized_status: false,
	//categories: "",
	//markdowncontent:e,
	//cover_images: [],
	//cover_type: 1,
	//id:132211147,
	//is_new: 1,
	//level: "1",
	//not_auto_saved: "1",
	//original_link: "",
	//pubStatus: "draft",
	//readType: "public",
	//resource_id: "",
	//resource_url: "",
	//source: "pc_mdeditor",
	//status: 2,
	//tags: "",
	//title: "【无标题】",
	//type: "original",
	//vote_id: 0
}
