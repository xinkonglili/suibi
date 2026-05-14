package middleware

// AuthMiddleware 认证中间件
// 从 Session/Token 中获取当前登录用户及其所属渠道，注入到 Context
//func AuthMiddleware(r *ghttp.Request) {
//	ctx := r.Context()
//
//	// 方式1: 从 Session 获取用户信息（推荐）
//	session := r.Session
//	if session == nil {
//		r.Middleware.Next()
//		return
//	}
//
//	// 获取登录用户ID
//	userId := session.GetInt64("user_id")
//	tenantId := session.GetInt64("tenant_id") // 渠道/租户ID
//
//	// 如果Session中没有，尝试从 Token/JWT 中获取
//	if userId == 0 {
//		token := r.GetHeader("Authorization")
//		if token != "" {
//			// TODO: 解析JWT token获取用户和租户信息
//			// claims, err := jwt.ParseToken(token)
//			// if err == nil {
//			//     userId = claims.UserId
//			//     tenantId = claims.TenantId
//			// }
//		}
//	}
//
//	// 如果获取到用户信息，注入到Context
//	if userId > 0 {
//		// 可以查询完整的用户信息
//		// user, _ := dao.User.Ctx(ctx).WherePri(userId).One()
//		// ctx = dao.WithUser(ctx, user)
//
//		// 至少注入租户ID用于数据隔离
//		if tenantId > 0 {
//			ctx = dao.WithTenant(ctx, tenantId)
//			g.Log().Info(ctx, "用户认证成功", g.Map{
//				"user_id":  userId,
//				"tenant_id": tenantId,
//			})
//		}
//	}
//
//	r.SetContext(ctx)
//	r.Middleware.Next()
//}
