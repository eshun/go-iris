package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"demo/services"
	"demo/controllers"
	"demo/controllers/api"
	"strconv"
	"github.com/kataras/iris/middleware/basicauth"
	"time"
	"fmt"
)
const Prefix  = "admin"

// Route 路由
func InitRoute(app *iris.Application) {
	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(c iris.Context) {
		user, err := services.GetUserById(1)
		fmt.Println(user, err)
		if err != nil {
			c.JSON(iris.Map{"status": false, "message": err})
		} else {
			c.JSON(iris.Map{"status": true, "message": ""})
		}
	})

	app.OnAnyErrorCode(func(c iris.Context) {
		statusCode := c.GetStatusCode()

		c.JSON(iris.Map{"status": false, "message": strconv.Itoa(statusCode)})
	})
	authConfig := basicauth.Config{
		Users:   map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	router := app.Party(Prefix, authentication)
	{
		router.Get("/categories", nil)
	}

	mvc.New(app.Party(Prefix + "/user")).Handle(new(controllers.UserController))

	mvc.New(app.Party("api/index")).Handle(new(api.IndexController))
	mvc.New(app.Party("api/address")).Handle(new(api.AddressController))
	mvc.New(app.Party("api/auth")).Handle(new(api.AuthController))
	mvc.New(app.Party("api/brand")).Handle(new(api.BrandController))
	mvc.New(app.Party("api/category")).Handle(new(api.CategoryController))
	mvc.New(app.Party("api/collect")).Handle(new(api.CollectController))
	mvc.New(app.Party("api/comment")).Handle(new(api.CommentController))
	mvc.New(app.Party("api/goods")).Handle(new(api.GoodsController))
	mvc.New(app.Party("api/keyword")).Handle(new(api.KeywordController))
	mvc.New(app.Party("api/pay")).Handle(new(api.PayController))
	mvc.New(app.Party("api/region")).Handle(new(api.RegionController))
	mvc.New(app.Party("api/topic")).Handle(new(api.TopicController))
	mvc.New(app.Party("api/user")).Handle(new(api.UserController))
}

func Authentication(ctx iris.Context) {
	if true {
		ctx.Next()
	}
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	// third parameter it will be always true because the middleware
	// makes sure for that, otherwise this handler will not be executed.

	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}