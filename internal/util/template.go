package util

import "github.com/gin-contrib/multitemplate"

func CreateRenderer(templatePath string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("car.index", templatePath+"/layout.html", templatePath+"/car/index.html")
	r.AddFromFiles("error", templatePath+"/layout.html", templatePath+"/error.html")

	return r
}
