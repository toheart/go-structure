# go-structure
golang clean structure。

# 初始化

```shell
# windows下先执行 
./init.ps1

make init 
```
> 安装make工具:  https://gnuwin32.sourceforge.net/packages/make.htm <br>
> 设置环境变量.

# 运行
```shell
make server
```

# 其他命令
* 添加命令
```shell
cobra-cli.exe add <command> 
```
* 编译api 
`make api`
* 编译配置
`make config`
* wire: `make generate`