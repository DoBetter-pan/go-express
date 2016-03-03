/**
* @file tabController.go
* @brief tab controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	_ "log"
	_ "fmt"
	_ "io/ioutil"
	"html/template"
)

type TabController struct {
}

func NewTabController() *TabController {
	return &TabController{}
}

func LoadTabIndexFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/tab/index.html"))
}

func (controller *TabController) IndexAction(w http.ResponseWriter, r *http.Request) {
	startup := `
	<script type="text/javascript">
    $(function() {
        $('#categoryTab a').click(function(e){
            e.preventDefault();
            $(this).tab('show');
        });

    });	
	</script>`

	mainConterller := NewMainController()
	//add new javascript and css
	//mainConterller.Stylesheets = append(mainConterller.Stylesheets, []string{"../extensions/bootstrap-treeview-master/dist/bootstrap-treeview.min.css"}...)
	//mainConterller.Javscripts = append(mainConterller.Javscripts, []string{"../extensions/bootstrap-treeview-master/dist/bootstrap-treeview.min.js"}...)
	mainConterller.Startup = template.HTML(startup)
	mainConterller.Content = LoadTabIndexFromTemplate()
	mainConterller.RenderMainFrame(w, r)
}
