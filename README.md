`pmsg`是一个发送消息的小工具。

### 支持的操作系统

* Windows
* Linux
* macOS
* FreeBSD

### 支持的消息

1. 微信公众号
    * 模板消息
    * 一次性订阅消息
    * 订阅通知
    * 客服消息

1. 微信小程序
    * 订阅消息
    * 客服消息

1. 企业微信
    * 应用消息
    * 群聊消息
    * 家校消息
    * 互联企业消息
    * 客服消息
    * 群机器人消息

1. 钉钉
    * 自定义机器人消息

1. 飞书
    * 自定义机器人消息

### 下载

#### 使用二进制发行版

1. 下载 `pmsg` [最新版本](https://github.com/lenye/pmsg/releases)

1. 开始运行它:

   linux

   ```shell
   ./pmsg --help
   
   Usage:
     pmsg [command]
   
   Available Commands:
     dingtalk    ding talk
     feishu      fei shu
     help        Help about any command
     weixin      weixin
     workweixin  work weixin
   
   Flags:
     -h, --help      help for pmsg
     -v, --version   version for pmsg
   
   Use "pmsg [command] --help" for more information about a command.
   ```

#### 源代码

```shell
git clone https://github.com/lenye/pmsg.git
```

### 文档

[https://github.com/lenye/pmsg/tree/main/docs](https://github.com/lenye/pmsg/tree/main/docs)

### License

`pmsg` is released under the [Apache 2.0 license](https://github.com/lenye/pmsg/blob/main/LICENSE). 