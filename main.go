package main
import (
	"strconv"
)
import "github.com/gofiber/fiber"
import "github.com/gofiber/fiber/middleware"
type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}
var todos=[]*Todo{{Id:1,Name:"name1",Completed:false},{Id:2,Name:"name2",Completed:false},}
func main(){
	app := fiber.New()
	app.Use(middleware.Logger())
	app.Get( "/" , func(ctx *fiber.Ctx){ ctx.Send("hello world")	})
	app.Get( "/todos" , GetTodos)
	app.Post( "/todos" , CreateTodo)
	app.Get( "/todos/:id" , GetTodo)
	app.Delete( "/todos/:id" , DeleteTodo)
	app.Patch( "/todos/:id" , UpdateTodo)
	err:=app.Listen(3000)
	if err != nil { panic(err) }
}
func UpdateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil { ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse id",})
		return
	}
	var body request
	err = ctx.BodyParser(&body)
	if err != nil { ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse body",})
		return
	}
	var todo *Todo
	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}
	if todo == nil {	ctx.Status(fiber.StatusNotFound) //no id can be =0 
		return	}
	if body.Name != nil {	todo.Name = *body.Name	}
	if body.Completed != nil {	todo.Completed = *body.Completed	}
	ctx.Status(fiber.StatusOK).JSON(todo)
}
func DeleteTodo(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{	"error": "cannot parse id",	})
		return
	}
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[0:i], todos[i+1:]...)
			ctx.Status(fiber.StatusNoContent)
			return
		}
	}
	ctx.Status(fiber.StatusNotFound)
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
	todo := &Todo{	Id:len(todos) + 1,	Name:body.Name,	Completed: false,	}
	todos = append(todos, todo)
	ctx.Status(fiber.StatusCreated).JSON(todo)
}
func GetTodo(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse id",})
		return
	}
	for _, todo := range todos {
		if todo.Id == id {	ctx.Status(fiber.StatusOK).JSON(todo)
			return
		}
	}
	ctx.Status(fiber.StatusNotFound)
}
