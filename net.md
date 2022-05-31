# http 各个版本区别
```
1. http 0.9  最基本的请求，只有GET
2. http 1.0  支持GET POST HEAD ，增加消息头，Cache，可以传输二进制等。
3. http 1.1  持久连接, 管道传输机制，增加HOST头
4. http 2.0  二级制协议，多路复用技术。
```

# websocket

### 采用的是http upgrade 机制
```
请求
GET /chat HTTP/1.1
Host: server.example.com
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
Origin: http://example.com
Sec-WebSocket-Protocol: chat, superchat
Sec-WebSocket-Version: 1


返回
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=
Sec-WebSocket-Protocol: chat
Sec-WebSocket-Version: 13

```

# TCP

```
TCP 给发送的每一个包进行编号，接收方对数据包进行排序，把有序数据传送给应用层。
校验和： TCP 将保持它首部和数据的检验和。这是一个端到端的检验和，目的是检测数据在传输过程中的任何变化。如果收到段的检验和有差错，TCP 将丢弃这个报文段和不确认收到此报文段。
滑动窗口实现流量控制
拥塞控制： 当网络拥塞时，减少数据的发送。
ARQ 协议： 也是为了实现可靠传输的，它的基本原理就是每发完一个分组就停止发送，等待对方确认。在收到确认后再发下一个分组。
超时重传： 当 TCP 发出一个段后，它启动一个定时器，等待目的端确认收到这个报文段。如果不能及时收到一个确认，将重发这个报文段。
```