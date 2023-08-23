# go 的 windows gui 开发包walk的使用

```shell

### 科学上网测试
Get-Alias -Definition Invoke-WebRequest | Format-Table -AutoSize
###
curl https://www.google.com | Select -ExpandProperty Content
### 安装walk
go get github.com/lxn/walk
### 安装rsrc
go get -u github.com/akavel/rsrc
go install github.com/akavel/rsrc
rsrc -manifest project.exe.manifest -o rsrc.syso
###

```

```shell 清理包
### 清理$GOPATH/src目录下的包
go clean -cache
### 清理$GOPATH/pkg/mod目录下的包
go clean -modcache

```



















































//----------------------------------------------------------------------------------


```shell 一些遇到的问题提示
### 添加项目到安全目录
### git config --global --add safe.directory E:/__Go_Blockchain_2023/GoProjects/src/github.com/lxn/walk

###
### git clone https://github.com/lxn/walk.git
```