package user

import (
    "net/http"
    "html/template"

    "go-websiteskeleton/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
        Active string
    }

    p := Page{
        Active: "user_home",
        Title: "User Page",
    }

    common.Templates = template.Must(template.ParseFiles("templates/user/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}