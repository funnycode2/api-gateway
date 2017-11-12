package web

//管理界面接口定义

//查询所有服务(按前缀模糊匹配/按过滤器名称/显示过滤器状态(是否启用)), 包括每个服务关联的过滤器
// /admin/service/list?prefix=a&filtername=b&status=1

//修改某个服务
// /admin/service/modify?serviceid=1&

//删除某个服务
// /admin/service/delete?serviceid=1&

//热配置(使网关配置生效)
// /admin/service/reload
