package router

import (
	"fmt"

	"github.com/andrewarrow/feedback/util"
)

func (r *Router) ResetDatabase() {
	for _, model := range r.Site.Models {
		tableName := util.Plural(model.Name)
		r.Db.Exec("drop table " + tableName)
	}
	r.Db.Exec("drop table feedback_schema")
	fmt.Println("done.")
}
