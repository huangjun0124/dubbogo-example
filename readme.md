## 一、项目说明
此项目为一个starter示例，介绍使用`springboot dubbo`项目通过`dubbo`协议访问`dubbogo`服务。
## 二、项目结构
```tree
├─dubbo-client                          // 定义java使用的dubbo dto和接口信息，用于消费者引用
│  ├─src
│  │  └─main
│  │      ├─java
│  │      │  └─com
│  │      │      └─demo
│  │      │          └─exp
│  │      │              ├─dto          // dubbo接口的请求及返回参数定义
│  │      │              └─service      // dubbo 接口定义
│  │      └─resources
├─dubbo-consumer                        // springboot 子项目，java消费者示例客户端
│  ├─src
│  │  ├─main
│  │  │  ├─java
│  │  │  │  └─com
│  │  │  │      └─example
│  │  │  │          └─demo
│  │  │  │              ├─controllers   // restcontroller http接口，内部调用引用【dubbo-client】调用注册到zk上的`go`服务
│  │  │  │              └─dto           // request dto
│  │  │  └─resources                    // 资源文件
└─dubbogodemo                           // go 语言的 dubbo 服务，为服务提供者
    ├─cmd                               // dubbogo 服务启动器
    ├─conf                              // dubbogo 服务配置文件
    └─pkg
        ├─dto                           // dubbo接口的请求及返回参数定义
        ├─service                       // dubbo 接口定义
        └─util                          // 项目工具库
```
## 三、类库
+ [pkg/errors](https://github.com/pkg/errors)
+ [dubbo-go-hessian2](https://github.com/apache/dubbo-go-hessian2)
+ [dubbo-go v3](https://github.com/apache/dubbo-go)
## 四、参考文档
+ [dubbo-go-samples](https://github.com/apache/dubbo-go-samples)
+ [dubbogo 基础使用示例---俊果果-简书](https://www.jianshu.com/p/3698adc67e20)