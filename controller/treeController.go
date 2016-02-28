/**
* @file treeController.go
* @brief tree controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	_ "log"
	"fmt"
	"io/ioutil"
	"html/template"
)

/*
type TreeNode struct {
	Text           string       `json:"text"`
	Icon           string       `json:"icon"`
	//SelectedIcon   string       `json:"selectedIcon"`
	Color          string       `json:"color"`
	BackColor      string       `json:"backColor"`
	Href           string       `json:"href"`
	Tags           []string     `json:"tags"`
	Nodes          []*TreeNode  `json:"nodes"`
}

type Tree struct {
	Tree           []*TreeNode `json:"tree"`
}
*/

func LoadTree(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

type TreeController struct {
}

func NewTreeController() *TreeController {
	return &TreeController{}
}

func LoadTreeIndexFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/tree/index.html"))
}

func (controller *TreeController) IndexAction(w http.ResponseWriter, r *http.Request) {
	data := LoadTree("models/tree/treenode.json")
	startup_format := `
	<script type="text/javascript">
    $(function() {

    var data = %s;

    $('#treeview').treeview({
    	levels: 1,
		color: "#428bca",
		expandIcon: "glyphicon glyphicon-plus-sign",
		collapseIcon: "glyphicon glyphicon-minus-sign",
		//nodeIcon: "glyphicon glyphicon-user",
		showTags: false,
		showBorder: false,
		enableLinks: true,
		data: data
    });

	$('#treeview').on('nodeSelected', function(event, data) {
		console.log(event);
		console.log(data);
	});

    });	
	</script>`

	startup := fmt.Sprintf(startup_format, data)

	mainConterller := NewMainController()
	//add new javascript and css
	mainConterller.Stylesheets = append(mainConterller.Stylesheets, []string{"../extensions/bootstrap-treeview-master/dist/bootstrap-treeview.min.css"}...)
	mainConterller.Javscripts = append(mainConterller.Javscripts, []string{"../extensions/bootstrap-treeview-master/dist/bootstrap-treeview.min.js"}...)
	mainConterller.Startup = template.HTML(startup)
	mainConterller.Content = LoadTreeIndexFromTemplate()
	mainConterller.RenderMainFrame(w, r)
}

func LoadTreeAboutFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/tree/about.html"))
}

func (controller *TreeController) AboutAction(w http.ResponseWriter, r *http.Request) {
	mainConterller := NewMainController()
	mainConterller.Content = LoadTreeAboutFromTemplate()
	mainConterller.RenderMainFrame(w, r)
}
