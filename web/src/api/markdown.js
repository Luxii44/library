import service from '@/utils/request'
// @Summary 用户登录 获取动态路由
// @Produce  application/json
// @Param 可以什么都不填 调一下即可
// @Router /org/getOrg [post]
export const saveMarkdown = (data) => {
  return service({
    url: '/markdown/saveMarkdown',
    method: 'post',
    data
  })
}

// @Summary 获取招聘账号列表
// @Produce  application/json
// @Param {
//  page     int
//	pageSize int
// }
// @Router /recruit/accountList [get]
export const getAccountList = (params) => {
  return service({
    url: '/recruit/accountList',
    method: 'get',
    params
  })
}

// @Summary 新增招聘账号
// @Produce  application/json
// @Param account Object
// @Router /recruit/account [post]
export const addAccount = (data) => {
  return service({
    url: '/recruit/account',
    method: 'post',
    data
  })
}

// // @Summary 新增招聘计划
// // @Produce  application/json
// // @Param plan Object
// // @Router /recruit/plan [post]
// export const addPlan = (data) => {
//   return service({
//     url: '/recruit/plan',
//     method: 'post',
//     data
//   })
// }

// @Summary 根据citycode获取地区
// @Produce  application/json
// @Param citycode int
// @Router /recruit/businessDistrict [get]
export const getBusinessDistrict = (params)=>{
  return service({
    url: '/recruit/businessDistrict',
    method: 'get',
    params
  })
}

// @Summary 新增招聘计划
// @Produce  application/json
// @Param plan Object
// @Router /recruit/plan [post]
export const createRecruitPlan = (data)=>{
  return service({
    url: '/recruit/plan',
    method: 'post',
    data
  })
}

// @Summary 更新招聘计划
// @Produce  application/json
// @Param plan Object
// @Router /recruit/plan [put]
export const updateRecruitPlan = (data)=>{
  return service({
    url: '/recruit/plan',
    method: 'put',
    data
  })
}

// @Summary 禁用招聘计划
// @Produce  application/json
// @Param ID float64
// @Router /recruit/plan [delete]
export const deleteRecruitPlan = (data)=>{
  return service({
    url: '/recruit/plan',
    method: 'delete',
    data
  })
}


// @Summary 获取招聘计划列表
// @Produce  application/json
// @Param {
//  page     int
//	pageSize int
// }
// @Router /recruit/planList [get]
export const getRecruitPlanList = (params)=>{
  return service({
    url: '/recruit/planList',
    method: 'get',
    params
  })
}


// @Summary 添加用户org关联关系
// @Produce  application/json
// @Param orgs Object authorityId string
// @Router /org/getOrgList [post]
export const addOrgAuthority = (data) => {
  return service({
    url: '/org/addOrgAuthority',
    method: 'post',
    data
  })
}

// @Summary 获取用户org关联关系
// @Produce  application/json
// @Param authorityId string
// @Router /org/getOrgAuthority [post]
export const getOrgAuthority = (data) => {
  return service({
    url: '/org/getOrgAuthority',
    method: 'post',
    data
  })
}

// @Summary 删除org
// @Produce  application/json
// @Param ID float64
// @Router /org/deleteBaseOrg [post]
export const deleteBaseOrg = (data) => {
  return service({
    url: '/org/deleteBaseOrg',
    method: 'post',
    data
  })
}

// @Summary 修改org列表
// @Produce  application/json
// @Param org Object
// @Router /org/updateBaseOrg [post]
export const updateBaseOrg = (data) => {
  return service({
    url: '/org/updateBaseOrg',
    method: 'post',
    data
  })
}

// @Tags org
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "根据id获取菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /org/getBaseOrgById [post]
export const getBaseOrgById = (data) => {
  return service({
    url: '/org/getBaseOrgById',
    method: 'post',
    data
  })
}
