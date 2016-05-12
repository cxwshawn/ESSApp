package essapp

import (
	"github.com/astaxie/beego"
)

// operations for Account
type AccountController struct {
	beego.Controller
}

func (c *AccountController) URLMapping() {
	// c.Mapping("Post", c.Post)
	// c.Mapping("GetOne", c.GetOne)
	// c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Account
// @Param	body		body 	models.Account	true		"body for Account content"
// @Success 201 {object} models.Account
// @Failure 403 body is empty
// @router / [post]
func (c *AccountController) Login() {
    if c.userid > 0 {
        c.Redirect("/admin", 302)
    }
    if c.GetString("dosubmit") == "yes" {
        account := strings.TrimSpace(c.GetString("account"))
        password := strings.TrimSpace(c.GetString("password"))
        remember := c.GetString("remember")
        if account != "" && password != "" {
            var user models.User
            user.UserName = account
            if user.Read("user_name") != nil || user.Password != util.Md5([]byte(password)) {
                c.Data["errmsg"] = "帐号或密码错误"
            } else if user.Active == 0 {
                c.Data["errmsg"] = "该帐号未激活"
            } else {
                user.LoginCount += 1
                user.LastIp = c.getClientIp()
                user.LastLogin = c.getTime()
                user.Update()
                authkey := util.Md5([]byte(c.getClientIp() + "|" + user.Password))
                if remember == "yes" {
                    c.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
                } else {
                    c.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
                }
                //TODO: chanage c to user's home page.
                c.Redirect("/admin", 302)
            }
        }
    }
    c.TplName = c.moduleName + "/account/login.html"
}

// @Title Get
// @Description get Account by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Account
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AccountController) GetOne() {

}

// @Title Get All
// @Description get Account
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Account
// @Failure 403
// @router / [get]
func (c *AccountController) GetAll() {

}

// @Title Update
// @Description update the Account
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Account	true		"body for Account content"
// @Success 200 {object} models.Account
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AccountController) Put() {

}

// @Title Delete
// @Description delete the Account
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AccountController) Delete() {

}
