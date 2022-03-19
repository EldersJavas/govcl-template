# govcl-template GitHub模板

govcl template repo

govcl repo: <https://github.com/ying32/govcl>

适配`Go 1.15`以上版本.  
CI构建使用`Go 1.17`.
## How to use

### Go project ([go.mod](./go.mod))

两种方法:  

1. 把`go.mod`文件里的包名改成自己的.(可能会出现未知错误)
2. 把`go.mod`删了,重新`go mod init 包名`.

### Govcl GUI ([project.lpi](./gui/project.lpi))

改个项目名行了.

### CI 构建 ([go-build.yml](./.github/workflows/go-build.yml))

把`.github\workflows\go-build.yml`里面的所有`govcl-template`替换成自己的项目名.  
自带忽略非代码文件和Artifacts.

### 本地构建

可以先用`init.sh`升级下包.  
然后,对应相应的系统运行批处理脚本

| System | File |
| :------------:  | :------------:  |
| Windows(默认64bit) | [build-win.bat](./build-win.bat) |
| Linux | [build-linux.sh](./build-linux.sh) |
| MacOS | 穷,谁帮我写个 |

### 其他

别忘了把`.github\FUNDING.yml`文件删了.

`.gitignore`文件可以增减,里面有注释.

`LICENSE`是本项目的许可证,可以在自己项目里替换.

其他随意.






