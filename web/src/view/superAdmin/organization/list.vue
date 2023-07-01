<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addOrg()">新增根组织</el-button>
      </div>
      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column align="left" label="组织编号" min-width="120" prop="organizationCode">
          <template #default="scope">
            <span>{{ scope.row.organizationCode }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="组织名称" min-width="140" prop="organizationName">
          <template #default="scope">
            <div class="icon-column">
              <span>{{ scope.row.organizationCode }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="组织全称" show-overflow-tooltip min-width="160" prop="organizationFullName" />
        <el-table-column align="left" label="ruji" show-overflow-tooltip min-width="160" prop="organizationFullName" />
        <el-table-column align="left" label="组织全称" show-overflow-tooltip min-width="160" prop="organizationFullName" />
        <!-- <el-table-column align="left" label="路由Path" show-overflow-tooltip min-width="160" prop="path" />
        <el-table-column align="left" label="是否隐藏" min-width="100" prop="hidden">
          <template #default="scope">
            <span>{{ scope.row.hidden?"隐藏":"显示" }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="父节点" min-width="90" prop="parentId" />
        <el-table-column align="left" label="排序" min-width="70" prop="sort" />
        <el-table-column align="left" label="文件路径" min-width="360" prop="component" /> -->
        <el-table-column align="left" fixed="right" label="操作" width="300">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="plus"
              @click="addMenu(scope.row.uuid)"
            >添加子组织</el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >编辑</el-button>
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
        ref="OrgFormRef"
        :inline="true"
        :model="form"
        :rules="rules"
        label-position="top"
        label-width="85px"
      >
        <el-form-item label="组织编号" prop="organizationCode" style="width:30%">
          <el-input
            v-model="form.organizationCode"
            autocomplete="off"
            placeholder="组织编号"
            @change="changeName"
          />
        </el-form-item>
        <el-form-item label="组织名称" prop="organizationName" style="width:30%">
          <el-input
            v-model="form.organizationName"
            autocomplete="off"
            placeholder="组织名称"
            @change="changeName"
          />
        </el-form-item>
        <!-- <el-form-item label="组织全称" prop="organizationFullName" style="width:30%">
          <el-input
            v-model="form.organizationFullName"
            autocomplete="off"
            placeholder="组织全称"
            @change="changeName"
          />
        </el-form-item> -->
        <el-form-item label="是否禁用" style="width:30%">
          <el-select v-model="form.enable" placeholder="是否公司不可用">
            <el-option :value="1" label="否" />
            <el-option :value="2" label="是" />
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
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog()">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  updateBaseOrg,
  getOrgList,
  createOrg,
  deleteBaseOrg,
  getBaseOrgById
} from '@/api/organization'
import icon from '@/view/superAdmin/menu/icon.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
import { onMounted, reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const validateOrgCode = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请输入组织编号'))
  } else {
    var gro =/gro-\d+/
    var com =/com-\d+/
    var dep =/dep-\d+/
    if (gro.test(value) || com.test(value) || dep.test(value)){
      callback()
    }else{
      callback(new Error('组织编号格式不匹配'))
    }
  }
}
const validateOrgName = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请输入组织名称'))
  }
  callback()
}

const validate = () => {
  console.log(form.value.organizationName.split(/[\t\r\f\n\s]*/g).join(''))
  if (form.value.organizationCode.split(/[\t\r\f\n\s]*/g).join('')=='' || form.value.organizationName.split(/[\t\r\f\n\s]*/g).join('')==''){
    return false
  }
  if (form.value.organizationCode.indexOf("gro")>-1 && form.value.organizationName.indexOf("集团")>-1){
    return false
  }else if (form.value.organizationCode.indexOf("com")>-1 && form.value.organizationName.indexOf("公司")>-1){
    return false
  }else if (form.value.organizationCode.indexOf("dep")>-1 && form.value.organizationName.indexOf("部门")>-1){
    return false
  }
  return true
}

const rules = reactive({
  organizationCode: [{ validator: validateOrgCode,required:true, trigger: 'blur' }],
  organizationName: [{ validator: validateOrgName,required:true, trigger: 'blur' }],
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(20)
const tableData = ref([])
const searchInfo = ref({})
// 查询
const getTableData = async() => {
  const table = await getOrgList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  if (tableData.value.length==0){
    
  }
}

getTableData()

// 新增参数
const addParameter = (form) => {
  if (!form.parameters) {
    form.parameters = []
  }
  form.parameters.push({
    type: 'query',
    key: '',
    value: ''
  })
}

const fmtComponent = () => {
  form.value.component = form.value.component.replace(/\\/g, '/')
}

// 删除参数
const deleteParameter = (parameters, index) => {
  parameters.splice(index, 1)
}

// 新增可控按钮
const addBtn = (form) => {
  if (!form.menuBtn) {
    form.menuBtn = []
  }
  form.menuBtn.push({
    name: '',
    desc: '',
  })
}
// 删除可控按钮
const deleteBtn = async(btns, index) => {
  const btn = btns[index]
  if (btn.ID === 0) {
    btns.splice(index, 1)
    return
  }
  const res = await canRemoveAuthorityBtnApi({ id: btn.ID })
  if (res.code === 0) {
    btns.splice(index, 1)
  }
}

const form = ref({
  organizationCode:'',
  organizationName:'',
  organizationFullName:'',
  parentUUID:'',
  parentName:'',
  companyUUID:'',
  companyName:'',
  topOrganizationUUID:'',
  topOrganizationName:'',
  enable: 1,
  // parentId: '',
  // component: '',
  // meta: {
  //   activeName: '',
  //   title: '',
  //   icon: '',
  //   defaultMenu: false,
  //   closeTab: false,
  //   keepAlive: false
  // },
  // parameters: [],
  // menuBtn: []
})
const changeName = () => {
  form.value.path = form.value.name
}

const handleClose = (done) => {
  initForm()
  done()
}
// 删除菜单
const deleteMenu = (ID) => {
  ElMessageBox.confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteBaseOrg({ ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}
// 初始化弹窗内表格方法
const OrgFormRef = ref(null)
const checkFlag = ref(false)
const initForm = () => {
  checkFlag.value = false
  OrgFormRef.value.resetFields()
  form.value = {
    organizationCode: '',
    organizationName: '',
    organizationFullName: '',
    parentUUID:'00000000-0000-0000-0000-000000000000',
    companyUUID:'00000000-0000-0000-0000-000000000000',
    topOrganizationUUID:'00000000-0000-0000-0000-000000000000',
    enable:1
    // hidden: false,
    // parentId: '',
    // component: '',
    // meta: {
    //   title: '',
    //   icon: '',
    //   defaultMenu: false,
    //   closeTab: false,
    //   keepAlive: false
    // }
  }
}
// 关闭弹窗

const dialogFormVisible = ref(false)
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}
// 添加menu
const enterDialog = async() => {
  OrgFormRef.value.validate(async valid => {
    if (valid) {
      validate
      let res
      if (isEdit.value) {
        res = await updateBaseOrg(form.value)
      } else {
        res = await createOrg(form.value)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? '编辑成功' : '添加成功!'
        })
        getTableData()
      }
      initForm()
      dialogFormVisible.value = false
    }
  })
}

const menuOption = ref([
  {
    ID: '0',
    title: '根菜单'
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: '0',
      title: '根目录'
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
        menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID,
              children: []
            }
            setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === form.value.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID
            }
            optionsData.push(option)
          }
        })
}

// 添加菜单方法，id为 0则为添加根菜单
const isEdit = ref(false)
const dialogTitle = ref('新增菜单')
const addMenu = (id) => {
  dialogTitle.value = '新增菜单'
  if (!id){
    form.value.organizationCode="gro-"
    form.value.organizationName="集团"
  }
  form.value.parentUUID = id
  form.value.parentName = id
  // isEdit.value = false
  // setOptions()
  dialogFormVisible.value = true
}
// 修改菜单方法
const editMenu = async(id) => {
  dialogTitle.value = '编辑菜单'
  const res = await getBaseOrgById({ id })
  form.value = res.data.menu
  isEdit.value = true
  setOptions()
  dialogFormVisible.value = true
}

</script>

<script>
export default {
  name: 'Menus',
}
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
.icon-column{
  display: flex;
  align-items: center;
  .el-icon{
    margin-right: 8px;
  }
}
</style>
