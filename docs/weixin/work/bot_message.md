### 推送企业微信群机器人消息

命令参数说明

```text
$ pmsg workweixin bot -h

-a, --user_agent string     http user agent

-k, --key string            企业微信群机器人key (必填)
-m, --msg_type string       消息类型 (必填)，text(文本消息)、markdown(markdown消息)、
                                           image(图片消息)、news(图文消息)、file(文件消息)、
                                           text_notice(文本通知模版卡片)、news_notice(图文展示模版卡片)
-o, --at_user string        文本消息时，提醒群中的指定成员(@某个成员)，多个接收者用‘|’分隔，@all表示提醒所有人。
                            如果开发者获取不到userid，可以使用at_mobile
-b, --at_mobile string      文本消息时，提醒手机号对应的群成员(@某个成员)，多个接收者用‘|’分隔，@all表示提醒所有人
    --raw                   消息内容是原始字符串字面值，没有任何转义处理

args                        参数：消息内容，默认是解释字符串，支持“\”转义
```

消息内容

1. 文本消息 --msg_type text
    ```text
    HelloWorld
    ```

1. markdown消息 --msg_type markdown
    ```text
    markdown
    ```

1. 图片消息 --msg_type image
    ```text
    文件名称含路径
    ```

1. 图文消息 --msg_type news
   ```json
   {
      "articles": [
         {
            "title": "中秋节礼品领取",
            "description": "今年中秋节公司有豪礼相送",
            "url": "www.qq.com",
            "picurl": "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"
         }
      ]
   }
   ```

1. 文件消息 --msg_type file
    ```text
    MEDIA_ID
    ```

1. 文本通知模版卡片 --msg_type text_notice
   ```json
   {
      "card_type": "text_notice",
      "source": {
         "icon_url": "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
         "desc": "企业微信",
         "desc_color": 0
      },
      "main_title": {
         "title": "欢迎使用企业微信",
         "desc": "您的好友正在邀请您加入企业微信"
      },
      "emphasis_content": {
         "title": "100",
         "desc": "数据含义"
      },
      "quote_area": {
         "type": 1,
         "url": "https://work.weixin.qq.com/?from=openApi",
         "appid": "APPID",
         "pagepath": "PAGEPATH",
         "title": "引用文本标题",
         "quote_text": "Jack：企业微信真的很好用~\nBalian：超级好的一款软件！"
      },
      "sub_title_text": "下载企业微信还能抢红包！",
      "horizontal_content_list": [
         {
            "keyname": "邀请人",
            "value": "张三"
         },
         {
            "keyname": "企微官网",
            "value": "点击访问",
            "type": 1,
            "url": "https://work.weixin.qq.com/?from=openApi"
         },
         {
            "keyname": "企微下载",
            "value": "企业微信.apk",
            "type": 2,
            "media_id": "MEDIAID"
         }
      ],
      "jump_list": [
         {
            "type": 1,
            "url": "https://work.weixin.qq.com/?from=openApi",
            "title": "企业微信官网"
         },
         {
            "type": 2,
            "appid": "APPID",
            "pagepath": "PAGEPATH",
            "title": "跳转小程序"
         }
      ],
      "card_action": {
         "type": 1,
         "url": "https://work.weixin.qq.com/?from=openApi",
         "appid": "APPID",
         "pagepath": "PAGEPATH"
      }
   }
   ```

1. 图文展示模版卡片 --msg_type news_notice
   ```json
   {
      "card_type": "news_notice",
      "source": {
         "icon_url": "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
         "desc": "企业微信",
         "desc_color": 0
      },
      "main_title": {
         "title": "欢迎使用企业微信",
         "desc": "您的好友正在邀请您加入企业微信"
      },
      "card_image": {
         "url": "https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0",
         "aspect_ratio": 2.25
      },
      "image_text_area": {
         "type": 1,
         "url": "https://work.weixin.qq.com",
         "title": "欢迎使用企业微信",
         "desc": "您的好友正在邀请您加入企业微信",
         "image_url": "https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0"
      },
      "quote_area": {
         "type": 1,
         "url": "https://work.weixin.qq.com/?from=openApi",
         "appid": "APPID",
         "pagepath": "PAGEPATH",
         "title": "引用文本标题",
         "quote_text": "Jack：企业微信真的很好用~\nBalian：超级好的一款软件！"
      },
      "vertical_content_list": [
         {
            "title": "惊喜红包等你来拿",
            "desc": "下载企业微信还能抢红包！"
         }
      ],
      "horizontal_content_list": [
         {
            "keyname": "邀请人",
            "value": "张三"
         },
         {
            "keyname": "企微官网",
            "value": "点击访问",
            "type": 1,
            "url": "https://work.weixin.qq.com/?from=openApi"
         },
         {
            "keyname": "企微下载",
            "value": "企业微信.apk",
            "type": 2,
            "media_id": "MEDIAID"
         }
      ],
      "jump_list": [
         {
            "type": 1,
            "url": "https://work.weixin.qq.com/?from=openApi",
            "title": "企业微信官网"
         },
         {
            "type": 2,
            "appid": "APPID",
            "pagepath": "PAGEPATH",
            "title": "跳转小程序"
         }
      ],
      "card_action": {
         "type": 1,
         "url": "https://work.weixin.qq.com/?from=openApi",
         "appid": "APPID",
         "pagepath": "PAGEPATH"
      }
   }
   ```

样例

linux

```shell
$ pmsg workweixin bot -k key -m text 'HelloWorld'

ok

$ pmsg wwx bot -k key -m image /img/app.png

ok
```

官方开发文档 [推送企业微信群机器人消息](https://developer.work.weixin.qq.com/document/path/91770)