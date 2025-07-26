`pmsg`是一个多平台消息推送的小工具、SDK、API。

## 支持的消息

### WebHook

* 企业微信群机器人消息
* 钉钉自定义机器人消息
* 飞书自定义机器人消息
* Slack机器人消息
* Discord机器人消息

### 微信

* 微信公众号
    * 模板消息
    * 一次性订阅消息
    * 订阅通知
    * 客服消息

* 微信小程序
    * 订阅消息
    * 客服消息

### 企业微信

* 应用消息
* 群聊消息
* 家校消息
* 互联企业消息
* 客服消息

## 支持的操作系统

* Windows
* Linux
* macOS
* FreeBSD

## 文档

https://github.com/lenye/pmsg/tree/main/docs

## 下载

### 使用二进制发行版

1. 下载 `pmsg` [最新版本](https://github.com/lenye/pmsg/releases)

1. 开始运行它:

   linux

   ```shell
    $ ./pmsg --help
    Usage:
      pmsg [command]
    
    Available Commands:
      dingtalk    钉钉
      discord     discord
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

### linux容器映像

1. `Packages` https://github.com/lenye/pmsg/pkgs/container/pmsg

1. 拉取容器映像
   ```shell
   $ docker pull ghcr.io/lenye/pmsg
   ```

1. 开始运行它
   ```shell
   $ docker run --rm ghcr.io/lenye/pmsg --help
   ```

### 源代码

```shell
$ git clone https://github.com/lenye/pmsg.git
```

## 贡献

欢迎创建 [Issue](https://github.com/lenye/pmsg/issues) 来帮助改进项目。

## License

`pmsg` is released under the [Apache 2.0 license](https://github.com/lenye/pmsg/blob/main/LICENSE).