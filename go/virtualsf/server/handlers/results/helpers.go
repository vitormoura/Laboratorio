package results

import (
	"fmt"
	"text/template"
)

var (

	//helpersFuncs é um mapa de funções do tipo helper que podem ser utilizadas nos templates
	helpersFuncs = make(template.FuncMap)
)

func init() {
	helpersFuncs["fmtFileSize"] = fmtFileSize
}

//fmtFileSize retorna tamanho de arquivo formatado adequadamente em GB, MB e KB
func fmtFileSize(size int64) string {

	if size > 1073741824 {
		return fmt.Sprintf("%.2f GB", float32(size)/1073741824.0)
	} else if size > 1048576 {
		return fmt.Sprintf("%.2f MB", float32(size)/1048576.0)
	} else {
		return fmt.Sprintf("%.2f KB", float64(size)/1024.0)
	}
}
