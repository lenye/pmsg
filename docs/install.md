### 下载

#### 支持的操作系统

* Windows
* Linux
* macOS
* FreeBSD

#### 使用二进制发行版

1. 下载 `pmsg` [最新版本](https://github.com/lenye/pmsg/releases)

1. 开始运行它:

   linux

   ```shell
    $ ./pmsg --help
    Usage:
      pmsg [command]
    
    Available Commands:
      dingtalk    钉钉
      feishu      飞书
      help        Help about any command
      slack       slack
      weixin      微信：公众号、小程序
      workweixin  企业微信
    
    Flags:
      -h, --help      help for pmsg
      -v, --version   version for pmsg
    
    Use "pmsg [command] --help" for more information about a command.   
   ```

#### linux容器映像

1. `Packages` https://github.com/lenye/pmsg/pkgs/container/pmsg

1. 拉取容器映像
   ```shell
   $ docker pull ghcr.io/lenye/pmsg
   ```

1. 开始运行它
   ```shell
   $ docker run --rm ghcr.io/lenye/pmsg --help
   ```

#### 源代码

```shell
$ git clone https://github.com/lenye/pmsg.git
```