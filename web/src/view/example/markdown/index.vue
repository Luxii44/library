<template>
    <div class="markdown-box">
        <div class="md-title-box">
            <el-input type="text" size="large" v-model="Title" placeholder="标题" style="border: none !important;"/>
        </div>
        <v-md-editor v-model="input" height="calc(100% - 54px)" @save="saveMD"></v-md-editor>
        <!-- <div>
            <video style="height: 500px;width: 100%;" src="http://47.106.153.221:8091/live/ch45.flv"></video>
        </div> -->
    </div>

</template>
    
<script setup>
    import {
        saveMarkdown,
    } from '@/api/markdown'
//   import WarningBar from '@/components/warningBar/warningBar.vue'
    import { connectWS,WebSocketTaskList,getMessage,sendMessage } from '@/api/chat'
//   import { getDB as getDBAPI } from '@/api/autoCode'
    import { ref, watch, onMounted,onBeforeUnmount } from 'vue'
    import { formatTimeToStr } from '@/utils/date'
    import { a } from '@/utils/preventDebug'

    let Title=ref("")
    const saveMD=async(e)=>{
        if (Title.value==""){
            Title.value="无标题"
        }
        const res = await saveMarkdown({
            Description: "",
            authorized_status: false,
            categories: "",
            markdowncontent:e,
            cover_images: [],
            cover_type: 1,
            id:132211147,
            is_new: 1,
            level: "1",
            not_auto_saved: "1",
            original_link: "",
            pubStatus: "draft",
            readType: "public",
            resource_id: "",
            resource_url: "",
            source: "pc_mdeditor",
            status: 2,
            tags: "",
            title: Title.value,
            type: "original",
            vote_id: 0
        })
        if (res.code === 0) {
            ElMessage({
                type: 'success',
                message: res.msg
            })
        }
    }

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
    .md-title-box{
        margin-bottom: 20px;
    }
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
    