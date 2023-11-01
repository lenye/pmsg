### 微信新增临时素材

命令参数说明

```text
$ pmsg weixin upload -h

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供 access_token，需要提供微信 app_id 和 app_secret 获取 access_token

-m, --media_type string     临时素材的格式类型 (必填)，image(图片)、voice(语音)、video(视频)、thumb(缩略图)、
    --raw                   文件名称含路径是原始字符串字面值，没有任何转义处理

args                        参数：文件名称含路径，默认是解释字符串，支持“\”转义
```

样例

linux

```shell
$ pmsg weixin upload -i app_id -s app_secret -m image --raw /img/app.png

ok; type: "image", media_id: "media_id", created_at: 1669210730 (2022-11-23T21:38:50+08:00)
```

官方开发文档 [微信公众号新增临时素材](https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html)
