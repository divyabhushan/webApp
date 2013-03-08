package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", signIn)
	http.ListenAndServe(":8090", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, signInPage)
}

const signInPage = `
<html>
    
    <body>
        <h1>User input form</h1>
        <form action="sign" method="post">
            <div>
                Name: <input type="text" name="Username">
                
            </div>
            <div>
                <br><input type="submit" value="Sign in">
            </div>
        </form>
    </body>

</html>
`

func signIn(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("Username"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var signTemplate = template.Must(template.New("signIn").Parse(signTemplateHTML))

const signTemplateHTML = `
<html>
  <body>
    <p>Welcome user:</p>
    <pre>{{.}}</pre>
  </body>
</html>
`
