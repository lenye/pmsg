### 获取企业微信接口调用凭证

命令参数说明

```text
$ pmsg workweixin token -h

-a, --user_agent string     http user agent

-i, --corp_id string        企业微信corp_id (必填)
-s, --corp_secret string    企业微信corp_secret (必填)
```

样例

linux

```shell
$ pmsg weixin work token -i corp_id -s corp_secret

ok; access_token: "access_token", expires_in: 7200, expire_at: "2022-09-20T15:00:20+08:00"
```

官方开发文档 [获取企业微信接口调用凭证](https://developer.work.weixin.qq.com/document/path/91039)