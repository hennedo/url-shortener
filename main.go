package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Link struct {
	bongo.DocumentBase `bson:",inline"`
	Name, Url string
	ClicksFacebook, ClicksInstagram, ClicksOther, ClicksNone, Clicks int
}

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	flag.String("mongodb", "localhost", "MongoDB Connection String")
	flag.String("virtual-host", "localhost", "Domain where the shortener is found")
	flag.Bool("ssl", false, "Enable SSL")
	flag.String("admin-password", "foobar2342", "Password for the admin endpoint")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	config := &bongo.Config{
		ConnectionString: viper.GetString("mongodb"),
	}
	connection, err := bongo.Connect(config)
	if err != nil {
		logrus.Fatal(err)
	}
	connection = connection
	logrus.Info("Connected to database")

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/", newShortUrl).Methods("POST")
	r.HandleFunc("/henne", henneHandler).Methods("GET")
	r.HandleFunc("/🐓", henneHandler).Methods("GET")
	r.PathPrefix("/").HandlerFunc(notFoundHandler)
	http.Handle("/", r)
	logrus.Fatal(http.ListenAndServe(":8000", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "index.html")
}

func henneHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "index_henne.html")
}

func newShortUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Test", "whoo2p")

}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "404 not found")
	logrus.Info(r.URL)
}

func serveFile(w http.ResponseWriter, filename string) {
	body, err := ioutil.ReadFile("htmlfiles/" + filename)
	if err != nil {
		logrus.Warn(err)
		fmt.Fprintf(w, "Etwas ging schief")
	}
	fmt.Fprintf(w, "%s", body)
}
