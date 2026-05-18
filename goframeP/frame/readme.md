
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


接口数据先更新表，再调用api，不再依赖于接口推送结果
推送失败要有一定的重试机制
轮询机制判断下游服务是否异常，隔一段时间check状态
最差的情况要有人工的介入
监控--告警

K8s内部健康检查(K8s自动调用，无需配置)
http://10.244.0.5:8080/health/

集群内其他服务(通过 Service 名称访问)
http://goframe-service:80/health/

外部用户(通过域名访问（需要DNS解析）)
http://api.example.com/health/


K8s会自动请求:http://<Pod-IP>:8080/health/
```dockerignore
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goframe-app
spec:
  replicas: 3
  template:
    spec:
      containers:
        - name: goframe-app
          image: your-app:latest
          ports:
            - containerPort: 8080  # ← K8s 知道容器监听 8080 端口
          
          livenessProbe:
            httpGet:
              path: /health/       # ← K8s 知道检查这个路径
              port: 8080           # ← K8s 知道访问这个端口
            periodSeconds: 5

```
外部访问你的服务，运维还需要配置 Service 和 Ingress。
service.yaml
```dockerignore
apiVersion: v1
kind: Service
metadata:
  name: goframe-service
spec:
  selector:
    app: goframe-app
  ports:
    - protocol: TCP
      port: 80           # 外部访问端口
      targetPort: 8080   # 转发到容器的 8080 端口
  type: ClusterIP        # 集群内部访问

```
ingress.yaml
```dockerignore
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: goframe-ingress
spec:
  rules:
    - host: api.example.com  # ← 域名
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: goframe-service
                port:
                  number: 80

```


```dockerignore
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
```

接口幂等性处理
项目中实现幂等性的接口完整列表
根据全面代码分析，以下是所有实现了幂等性的接口
1️⃣ 发票创建接口 - InvoiceCreate
```dockerignore
- 文件: internal/controller/tencent/tencent_v1_invoice_create.go 
- 上游渠道: 腾讯、阿里商旅等 幂等 Key: invoice_create: + MD5(排序后的订单号列表) 
- 实现方式:
✅ Redis 分布式锁（5秒）
✅ Redis 缓存结果（30分钟）
lockKey := consts.InvoiceCreateLockKey + orderNosStrMd5
redisKey := consts.InvoiceCreateKey + orderNosStrMd5

```
2️⃣ 取消订单接口 - CancelOrder (阿里商旅)
```dockerignore
文件: internal/controller/alitrip/alitrip_v1_cancel_order.go 
上游渠道: 阿里商旅 (alitrip) 
幂等 Key: alitrip:order:cancel:{OrderId} 实现方式:
✅ Redis 分布式锁（3秒）
✅ Redis 缓存结果（5分钟）

lockKey := const_alitrip.RedisCancelOrderLock + data.OrderId
redisKey := const_alitrip.RedisCancelOrder + data.OrderId

```
3️⃣ 取消订单接口 - CancelOrder (美团)
```dockerignore
文件: internal/controller/meituan/meituan_v1_cancel_order.go 
上游渠道: 美团 (meituan) 幂等 Key: meituan:order:cancel:{PartnerOrderId}
 实现方式:
✅ Redis 分布式锁（3秒）
✅ Redis 缓存结果（5分钟）
lockKey := const_meituan.RedisCancelOrderLock + req.PartnerOrderId
redisKey := const_meituan.RedisCancelOrder + req.PartnerOrderId

```

4️⃣ 取消订单接口 - CancelOrder (腾讯)
```dockerignore
文件: internal/controller/tencent/tencent_v1_cancel_order.go 
上游渠道: 腾讯 (tencent) 幂等 Key: tencent:order:cancel:{SpOrderId} 
实现方式:

✅ Redis 分布式锁（3秒）
✅ Redis 缓存结果（5分钟）
lockKey := const_tencent.RedisCancelOrderLock + req.Data.SpOrderId
redisKey := const_tencent.RedisCancelOrder + req.Data.SpOrderId

```
5️⃣ 取消订单接口 - CancelOrder (OpenAPI V2)
```dockerignore
文件: internal/controller/openapi/openapi_v2_cancel_order.go 
上游渠道: OpenAPI (支持 yiqi、alitrip 等多渠道) 幂等 Key: tencent:order:cancel:{SubOrderNumber} 
实现方式:
✅ Redis 分布式锁（3秒）
✅ Redis 缓存结果（10分钟）
lockKey := const_openapi.RedisCancelOrderLock + req.SubOrderNumber
redisKey := const_openapi.RedisCancelOrder + req.SubOrderNumber

```

6️⃣ 支付通知接口 - PayOrder
文件: internal/application/pay_order/pay_order.go 功能: 处理上游支付通知 幂等 Key: pay_notify_lock:{SwitchOrderNumber} 实现方式:
✅ Redis SetNX 锁（10秒）
✅ 检查订单状态（已支付直接返回）
```dockerignore
redisKey := consts.RedisPayNotifyLockKey + req.SwitchOrderInfo.SwitchOrderNumber
if payNotifyLock, err := redis.SetNX(ctx, redisKey, "1", 10); !payNotifyLock {
    return result, gerror.New("重复的支付通知")
}

```

7️⃣ 举手/PK 接口 - HandUp
文件: internal/application/hand_up/hand_up.go 功能: 司机举手抢单 幂等 Key:
PK 等待锁：hand_up_pk_wait:{orderNo}
举手锁：使用 SetNX 实现方式:
✅ Redis Lock（PK 等待时间+2秒）
✅ Redis SetNX（30秒）
```dockerignore
// PK 等待锁
if isLock, err := redis.Lock(ctx, lockKey, "1", time.Duration(pkWaitTime/1000+2)); err != nil { ... }

// 举手锁
if handUpLock, err := redis.SetNX(ctx, handUpLockKey, "1", 30); err != nil { ... }

```

退款接口 - DownRefund
文件: internal/service/order/down_refund.go 功能: 下游退款操作 幂等 Key: refund_lock:{orderNo} 实现方式:
✅ Redis 分布式锁（4秒）
```dockerignore
if locked, err := redis.Lock(s.ctx, lockKey, "1", 4); err != nil { ... }

```

9️⃣ 查询发票详情定时任务 - QueryInvoiceDetail
文件: internal/crontab/query_invoice_detail.go 功能: 定时查询发票详情 幂等 Key: query_invoice_detail_lock:{invoiceId} 实现方式:
✅ Redis 分布式锁
```dockerignore
locked, err := redis.Lock(c.ctx, lockKey, lockValue, expire)

```

🔟 查询订单详情定时任务 - QueryOrderDetail

文件: internal/crontab/query_order_detail.go 功能: 定时同步订单详情 幂等 Key: schedule_query_order_lock 实现方式:
✅ Redis 分布式锁
```dockerignore
locked, err := redis.Lock(c.ctx, lockKey, lockToken, lockExpiration)

```

1️⃣1️⃣ 上游取消订单服务 - UpCancel
文件: internal/service/order/up_cancel.go 功能: 处理上游取消请求 幂等 Key: up_cancel_lock:{orderNo} 实现方式:
✅ Redis 分布式锁（2秒）
1️⃣1️⃣ 上游取消订单服务 - UpCancel
文件: internal/service/order/up_cancel.go 功能: 处理上游取消请求 幂等 Key: up_cancel_lock:{orderNo} 实现方式:
✅ Redis 分布式锁（2秒）
```dockerignore
if locked, err := redis.Lock(s.ctx, lockKey, "1", 2); err != nil { ... }

```

1️⃣2️⃣ 切换订单状态服务 - SwitchOrderStatus
文件: internal/service/order/switch_order_status.go 功能: 订单状态流转 幂等 Key:
订单号锁：order_number_lock:{orderNumber}（360秒）
状态切换锁：switch_order_status_lock:{orderNo}（5秒） 实现方式:
✅ Redis SetNX（订单号唯一性校验）
✅ Redis Lock（状态切换防并发）
```dockerignore
// 订单号唯一性
if orderNumberExists, _ := redis.SetNX(s.ctx, lockKey, orderNumber, 360); !orderNumberExists { ... }

// 状态切换锁
if locked, err = redis.Lock(s.ctx, lockKey, "1", 5); err != nil { ... }

```

1️⃣3️⃣ 获取司机信息 - GetDriverInfo
文件: internal/service/order/driver.go 功能: 获取司机位置信息 幂等 Key: driver_info_lock:{orderNo} 实现方式:
✅ Redis 分布式锁（120秒）
````dockerignore
if r, e := redis.Lock(s.ctx, lk, "1", 120); e != nil { ... }

````

1️⃣4️⃣ 举手缓存 - HandUp Cache
文件: internal/service/order/hand_up.go 功能: 缓存举手司机信息 幂等 Key: hand_up_cache:{orderNo} 实现方式:
✅ Redis SetNX（900秒）
```dockerignore
cacheRes, cacheErr := redis.SetNX(s.ctx, redisKey, driverInfo, 900)

```

1️⃣5️⃣ 发单缓存 - Create Order Cache
文件: internal/service/cache.go 功能: 防止重复发单 幂等 Key: create_order_lock:{orderNo} 实现方式:
✅ Redis SetNX（600秒）
```
if lock, err := redis.SetNX(ctx, redisKey, toUpOrderNumber, 600); err != nil { ... }
```

取消订单 4个
alitrip、meituan、tencent、openapi
支付相关 1个
支付通知
订单操作 4个
举手、退款、状态切换、发单
定时任务 2个
发票查询、订单同步
其他 2个
司机信息、举手缓存