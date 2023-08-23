# Go代理配置
http://www.taodudu.cc/news/show-3681779.html?action=onClick
https://goproxy.io/zh/

```shell -------------Bash (Linux or macOS)

# 配置 GOPROXY 环境变量
export GOPROXY=https://proxy.golang.com.cn,direct
# 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
export GOPRIVATE=git.mycompany.com,github.com/my/private

```

```shell -------------PowerShell (Windows)

# 配置 GOPROXY 环境变量
$env:GOPROXY = "https://proxy.golang.com.cn,direct"
# 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
$env:GOPRIVATE = "git.mycompany.com,github.com/my/private"

```