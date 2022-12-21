### 企业微信群机器人文件上传

命令参数说明

```text
$ pmsg workweixin bot upload -h

-a, --user_agent string     http user agent

-k, --key string            企业微信群机器人key (必填)

args                        参数：文件名称含路径
```

样例

linux

```shell
$ pmsg workweixin bot upload -k key /img/app.png

ok; type: "file", media_id: "media_id", created_at: 1670472324 (2022-12-08T12:05:24+08:00)
```

官方开发文档 [企业微信群机器人文件上传](https://developer.work.weixin.qq.com/document/path/91770#文件上传接口)