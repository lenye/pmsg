### 企业微信上传临时素材

命令参数说明

```text
$ pmsg workweixin upload -h

-a, --user_agent string     http user agent

-t, --access_token string   企业微信接口调用凭证
-i, --corp_id string        企业微信corp_id
-s, --corp_secret string    企业微信corp_secret

如果没有提供 access_token，需要提供微信 app_id 和 app_secret 获取 access_token

-m, --media_type string     临时素材的格式类型 (必填)，image(图片)、voice(语音)、video(视频)、file(普通文件)、
    --raw                   文件名称含路径是原始字符串字面值，没有任何转义处理

args                        参数：文件名称含路径，默认是解释字符串，支持“\”转义
```

样例

linux

```shell
$ pmsg workweixin upload -i corp_id -s corp_secret -m image --raw /img/app.png

ok; type: "image", media_id: "media_id", created_at: 1670301152 (2022-12-06T12:32:32+08:00)
```

官方开发文档 [企业微信上传临时素材](https://developer.work.weixin.qq.com/document/path/90389)
