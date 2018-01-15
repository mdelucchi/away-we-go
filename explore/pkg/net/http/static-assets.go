// serve static assets like JavaScript, CSS and images
// use the inbuilt http.FileServer and point it to a url path
// to work properly it needs to know, where to serve files from.

fs := http.FileServer(http.Dir("static/"))

// point a url path at it
// In order to serve files correctly, strip away a part of the url path
// typically the name of the directory our files live in.

http.Handle("/static/", http.StripPrefix("/static/", fs))