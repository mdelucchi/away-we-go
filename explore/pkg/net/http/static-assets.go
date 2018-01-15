// To serve static assets like JavaScript, CSS and images, we use the inbuilt http.FileServer and point it to a url path. For the file server to work properly it needs to know, where to serve files from. We can do this like so:  

fs := http.FileServer(http.Dir("static/"))

// Once our file server is in place, we just need to point a url path at it, just like we did with the dynamic requests. One thing to note: In order to serve files correctly, we need to strip away a part of the url path. Usually this is the name of the directory our files live in.

http.Handle("/static/", http.StripPrefix("/static/", fs))