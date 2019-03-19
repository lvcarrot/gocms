# GoCMS
基于 [AdminLTE](https://adminlte.io)、[Gorilla](http://www.gorillatoolkit.org) 和 [Gorm](http://gorm.io) 实现的内容管理系统

## 获取安装

执如下命令，就能够在你的 `$GOPATH/bin` 目录下发现 gocms
```bash
git clone https://github.com/dragonflylee/gocms.git $GOPATH/gocms
go get -v gocms
```

## 目录结构

 ├── handler    Web业务逻辑  
 ├── model      数据操作层  
 ├── static     前端静态资源  
 ├── util       工具函数  
 ├── views      模板页面  
 ├── main.go    路由入口  
 └── nodes.json 节点初始化数据  

