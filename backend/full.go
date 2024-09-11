package full

// // import (
// // 	"github.com/KUNSALISA/Login_SA/database"
// // 	"github.com/KUNSALISA/Login_SA/handlers"
// // 	"github.com/gin-gonic/gin"
// // )

// // func main() {
// // 	r := gin.Default()

// // 	// Connect to database
// // 	database.ConnectDatabase()

// // 	// Routes
// // 	r.POST("/login", handlers.Login)
// // 	r.POST("/register", handlers.Register)

// // 	// Start the server
// // 	err := r.Run(":8080") // Run on port 8080
// // 	if err != nil {
// // 		panic("Failed to start server: " + err.Error())
// // 	}

// // }

// // package main

// // import (
// // 	"github.com/KUNSALISA/Login_SA/controller"
// // 	"github.com/KUNSALISA/Login_SA/entity"
// // 	"github.com/gin-gonic/gin"
// // )

// // const PORT = "8000"

// // func main() {



// // 	entity.SetupDatabase()
// // 	r := gin.Default()
// // 	r.Use(CORSMiddleware())

// // 	// การจัดการสมาชิก
// // 	r.GET("/members", controller.ListMember)
// // 	r.GET("/member/:ID", controller.GetMember)
// // 	r.PATCH("/members", controller.UpdateMember)
// // 	r.DELETE("/members/:id", controller.DeleteMember)
// // 	r.POST("/members", controller.CreateMember)

// // 	// การเข้าสู่ระบบ
// // 	r.POST("/login", controller.LoginByUsername)

// // 	r.Run("localhost:" + PORT)
// // 	// รันเซิร์ฟเวอร์

// // }

// // func CORSMiddleware() gin.HandlerFunc {
// // 	return func(c *gin.Context) {
// // 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// // 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// // 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// // 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
// // 		if c.Request.Method == "OPTIONS" {
// // 			c.AbortWithStatus(204)
// // 			return
// // 		}
// // 		c.Next()
// // 	}
// // }

// // // package main

// // // import (
// // // 	"github.com/KUNSALISA/Login_SA/entity"
// // // 	"gorm.io/driver/sqlite"
// // // 	"gorm.io/gorm"
// // // )

// // // func main() {
// // // 	db, err := gorm.Open(sqlite.Open("G11_PROJECT.db"), &gorm.Config{})
// // // 	if err != nil {
// // // 		panic("failed to connect database")
// // // 	}

// // // 	db.AutoMigrate(&entity.Member{})

// // // }


// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/KUNSALISA/Login_SA/config"
// 	"github.com/KUNSALISA/Login_SA/controller"
// )

// const PORT = "8000"

// func main() {

// 	// open connection database
// 	config.ConnectionDB()

// 	// Generate databases
// 	config.SetupDatabase()

// 	r := gin.Default()

// 	r.Use(CORSMiddleware())

// 	router := r.Group("")
// 	{

// 		// User Routes
// 		router.GET("/users", controller.ListUsers)
// 		router.GET("/user/:id", controller.GetUser)
// 		router.POST("/users", controller.CreateUser)
// 		router.PATCH("/users", controller.UpdateUser)
// 		router.DELETE("/users/:id", controller.DeleteUser)
// 		// // Gender Routes
// 		// router.GET("/genders", controller.ListGenders)
// 	}

// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
// 	})

// 	// Run the server

// 	r.Run("localhost:" + PORT)

// }

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }