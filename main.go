package main
import "github.com/gofiber/fiber"
import "github.com/gofiber/fiber/middleware"
type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}
var todos=[]Todo{
	{Id:1,Name:"first one",Completed:false},
	{Id:2,Name:"first one",Completed:false},
}
func main(){
	app := fiber.New()
	app.Use(middleware.Logger())
	app.Get( "/" , func(ctx *fiber.Ctx){ ctx.Send("hello world")	})
	err:=app.Listen(3000)
	if err != nil { panic(err) }
}