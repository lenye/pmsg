### 撤回企业微信应用消息

命令参数说明

```text
$ pmsg weixin work app undo -h

-a, --user_agent string     http user agent

-t, --access_token string   企业微信接口调用凭证
-i, --corp_id string        企业微信corp_id
-s, --corp_secret string    企业微信corp_secret

如果没有提供 access_token，需要提供企业微信 corp_id 和 corp_secret 来获取 access_token

-c, --msg_id string         消息ID。从应用发送消息接口处获得
```

样例

linux

发送文本消息

```shell
$ pmsg weixin work app undo -i corp_id -s corp_secret -c msg_id

ok
```

[撤回企业微信应用消息开发文档](https://developer.work.weixin.qq.com/document/path/94867)