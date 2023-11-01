### 推送钉钉自定义机器人消息

命令参数说明

```text
$ pmsg dingtalk bot -h

-a, --user_agent string     http user agent

-t, --access_token string   钉钉自定义机器人 access token (必填)
-s, --secret string         签名密钥
-m, --msg_type string       消息类型 (必填)，text(文本消息)、link(链接)、markdown(markdown消息)、
                                           single_actionCard(整体跳转actionCard)、actionCard(独立跳转actionCard)、
                                           feedCard
-o, --at_user string        文本或markdown消息时，被@人的用户userid，多个接收者用‘|’分隔。
-b, --at_mobile string      文本或markdown消息时，被@人的手机号，多个接收者用‘|’分隔
-i, --is_at_all             文本或markdown消息时，是否@所有人
    --raw                   消息内容是原始字符串字面值，没有任何转义处理

args                        参数：消息内容，默认是解释字符串，支持“\”转义
```

消息内容

1. 文本消息 --msg_type text
    ```text
    HelloWorld
    ```
1. 链接消息 --msg_type link
   ```json
   {
     "text": "这个即将发布的新版本，创始人xx称它为红树林。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是红树林",
     "title": "时代的火车向前开",
     "picUrl": "",
     "messageUrl": "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
   }
   ```

1. markdown消息 --msg_type markdown
   ```json
   {
     "title": "杭州天气",
     "text": "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n"
   }
   ```

1. actionCard整体跳转消息 --msg_type single_actionCard
   ```json
   {
      "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
      "text": "![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)### 乔布斯 20 年前想打造的苹果咖啡厅Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
      "btnOrientation": "0",
      "singleTitle": "阅读全文",
      "singleURL": "https://www.dingtalk.com/"
   }
   ```

1. actionCard独立跳转消息 --msg_type actionCard
   ```json
   {
      "title": "我 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
      "text": "![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
      "btnOrientation": "0",
      "btns": [
         {
            "title": "内容不错",
            "actionURL": "https://www.dingtalk.com/"
         },
         {
            "title": "不感兴趣",
            "actionURL": "https://www.dingtalk.com/"
         }
      ]
   }
   ```

1. feedCard消息 --msg_type feedCard
   ```json
   {
     "links": [
       {
         "title": "时代的火车向前开1",
         "messageURL": "https://www.dingtalk.com/",
         "picURL": "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png"
       },
       {
         "title": "时代的火车向前开2",
         "messageURL": "https://www.dingtalk.com/",
         "picURL": "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png"
       }
     ]
   }
   ```

样例

linux

```shell
$ pmsg dingtalk bot -t access_token -m text 'HelloWorld'

ok
```

官方开发文档 [推送钉钉自定义机器人消息](https://open.dingtalk.com/document/robots/custom-robot-access)