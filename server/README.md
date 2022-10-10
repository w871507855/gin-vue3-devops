# 本项目为调度服务器的面板系统

## 启动项目

### 启动前需要修改conf/config.yaml中的数据库配置
```
go mod init server

go mod tidy

# 生成swagger
~/go/bin/swag init

go run main.go
```
