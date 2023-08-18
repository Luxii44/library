<template>
  <div class="chat-box">
    <!-- 123 -->
    <div style="height:100%">
      <div v-for="item in List" :key="item.type">
        <div v-if="item.type=='s'" style="display: flex;justify-content: end;">
          <span class="sent">{{item.messageContent}}</span>
          <img class="avatar " src="@/assets/logo.png" alt="">
        </div>
        <div v-else style="display: flex;justify-content:start;">
          <img class="avatar " src="@/assets/logo.png" alt="">
          <span class="received">{{item.messageContent}}</span>
        </div>
      </div>
    </div>
    <el-input style="position: sticky;bottom: 0;" v-model="input" placeholder="Please input">
      <template #append>
        <el-button @click="send">发送</el-button>
      </template>
    </el-input>
  </div>
      <!-- <warning-bar title="使用GPT-3.5模型，存在一定不稳定性，成功率为50%左右，使用GPT-4可以极大提升成功率，但是费用较高。" />
      <div v-if="!chatToken">
        <el-input v-model="skObj.sk" class="query-ipt" placeholder="请输入您的ChatGpt SK" clearable />
        <el-button type="primary" @click="save">保存</el-button>
        <div class="secret">
          <el-empty description="请到gpt网站获取您的sk：https://platform.openai.com/account/api-keys" />
        </div>
      </div>
      <div v-else>
        <el-form :model="form" label-width="120px">
          <el-form-item label="删除当前sk：">
            <el-popover placement="top" width="160">
              <p>确定要删除并返回吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" @click="deleteSK">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete">删除</el-button>
              </template>
            </el-popover>
          </el-form-item>
          <el-form-item label="查询db名称：">
            <el-select v-model="form.dbname" placeholder="请选择库" style="width: 115px">
              <el-option
                v-for="(item, index) in dbArr"
                :key="index"
                :label="item.database"
                :value="item.database"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="查询db描述：">
            <el-input
              v-model="form.chat"
              :autosize="{ minRows: 2, maxRows: 4 }"
              type="textarea"
              clearable
              placeholder="请输入对话"
            />
          </el-form-item>
          <el-form-item label="GPT生成SQL:">
            <el-input
                v-model="sql"
                :autosize="{ minRows: 2, maxRows: 4 }"
                type="textarea"
                disabled
                placeholder="此处展示自动生成的sql"
            />
          </el-form-item>
          <el-button type="primary" @click="handleQueryTable">查询</el-button>
        </el-form>
        <div class="tables">
          <el-table
            v-if="tableData.length"
            ref="multipleTable"
            :data="tableData"
            style="width: 100%"
            tooltip-effect="dark"
            height="400px"
          >
            <el-table-column
              v-for="(item, index) in tableData[0]"
              :key="index"
              :prop="index"
              :label="index"
              min-width="200"
              show-overflow-tooltip
            />
          </el-table>
          <p v-else class="text">请在对话框输入你需要AI帮你查询的内容：）</p>
        </div>
      </div> -->
  </template>
  
  <script setup>
//   import WarningBar from '@/components/warningBar/warningBar.vue'
    import { connectWS,WebSocketTaskList,getMessage,sendMessage } from '@/api/chat'
  //   import { getDB as getDBAPI } from '@/api/autoCode'
    import { ref, reactive, onMounted } from 'vue'
    import { useUserStore } from '@/pinia/modules/user'
    import { formatTimeToStr } from '@/utils/date'
    import { a } from '@/utils/preventDebug'
    
    // a()
    let userStore=useUserStore()
    const input = ref('')
    let List=ref([])
    let socketTask=ref({})
    let socketId=ref(null)
    // 连接建立时触发
    const send=async()=>{
      sendMessage(socketId,input.value,userStore.userInfo.uuid)
      List.value.push({type:'s',messageContent:input.value})
      input.value=""
    }

    // // 连接关闭时触发
    // socket.onclose = function(event) {
    //   console.log('WebSocket连接已关闭');
    // };

    // // 发生错误时触发
    // socket.onerror = function(error) {
    //   console.error('WebSocket发生错误:', error);
    // };
    onMounted(async()=>{
      await connectWS({
        url:'ws://localhost:3000/chat/ping',
        success:(res)=>{
          socketId.value=res.socketTaskId
        },
        fail:(res)=>{
          console.log(res)
        }
      })
      getMessage(socketId,(event)=>{
        List.value.push({type:'r',messageContent:JSON.parse(event.data).messageContent})
      })
      // socket.onopen = function(event) {
      //   console.log('WebSocket连接已建立');
        
      //   // 发送消息给服务器
      //   // socket.send(JSON.stringify({senderId:userStore.userInfo.uuid,messageContent:'Hello, server!',messageType:0,receiverId:userStore.userInfo.uuid,time:new Date().getTime()}));
      // };
      // socket.onmessage = function(event) {
      //   const res = JSON.parse(event.data);
      //   if (res.receiverId==userStore.userInfo.uuid){
      //     res.type="r"
      //     console.log('收到消息了',res.messageContent);
      //     List.value.push(res)
      //   }
        // List=JSON.parse(res)
      // };
    })
    // var videoPlayer = document.getElementById('videoPlayer');
    // var mediaSource = new MediaSource();
    // var videoUrl = URL.createObjectURL(mediaSource);

    // // 设置 video 元素的 src 属性为经过 URL 创建的视频地址
    // videoPlayer.src = videoUrl;

    // // 监听 MediaSource 打开事件
    // mediaSource.addEventListener('sourceopen', handleSourceOpen, false);

    // function handleSourceOpen() {
    //   // 创建一个 MediaSource 资源缓冲对象
    //   var sourceBuffer = mediaSource.addSourceBuffer('video/mp4; codecs="avc1.42E01E, mp4a.40.2"');

    //   // 请求视频流分片
    //   fetch('your_video_stream_url')
    //     .then(function(response) {
    //       return response.arrayBuffer();
    //     })
    //     .then(function(data) {
    //       // 将获取的视频数据追加到缓冲区
    //       sourceBuffer.appendBuffer(data);
    //     });
    // }
  </script>
  
  <style scoped lang="scss">
  .chat-box {
  box-shadow: 0 0 10px rgba(0,0,0,0.3);
  border-radius: 10px;
  overflow: hidden;
  position: relative;
  height: 100%;
  width: 100%;
  padding: 24px;
  background-color: #fff;
  box-sizing: border-box;
  /* 头像 */
    .avatar {
      border-radius: 50%;
      height: 50px;
      width: 50px;
    }
    /* 消息内容 */
    .message {
      padding: 10px;
      border-radius: 5px;
      margin: 10px;
      max-width: 60%;
    }
    /* 发送方消息 */
    .sent {
      background-color: #0077FF;
      color: #FFF;
      float: right;
    }
    /* 接收方消息 */
    .received {
      background-color: #FFF;
      color: #333;
      float: left;
    }
  }
  .buddha{
    width: 30px;
    height: 30px;
    margin-top: 20px;
    background: #F5F5F5;
    p {
      line-height: 30px;
    }
  }
  .query-ipt{
    width: 300px;
    margin-right: 30px;
  }
  .content{
    p {
      font-size: 16px;
      line-height: 20px;
    }
    padding: 10px;
    width: 100%;
    background: #F5F5F5;
    margin-top: 30px;
  }
  .tables{
    width: 100%;
    margin-top: 30px;
    .text{
      font-size: 18px;
      color: #6B7687;
      text-align: center;
    }
  }
  </style>
  