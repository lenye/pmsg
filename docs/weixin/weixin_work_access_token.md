### 获取企业微信接口调用凭证

命令参数说明

```text
$ pmsg weixin work token -h

-a, --user_agent string     http user agent

-i, --corp_id string        企业微信corp_id (必填)
-s, --corp_secret string    企业微信corp_secret (必填)
```

样例

```shell
$ pmsg weixin work token -i app_id -s app_secret

使用命令别名
$ pmsg wx token -i corp_id -s corp_secret

ok; access_token: "access_token", expires_in: 7200, expire_at: "2022-09-20T15:00:20+08:00"
```

[获取企业微信接口调用凭证开发文档](https://developer.work.weixin.qq.com/document/path/91039)