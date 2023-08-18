<template>
    <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addNewAccount()">新增账号</el-button>
      </div>
      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column align="left" label="账号" min-width="120" prop="account">
          <template #default="scope">
            <span>{{ scope.row.account }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="平台" min-width="140" prop="platform">
          <template #default="scope">
            <div class="icon-column">
              <span>{{ scope.row.platform }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="账号类型" show-overflow-tooltip min-width="160" prop="type" />
        <el-table-column align="left" fixed="right" label="操作" width="300">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="tickets"
              @click="openDetail(scope.row)"
            >详情</el-button>
            <!-- <el-button
              type="primary"
              link
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >编辑</el-button> -->
            <el-button

              type="primary"
              link
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="handleClose" :title="dialogTitle">
      <warning-bar title="新增组织，需要在用户管理内配置给用户才可使用" />
      <el-form
        v-if="dialogFormVisible"
        ref="AccountFormRef"
        :inline="true"
        :model="form"
        :rules="AccountRules"
        label-position="top"
        label-width="85px"
      >
        <el-form-item label="账号" prop="account" style="width:30%">
          <el-input
            v-model="form.account"
            autocomplete="off"
            placeholder="账号"
            @change="changeName"
          />
        </el-form-item>
        <el-form-item label="平台" prop="platform" style="width:30%">
          <el-select v-model="form.platform" placeholder="">
            <el-option v-for="item in PlatformList" :value="item" :label="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="账号类型" prop="type" style="width:30%">
          <el-select v-model="form.type" placeholder="">
            <el-option :value="1" label="招聘者" />
            <el-option :value="2" label="企业" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否启用" style="width:30%">
          <el-select v-model="form.enable" placeholder="账号是否启用">
            <el-option :value="1" label="是" />
            <el-option :value="2" label="否" />
          </el-select>
        </el-form-item>
        <!-- <el-form-item prop="path" style="width:30%">
          <template #label>
            <span style="display: inline-flex;align-items: center;">
              <span>路由Path</span>
              <el-checkbox v-model="checkFlag" style="margin-left:12px;height: auto">添加参数</el-checkbox>
            </span>
          </template>

          <el-input
            v-model="form.path"
            :disabled="!checkFlag"
            autocomplete="off"
            placeholder="建议只在后方拼接参数"
          />
        </el-form-item> -->
        <el-form-item label="父级组织" v-if="form.parentUUID" style="width:30%">
          <el-tree-select v-model="form.parentId" :data="data" :render-after-expand="false" />
          <!-- <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          /> -->
        </el-form-item>
        <el-form-item label="所属公司" v-if="form.parentUUID" style="width:30%">
          <el-tree-select v-model="form.parentId" :data="data" :render-after-expand="false" />
          <!-- <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          /> -->
        </el-form-item>
      </el-form>
      <div>
        <el-button
          type="primary"
          icon="edit"
          @click="openPLan(form.account)"
        >新增计划</el-button>
        <el-table :data="PlanTableData" style="width: 100%;margin-top: 12px;">
          <el-table-column align="left" prop="type" label="query" width="180">
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  icon="delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >查看</el-button>
              </div>
              <div>
                <el-button
                  type="danger"
                  icon="delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >修改</el-button>
              </div>
              <div>
                <el-button
                  type="danger"

                  icon="delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >删除</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

        <el-button
          style="margin-top:12px"
          type="primary"
          icon="edit"
          @click="addBtn(form)"
        >新增可控按钮</el-button>
        <el-table :data="form.menuBtn" style="width: 100%;margin-top: 12px;">
          <el-table-column align="left" prop="name" label="按钮名称" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.name" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="name" label="备注" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.desc" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"

                  icon="delete"
                  @click="deleteBtn(form.menuBtn,scope.$index)"
                >删除</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog(0)">取 消</el-button>
          <el-button type="primary" @click="enterAccountDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    
    <el-dialog v-model="dialogPlanVisible" :before-close="handleClose" :title="true ? '新增计划' : ''">
      <el-form
        v-if="dialogPlanVisible"
        ref="PlanFormRef"
        :inline="true"
        :model="Plan"
        :rules="PlanRules"
        label-position="top"
        label-width="85px"
      >
        <el-form-item label="关键字" prop="query" style="width:100%">
          <el-input
            v-model="Plan.query"
            autocomplete="off"
            placeholder="关键字"
          />
        </el-form-item>
        <el-form-item label="城市" prop="city" style="width:100%">
          <el-button :type="Plan.city==item.code ? 'primary' : 'default'" v-for="item in HotCityList" :key="item.name" @click="selectCity(item)" style="margin-bottom: 10px;">{{ item.name }}</el-button>
        </el-form-item>
        <el-form-item label="区域" prop="multiBusinessDistrict" style="width:100%">
          <div>
            <el-button :type="Plan.multiBusinessDistrict.split(':')[0]==item.code ? 'primary' : 'default'" v-for="item in DistrictList" :key="item.code" @click="selectDistrict(item,0)" style="margin-bottom: 10px;">{{ item.name }}</el-button>
          </div>
          <div>
            <el-button :type="Plan.multiBusinessDistrict.split(':')[1]==item.code ? 'primary' : 'default'" v-for="item in SpecificDistrictList" :key="item.code" @click="selectDistrict(item,1)" style="margin-bottom: 10px;">{{ item.name }}</el-button>
          </div>
        </el-form-item>
        <el-form-item label="行业" prop="position" style="width:100%">
          <el-cascader v-model="value" :props="{value:'code', label: 'name', children: 'subLevelModelList'}" :options="IndustryList" @change="selectPosition" />
        </el-form-item>
        <el-form-item label="是否投简历" prop="delivery" style="width:30%">
          <el-select v-model="Plan.delivery" placeholder="是否投简历">
            <el-option :value="1" label="是" />
            <el-option :value="2" label="否" />
          </el-select>
        </el-form-item>
        <el-form-item label="开始时间" prop="beginTime" style="width:30%">
          <el-time-picker v-model="Plan.beginTime" placeholder="选择开始时间" />
        </el-form-item>
        <el-form-item label="间隔时间" prop="intervalTime" style="width:30%">
          <el-input v-model="Plan.intervalTime" types="number" placeholder="请输入间隔时间" >
            <template #append>分</template>
          </el-input>
        </el-form-item>
        <el-form-item label="投递数量" prop="deliverCount" style="width:30%">
          <el-input v-model="Plan.deliverCount" types="number" placeholder="请输入投递数量" >
            <template #append>个</template>
          </el-input>
        </el-form-item>
        <el-form-item label="账号" prop="account" style="width:30%">
          <span>{{ Plan.account }}</span>
        </el-form-item>
        <el-form-item label="是否启用" style="width:30%">
          <el-select v-model="Plan.enable" placeholder="账号是否启用">
            <el-option :value="1" label="是" />
            <el-option :value="2" label="否" />
          </el-select>
        </el-form-item>
        <!-- <el-form-item prop="path" style="width:30%">
          <template #label>
            <span style="display: inline-flex;align-items: center;">
              <span>路由Path</span>
              <el-checkbox v-model="checkFlag" style="margin-left:12px;height: auto">添加参数</el-checkbox>
            </span>
          </template>

          <el-input
            v-model="form.path"
            :disabled="!checkFlag"
            autocomplete="off"
            placeholder="建议只在后方拼接参数"
          />
        </el-form-item> -->
        <el-form-item label="父级组织" v-if="form.parentUUID" style="width:30%">
          <el-tree-select v-model="form.parentId" :data="data" :render-after-expand="false" />
          <!-- <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          /> -->
        </el-form-item>
        <el-form-item label="所属公司" v-if="form.parentUUID" style="width:30%">
          <el-tree-select v-model="form.parentId" :data="data" :render-after-expand="false" />
          <!-- <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          /> -->
        </el-form-item>
      </el-form>
      <template #footer>
          <div class="dialog-footer">
            <el-button @click="closeDialog(1)">取 消</el-button>
            <el-button type="primary" @click="enterPlanDialog">确 定</el-button>
          </div>
        </template>
    </el-dialog>
  </div>
</template>
  
<script setup>
  
  import {
    getAccountList,
    addAccount,
    // addPlan,
    getBusinessDistrict,
    createRecruitPlan,
    updateRecruitPlan,
    deleteRecruitPlan,
    getRecruitPlanList
  } from '@/api/recruit'
  import HotCityList from "@/assets/hotCity.json"
  import IndustryList from "@/assets/industry.json"
  import EchartsLine from '@/view/dashboard/dashboardCharts/echartsLine.vue'
  import DashboardTable from '@/view/dashboard/dashboardTable/dashboardTable.vue'
  import { onMounted, ref,reactive  } from 'vue'
  import { useRouter } from 'vue-router'
  import { useWeatherInfo } from '@/view/dashboard/weather.js'
  defineOptions({
    name:"Recruit"
  })
  const queryCity=()=>{
    
  }
  onMounted(()=>{
    getAccountData()
    console.log(IndustryList)
    // queryCity()
  })
  let PageRequest=ref({
    pageSize:20,
    pageIndex:1,
    total:0
  })
  let tableData=ref([])
  // 查询
  const getAccountData = async() => {
    const table = await getAccountList({ page: PageRequest.value.pageIndex, pageSize: PageRequest.value.pageSize })
    if (table.code === 0) {
      tableData.value = table.data.list
      PageRequest.value.total = table.data.total
    }
    if (tableData.value.length==0){
      
    }
  }
  let dialogTitle=ref("")
  let dialogFormVisible=ref(false)
  // 初始化弹窗内表格方法
  let AccountFormRef=ref(null)
  const checkFlag = ref(false)
  const initAccountForm = () => {
    checkFlag.value = false
    AccountFormRef.value.resetFields()
    form.value = {
      account:"19875606386",
      platform:"boss",
      type:1,
      enable:1,
      plans:[]
    }
  }
  const AccountRules = reactive({
    account: [{ required:true,message:"请填写账号", trigger: 'blur' }],
    platform: [{ required:true,message:"请填写平台",trigger: 'blur' }],
  })
  let form=ref({
    account:"19875606386",
    platform:"boss",
    type:1,
    enable:1,
    plans:[]
  })
  let PlanTableData=ref([])
  let isEdit=ref(false)
  const PlatformList=(["boss","51job","猎聘"])
  // 打开新增/详情弹窗
  const addNewAccount = async(id) => {
    dialogTitle.value = '新增账号'
    form.value.parentUUID = id
    isEdit.value = false
    // setOptions()
    dialogFormVisible.value = true
    // const result = await addAccount(form.value)
  }
  // 添加账号
  const enterAccountDialog = async() => {
    AccountFormRef.value.validate(async valid => {
      if (valid) {
        let res
        if (isEdit.value) {
          res = await updateBaseOrg(form.value)
        } else {
          res = await addAccount(form.value)
        }
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: isEdit.value ? '编辑成功' : '添加成功!'
          })
          getAccountData()
        }
        initAccountForm()
        dialogFormVisible.value = false
      }
    })
  }
  // 打开详情
  const closeDialog = async(code) => {
    if (code==0){
      dialogFormVisible.value=false
      return
    }
    dialogPlanVisible.value=false
  }
  // 打开详情
  const openDetail = async(item) => {
    isEdit.value=true
    dialogFormVisible.value=true
    Plan.value.account=item.account
    getPlanData()
  }

  let dialogPlanVisible=ref(false)
  // 初始化弹窗内表格方法
  let PlanFormRef=ref(null)
  const initPlanForm = () => {
    PlanFormRef.value.resetFields()
    Plan.value = {
      query:"",
      cityCode:"",
      cityName:"",
      districtCode:"",
      districtName:"",
      position:"",
      delivery:1,
      beginTime:"",
      intervalTime:30,
      account:"",
      enable:1
    }
  }
  //关键字、城市、区域、公司行业、职业、
  let Plan=ref({
    query:"",
    city:"",
    cityName:"",
    multiBusinessDistrict:"",
    multiBusinessDistrictName:"",
    industry:"",
    position:"",
    positionName:"",
    experience:"",
    payType:"",
    partTime:"",
    degree:"",
    scale:"",
    stage:"",
    jobType:"",
    salary:"",
    multiSubway:"",
    delivery:1,
    beginTime:"",
    intervalTime:30,
    deliverCount:30,
    account:"",
    enable:1
  })
  const PlanRules = reactive({
    query:[{ required:true,message:"请填写平台",trigger: 'blur' }],
    beginTime:[{ required:true,message:"请选择开始时间",trigger: 'blur' }],
    intervalTime:[{ required:true,message:"请输入间隔时间",trigger: 'blur' }],
  })
  // 新增计划
  const openPLan = (account) => {
    dialogPlanVisible.value=true
    Plan.value.account=account
  }
  const selectCity = async(item) => {
    Plan.value.cityName=item.name
    Plan.value.city=item.code
    if (item.code==100010000){
      return
    }
    let res=await getBusinessDistrict({cityCode:item.code})
    if (res.data.code==item.code && res.code==0){
      DistrictList.value=res.data.subLevelModelList
      console.log(DistrictList.value)
    }
  }
  let DistrictList=ref([])
  let SpecificDistrictList=ref([])
  const selectDistrict = (item,code) => {
    if (code==0){
      Plan.value.multiBusinessDistrict=String(item.code)
      Plan.value.multiBusinessDistrictName=item.name
      SpecificDistrictList.value=item.subLevelModelList
      return
    }
    if(Plan.value.multiBusinessDistrict.indexOf(":")<0){
      Plan.value.multiBusinessDistrict=Plan.value.multiBusinessDistrict+":"+item.code
      Plan.value.multiBusinessDistrictName=Plan.value.multiBusinessDistrictName+":"+item.name
    }else{
      Plan.value.multiBusinessDistrict=Plan.value.multiBusinessDistrict.split(":")[0]+":"+item.code
      Plan.value.multiBusinessDistrictName=Plan.value.multiBusinessDistrictName.split(":")[0]+":"+item.name
    }
  }
  let PositionList=ref([])
  let SpecificPositionList=ref([])
  const selectPosition = (item) => {
    Plan.value.position=item[2]
  }
  
  const getPlanData = async(item) => {
    let res=await getRecruitPlanList({account:form.value.account})
    if (res.code==0){
      DistrictList.value=res.data.subLevelModelList
      console.log(DistrictList.value)
    }
  }
  
  let isEditPlan=ref(false)
  // 添加计划
  const enterPlanDialog = async() => {
    PlanFormRef.value.validate(async valid => {
      if (valid) {
        let res
        if (isEditPlan.value) {
          res = await updateRecruitPlan(Plan.value)
        } else {
          res = await createRecruitPlan(Plan.value)
        }
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: isEdit.value ? '编辑成功' : '添加成功!'
          })
          getPlanData()
        }
        initPlanForm()
        dialogFormVisible.value = false
      }
      
    })
  }
  // const createPlan = async() => {
  //   let res = await addPlan(Plan.value)
  // }
  
  


  const weatherInfo = useWeatherInfo()
  
  const toolCards = ref([
    {
      label: '用户管理',
      icon: 'monitor',
      name: 'user',
      color: '#ff9c6e',
      bg: 'rgba(255, 156, 110,.3)'
    },
    {
      label: '角色管理',
      icon: 'setting',
      name: 'authority',
      color: '#69c0ff',
      bg: 'rgba(105, 192, 255,.3)'
    },
    {
      label: '菜单管理',
      icon: 'menu',
      name: 'menu',
      color: '#b37feb',
      bg: 'rgba(179, 127, 235,.3)'
    },
    {
      label: '代码生成器',
      icon: 'cpu',
      name: 'autoCode',
      color: '#ffd666',
      bg: 'rgba(255, 214, 102,.3)'
    },
    {
      label: '表单生成器',
      icon: 'document-checked',
      name: 'formCreate',
      color: '#ff85c0',
      bg: 'rgba(255, 133, 192,.3)'
    },
    {
      label: '关于我们',
      icon: 'user',
      name: 'about',
      color: '#5cdbd3',
      bg: 'rgba(92, 219, 211,.3)'
    }
  ])
  
  const router = useRouter()
  
  const toTarget = (name) => {
    router.push({ name })
  }
  
  </script>
  
  <style lang="scss" scoped>
  @mixin flex-center {
      display: flex;
      align-items: center;
  }
  .page {
      background: #f0f2f5;
      padding: 0;
      .gva-card-box{
        padding: 12px 16px;
        &+.gva-card-box{
          padding-top: 0px;
        }
      }
      .gva-card {
        box-sizing: border-box;
          background-color: #fff;
          border-radius: 2px;
          height: auto;
          padding: 26px 30px;
          overflow: hidden;
          box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
      }
      .gva-top-card {
          height: 260px;
          @include flex-center;
          justify-content: space-between;
          color: #777;
          &-left {
            height: 100%;
            display: flex;
            flex-direction: column;
              &-title {
                  font-size: 22px;
                  color: #343844;
              }
              &-dot {
                  font-size: 16px;
                  color: #6B7687;
                  margin-top: 24px;
              }
              &-rows {
                  // margin-top: 15px;
                  margin-top: 18px;
                  color: #6B7687;
                  width: 600px;
                  align-items: center;
              }
              &-item{
                +.gva-top-card-left-item{
                  margin-top: 24px;
                }
                margin-top: 14px;
              }
          }
          &-right {
              height: 600px;
              width: 600px;
              margin-top: 28px;
          }
      }
       ::v-deep(.el-card__header){
            padding:0;
            border-bottom: none;
          }
          .card-header{
            padding-bottom: 20px;
            border-bottom: 1px solid #e8e8e8;
          }
      .quick-entrance-title {
          height: 30px;
          font-size: 22px;
          color: #333;
          width: 100%;
          border-bottom: 1px solid #eee;
      }
      .quick-entrance-items {
          @include flex-center;
          justify-content: center;
          text-align: center;
          color: #333;
          .quick-entrance-item {
            padding: 16px 28px;
            margin-top: -16px;
            margin-bottom: -16px;
            border-radius: 4px;
            transition: all 0.2s;
            &:hover{
              box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
            }
              cursor: pointer;
              height: auto;
              text-align: center;
              // align-items: center;
              &-icon {
                  width: 50px;
                  height: 50px !important;
                  border-radius: 8px;
                  @include flex-center;
                  justify-content: center;
                  margin: 0 auto;
                  i {
                      font-size: 24px;
                  }
              }
              p {
                  margin-top: 10px;
              }
          }
      }
      .echart-box{
        padding: 14px;
      }
  }
  .dashboard-icon {
      font-size: 20px;
      color: rgb(85, 160, 248);
      width: 30px;
      height: 30px;
      margin-right: 10px;
      @include flex-center;
  }
  .flex-center {
      @include flex-center;
  }
  
  //小屏幕不显示右侧，将登录框居中
  @media (max-width: 750px) {
      .gva-card {
          padding: 20px 10px !important;
          .gva-top-card {
              height: auto;
              &-left {
                  &-title {
                      font-size: 20px !important;
                  }
                  &-rows {
                      margin-top: 15px;
                      align-items: center;
                  }
              }
              &-right {
                  display: none;
              }
          }
          .gva-middle-card {
              &-item {
                  line-height: 20px;
              }
          }
          .dashboard-icon {
              font-size: 18px;
          }
      }
  }
  </style>
  