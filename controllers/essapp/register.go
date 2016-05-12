package essapp

import (
	"github.com/astaxie/beego"
    "fmt"
)

// operations for Register
type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) URLMapping() {
	c.Mapping("register/post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Index
// @Description create Register
// @Param   body        body    models.Register true        "body for Register content"
// @Success 201 {object} models.Register
// @Failure 403 body is empty
// @router / [Index]
func (c *RegisterController) Index() {
    var id int
    c.Ctx.Input.Bind(&id, "id")
    c.Ctx.Output.Body([]byte(fmt.Sprintf("price id is:%d\n", id)))
    // c.TplName = "register.html"
}


// @Title Post
// @Description create Register
// @Param	body		body 	models.Register	true		"body for Register content"
// @Success 201 {object} models.Register
// @Failure 403 body is empty
// @router / [post]
func (c *RegisterController) Post() {
    // c.Ctx.Output.WriteString("I am Post")
}

// @Title Get
// @Description get Register by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Register
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RegisterController) GetOne() {
    // c.Ctx.Output.WriteString("I am GetOne")
}

// @Title Get All
// @Description get Register
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Register
// @Failure 403
// @router / [get]
func (c *RegisterController) GetAll() {
    // c.Ctx.Output.WriteString("I am GetAll")
}

// @Title Update
// @Description update the Register
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Register	true		"body for Register content"
// @Success 200 {object} models.Register
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RegisterController) Put() {
    // c.Ctx.Output.WriteString("I am Put")
}

// @Title Delete
// @Description delete the Register
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RegisterController) Delete() {
    // c.Ctx.Output.WriteString("I am Delete")
}
