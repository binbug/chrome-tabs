

## 说明

一个简单的通过cdp协议控制chrome的tab的工具，用于通过alfred来控制chrome的tab。

除了控制已有tab的功能外，还支持：
1. 打开新的tab（通过自定义tab的url来实现）
2. google搜索
3. github搜索

## 使用

1. 下载[chrome-tabs-v1.0.1.zip](https://github.com/binbug/chrome-tabs/releases/download/v1.0.0/chrome-tabs-v1.0.1.zip)
2. 双击chrome-tabs-v1.0.1.alfredworkflow安装
3. 运行服务端程序：
```shell
# 授予权限
chmod +x chrome-tabs-darwin 
# 运行
./chrome-tabs-darwin -rod-bin="/Applications/Microsoft\ Edge.app/Contents/MacOS/Microsoft\ Edge"

# 

```
