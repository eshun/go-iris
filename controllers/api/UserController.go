package api

import (
	"demo/controllers"
	"os"
	"io"
)

type UserController struct{
	controllers.ApiController
}

func (c *UserController) Get() {

}

func (c *UserController) PostSaveAvatar() {
	file, info, err := c.Ctx.FormFile("uploadfile")

	if err != nil {
		c.Fail(err)
		return
	}

	defer file.Close()
	fileName := info.Filename

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile("./uploads/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		c.Fail(err)
		return
	}
	defer out.Close()

	io.Copy(out, file)
	c.Success("")
}
