### 发送微信小程序客服消息

命令参数说明

```text
$ pmsg weixin mini customer -h

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供access_token，需要提供微信app_id和app_secret来获取access_token

-o, --open_id string        接收人的open_id (必填)
    --type string           消息类型 (必填)，text(文本消息)、image(图片消息)、link(图文链接)、miniprogrampage(小程序卡片)
```

消息内容
1. 文本消息 --type text
    ```json
    {
      "content": "HelloWorld"
    }
    ```

1. 图片消息 --type image
    ```json
    {
      "media_id": "MEDIA_ID"
    }
    ```

1. 图文链接 --type link
    ```json
    {
      "title": "Happy Day",
      "description": "Is Really A Happy Day",
      "url": "URL",
      "thumb_url": "THUMB_URL"
    }
    ```

1. 小程序卡片 --type miniprogrampage
    ```json
    {
      "title": "title",
      "pagepath": "pagepath",
      "thumb_media_id": "thumb_media_id"
    }
    ```

样例

windows

```shell
pmsg.exe weixin mini customer -i app_id -s app_secret -o open_id --type text "{\"content\":\"HelloWorld\"}"

ok
```

linux

```shell
$ pmsg weixin mini customer -i app_id -s app_secret -o open_id --type text '{"content":"HelloWorld"}'

ok
```

[微信小程序客服消息开发文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html)