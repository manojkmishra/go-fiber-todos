package main
import "github.com/gofiber/fiber"
import "github.com/gofiber/fiber/middleware"
type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}
var todos=[]Todo{{Id:1,Name:"name1",Completed:false},{Id:2,Name:"name2",Completed:false},}
func main(){
	app := fiber.New()
	app.Use(middleware.Logger())
	app.Get( "/" , func(ctx *fiber.Ctx){ ctx.Send("hello world")	})
	app.Get( "/todos" , GetTodos)
	app.Post( "/todos" , CreateTodo)
	err:=app.Listen(3000)
	if err != nil { panic(err) }
}
func GetTodos(ctx *fiber.Ctx){	ctx.Status(fiber.StatusOK).JSON(todos) }
func CreateTodo(ctx *fiber.Ctx) {
	type request struct {Name string `json:"name"`}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {	
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json",	})
		return
	}
	todo := Todo{	Id:len(todos) + 1,	Name:body.Name,	Completed: false,	}
	todos = append(todos, todo)
	ctx.Status(fiber.StatusCreated).JSON(todo)
}