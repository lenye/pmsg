### 发送企业微信应用消息

命令参数说明

```text
$ pmsg workweixin app -h

-a, --user_agent string     http user agent

-t, --access_token string   企业微信接口调用凭证
-i, --corp_id string        企业微信corp_id
-s, --corp_secret string    企业微信corp_secret

如果没有提供 access_token，需要提供企业微信 corp_id 和 corp_secret 获取 access_token

-o, --to_user string        指定接收消息的成员，成员ID列表，最多支持1000个，多个接收者用‘|’分隔。指定为"@all"，则向该企业应用的全部成员发送
-p, --to_party string       指定接收消息的部门，部门ID列表，最多支持100个，多个接收者用‘|’分隔。to_user"@all"时忽略本参数
-g, --to_tag string         指定接收消息的标签，标签ID列表，最多支持100个，多个接收者用‘|’分隔。to_user"@all"时忽略本参数
-e, --agent_id int          企业应用的id (必填)
-m, --msg_type string       消息类型 (必填)，text(文本消息)、image(图片消息)、
                                           voice(语音消息)、video(视频消息)、file(文件消息)、
                                           textcard(文本卡片消息)、news(图文消息)、mpnews(图文消息)、
                                           markdown(markdown消息)、miniprogram_notice(小程序通知消息)、
                                           template_card(模板卡片消息)

-d, --duplicate_check_interval int   表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时(14400)
-c, --enable_duplicate_check int     表示是否开启重复消息检查，0表示否，1表示是，默认0
-r, --enable_id_trans int            表示是否开启id转译，0表示否，1表示是，默认0。仅第三方应用需要用到，企业自建应用可以忽略。
    --safe int                       表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0
    --raw                   消息内容是原始字符串字面值，没有任何转义处理

args                        参数：消息内容，默认是解释字符串，支持“\”转义    
```

消息内容

1. 文本消息 --msg_type text
    ```text
    HelloWorld
    ```

1. 图片消息 --msg_type image
    ```text
    MEDIA_ID
    ```

1. 语音消息 --msg_type voice
    ```text
    MEDIA_ID
    ```

1. 视频消息 --msg_type video
    ```json
    {
      "media_id": "MEDIA_ID",
      "title": "TITLE",
      "description": "DESCRIPTION"
    }
    ```

1. 文件消息 --msg_type file
    ```text
    MEDIA_ID
    ```

1. 文本卡片消息 --msg_type textcard
    ```json
    {
      "title": "title",
      "description": "description",
      "url": "URL",
      "btntxt": "更多"
    }
    ```

1. 图文消息 --msg_type news
   ```json
   {
     "articles": [
       {
         "title": "中秋节礼品领取",
         "description": "今年中秋节公司有豪礼相送",
         "url": "URL",
         "picurl": "http://res.mail.qq.com/msg.png",
         "appid": "wx123123123123123",
         "pagepath": "pages/index?userid=zhangsan&orderid=123123123"
       }
     ]
   }
   ```

1. 图文消息 --msg_type mpnews
   ```json
   {
     "articles": [
       {
         "title": "Title",
         "thumb_media_id": "MEDIA_ID",
         "author": "Author",
         "content_source_url": "URL",
         "content": "Content",
         "digest": "Digest description"
       }
     ]
   }
   ```

1. markdown消息 --msg_type markdown
    ```text
    markdown
    ```

1. 小程序通知消息 --msg_type miniprogram_notice
   ```json
   {
     "appid": "wx123123123123123",
     "page": "pages/index?userid=zhangsan&orderid=123123123",
     "title": "会议室预订成功通知",
     "description": "4月27日 16:16",
     "emphasis_first_item": true,
     "content_item": [
       {
         "key": "会议室",
         "value": "402"
       },
       {
         "key": "会议地点",
         "value": "广州TIT-402会议室"
       },
       {
         "key": "会议时间",
         "value": "2018年8月1日 09:00-09:30"
       },
       {
         "key": "参与人员",
         "value": "周剑轩"
       }
     ]
   }
   ```

1. 模板卡片消息 --msg_type template_card
   ```json
   {
     "card_type": "text_notice",
     "source": {
       "icon_url": "图片的url",
       "desc": "企业微信",
       "desc_color": 1
     },
     "action_menu": {
       "desc": "卡片副交互辅助文本说明",
       "action_list": [
         {
           "text": "接受推送",
           "key": "A"
         },
         {
           "text": "不再推送",
           "key": "B"
         }
       ]
     },
     "task_id": "task_id",
     "main_title": {
       "title": "欢迎使用企业微信",
       "desc": "您的好友正在邀请您加入企业微信"
     },
     "quote_area": {
       "type": 1,
       "url": "https://work.weixin.qq.com",
       "title": "企业微信的引用样式",
       "quote_text": "企业微信真好用呀真好用"
     },
     "emphasis_content": {
       "title": "100",
       "desc": "核心数据"
     },
     "sub_title_text": "下载企业微信还能抢红包！",
     "horizontal_content_list": [
       {
         "keyname": "邀请人",
         "value": "张三"
       },
       {
         "type": 1,
         "keyname": "企业微信官网",
         "value": "点击访问",
         "url": "https://work.weixin.qq.com"
       },
       {
         "type": 2,
         "keyname": "企业微信下载",
         "value": "企业微信.apk",
         "media_id": "文件的media_id"
       },
       {
         "type": 3,
         "keyname": "员工信息",
         "value": "点击查看",
         "userid": "zhangsan"
       }
     ],
     "jump_list": [
       {
         "type": 1,
         "title": "企业微信官网",
         "url": "https://work.weixin.qq.com"
       },
       {
         "type": 2,
         "title": "跳转小程序",
         "appid": "小程序的appid",
         "pagepath": "/index.html"
       }
     ],
     "card_action": {
       "type": 2,
       "url": "https://work.weixin.qq.com",
       "appid": "小程序的appid",
       "pagepath": "/index.html"
     }
   }
   ```

样例

linux

发送文本消息

```shell
$ pmsg workweixin app -i corp_id -s corp_secret -e agent_id -o '@all' -m text 'HelloWorld'

ok; msgid: "msgid"
```

官方开发文档 [企业微信应用消息](https://developer.work.weixin.qq.com/document/path/90236)