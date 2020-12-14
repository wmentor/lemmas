package generic

import (
	"time"

	"github.com/wmentor/lemmas/engine/forms"
	"github.com/wmentor/serv"
)

var (
	version int64 = time.Now().Unix()
)

func DefaultVars(c *serv.Context) map[string]interface{} {
	vars := make(map[string]interface{})

	vars["title"] = "Lemmas"
	vars["version"] = version
	vars["totalForms"] = forms.TotalForms()
	vars["totalFixed"] = forms.TotalFixed()

	return vars
}
