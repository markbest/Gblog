## Gblog
用beego实现的一个博客程序

## feature
- 文章管理
- 分类管理
- 用户管理
- 使用Redis缓存数据
- 数据库迁移管理
- 图片上传管理
- 博客主题：[zanblog](http://www.yeahzan.com/zanblog)

## usage
- 从[Github](https://github.com/markbest/Gblog)上下载源代码
- 进入conf文件夹，复制app.conf.example为app.conf
- 编译可执行文件：
```
//数据库迁移工具
go build -o bin/migrate migrate.go
//程序入口
go build -o bin/blog main.go
```
- 配置数据库信息，执行数据库迁移：bin/migrate up
- 启动服务bin/blog，博客后台登录地址：http://127.0.0.1:8080/admin

## database migrate
- bin/migrate create filename(创建迁移文件)
- bin/migrate up(执行迁移)
- bin/migrate down(回滚迁移）
- bin/migrate status(查看迁移文件的状态) 