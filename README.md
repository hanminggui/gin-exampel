# gin-exampel
### 开发环境配置
1. 设置github token  可以提升从github下载项目的速度
- 打开 https://github.com/settings/tokens 点击 "Generate New Token" 生成 有 “public_repo” 权限的token
- 编辑 ～/.netrc 写入 
```
machine api.github.com
  login <username>
  password <token>
```
2. 下载项目/编译/下载第三方包
```
go get github.com/hanminggui/gin-exampel
```
Linux
```
cd ~/go/src/github.com/hanminggui/gin-exampel
./install.sh
```
Windows
```
cd C:
cd C:\Users\song\go\src\github.com\hanminggui\gin-exampel
install.bat
```
### 运行环境配置
- Linux
```
export GIN_MODE=release
```
- Windows
```
set GIN_MODE=release
```

### 启动
- build
```
go build
```
- run
Linux
```
go run #开发环境 编译并运行
./exampel #编译后运行
```
Windows
```
gin-exampel.exe #编译后运行
```

### 文件说明
- main.go  程序入口
- install.sh  包管理文件
- routers/  路由
- apis/  接口逻辑
- moudles/  对象模型
- lib/  工具包
- database/  初始化mysql
- config/  初始化配置文件  生产环境部署时需要把配置文件拿过去
