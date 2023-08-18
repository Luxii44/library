package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"nhooyr.io/websocket"
	"strings"
)

type ChatApi struct{}

//var upGrader = websocket.Upgrader{
//	ReadBufferSize:  1124,
//	WriteBufferSize: 1124,
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}

// CreateWS
// @Tags      ExaChat
// @Summary   创建聊天室
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaRecruitAccount  true  "消息"
// @Success   200   {object}  response.Response  "创建聊天室"
// @Router    /chat/ping [get]
func (e *ChatApi) CreateWS(c *gin.Context) {
	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	defer conn.Close(websocket.StatusInternalError, "连接意外关闭")

	for {
		msgType, data, err := conn.Read(c)
		if err != nil {
			log.Println("WebSocket read error:", err)
			return
		}

		if msgType == websocket.MessageText {
			message := strings.TrimSpace(string(data))

			// 处理聊天消息
			fmt.Println("收到消息:", message)

			// 回复消息
			err = conn.Write(c, websocket.MessageText, []byte(message))
			if err != nil {
				log.Println("WebSocket write error:", err)
				return
			}
		}
	}
	//升级get请求为webSocket协议
	//ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	//fmt.Println(ws)
	//if err != nil {
	//	return
	//}
	//global.SocketTask[len(global.SocketTask)+1] = ws
	//defer ws.Close() //返回前关闭
	//for {
	//	//读取ws中的数据
	//	mt, message, err := ws.ReadMessage()
	//	//获取ip
	//	//ip := c.Request.Header.Get("X-Real-Ip")
	//	//if ip == "" {
	//	//	ip = strings.Split(c.Request.Header.Get("x-Forwarded-For"), ",")[0]
	//	//}
	//	//if ip == "" {
	//	//	ip = c.Request.RemoteAddr
	//	//}
	//	//fmt.Println("Client IP Address:", ip) // to do something
	//	//ip := c.ClientIP()
	//	//c.JSON(http.StatusOK, gin.H{"ip": ip})
	//	if err != nil {
	//		break
	//	}
	//	a := example.ChatMessage{}
	//	fmt.Println(mt)
	//	err = json.Unmarshal(message, &a)
	//	if err != nil {
	//		break
	//	}
	//	chatService.ReceiveMessages(a)
	//	//写入ws数据
	//	//chatService.StartTimer(func() {
	//	err = ws.WriteMessage(mt, message)
	//	//})
	//	if err != nil {
	//		break
	//	}
	//}
}
