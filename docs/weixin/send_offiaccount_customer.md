### 发送微信公众号客服消息

命令参数说明

```text
$ pmsg weixin offiaccount customer -h

Aliases:
  customer, kf

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供access_token，需要提供微信app_id和app_secret来获取access_token

-k, --kf_account string     客服帐号
-o, --open_id string        接收人的open_id (必填)
    --type string           消息类型 (必填)，text(文本消息)、image(图片消息)、
                                            voice(语音消息)、video(视频消息)、music(音乐消息)、
                                            news(图文消息)、mpnews(图文消息)、mpnewsarticle(图文消息)、
                                            msgmenu(菜单消息)、wxcard(卡券)、miniprogrampage(小程序卡片)
```

消息内容

1. 文本消息 --type text
    ```text
    "HelloWorld"
    ```

1. 图片消息 --type image
    ```text
    "MEDIA_ID"
    ```

1. 语音消息 --type voice
    ```text
    "MEDIA_ID"
    ```

1. 视频消息 --type video
    ```json
    {
      "media_id": "MEDIA_ID",
      "thumb_media_id": "MEDIA_ID",
      "title": "TITLE",
      "description": "DESCRIPTION"
    }
    ```

1. 音乐消息 --type music
    ```json
    {
      "title": "MUSIC_TITLE",
      "description": "MUSIC_DESCRIPTION",
      "musicurl": "MUSIC_URL",
      "hqmusicurl": "HQ_MUSIC_URL",
      "thumb_media_id": "THUMB_MEDIA_ID"
    }
    ```

1. 图文消息 --type news
    ```json
    {
      "articles": [
        {
          "title": "Happy Day",
          "description": "Is Really A Happy Day",
          "url": "URL",
          "picurl": "PIC_URL"
        }
      ]
    }
    ```

1. 图文消息 --type mpnews
    ```text
    "MEDIA_ID"
    ```

1. 图文消息 --type mpnewsarticle
    ```text
    "ARTICLE_ID"
    ```

1. 菜单消息 --type msgmenu
    ```json
    {
      "head_content": "您对本次服务是否满意呢? ",
      "list": [
        {
          "id": "101",
          "content": "满意"
        },
        {
          "id": "102",
          "content": "不满意"
        }
      ],
      "tail_content": "欢迎再次光临"
    }
    ```

1. 卡券 --type wxcard
    ```text
    "CARD_ID"
    ```

1. 小程序卡片 --type miniprogrampage
    ```json
    {
      "title": "title",
      "appid": "appid",
      "pagepath": "pagepath",
      "thumb_media_id": "thumb_media_id"
    }
    ```

样例

windows

发送文本消息

```shell
pmsg.exe weixin offiaccount customer -i app_id -s app_secret -o open_id --type text "HelloWorld"

ok
```

发送小程序卡片

```shell
pmsg.exe weixin offiaccount customer -i app_id -s app_secret -o open_id --type miniprogrampage "{\"title\":\"title\",\"appid\":\"appid\",\"pagepath\":\"pagepath\",\"thumb_media_id\":\"thumb_media_id\"}"

ok
```

linux

发送文本消息

```shell
$ pmsg weixin offiaccount customer -i app_id -s app_secret -o open_id --type text 'HelloWorld'

ok
```

发送小程序卡片

```shell
pmsg weixin offiaccount customer -i app_id -s app_secret -o open_id --type miniprogrampage '{"title":"title","appid":"appid","pagepath":"pagepath","thumb_media_id":"thumb_media_id"}'

ok
```

[微信公众号客服消息开发文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html#7)