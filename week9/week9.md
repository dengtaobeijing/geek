总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用

## fix length
* 使用包定长的方式来解决粘包问题,每次发送固定长度的数据包。

## delimiter based
* 通过自定义分隔符解决粘包问题。添加特殊符号，接收方通过这个特殊符号将接收到的数据包拆分开。

## length field based frame decoder
* 自定义长度解码器解决粘包问题。在消息头中定义长度字段，来标识消息的总长度。
