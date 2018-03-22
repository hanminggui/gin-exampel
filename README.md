# gin-exampel

### 设置环境变量
- debug
```
export GIN_MODE=debug
```
- release
```
export GIN_MODE=release
```

### 启动
- build
```
go build main.go -o exampel
```
- run
```
go run main.go #开发环境 编译并运行
./exampel #编译后运行
```

### 文件说明
- main.go  程序入口
- routers/  路由
- apis/  接口逻辑
- moudles/  对象模型
- lib/  工具包
- database/  初始化mysql
- config/  初始化配置文件  生产环境部署时需要把配置文件拿过去