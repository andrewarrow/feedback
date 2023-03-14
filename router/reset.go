package router

import (
	"fmt"
	"os"
)

func (r *Router) ResetDatabase() {
	for _, model := range r.Site.Models {
		r.Db.Exec("drop table " + model.TableName())
	}
	prefix := os.Getenv("FEEDBACK_NAME")
	r.Db.Exec(fmt.Sprintf("drop table %s_feedback_schema", prefix))
	fmt.Println("done.")
}
