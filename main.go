package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var charset = []string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "U", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var emojiCharset = []string {"😀","😃","😄","😁","😆","😊","😂","😅","😇","🙂","🙃","😉","😌","😚","😙","😗","😘","😍","😋","😛","😝","😜","🤪","🤩","😎","🤓","🧐","🤨","😏","😒","😞","😔","😟","😖","😣","🙁","😕","😫","😩","😢","😭","😤","😠","😡","🤬","🤯","😳","😓","😥","😰","😨","😱","🤗","🤔","🤭","🤫","🤥","🙄","😬","😑","😐","😶","😯","😦","😧","😮","😲","🤐","😵","😪","🤤","😴","🤢","🤮","🤧","😷","🤒","🤕","🤑","🤠","😈","👿","👻","💩","🤡","👺","👹","💀","👽","👾","🤖","😻","😹","😸","😺","🎃","😼","😽","🙀","😿","😾","🙌","🤲","👐","👏","🤝","🤛","✊","👊","👎","👍","🤜","🤞","🤟","🤘","👇","👆","👉","👈","👌","✋","🤚","🖐","🖖","🖕","💪","🤙","👋","🙏","💍","💄","💋","👄","👁","👣","👃","👂","👅","👀","🧠","🗣","👤","👥","👩","👦","🧒","👧","👶","🧑","👨","👱","🧔","👳","👲","👴","🧓","👵","🧕","👮","👷","💂","🕵","🍳","🌾","🎓","🎤","💻","🏭","🏫","💼","🔧","🚒","🎨","🔬","🚀","👸","🤵","👰","🤴","🤶","🎅","🧙","🧟","🧛","🧝","🧜","👗","👔","👖","👕","👚","👙","👘","👠","👡","👢","🧣","🧤","🧦","👟","👞","🎩","🧢","👒","⛑","👜","👛","👝","👑","🎒","👓","🕶","🌂","🐶","🐱","🐭","🐹","🐰","🐯","🐨","🐼","🐻","🦊","🦁","🐮","🐷","🐽","🐸","🐒","🙊","🙉","🙈","🐵","🐔","🐧","🐦","🐤","🐣","🦇","🦉","🦅","🦆","🐥","🐺","🐗","🐴","🦄","🐝","🐞","🐚","🐌","🦋","🐛","🐜","🦗","🕷","🕸","🦂","🦕","🦖","🦎","🐍","🐢","🐙","🦑","🦐","🦀","🐡","🐋","🐳","🐬","🐟","🐠","🦈","🐊","🐅","🐆","🦓","🐫","🐪","🦏","🐘","🦍","🦒","🐃","🐂","🐄","🐎","🦌","🐐","🐑","🐏","🐖","🐕","🐩","🐈","🐓","🦃","🐿","🐀","🐁","🐇","🕊","🦔","🐾","🐉","🐲","🌵","🌱","🌴","🌳","🌲","🎄","🌿","☘","🍀","🎍","🎋","🍄","🍁","🍂","🍃","💐","🌷","🌹","🥀","🌺","🌝","🌞","🌻","🌼","🌸","🌛","🌜","🌚","🌕","🌖","🌓","🌒","🌑","🌘","🌗","🌔","🌙","🌎","🌍","🌏","✨","🌟","⭐","💫","💥","🔥","🌪","🌈","🌥","⛅","🌤","🌦","🌧","⛈","🌩","🌨","💨","🌬","⛄","💧","💦","🌊","🌫","🍏","🍎","🍐","🍊","🍋","🍈","🍓","🍇","🍉","🍌","🍒","🍑","🍍","🥥","🥝","🥒","🥦","🥑","🍆","🍅","🌶","🌽","🥕","🥔","🍠","🧀","🥨","🥖","🍞","🥐","🥚","🥞","🥓","🥩","🍟","🍔","🌭","🍖","🍗","🍕","🥪","🥙","🌮","🌯","🥗","🥘","🥫","🍝","🍜","🥟","🍱","🍣","🍛","🍲","🍤","🍙","🍚","🍘","🍥","🍨","🍧","🍡","🍢","🥠","🍦","🥧","🍰","🎂","🍮","🍩","🍿","🍫","🍬","🍭","🍪","🌰","🥜","🍯","🥛","🍶","🥤","🍼","🍺","🍻","🥂","🍷","🥃","🍴","🥄","🍾","🍹","🍸","🍽","🥣","🥡","🥢","🎾","🏈","🏀","⚽","🏐","🏉","🎱","🏓","🏸","⛳","🏏","🏑","🏒","🥅","🏹","🎣","🥊","🎽","⛷","🎿","🛷","🥌","⛸","🏂","🤼","🏆","🥇","🥈","🥉","🎫","🎗","🏵","🎖","🏅","🎟","🎪","🎭","🎼","🎧","🎬","🎹","🥁","🎷","🎺","🎸","🎮","🎳","🎯","🎲","🎻","🎰","🚗","🚕","🚙","🚌","🚎","🚐","🚑","🚓","🏎","🚚","🚛","🚜","🛴","🚲","🚍","🚔","🚨","🏍","🛵","🚘","🚖","🚡","🚠","🚄","🚝","🚞","🚋","🚃","🚅","🚈","🚂","🚇","🛬","🛫","🚉","🚊","🛩","💺","🛰","🛸","🛳","⛴","🚢","⛽","🗺","🚏","🚥","🚦","🚧","🗿","🗽","🗼","🏰","🏯","⛲","🎠","🎢","🎡","🏟","⛱","🏖","🏝","🏜","⛺","🏕","🗻","🏔","⛰","⌚","📱","🖥","🖨","🗜","💽","📼","📷","📽","🎞","📠","📺","🎛","⏱","🔦","💡","📞","📸","💶","💴","💵","💸","🔫","💣","🚬","🎊","🎉","🏮","🎏","💊","🚿","🗝","🚰","🔑","🛎","🛏","🛒", "❤️", "🧡","💛","💚","💙","💜","🖤","💯","✅","❎","💤","❔","❓","❕","❗"}
var connection *bongo.Connection
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
	var err error
	connection, err = bongo.Connect(config)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Connected to database")

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/", newShortUrl).Methods("POST")
	r.HandleFunc("/henne", henneHandler).Methods("GET")
	r.HandleFunc("/🐓", henneHandler).Methods("GET")
	r.HandleFunc("/{name}", redirectHandler).Methods("GET")
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

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	link := &Link{}
	err := connection.Collection("links").FindOne(bson.M{"name": params["name"]}, link)
	if err != nil {
		logrus.Warn(err)
		fmt.Fprintf(w, "Etwas ging schief")
		return
	}
	http.Redirect(w, r, link.Url, 302)
}

func newShortUrl(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logrus.Warn(err)
		fmt.Fprintf(w, "Etwas ging schief")
	}

	url := r.FormValue("url")
	emoji := r.FormValue("emoji")
	password := r.FormValue("password")
	name := randomString(6, emoji == "1")
	if password != "" && password == viper.GetString("admin-password") {
		if name != "" {
			name = r.FormValue("name")
		}
	}
	newLink := &Link{
		Url:url,
		Name:name,
	}
	err := connection.Collection("links").Save(newLink)
	if err != nil {
		logrus.Warn(err)
	}
	fmt.Fprintf(w, "http://" + viper.GetString("virtual-host") + "/" + name)
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

func randomString(length int, emoji bool) string {
	randomString := ""
	if emoji {
		for i := 0; i < length; i++ {
			randomString += emojiCharset[rand.Intn(len(emojiCharset))]
		}
	} else {
		for i := 0; i < length; i++ {
			randomString += charset[rand.Intn(len(charset))]
		}
	}
	return randomString
}
