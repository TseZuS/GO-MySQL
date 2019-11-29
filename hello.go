package main

import (
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type sbase struct {
    Id    int
    Name  string
    Autor string
    Date string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "Oh140803"
    dbName := "sbase"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM book ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
    emp := sbase{}
    res := []sbase{}
    for selDB.Next() {
        var id int
        var name, autor, date string
        err = selDB.Scan(&id, &name, &autor, &date)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Autor = autor
        emp.Date = date
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM book WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := sbase{}
    for selDB.Next() {
        var id int
        var name, autor, date string
        err = selDB.Scan(&id, &name, &autor, &date)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Autor = autor
        emp.Date = date
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM book WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := sbase{}
    for selDB.Next() {
        var id int
        var name, autor, date string
        err = selDB.Scan(&id, &name, &autor, &date)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Autor = autor
        emp.Date = date
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        autor := r.FormValue("autor")
        date := r.FormValue("date")
        insForm, err := db.Prepare("INSERT INTO book(name, autor, date) VALUES(?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, autor, date)
        log.Println("INSERT: Name: " + name + " | Autor: " + autor + " | Date: " + date)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        autor := r.FormValue("autor")
        date := r.FormValue("date")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE book SET name=?, autor=?, date=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(id, name, autor, date)
        log.Println("UPDATE: Name: " + name + " | Autor: " + autor + " | Date: " + date)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM book WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {
    log.Println("Server started on: http://192.168.1.54:8181")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe("192.168.1.54:8181", nil)
}