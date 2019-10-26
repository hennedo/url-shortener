package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	bson2 "gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var charset = []string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "U", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var emojiCharset = []string {"😀","😃","😄","😁","😆","😊","😂","😅","😇","🙂","🙃","😉","😌","😚","😙","😗","😘","😍","😋","😛","😝","😜","🤪","🤩","😎","🤓","🧐","🤨","😏","😒","😞","😔","😟","😖","😣","🙁","😕","😫","😩","😢","😭","😤","😠","😡","🤬","🤯","😳","😓","😥","😰","😨","😱","🤗","🤔","🤭","🤫","🤥","🙄","😬","😑","😐","😶","😯","😦","😧","😮","😲","🤐","😵","😪","🤤","😴","🤢","🤮","🤧","😷","🤒","🤕","🤑","🤠","😈","👿","👻","💩","🤡","👺","👹","💀","👽","👾","🤖","😻","😹","😸","😺","🎃","😼","😽","🙀","😿","😾","🙌","🤲","👐","👏","🤝","🤛","✊","👊","👎","👍","🤜","🤞","🤟","🤘","👇","👆","👉","👈","👌","✋","🤚","🖐","🖖","🖕","💪","🤙","👋","🙏","💍","💄","💋","👄","👁","👣","👃","👂","👅","👀","🧠","🗣","👤","👥","👩","👦","🧒","👧","👶","🧑","👨","👱","🧔","👳","👲","👴","🧓","👵","🧕","👮","👷","💂","🕵","🍳","🌾","🎓","🎤","💻","🏭","🏫","💼","🔧","🚒","🎨","🔬","🚀","👸","🤵","👰","🤴","🤶","🎅","🧙","🧟","🧛","🧝","🧜","👗","👔","👖","👕","👚","👙","👘","👠","👡","👢","🧣","🧤","🧦","👟","👞","🎩","🧢","👒","⛑","👜","👛","👝","👑","🎒","👓","🕶","🌂","🐶","🐱","🐭","🐹","🐰","🐯","🐨","🐼","🐻","🦊","🦁","🐮","🐷","🐽","🐸","🐒","🙊","🙉","🙈","🐵","🐔","🐧","🐦","🐤","🐣","🦇","🦉","🦅","🦆","🐥","🐺","🐗","🐴","🦄","🐝","🐞","🐚","🐌","🦋","🐛","🐜","🦗","🕷","🕸","🦂","🦕","🦖","🦎","🐍","🐢","🐙","🦑","🦐","🦀","🐡","🐋","🐳","🐬","🐟","🐠","🦈","🐊","🐅","🐆","🦓","🐫","🐪","🦏","🐘","🦍","🦒","🐃","🐂","🐄","🐎","🦌","🐐","🐑","🐏","🐖","🐕","🐩","🐈","🐓","🦃","🐿","🐀","🐁","🐇","🕊","🦔","🐾","🐉","🐲","🌵","🌱","🌴","🌳","🌲","🎄","🌿","☘","🍀","🎍","🎋","🍄","🍁","🍂","🍃","💐","🌷","🌹","🥀","🌺","🌝","🌞","🌻","🌼","🌸","🌛","🌜","🌚","🌕","🌖","🌓","🌒","🌑","🌘","🌗","🌔","🌙","🌎","🌍","🌏","✨","🌟","⭐","💫","💥","🔥","🌪","🌈","🌥","⛅","🌤","🌦","🌧","⛈","🌩","🌨","💨","🌬","⛄","💧","💦","🌊","🌫","🍏","🍎","🍐","🍊","🍋","🍈","🍓","🍇","🍉","🍌","🍒","🍑","🍍","🥥","🥝","🥒","🥦","🥑","🍆","🍅","🌶","🌽","🥕","🥔","🍠","🧀","🥨","🥖","🍞","🥐","🥚","🥞","🥓","🥩","🍟","🍔","🌭","🍖","🍗","🍕","🥪","🥙","🌮","🌯","🥗","🥘","🥫","🍝","🍜","🥟","🍱","🍣","🍛","🍲","🍤","🍙","🍚","🍘","🍥","🍨","🍧","🍡","🍢","🥠","🍦","🥧","🍰","🎂","🍮","🍩","🍿","🍫","🍬","🍭","🍪","🌰","🥜","🍯","🥛","🍶","🥤","🍼","🍺","🍻","🥂","🍷","🥃","🍴","🥄","🍾","🍹","🍸","🍽","🥣","🥡","🥢","🎾","🏈","🏀","⚽","🏐","🏉","🎱","🏓","🏸","⛳","🏏","🏑","🏒","🥅","🏹","🎣","🥊","🎽","⛷","🎿","🛷","🥌","⛸","🏂","🤼","🏆","🥇","🥈","🥉","🎫","🎗","🏵","🎖","🏅","🎟","🎪","🎭","🎼","🎧","🎬","🎹","🥁","🎷","🎺","🎸","🎮","🎳","🎯","🎲","🎻","🎰","🚗","🚕","🚙","🚌","🚎","🚐","🚑","🚓","🏎","🚚","🚛","🚜","🛴","🚲","🚍","🚔","🚨","🏍","🛵","🚘","🚖","🚡","🚠","🚄","🚝","🚞","🚋","🚃","🚅","🚈","🚂","🚇","🛬","🛫","🚉","🚊","🛩","💺","🛰","🛸","🛳","⛴","🚢","⛽","🗺","🚏","🚥","🚦","🚧","🗿","🗽","🗼","🏰","🏯","⛲","🎠","🎢","🎡","🏟","⛱","🏖","🏝","🏜","⛺","🏕","🗻","🏔","⛰","⌚","📱","🖥","🖨","🗜","💽","📼","📷","📽","🎞","📠","📺","🎛","⏱","🔦","💡","📞","📸","💶","💴","💵","💸","🔫","💣","🚬","🎊","🎉","🏮","🎏","💊","🚿","🗝","🚰","🔑","🛎","🛏","🛒", "❤️", "🧡","💛","💚","💙","💜","🖤","💯","✅","❎","💤","❔","❓","❕","❗"}
var connection *bongo.Connection

var base string
type Link struct {
	bongo.DocumentBase `bson:",inline"`
	Name string
	Url string
	ClicksFacebook int `bson:"clicksFacebook"`
	ClicksInstagram int `bson:"clicksInstagram"`
	ClicksOther int `bson:"clicksOther"`
	ClicksNone int `bson:"clicksNone"`
	Clicks int `bson:"clicks"`
}

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		FullTimestamp:             true,
	})
	flag.String("mongodb", "localhost", "MongoDB Connection String")
	flag.String("virtual-host", "localhost", "Domain where the shortener is found")
	flag.Int("virtual-host-port", 80, "Port where the shortener listens")
	flag.Bool("ssl", false, "Enable SSL")
	flag.String("admin-password", "foobar2342", "Password for the admin endpoint")
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

	if viper.GetBool("ssl") {
		base = "https://"
	} else {
		base = "http://"
	}
	base += viper.GetString("virtual-host")
	if viper.GetInt("virtual-host-port") != 80 {
		base = fmt.Sprintf("%s:%d", base, viper.GetInt("virtual-host-port"))
	}

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/", newShortUrl).Methods("POST")
	r.HandleFunc("/delete", deleteHandler).Methods("POST")
	r.HandleFunc("/getcount", countHandler).Methods("POST")
	r.HandleFunc("/henne", henneHandler).Methods("GET")
	r.HandleFunc("/favicon.ico", faviconHandler).Methods("GET")
	r.HandleFunc("/🐓", henneHandler).Methods("GET")
	r.HandleFunc("/{name}", redirectHandler).Methods("GET")
	r.PathPrefix("/").HandlerFunc(notFoundHandler)
	http.Handle("/", r)
	logrus.Fatal(http.ListenAndServe(":8000", nil))
}
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "favicon.ico")
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "index.html")
}

func henneHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "index_henne.html")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logrus.Warn(err)
		_, _ = fmt.Fprintf(w, "Etwas ging schief")
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	if password != viper.Get("admin-password") {
		logrus.Info(fmt.Sprintf("Auth failed on delete for \"%s\"... tried \"%s\"", name, password))
		_, _ = fmt.Fprintf(w, "Nicht authorisiert")
		return
	}
	if name == "" {
		_, _ = fmt.Fprintf(w, "Kein Name angegeben")
		return
	}
	err := connection.Collection("links").DeleteOne(bson2.M{"name": name})
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		logrus.Info(fmt.Sprintf("Short \"%s\" not found", name))
		_, _ = fmt.Fprintf(w, "Link nicht gefunden")
		return
	// crappingfuckfuckers warum gibt es denn nen string zurück mit not found.. wie dumm.
	} else if err != nil {
		if err.Error() == "not found" {
			logrus.Info(fmt.Sprintf("Short \"%s\" not found", name))
			_, _ = fmt.Fprintf(w, "Link nicht gefunden")
			return
		}
		logrus.Warn("test")
		logrus.Warn(err)
		_, _ = fmt.Fprintf(w, "Etwas ging schief")
		return
	}
	logrus.Info(fmt.Sprintf("Deleting \"%s\"", name))
	_, _ = fmt.Fprintf(w, "Link gelöscht!")
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logrus.Warn(err)
		_, _ = fmt.Fprintf(w, "Etwas ging schief")
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	if password != viper.Get("admin-password") {
		logrus.Info(fmt.Sprintf("Auth failed on getcount for \"%s\"... tried \"%s\"", name, password))
		_, _ = fmt.Fprintf(w, "Nicht authorisiert")
		return
	}
	if name == "" {
		_, _ = fmt.Fprintf(w, "Kein Name angegeben")
		return
	}
	link := &Link{}
	err := connection.Collection("links").FindOne(bson.M{"name": name}, link)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		logrus.Info(fmt.Sprintf("Short \"%s\" not found", name))
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "404 not found")
		return
	} else if err != nil {
		logrus.Warn(err)
		_, _ = fmt.Fprintf(w, "Etwas ging schief")
		return
	}
	logrus.Info(fmt.Sprintf("Getting counts for \"%s\"", name))
	fmt.Fprintf(w, "<html><body><h1>h%s/%s</h1><ul><li>click count: %d</li><li>facebook: %d</li><li>instagram: %d</li><li>other: %d</li><li>none: %d</li></ul></body></html>", base, link.Name, link.Clicks, link.ClicksFacebook, link.ClicksInstagram, link.ClicksOther, link.ClicksNone)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	link := &Link{}
	err := connection.Collection("links").FindOne(bson.M{"name": params["name"]}, link)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		logrus.Info(fmt.Sprintf("Short \"%s\" not found", link.Name))
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "404 not found")
		return
	} else if err != nil {
		logrus.Warn(err)
		_, _ = fmt.Fprintf(w, "Etwas ging schief")
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
		logrus.Warn(err)
	}
	logrus.Info(fmt.Sprintf("Redirecting %s to %s", link.Name, link.Url))
	http.Redirect(w, r, link.Url, 302)
}

func newShortUrl(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logrus.Warn(err)
		_, _ = fmt.Fprintf(w, "Etwas ging schief")
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
			logrus.Warn(err)
			_, _ = fmt.Fprintf(w, "Etwas ging schief")
			return
		}
		link = &Link{
			Url:url,
			Name:name,
		}
	}
	err := connection.Collection("links").Save(link)
	if err != nil {
		logrus.Warn(err)
		_,_ = fmt.Fprintf(w, "Etwas ging schief")
		return
	}
	logrus.Info(fmt.Sprintf("New Shorturl: %s redirects to %s", name, url))
	_, _ = fmt.Fprintf(w, "%s/%s", base, name)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = 	fmt.Fprintf(w, "404 not found")
	logrus.Info(fmt.Sprintf("URL not found: %s", r.URL))
}

func serveFile(w http.ResponseWriter, filename string) {
	body, err := ioutil.ReadFile("htmlfiles/" + filename)
	if err != nil {
		logrus.Warn(err)
		_, _ = fmt.Fprintf(w, "Etwas ging schief")
	}
	_, _ = fmt.Fprintf(w, "%s", body)
}

func addHttp(url string) string {
	r := regexp.MustCompile("^(((f|ht)tps?)|tg)://")
	if !r.MatchString(url) {
		url = "http://" + url
	}
	return url
}

func getUniqueRandomString(length int, emoji bool) (string, error) {
Start:
	name := randomString(length, emoji)
	link := &Link{}
	err := connection.Collection("links").FindOne(bson.M{"Name": name}, link)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		return name, nil
	} else if err != nil {
		return "", err
	} else {
		goto Start
	}
}

func randomString(length int, emoji bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomString := ""
	if emoji {
		for i := 0; i < length; i++ {
			randomString += emojiCharset[r.Intn(len(emojiCharset))]
		}
	} else {
		for i := 0; i < length; i++ {
			randomString += charset[r.Intn(len(charset))]
		}
	}
	return randomString
}
