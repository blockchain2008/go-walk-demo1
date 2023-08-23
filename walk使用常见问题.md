# walk使用常见问题

1. MainWindow返回错误“TTM_ADDTOOL failed”。
解决方法：https://github.com/lxn/walk/issues/733
```shell
### tc-hib不建议使用，说他可能删掉这个fork
go get -u github.com/tc-hib/rsrc
### 建议使用
go get -u github.com/akavel/rsrc
go install github.com/akavel/rsrc
###
```

```shell
### 查看所有go mod
go list -m all
### 卸载包
go clean -i -r github.com/akavel/rsrc
```