package main

import (
    "fmt"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Fyd struct {
	Id string
	Name string
	Address string
	City string
	Country string
	Description string
	Music string
}

type Page struct {
	Id string 
	Body []byte
}

type FydList struct {
	fyds []Fyd
}

func (p *Page) save() error {
	filename := p.Id + ".json"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*func addFyd(fyd Fyd) {
	var fyds []Fyd
	var data []byte
	body, err := ioutil.ReadFile("fyd-list.json")

	if err != nil {
		return
	}

	if body == nil {
		data, err = json.Marshal(fyd)
	} else {
		json.Unmarshal(body, &fyds)
		data, err = json.Marshal(fyds)
	}
	
	ioutil.WriteFile("fyd-list.json", data, 0600)
}*/

func addFyd(fyd Fyd) {
	var fyds []Fyd
	err := json.Unmarshal(ioutil.ReadFile("fyds.json"), &fyds)
	
	if err != nil {
		return
	}
	
	var fydList []Fyd
	
	//Check if exists the corresponding field fyd to replace
	for i := 0; i < len(fyds); i++ {
		if fyds[i].Id == fyd.Id {
			removeFyd(fyds, i)
		}
	}
	//Lets add a new or existing fyd
	fydList = append(fyds, fyd)
	
	ioutil.WriteFile("fyds.json", []byte(fydList), 0600)

}

func searchFyd(val string) {
	var fyds []Fyd
	err := json.Unmarshal(ioutil.ReadFile("fyds.json"), &fyds)
	
	if err != nil {
		return
	}
	
	var fydList []Fyd
	
	//Check if exists the corresponding field fyd to replace
	for i := 0; i < len(fyds); i++ {
		if fyds[i].Id == val || fyds[i].Name == val || fyds[i].City == val || fyds[i].Country == val || strings.Contains(fyds[i].Music, val) {
			//Lets add the matching fyd
			fydList = append(fydList, fyds[i])
		}
	}
	os.Remove("search.json")
	ioutil.WriteFile("search.json", []byte(fydList), 0600)
}

func removeFyd(fyds []Fyd, index Int) {
	lastIndex := len(fyds) - 1
	//swap the last value and the value we want to remove
	fyds[index], fyds[lastIndex] = fyds[lastIndex], fyds[index]
	//return fyds[:lastIndex]
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkHttp(err error, status http.StatusInternalServerError) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func load(id string) (*Page, error) {
	filename := id + ".json"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Id: id, Body: body}, nil
}

func byteContent(id string) ([]byte) {
	filename := id + ".json"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil
	}

	return body
}

func get(w http.ResponseWriter, r *http.Request, id string) {
	var fyds []Fyd
	var fl *FydList
	var body []byte
	
	if strings.Contains(id, "search") {
		body := byteContent(id)
	} else {
		body := byteContent(id)
	}
	
	json.Unmarshal(body, &fyds)
	fl = &FydList{fyds: fyds}
	renderViewList(w, "view", fl)
}

func searchById(w http.ResponseWriter, r *http.Request, id string) {
	
	s := r.FormValue("search")
	valId = "?id="+s
	searchVal = "search" + valId
	searchFyd(s)
	get(w, r, searchVal)
}

func getById(w http.ResponseWriter, r *http.Request, id string) {
	var fyd *Fyd
	p, err := load(id)
	if err != nil {
		return
	}
	json.Unmarshal(p.Body, &fyd)
	//how to use reference as type *Fyd
	renderView(w, "view", fyd)
}

func createView(w http.ResponseWriter, r *http.Request, id string) {
	var fyd *Fyd
	p, err := load(id)
	json.Unmarshal(p.Body, &fyd)
	if err != nil {
		p = &Page{Id: id}
	}
	renderView(w, "create", fyd)
}

func create(w http.ResponseWriter, r *http.Request, id string) {
	
	name := r.FormValue("name")
	city := r.FormValue("city")
	country := r.FormValue("country")
	address := r.FormValue("address")
	description := r.FormValue("description")
	music := r.FormValue("music")

	id_ := time.Now().String()
	fyd := Fyd {
		Id: id,
		Name: name,
		Address: address,
		City: city,
		Country: country,
		Description: description,
		Music: music,
	}
	fmt.Println("redirection vers la view")
	data, err := json.Marshal(fyd)
	p := &Page{Id: id, Body: []byte(data)}
	_err := p.save()
	//addFyd(fyd)

	if _err != nil {
		fmt.Println("redirection vers la view", err, id_)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	ioutil.WriteFile(id+".json", data, 0600)
	fmt.Println("redirection vers la view", http.StatusFound)
	statusFound := http.StatusFound
	http.Redirect(w, r, "/view/"+id, statusFound)
}

func update(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Println("redirection vers la view")
	name := r.FormValue("name")
	city := r.FormValue("city")
	country := r.FormValue("country")
	address := r.FormValue("address")
	description := r.FormValue("description")
	music := r.FormValue("music")

	fyd := Fyd {
		Id: id,
		Name: name,
		Address: address,
		City: city,
		Country: country,
		Description: description,
		Music: music,
	}
	fmt.Println("redirection vers la view")
	data, err := json.Marshal(fyd)
	p := &Page{Id: id, Body: []byte(data)}
	_err := p.save()
	//addFyd(fyd)

	if _err != nil {
		fmt.Println("redirection vers la view", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	os.Remove(id+".json")
	ioutil.WriteFile(id+".json", data, 0600)
	fmt.Println("redirection vers la view")
	statusFound := http.StatusFound
	http.Redirect(w, r, "/view/"+id, statusFound)
}

func updateView(w http.ResponseWriter, r *http.Request, id string) {
	var fyd *Fyd
	p, err := load(id)
	json.Unmarshal(p.Body, &fyd)
	if err != nil {
		p = &Page{Id: id}
	}
	renderView(w, "edit", fyd)
}

func deleteView(w http.ResponseWriter, r *http.Request, id string) {
	var fyd *Fyd
	p, err := load(id)
	json.Unmarshal(p.Body, &fyd)
	if err != nil {
		p = &Page{Id: id}
	}
	renderView(w, "delete", fyd)
}

func delete(w http.ResponseWriter, r *http.Request, id string) {
	os.Remove(id+".json") 
	http.Redirect(w, r, "/list/fyds", http.StatusFound)
}

var templates = template.Must(template.ParseFiles("create.html", "edit.html", "list.html", "delete.html", "view.html"))

func renderView(w http.ResponseWriter, tmpl string, f *Fyd) {
	err := templates.ExecuteTemplate(w, tmpl+".html", f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderViewList(w http.ResponseWriter, tmpl string, fl *FydList) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		log.Println("parsing files:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Execute the template for each fyd.
	for _, r := range fl.fyds {
		err := t.Execute(w, r)
		if err != nil {
			log.Println("executing template:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

var validPath = regexp.MustCompile("^/(create|edit|save|update|list|delete|view)/([a-zA-Z0-9]+)$")

func requestHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		fmt.Println(validPath.FindStringSubmatch(r.URL.Path))
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", requestHandler(getById))
	http.HandleFunc("/edit/", requestHandler(updateView))
	http.HandleFunc("/update/", requestHandler(update))
	http.HandleFunc("/create/", requestHandler(createView))
	http.HandleFunc("/delete/", requestHandler(deleteView))
	http.HandleFunc("/save/", requestHandler(create))
	http.HandleFunc("/remove/", requestHandler(delete))
	http.HandleFunc("/list/", requestHandler(get))
	http.HandleFunc("/search", requestHandler(searchById))
	log.Fatal(http.ListenAndServe(":8083", nil))
}
