## 一、查看指定端口的进程
sudo lsof -i :27017

COMMAND   PID    USER        FD      TYPE             DEVICE             SIZE/OFF      NODE       NAME
[mongod]  859   zhangsan    313u      IPv6            0x1111111111111     0t0

然后根据PID杀进程：
sudo kill -9 859

## 二、protoc修改
```objc
protoc --go_out=. hello.proto
protoc --go-grpc_out=. hello.proto
```
## 参考文档
https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-04-grpc.html