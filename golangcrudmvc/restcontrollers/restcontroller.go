package restcontrollers

import (
	"aksan/golangcrudmvc/config"
	"aksan/golangcrudmvc/models"
	"aksan/golangcrudmvc/service"
	"aksan/golangcrudmvc/serviceimpl"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var db, err = config.GetPostgresDB()

func RestPost(w http.ResponseWriter, r *http.Request) {

	var custom models.DataRest
	userRestService := service.NewRestService(db)
	userRest, err := serviceimpl.FindByUsernameRestImpl(r.FormValue("postUsername"), userRestService)
	if err != nil {
		fmt.Println(err)
	}

	//error handling
	err = config.CheckHash(r.FormValue("postKey"), userRest.Hash)
	if err != nil {

		// create response json
		custom.Status = "Error Authentication"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(custom)
		return
	}
	// find data
	bukuService := service.NewBukuService(db)
	bukus, err := serviceimpl.FindAllBuku(bukuService)

	if err != nil {
		fmt.Println(err)
	}
	// assign data
	var arrData models.DataDetail
	for _, v := range bukus {
		arrData.Judul = v.Judul
		arrData.Harga = v.Harga
		custom.Data = append(custom.Data, arrData)
	}
	custom.Status = "Success"

	// create response json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(custom)
	return
}

func AddRest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/rest", http.StatusSeeOther)
		return
	}
	// initial
	userRest := models.SetUserRest()
	userRest.Username = r.FormValue("addUsername")
	userRest.Key = r.FormValue("addKey")
	hashPass, err := config.HashPass(r.FormValue("addKey"))
	if err != nil {
		fmt.Println(err)
	}
	userRest.Hash = string(hashPass)

	userRestService := service.NewRestService(db)
	err = serviceimpl.SaveRest(userRest, userRestService)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/rest", http.StatusSeeOther)
	return
}

func FindRestId(id int) (models.UserRest, error) {
	userRestService := service.NewRestService(db)
	userRest, err := serviceimpl.FindByIDUserRest(id, userRestService)
	if err != nil {
		return userRest, err
	}
	return userRest, nil
}

func HapusRest(w http.ResponseWriter, r *http.Request) {
	id, _ := r.URL.Query()["id"]
	idUser, err := strconv.Atoi(id[0])
	if err != nil {
		fmt.Println(err)
	}
	userRestService := service.NewRestService(db)
	err = serviceimpl.DeleteUserRest(idUser, userRestService)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/rest", http.StatusSeeOther)
	return
}
