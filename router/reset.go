package router

import (
	"fmt"

	"github.com/andrewarrow/feedback/sqlgen"
)

func (r *Router) ResetDatabase() {
	for _, model := range r.Site.Models {
		r.Db.Exec("drop table " + model.TableName())
	}
	r.Db.Exec(fmt.Sprintf("drop table %s", sqlgen.FeedbackSchemaTable()))
	fmt.Println("done.")
}
