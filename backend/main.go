package main

import (
		"github.com/gin-gonic/gin"
		"github.com/gin-contrib/cors"
		_ "github.com/go-sql-driver/mysql"
		"fmt"
		"github.com/joho/godotenv"
		"os"
		// "handmade_mask_shop/infrastructure/database"
		"handmade_mask_shop/routes"
		"github.com/gin-contrib/sessions"
		"github.com/gin-contrib/sessions/cookie"
)

func main() {
	fmt.Println()
	r := gin.Default()
	r = CORS(r)

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r = routes.GetAdminRoutes(r)
	r = routes.GetRoutes(r)

	// db := database.GormConnect()
	// database.Migrations(db)
	// database.Seeds(db)

		files := []string{ 
			"./templates/top/index.html", "./templates/front/item/detail.html",
			"./templates/admin/dashboards/index.html", "./templates/admin/items/index.html", "./templates/admin/items/detail.html", "./templates/admin/items/create.html", "./templates/admin/items/edit.html", "./templates/admin/items/category.html",
			"./templates/admin/users/regist.html", "./templates/admin/users/index.html",
			"./templates/admin/adminUsers/regist.html", "./templates/admin/adminUsers/edit.html",
      "./templates/layout/dafault.html", "./templates/layout/admin_default.html",
			"./templates/admin/elements/header.html", "./templates/admin/elements/footer.html", "./templates/admin/elements/sidebar.html",
			"./templates/admin/logins/index.html", "./templates/admin/logins/reset_password.html", "./templates/admin/logins/reset_password_complete.html",
		}

		r.LoadHTMLFiles(files...)
		
		r.Static("/src", "./src")
		r.Static("/public", "./public")

	r.Run(":80")
};

func CORS(r *gin.Engine) (*gin.Engine) {
	err := godotenv.Load((".env"))
	if err != nil {
		panic ("envファイルの読み込みに失敗しました。")
	}

	httpUrl :=  os.Getenv("HTTP_URL")
	r.Use( cors.New(cors.Config{
    // アクセスを許可したいアクセス元
    AllowOrigins:  []string{httpUrl},
    // アクセスを許可したいHTTPメソッド
    AllowMethods: []string{
        "POST","GET", "PUT", "OPTIONS",
    },
		AllowCredentials: true,
    // 許可したいHTTPリクエストヘッダ
    AllowHeaders: []string{
        "Access-Control-Allow-Credentials",
        "Access-Control-Allow-Headers",
				"Access-Control-Allow-Origin",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
    },
	}))
	return r
}

// func (c *booksController) Update(ctx *gin.Context) (interface{}, *errors.Error) {
//   req := &BooksApiUpdateRequest{}
//   if err := ctx.ShouldBind(req); err != nil {
//     // 独自エラーへ変換
//     return nil, errors.NewError(errors.InvalidRequest, err)
//   }

//   return c.BooksUseCase.Update(req.Id, req.Name)
// }
// func handler(fn func(c *gin.Context) (interface{}, *errors.Error)) gin.HandlerFunc {
//   return func(c *gin.Context) {
//     result, err := fn(c)
//     if err != nil {
//       c.JSON(toHttpStatusCode(err.Code), gin.H{"msg": err.Msg()})
//       c.Abort()
//     } else {
//       c.JSON(http.StatusOK, gin.H{"msg": "success", "result": result})
//     }
//   }
// }

// // 独自エラーコードに基づいてHttpレスポンスコードに変換する
// func toHttpStatusCode(code errors.ErrorCode) int {
//   switch code {
//   case errors.InvalidRequest:
//     return http.StatusBadRequest
//   case errors.NotFound:
//     return http.StatusNotFound
//   case errors.InternalError:
//     return http.StatusNotFound
//   }
// }