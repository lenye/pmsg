### 撤回企业微信应用消息

命令参数说明

```text
$ pmsg workweixin app undo -h

-a, --user_agent string     http user agent

-t, --access_token string   企业微信接口调用凭证
-i, --corp_id string        企业微信corp_id
-s, --corp_secret string    企业微信corp_secret

如果没有提供 access_token，需要提供企业微信 corp_id 和 corp_secret 获取 access_token

args                        参数：消息ID。从应用发送消息接口处获得
```

样例

linux

```shell
$ pmsg workweixin app undo -i corp_id -s corp_secret msg_id

ok
```

官方开发文档 [撤回企业微信应用消息](https://developer.work.weixin.qq.com/document/path/94867)