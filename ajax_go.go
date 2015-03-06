/* tick-tock.go */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Content for the main html page..
var page = `<html>
           <head>
             <script type="text/javascript"
               src="http://localhost:9999/jquery.min.js">
             </script>
             <style> 
               div {
                 font-family: "Times New Roman", Georgia, Serif;
                 font-size: 1em;
                 width: 13.3em;
                 padding: 8px 8px; 
                 border: 2px solid #2B1B17;
                 color: #2B1B17;
                 text-shadow: 1px 1px #E5E4E2;
                 background: #FFFFFF;
               }
             </style>
           </head>
           <body>
             <h2 align=center>Go Timer </h2>
             <div id="output" style="width: 30%; height: 63%; overflow-y: scroll; float:left;"></div>
             <div id="v1" style="width: 50%; height: 30%; overflow-y: scroll; float:left;"></div>
             <div id="v2" style="width: 50%; height: 30%; overflow-y: scroll; float:left;"></div>
             <input id="sett" type="submit" name="sett" value="Settings" onclick="changeUrl()">

             <script type="text/javascript">

               var myDelay;

               $(document).ready(function () 
               {
                   $("#output").append("Waiting for system time..");
                   
                   myDelay = setInterval("delayedPost()", 1000);               

                });

               function delayedPost() 
               {
                 $.post("http://localhost:9999/dev", "", function(data, status) 
                 {
                    //$("#output").empty();
                    $("#output").prepend(data);
                 });

                 $.post("http://localhost:9999/v1", "", function(data, status) {
                    //$("#output").empty();
                    $("#v1").prepend(data);
                 });

                 $.post("http://localhost:9999/v2", "", function(data, status) {
                    //$("#output").empty();
                    $("#v2").prepend(data);
                 });
               }

               function delayedPost1() 
               {
                 $.post("http://localhost:9999/dev", "", function(data, status) 
                 {                    
                    $("#output").prepend(data);
                 });

                 $.post("http://localhost:9999/v1", "", function(data, status) 
                 {                   
                    $("#v1").prepend(data);
                 });

                 $.post("http://localhost:9999/v3", "", function(data, status) 
                 {                    
                    $("#v2").prepend(data);
                 });
               }

               function changeUrl()
               {
                  alert('salom');                      
                  clearInterval(myDelay);

                  window.location.replace("http://beta.biotrack.uz");

               }





             </script>
           </body>
         </html>`

// handler for the main page.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}

// handler to cater AJAX requests
func handlerDevs(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
	fmt.Fprint(w, "<font color=red>Dev1<br></font>")
}

func handlerV1(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
	fmt.Fprint(w, "<font color=blue>Vertical1<br></font>")
}

func handlerV2(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
	fmt.Fprint(w, "<font color=green>Vertical2<br></font>")
}

func main() {
	http.HandleFunc("/jquery.min.js", SendJqueryJs)
	http.HandleFunc("/", handler)
	http.HandleFunc("/dev", handlerDevs)
	http.HandleFunc("/v1", handlerV1)
	http.HandleFunc("/v2", handlerV2)
	log.Fatal(http.ListenAndServe(":9999", nil))
	//panic(http.ListenAndServe(":8081", nil))
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
