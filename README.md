# gostars
learn and share 'go'

### 1.1 目录结构

```
    ├── server
        ├── api             (api层)
        │   └── v1          (v1版本接口)
        ├── config          (配置包--config.ini)
        ├── core            (核心文件)
        ├── docs            (swagger文档目录)
        ├── global          (全局对象)                    
        ├── initialize      (初始化)                        
        │   └── base        (初始化函数)                            
        ├── middlewares     (中间件层)                        
        ├── models          (模型层)                    
        │   ├── request     (入参结构体)                        
        │   └── response    (出参结构体)                            
        ├── packfile        (静态文件打包)                        
        ├── resource        (静态资源文件夹)                                                   
        ├── routers          (路由层)                    
        ├── service         (service层)                    
        ├── source          (source层)                    
        └── utils           (工具包及参数)                    
            ├── timer       (定时器接口封装)                        
            └── upload      (oss接口封装)                        
            └── jwt         (jwt令牌设置)
            └── validator   (验证模块)
            └── code        (错误码)
            └── ipsource    (用户ip属地获取)
```