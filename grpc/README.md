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
|proto类型|C++类型|Go类型|备注|
|double|double|float64||
|float|float|float32||
|int32|int32|int32||