## reflect 反射  

1. 方法的反射：常用于自定义结构的方法    
Type:  
+ NumMethod(): 返回方法集的数目  
+ Method(i): 返回第i个方法  
+ NumIn(): 传入参数个数 
+ In(i): 返回第i个传入参数类型 `Type`      
+ NumOut(): 输出参数个数  
+ Out(i): 返回第i个输出参数类型 `Type`  

2. channel的反射：   
+ Elem(): 获取反射类型 
+ TrySend(v): 向channel发送数据  
+ TryRecv(): 接受数据 
+ Recv(): 接受channel返回的值  
+ Send(v): 发送channel数据 
+ ChanDir: 类型： `RecvDir`, `SendDir`, `BothDir`  

