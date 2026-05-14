
1、配置config文件,数据库建好表
2、执行，生成dao的数据库模型
```
gf gen dao
```

3、新建api文件，添加访问路由
4、生成controller层代码

![执行命令gf gen ctrl之后生成的代码模版](./images/img.png)

-   yaml配置文件的相对路径
```
    srcFolder: "./table/api"
    dstFolder: "./table/internal/controller"
```
运行gf gen ctrl的前提是已经写好了api层的请求响应路由

![api层代码](./images/img_3.png)

-   运行命令gf gen ctrl生成controller层代码
```
    gf gen ctrl
```

如果想生成带有api/hello/v1/user的带有user前缀的代码，需要请求前面有完整的user前缀
![第一步](./images/img_1.png)
![第二步](./images/img_2.png)



package dao


//// UserKey 登录用户信息的context key
//type UserKey struct{}
//
//// TenantKey 租户（渠道）信息的context key
//type TenantKey struct{}
//
//// WithUser 将登录用户信息存入context
//func WithUser(ctx context.Context, user *entity.User) context.Context {
//	return context.WithValue(ctx, UserKey{}, user)
//}
//
//// GetUser 从context获取登录用户信息
//func GetUser(ctx context.Context) (*entity.User, bool) {
//	user, ok := ctx.Value(UserKey{}).(*entity.User)
//	return user, ok
//}
//
//// WithTenant 将租户（渠道）信息存入context
//func WithTenant(ctx context.Context, tenantId int64) context.Context {
//	return context.WithValue(ctx, TenantKey{}, tenantId)
//}
//
//// GetTenant 从context获取租户（渠道）信息
//func GetTenant(ctx context.Context) (int64, bool) {
//	tenantId, ok := ctx.Value(TenantKey{}).(int64)
//	return tenantId, ok
//}