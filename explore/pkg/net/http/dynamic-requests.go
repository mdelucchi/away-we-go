// The net/http package contains all utilities needed to accept requests and handle them dynamically. We can register a new handler with the http.HandleFunc function. It’s first parameter takes a path to match and a function to execute as a second. In this example: When someone browses your websites (http://example.com/), he or she will be greeted with a nice message.

http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to my website!")
})

// For the dynamic aspect, the http.Request contains all information about the request and it’s parameters. You can read GET parameters with r.URL.Query().Get("token") or POST parameters (fields from an HTML form) with r.FormValue("email").