package example

import "github.com/Luxii44/library/server/service"

type ApiGroup struct {
	CustomerApi
	RecruitApi
	ChatApi
	FileUploadAndDownloadApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
	recruitService               = service.ServiceGroupApp.ExampleServiceGroup.RecruitService
	chatService                  = service.ServiceGroupApp.ExampleServiceGroup.ChatService
)
