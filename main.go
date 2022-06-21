package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	// app := gin.Default()
	// app.Use(ginsession.New())
	// app.GET("/login/:name", func(ctx *gin.Context) {
	// 	name := ctx.Param("name")
	// 	store := ginsession.FromContext(ctx)
	// 	store.Set("name", name)
	// 	err := store.Save()
	// 	if err != nil {
	// 		ctx.AbortWithError(500, err)
	// 		return
	// 	}

	// 	ctx.String(http.StatusOK, "session : %s", store.SessionID())
	// })
	// app.GET("/autologin", func(ctx *gin.Context) {
	// 	store := ginsession.FromContext(ctx)
	// 	name, ok := store.Get("name")
	// 	key := store.SessionID()
	// 	if !ok {
	// 		ctx.AbortWithStatus(404)
	// 		return
	// 	}
	// 	ctx.String(http.StatusOK, "name:%s key:%s", name, key)
	// })
	// app.Run(":9000")

	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("gamepub", store))

	r.GET("/login/:name", func(c *gin.Context) {
		name := c.Param("name")
		session := sessions.Default(c)
		session.Set("name", name)
		session.Save()
		c.JSON(200, gin.H{"name": name})
	})
	r.GET("/autologin", func(c *gin.Context) {
		session := sessions.Default(c)
		name := session.Get("name")
		c.JSON(200, gin.H{"name": name, "session": session.ID()})
	})
	r.Run(":9100")
}
