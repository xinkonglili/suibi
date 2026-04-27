
1、配置config文件,数据库建好表
2、执行，生成dao的数据库模型
```
gf gen dao
```

3、新建api文件，添加访问路由
4、生成controller层代码

![执行命令gf gen ctrl之后生成的代码模版](./image/image.yaml)

-   yaml配置文件的相对路径
```
    srcFolder: "./table/api"
    dstFolder: "./table/internal/controller"
```
-   运行命令gf gen ctrl生成controller层代码
```
    gf gen ctrl
```

如果想生成带有api/hello/v1/user的带有user前缀的代码，需要请求前面有完整的user前缀
![第一步](./image/img_1.png)
![第二步](./image/img_2.png)