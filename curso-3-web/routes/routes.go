package routes

import (
	"net/http"
	"alura-Go/sites/controllers"
)

func CarregaRotas()  {
	http.HandleFunc("/", controllers.Index)	
	http.HandleFunc("/new",  controllers.New)
	http.HandleFunc("/insert",  controllers.Insert)
	http.HandleFunc("/delete",  controllers.Delete)
	http.HandleFunc("/edit",  controllers.Edit)
	http.HandleFunc("/update",  controllers.Update)
}