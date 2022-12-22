package routes

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/controllers"
	"go_code/gintest/app/controllers/post"
	"go_code/gintest/app/controllers/user"
	"go_code/gintest/app/middlewares"
	"go_code/gintest/app/validators"
	"go_code/gintest/bootstrap"
	"go_code/gintest/bootstrap/glog"
	"time"
)

func RegisterRouters() *gin.Engine {
	// 翻译参数校验提示
	validators.InitTrans("zh")
	// 设置环境
	if bootstrap.Config.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(glog.GinLogger(), glog.GinRecovery(true))

	v1 := r.Group("api/v1")
	{
		v1.POST("register", user.SignUp)
		v1.POST("login", user.SignIn)
		v1.POST("refresh-token", user.RefreshToken)
		// 登录认证、
		v1.Use(middlewares.JWTAuth())
		{
			v1.GET("communities", controllers.Communities)
			v1.GET("communities/detail", controllers.CommunityDetail)
			v1.GET("community/posts", post.OrderedPostsForCommunity)
			v1.POST("posts", post.StorePost)
			v1.GET("posts/detail", post.ShowPost)
			v1.GET("posts", middlewares.RateLimitMiddleware(2 * time.Second, 1), post.OrderedPosts)
			v1.POST("vote", post.VotePost)
		}
	}

	pprof.Register(r)
	return r
}
