## Gblog
用beego实现的一个博客程序

## 博客特色
- 简单的文章管理
- 简单的分类管理
- 使用Redis管理博客的缓存
- 数据库迁移管理

## 数据库迁移用法
- 先编译为可执行文件:go build -o bin/migrate migrate.go
- bin/migrate create filename(创建迁移文件)
- bin/migrate up(执行迁移)
- bin/migrate down(回滚迁移）
- bin/migrate status(查看迁移文件的状态) 

## 安装使用
- 从[Github](https://github.com/markbest/Gblog)上下载源代码
- 进入conf文件夹，复制app.conf.example为app.conf
- 配置数据库信息，执行数据库迁移：bin/migrate up
- 安装bee工具：go get github.com/beego/bee，执行bee run -main=main.go
- 博客后台登录地址：http://127.0.0.1:8080/admin