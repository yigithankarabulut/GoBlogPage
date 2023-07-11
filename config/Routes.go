package config

import (
	"github.com/julienschmidt/httprouter"
	admin "goblog/admin/controllers"
	site "goblog/site/controllers"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//Admin
	r.GET("/admin", admin.Dashboard{}.Index)

	//Blog Posts
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)

	//Upload
	r.POST("/admin/add", admin.Dashboard{}.Add)

	//Delete
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)

	//Edit
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)

	//Update
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//Login
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)

	//Admin Logout
	r.GET("/admin/logout", admin.Userops{}.LogOut)

	//Categories
	r.GET("/admin/kategoriler", admin.Categories{}.Index)
	r.POST("/admin/kategoriler/add", admin.Categories{}.Add)
	r.GET("/admin/kategoriler/delete/:id", admin.Categories{}.Delete)

	//Homepage
	r.GET("/", site.Homepage{}.Index)
	r.GET("/yazilar/:slug", site.Homepage{}.Detail)

	//Serve Files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/assets/*filepath", http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
