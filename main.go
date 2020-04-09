package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	bson2 "gopkg.in/mgo.v2/bson"
)

var connection *bongo.Connection

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		FullTimestamp:             true,
	})
	flag.String("mongodb", "localhost", "MongoDB Connection String")
	flag.Int("port", 8000, "Port where the url shortener listens")
	flag.String("admin-password", "foobar2342", "Password for the admin endpoint")
	flag.String("base-url", "http://localhost:8000", "Baseurl of the URL shortener")
	_ = viper.BindPFlags(flag.CommandLine)
	flag.Parse()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	config := &bongo.Config{
		ConnectionString: viper.GetString("mongodb"),
	}
	var err error
	connection, err = bongo.Connect(config)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Connected to database")
	_ = connection.Collection("links").Collection().EnsureIndex(mgo.Index{
		Key:              []string{"name"},
		Unique:           true,
	})

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/", newShortUrl).Methods("POST")
	r.HandleFunc("/delete", deleteHandler).Methods("POST")
	r.HandleFunc("/getcount", countHandler).Methods("POST")
	r.HandleFunc("/scam", scamHandler).Methods("POST")
	r.HandleFunc("/admin", adminHandler).Methods("GET")
	r.HandleFunc("/favicon.ico", faviconHandler).Methods("GET")
	r.HandleFunc("/{name}", redirectHandler).Methods("GET")
	r.HandleFunc("/{name}", redirectHeadHandler).Methods("HEAD")
	r.HandleFunc("/{name}", redirectHandler).Methods("POST")
	r.PathPrefix("/").HandlerFunc(notFoundHandler)
	http.Handle("/", r)
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), nil))
}
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "favicon.ico")
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "index.html")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "index_admin.html")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		returnError500(err, w);
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	if password != viper.Get("admin-password") {
		logrus.Info(fmt.Sprintf("Auth failed on delete for \"%s\"... tried \"%s\"", name, password))
		returnError401(w)
		return
	}
	if name == "" {
		_, _ = fmt.Fprintf(w, "We need a link name..")
		return
	}
	err := connection.Collection("links").DeleteOne(bson2.M{"name": name})
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		logrus.Info(fmt.Sprintf("Short \"%s\" not found", name))
		returnError404(w)
		return
		// crappingfuckfuckers warum gibt es denn nen string zurück mit not found.. wie dumm.
	} else if err != nil {
		if err.Error() == "not found" {
			logrus.Info(fmt.Sprintf("Short \"%s\" not found", name))
			returnError404(w)
			return
		}
		returnError500(err, w)
		return
	}
	logrus.Info(fmt.Sprintf("Deleting \"%s\"", name))
	_, _ = fmt.Fprintf(w, "Link deleted!")
}

func scamHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		returnError500(err, w);
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	if password != viper.Get("admin-password") {
		logrus.Info(fmt.Sprintf("Auth failed on delete for \"%s\"... tried \"%s\"", name, password))
		returnError401(w)
		return
	}
	if name == "" {
		_, _ = fmt.Fprintf(w, "We need a link name..")
		return
	}
	err, link := getLink(name)
	if err != nil {
		if err.Error() == "404" {
			returnError404(w)
			return
		}
		returnError500(err, w)
		return
	}
	link.Scam = true
	err = connection.Collection("links").Save(link)
	if err != nil {
		returnError500(err, w)
	} else {
		logrus.Info(fmt.Sprintf("Scamming \"%s\"", name))
		_, _ = fmt.Fprintf(w, "Link Scammed!")
	}
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		returnError500(err, w)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	if password != viper.Get("admin-password") {
		logrus.Info(fmt.Sprintf("Auth failed on getcount for \"%s\"... tried \"%s\"", name, password))
		returnError401(w)
		return
	}
	if name == "" {
		_, _ = fmt.Fprintf(w, "We need a link name.")
		return
	}
	err, link := getLink(name)
	if err != nil {
		if err.Error() == "404" {
			returnError404(w)
			return
		}
		returnError500(err, w)
		return
	}
	logrus.Info(fmt.Sprintf("Getting counts for \"%s\"", name))
	fmt.Fprintf(w, "<html><body><h1>h%s/%s</h1><ul><li>click count: %d</li><li>facebook: %d</li><li>instagram: %d</li><li>other: %d</li><li>none: %d</li></ul></body></html>", viper.GetString("base-url"), link.Name, link.Clicks, link.ClicksFacebook, link.ClicksInstagram, link.ClicksOther, link.ClicksNone)
}
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err, link := getLink(params["name"])
	if err != nil {
		if err.Error() == "404" {
			returnError404(w)
			return
		}
		returnError500(err, w)
		return
	}
	if link.Scam && r.Method != "POST" {
		tmpl := template.Must(template.ParseFiles("htmlfiles/scam.html"))
		err := tmpl.ExecuteTemplate(w, "scam.html", link)
		if err != nil {
			logrus.Error(err)
		}
		return
	}
	referer := r.Header.Get("referer")
	link.Clicks++
	if referer == "" {
		link.ClicksNone++
	} else if strings.Contains(referer, "facebook.com") {
		link.ClicksFacebook++
	} else if strings.Contains(referer, "instagram.com") {
		link.ClicksInstagram++
	} else {
		link.ClicksOther++
	}
	err = connection.Collection("links").Save(link)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(fmt.Sprintf("Redirecting %s to %s", link.Name, link.Url))
	http.Redirect(w, r, link.Url, 302)
}
func redirectHeadHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err, link := getLink(params["name"])
	if err != nil {
		if err.Error() == "404" {
			returnError404(w)
			return
		}
		returnError500(err, w)
		return
	}
	if link.Scam && r.Method != "POST" {
		tmpl := template.Must(template.ParseFiles("htmlfiles/scam.html"))
		err := tmpl.ExecuteTemplate(w, "scam.html", link)
		if err != nil {
			logrus.Error(err)
		}
		return
	}
	referer := r.Header.Get("referer")
	link.Clicks++
	if referer == "" {
		link.ClicksNone++
	} else if strings.Contains(referer, "facebook.com") {
		link.ClicksFacebook++
	} else if strings.Contains(referer, "instagram.com") {
		link.ClicksInstagram++
	} else {
		link.ClicksOther++
	}
	err = connection.Collection("links").Save(link)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(fmt.Sprintf("Redirecting %s to %s", link.Name, link.Url))
	http.Redirect(w, r, link.Url, 302)
}

func newShortUrl(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		returnError500(err, w)
		return
	}

	url := addHttp(r.FormValue("url"))
	emoji := r.FormValue("emoji")
	password := r.FormValue("password")
	var name string
	link := &Link{}
	if password != "" && password == viper.GetString("admin-password") && r.FormValue("name") != "" {
		name = r.FormValue("name")
		err := connection.Collection("links").FindOne(bson.M{"name": name}, link)
		if err != nil {
			link = &Link{
				Url:url,
				Name:name,
			}
		}
		link.Url = url
	} else {
		var err error
		name, err = getUniqueRandomString(6, emoji == "1")
		if err != nil {
			returnError500(err, w)
			return
		}
		link = &Link{
			Url:url,
			Name:name,
		}
	}
	err := connection.Collection("links").Save(link)
	if err != nil {
		returnError500(err, w)
		return
	}
	logrus.Info(fmt.Sprintf("New Shorturl: %s redirects to %s", name, url))
	_, _ = fmt.Fprintf(w, "%s/%s", viper.GetString("base-url"), name)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	returnError404(w)
	logrus.Info(fmt.Sprintf("URL not found: %s %s", r.URL, r.Method))
}
