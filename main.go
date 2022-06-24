//documentacao do componente do mysql: https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-important-settings  pego em : https://pkg.go.dev/?utm_source=godoc
package main

import (		
	"net/http"
	"loja/routes"
)

func main() {	
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}



