## grpc

### 安装 
1. 安装protobuf
源码下载： https://github.com/google/protobuf  
编译: 
``` 
./configure
make & make install
```   

2. go 插件  
```
go get -u github.com/golang/protobuf/protoc-gen-go
```  

3. 依赖包 
由于墙的原因，没办法通过go get直接获取，好在github可以访问，通过clone github到本地，然后修改本地目录名：
+ golang.org/x/net  
+ golang.org/x/text  
+ golang.org/x/oauth2  
```
git clone github.com/grpc/grpc-go  
git clone github.com/golang/net 
``` 

### proto使用 
1. 编译 
```
protoc --go_out=plugins=grpc:. test.proto
```



### proto标量类型列表 
proto类型|C++类型|Go类型|备注
-|-|-|-
double|double|float64|浮点
float|float|float32|
int32|int32|int32|
int64|int64|int64|
uint32|uint32|uin32|
uint64|uint64|uint64|
sint32|sint32||
sint64|sint64||
fixed32|fixed32||4个字节，如果数值总是比2^28大的话，这个类型会比uint32高效
fixed64|fixed64||8个字节，如果数值总是比2^56大的话，这个类型会比uint64高效
sfixed32|sfixed32||
sfixed64|sfixed64||
bool|bool|bool|
string|string|string|字符串必须是UTF-8编码或者7-bit ASCII编码的文本
bytes|bytes||