package main

import (
	"aksan/golangcrudmvc/config"
	"aksan/golangcrudmvc/controllers"
	"aksan/golangcrudmvc/models"
	"aksan/golangcrudmvc/restcontrollers"
	"aksan/golangcrudmvc/service"
	"aksan/golangcrudmvc/serviceimpl"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// route
	http.HandleFunc("/", pertama)
	http.HandleFunc("/form", form)
	http.HandleFunc("/proses", proses)
	http.HandleFunc("/list", list)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/hapus", hapus)
	// rest
	http.HandleFunc("/rest", rest)
	http.HandleFunc("/restPost", restcontrollers.RestPost)
	http.HandleFunc("/addRest", restcontrollers.AddRest)
	http.HandleFunc("/formRest", formRest)
	http.HandleFunc("/userRestHapus", restcontrollers.HapusRest)

	// handle resources
	http.Handle("/resources/images/", http.StripPrefix("/resources/images", http.FileServer(http.Dir("./resources/images"))))
	http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./resources/styles"))))
	// port
	http.ListenAndServe(":8085", nil)
}

func pertama(w http.ResponseWriter, r *http.Request) {

	bukus, err := controllers.List()
	if err != nil {
		fmt.Println(err)
	}
	type custom struct {
		Data models.Bukus
	}
	var dataCustom custom
	dataCustom.Data = bukus
	tpl.ExecuteTemplate(w, "pertama.gohtml", dataCustom)
}

func form(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "form.gohtml", nil)
}

func proses(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/form", http.StatusSeeOther)
		return
	}
	// copy to resource
	alreadyPath := r.FormValue("alreadyPath")
	file, h, err := r.FormFile("path")
	if err != nil {
		alreadyPath = "/resources/images/NoData.png"
	} else {
		f, err := os.OpenFile("./resources/images/"+h.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
		io.Copy(f, file)
		alreadyPath = "/resources/images/" + h.Filename
	}
	//initial entity
	buku := models.SetBuku()
	buku.Judul = r.FormValue("judul")
	buku.Harga = r.FormValue("harga")
	buku.Path = alreadyPath

	// update or save
	if r.FormValue("id") != "" {
		idbuku, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			fmt.Println(err)
		}
		buku.Idbuku = idbuku
		err = controllers.Edit(idbuku, buku)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := controllers.Proses(buku)
		if err != nil {
			fmt.Println(err)
		}
	}
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func list(w http.ResponseWriter, r *http.Request) {

	bukus, err := controllers.List()
	if err != nil {
		fmt.Println(err)
	}
	type custom struct {
		Data models.Bukus
	}
	var dataCustom custom
	dataCustom.Data = bukus
	tpl.ExecuteTemplate(w, "list.gohtml", dataCustom)
}

func edit(w http.ResponseWriter, r *http.Request) {

	id, errs := r.URL.Query()["id"]
	if !errs {
		fmt.Println(errs)
	}
	idbuku, err := strconv.Atoi(id[0])
	if err != nil {
		fmt.Println(err)
	}
	data, err := controllers.FindByID(idbuku)
	if err != nil {
		fmt.Println(err)
	}
	tpl.ExecuteTemplate(w, "form.gohtml", data)
}

func hapus(w http.ResponseWriter, r *http.Request) {

	id, errs := r.URL.Query()["id"]
	if !errs {
		fmt.Println(errs)
	}
	idbuku, err := strconv.Atoi(id[0])
	if err != nil {
		fmt.Println(err)
	}
	err = controllers.Hapus(idbuku)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func rest(w http.ResponseWriter, r *http.Request) {
	type customRest struct {
		Data models.UserRests
	}
	db, err := config.GetPostgresDB()
	if err != nil {
		fmt.Println(err)
	}

	userRestService := service.NewRestService(db)
	userRests, err := serviceimpl.FindAllUserRest(userRestService)
	if err != nil {
		fmt.Println(err)
	}

	var dataCustomRest customRest
	dataCustomRest.Data = userRests
	tpl.ExecuteTemplate(w, "rest.gohtml", dataCustomRest)
}

func formRest(w http.ResponseWriter, r *http.Request) {
	id, _ := r.URL.Query()["id"]
	var data models.UserRest
	if id != nil {
		idUser, err := strconv.Atoi(id[0])
		if err != nil {
			fmt.Println(err)
		}

		userRest, err := restcontrollers.FindRestId(idUser)
		if err != nil {
			fmt.Println(err)
		}
		data = userRest
	}
	tpl.ExecuteTemplate(w, "formRest.gohtml", data)
}
