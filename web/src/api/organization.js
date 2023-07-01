import service from '@/utils/request'
// @Summary 用户登录 获取动态路由
// @Produce  application/json
// @Param 可以什么都不填 调一下即可
// @Router /org/getOrg [post]
export const asyncOrg = () => {
  return service({
    url: '/org/getOrg',
    method: 'post'
  })
}

// @Summary 获取org列表
// @Produce  application/json
// @Param {
//  page     int
//	pageSize int
// }
// @Router /org/getOrgList [get]
export const getOrgList = (params) => {
  return service({
    url: '/org/getOrgList',
    method: 'get',
    params
  })
}

// @Summary 新增基础org
// @Produce  application/json
// @Param org Object
// @Router /org/createOrg [post]
export const createOrg = (data) => {
  return service({
    url: '/org/createOrg',
    method: 'post',
    data
  })
}

// @Summary 获取基础路由列表
// @Produce  application/json
// @Param 可以什么都不填 调一下即可
// @Router /org/getBaseOrgTree [post]
export const getBaseOrgTree = () => {
  return service({
    url: '/org/getBaseOrgTree',
    method: 'post'
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
