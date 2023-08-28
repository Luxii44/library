package example

import "github.com/Luxii44/library/server/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
	RecruitApi
	ChatApi
	MarkdownApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
	recruitService               = service.ServiceGroupApp.ExampleServiceGroup.RecruitService
	chatService                  = service.ServiceGroupApp.ExampleServiceGroup.ChatService
	markdownService              = service.ServiceGroupApp.ExampleServiceGroup.MarkdownService
)
