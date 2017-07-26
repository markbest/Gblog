## Gblog
用beego实现的一个博客程序

## 博客特色
- 简单的文章管理
- 简单的文章分类管理
- 前台客户注册以及客户中心管理
- 使用Redis管理博客的缓存
- 资料后台上传、前台下载
- 数据库迁移管理

## 数据库迁移用法
- 先编译为可执行文件:go build -o m migrate.go
- ./m create filename(创建迁移文件)
- ./m up(执行迁移)
- ./m down(回滚迁移）
- ./m status(查看迁移文件的状态) 

## 安装使用
- 从[github](https://github.com/markbest/golang_blog)上下载源代码
- 创建conf文件夹，将app.conf.example重命名为app.conf放置conf下
- 配置数据库信息，执行数据库迁移：./m up
- 博客后台登录地址：http://xxx/admin
