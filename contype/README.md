<h1>Contype </h1>
<code>go get github.com/kampsy/Go/contype</code>
<hr>
<p>Contype abstracts away all the boilerplate needed to serve the right media file in an Http server</p>
<p>Usage</p>
<pre>
  func home(w http.ResponseWriter, r *http.Request)  {
    path := r.URL.Path[1:]
    log.Printf("%s %s\n", r.Method, path)

    data, err := ioutil.ReadFile(path)
    if err != nil {
      log.Printf("err : %v\n", err)
    }
    <b>cont := contype.FileType(r.URL.Path)</b>
    w.Header().Set("Content-Type", cont)
    w.Header().Set("Server", "kampsy")
    w.Write(data)
  }
</pre>
