

## 说明

一个简单的通过cdp协议控制chrome的tab的工具，用于通过alfred来控制chrome的tab。

除了控制已有tab的功能外，还支持：
1. 打开新的tab（通过自定义tab的url来实现）
2. google搜索
3. github搜索

## 使用

1. 下载[chrome-tabs-v1.0.1.zip](https://github.com/binbug/chrome-tabs/releases/download/v1.0.1/chrome-tabs-v1.0.1.zip)
2. 双击chrome-tabs-v1.0.1.alfredworkflow安装
3. 运行服务端程序：

```shell
# 授予权限
chmod +x chrome-tabs-darwin 
# 运行
./chrome-tabs-darwin

# 指定浏览器运行
./chrome-tabs-darwin -rod-bin="/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge"

```

运行前需要确保浏览器已经关闭，由服务端程序来启动浏览器。服务端程序关闭时浏览器也会关闭。
可以使用tmux或者nohup来运行服务端程序。


## 添加自定义tab

如果需要添加自定义tab，服务端启动后，可以调用相关接口来添加自定义tab。

```shell
# 添加自定义tab
curl  -X POST \
  'http://127.0.0.1:8787/extra-page/add' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
  --data-raw '{
        "url": "https://github.com/binbug/chrome-tabs",
        "title": "github chrome-tabs",
        "match_regexp": "https://github.com/binbug/chrome-tabs.*",
        "overwrite_title": true
    }'
# 删除自定义tab key 为url

curl  -X POST \
  'http://127.0.0.1:8787/extra-page/delete' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'key=https://github.com/binbug/chrome-tabs'
```