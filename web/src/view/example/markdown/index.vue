<template>
    <div class="markdown-box">
        <!-- <div style="height:100%;display: flex;flex: 1;">
            <div style="height: 100%;width:calc((100% - 20px)/2);padding:10px;box-sizing: border-box;background:linear-gradient(180deg,rgba(255, 0, 0,0.2) 0%,rgba(255, 255, 0, 0.5) 50%,rgba(0, 0, 255,0.8) 100%,transparent);font-size: 12px;color:rgb(100, 19, 251);">
                <pre contenteditable="true" style="">
                    <div>21321</div>
                </pre>
            </div>
            <div style="height: 100%;width:20px;background-color: #ccc;">

            </div>
            <div style="height: 100%;width:calc((100% - 20px)/2)">
                <pre>
                    <div>21321</div>
                </pre>
            </div>
        </div> -->
        <v-md-editor v-model="input" height="100%"></v-md-editor>
        <!-- <div style="border: 1px solid #ccc">
            <Toolbar
                style="border-bottom: 1px solid #ccc"
                :editor="editorRef"
                :defaultConfig="toolbarConfig"
                :mode="mode"
            />
            <Editor
                style="min-height: 250px; overflow-y: hidden;"
                v-model="valueHtml"
                :defaultConfig="editorConfig"
                :mode="mode"
                @onCreated="handleCreated"
            />
        </div> -->
    </div>

</template>
    
<script setup>
//   import WarningBar from '@/components/warningBar/warningBar.vue'
    import { connectWS,WebSocketTaskList,getMessage,sendMessage } from '@/api/chat'
//   import { getDB as getDBAPI } from '@/api/autoCode'
    import { ref, watch, onMounted,onBeforeUnmount } from 'vue'
    import { formatTimeToStr } from '@/utils/date'
    import { a } from '@/utils/preventDebug'
   

    //   a()
    let input = ref('')
    watch(input,(newVal)=>{
        console.log(newVal)
    })

    onBeforeUnmount(() => {
    })


    defineExpose({
    })
    onMounted(async()=>{


    })

    </script>
    
    <style scoped lang="scss">
    .markdown-box {
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
    .md{
        contenteditable:true;
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
    