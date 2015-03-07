/* tick-tock.go */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// handler for the main page.
func handler(w http.ResponseWriter, r *http.Request) {
	// Content for the main html page..

	page, _ := ioutil.ReadFile("html/index.html")
	fmt.Fprint(w, string(page))
}

// handler to cater AJAX requests
func handlerDevs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<font color=red>Dev1<br></font>")
}

func handlerV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<font color=blue>Vertical1<br></font>")
}

func handlerV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<font color=green>Vertical2<br></font>")
}

func handlerSett(w http.ResponseWriter, r *http.Request) {
	setpage := `
  <html>

  <head>
  <title>Настройки</title>
  <meta name="generator" content="Namo WebEditor v5.0">
  </head>

  <body bgcolor="white" text="black" link="blue" vlink="purple" alink="red">
  <p>&nbsp;</p>
  <div id="layer1" style="background-color:rgb(0,204,255); width:832px; height:398px; position:absolute; left:153px; top:25px; z-index:1; layer-background-color:rgb(0,204,255); ">
    <p align="center"><font color="blue"><b>Настройки программы</b></font></p>
    <br>
      <form name="form1" action="/setting" method="Post">
    URL сервера &nbsp;<input type="text" name="srv_url"> &nbsp;Логин &nbsp;<input type="text" name="login">
    &nbsp;Пароль &nbsp;<input type="text" name="pass">
          <p align="center"><input type="submit" name="Сохранить"></p>
      </form>
  </div>
  </body>

  </html>
  `
	fmt.Fprint(w, setpage)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/jquery.min.js", SendJqueryJs)
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/dev", handlerDevs)
	mux.HandleFunc("/v1", handlerV1)
	mux.HandleFunc("/v2", handlerV2)
	mux.HandleFunc("/sett", handlerSett)
	mux.HandleFunc("/setting", handlerSetting)
	http.ListenAndServe(":9999", mux)
}

func handlerSetting(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("srv_url")
	fmt.Fprint(w, name)
}

func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("jquery.min.js")
	if err != nil {
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(data)
}
