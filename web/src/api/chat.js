import { useUserStore } from '@/pinia/modules/user'

let userStore=useUserStore()
export class Task {
  SocketTaskId
  SocketTask
  isConnected
  success(){

  }
}
export var WebSocketTaskList=new Array()
export const connectWS=(obj)=>{
  return new Promise(function(resolve,reject){
    if("WebSocket" in window){
      console.log("您的浏览器支持WebSocket");
    }else{
      console.log("您的浏览器不支持WebSocket");
      return
    }
    let task = new WebSocket(obj.url);
    task.onclose=new Function
    task.onmessage=new Function
    task.onopen=function(event) {
      WebSocketTaskList.push({SocketTaskId:WebSocketTaskList.length+1,SocketTask:task})
      console.log(WebSocketTaskList)
      obj.success({errMsg: "connectSocket:ok",socketTaskId:WebSocketTaskList.length})
      resolve()
    }
    task.onerror= function(event) {
      console.log("连接已关闭或者不能打开")
      if (event.currentTarget.readyState==3){
        console.log("连接已关闭或者不能打开")
        obj.fail({errMsg: "连接已关闭或者不能打开",socketTaskId:WebSocketTaskList.length})
      }
      reject()
    }
  })
}
export const getMessage=(id,callback)=>{
  console.log(WebSocketTaskList)
  WebSocketTaskList.filter(item=>{return item.socketTaskId=id})[0].SocketTask.onmessage=callback
}
export const sendMessage=(id,Content,receiverId)=>{
  let a=WebSocketTaskList.filter(item=>{return item.socketTaskId=id})[0].SocketTask
  a.send(JSON.stringify({senderId:userStore.userInfo.uuid,messageContent:Content,messageType:0,receiverId:receiverId,time:new Date().getTime()}));
}
export const close=(id)=>{
  let a=WebSocketTaskList.filter(item=>{return item.socketTaskId=id})[0].SocketTask
  a.close()
  a.onclose=function(event){
    console.log("WebSocket"+id+"已关闭")
  }
}