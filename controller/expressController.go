/**
* @file expressController.go
* @brief express controller
* @author yingx
* @date 2015-12-12
 */

package controller

import (
	"net/http"
	_ "log"
	_ "fmt"
	"html/template"
)

type ExpressController struct {
}

func NewExpressController() *ExpressController {
	return &ExpressController{}
}

func LoadExpressIndexFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/express/index.html"))
}

func (controller *ExpressController) IndexAction(w http.ResponseWriter, r *http.Request) {
	startup := `
	<script>
  	$(document).ready(function(){
  		prettyPrint();
  	});
	</script>`
	mainConterller := NewMainController()
	//add new javascript and css
	mainConterller.Stylesheets = append(mainConterller.Stylesheets, []string{"../extensions/google-code-prettify/prettify.css"}...)
	mainConterller.Javscripts = append(mainConterller.Javscripts, []string{"../extensions/google-code-prettify/prettify.js"}...)
	mainConterller.Startup = template.HTML(startup)
	mainConterller.Content = LoadExpressIndexFromTemplate()
	mainConterller.RenderMainFrame(w, r)
}

func LoadExpressAboutFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/express/about.html"))
}

func (controller *ExpressController) AboutAction(w http.ResponseWriter, r *http.Request) {
	mainConterller := NewMainController()
	mainConterller.Content = LoadExpressAboutFromTemplate()
	mainConterller.RenderMainFrame(w, r)
}
