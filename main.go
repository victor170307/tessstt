package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

var artists []models.Artist

func main() {
var err error
artists, err = api.FetchArtists()
if err != nil {
log.Fatal("Error:", err)
}

fmt.Printf("Loaded %d artists\n", len(artists))

http.HandleFunc("/", handleHome)
http.HandleFunc("/api/artists", handleArtistsAPI)

url := "http://localhost:8080"
fmt.Printf("Opening %s in browser...\n", url)

// Auto-open browser
switch runtime.GOOS {
case "windows":
exec.Command("cmd", "/c", "start", url).Run()
case "darwin":
exec.Command("open", url).Run()
case "linux":
exec.Command("xdg-open", url).Run()
}

fmt.Println("Server running on http://localhost:8080 (Press Ctrl+C to stop)")
log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
html := `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Groupie Tracker</title>
<style>
body { font-family: Arial, sans-serif; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: #333; padding: 20px; margin: 0; }
.container { max-width: 1200px; margin: 0 auto; }
h1 { color: white; text-align: center; margin-bottom: 30px; }
.grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); gap: 20px; }
.card { background: white; padding: 20px; border-radius: 8px; cursor: pointer; box-shadow: 0 4px 6px rgba(0,0,0,0.1); transition: all 0.3s; }
.card:hover { box-shadow: 0 8px 16px rgba(0,0,0,0.2); transform: translateY(-5px); }
.modal { display: none; position: fixed; top: 0; left: 0; width: 100%; height: 100%; background: rgba(0,0,0,0.6); z-index: 1; }
.modal.show { display: flex; align-items: center; justify-content: center; }
.modal-content { background: white; padding: 30px; border-radius: 8px; max-width: 500px; max-height: 80vh; overflow-y: auto; }
.close { float: right; cursor: pointer; font-size: 28px; font-weight: bold; color: #667eea; }
.close:hover { color: #764ba2; }
h2 { color: #667eea; margin-top: 0; }
ul { padding-left: 20px; }
</style>
</head>
<body>
<div class="container">
<h1>Groupie Tracker</h1>
<div class="grid" id="grid"></div>
<div id="modal" class="modal">
<div class="modal-content">
<span class="close" onclick="closeModal()">&times;</span>
<h2 id="modal-title"></h2>
<p id="modal-year"></p>
<p id="modal-album"></p>
<h4>Members:</h4>
<ul id="modal-members"></ul>
</div>
</div>
</div>
<script>
async function load() {
const r = await fetch('/api/artists');
const data = await r.json();
const grid = document.getElementById('grid');
data.forEach((a, i) => {
const card = document.createElement('div');
card.className = 'card';
card.innerHTML = '<strong>' + a.name + '</strong><br>' + a.creationDate;
card.onclick = () => showModal(a);
grid.appendChild(card);
});
}

function showModal(a) {
document.getElementById('modal-title').textContent = a.name;
document.getElementById('modal-year').textContent = 'Year: ' + a.creationDate;
document.getElementById('modal-album').textContent = 'Album: ' + a.firstAlbum;
const ul = document.getElementById('modal-members');
ul.innerHTML = '';
a.members.forEach(m => {
const li = document.createElement('li');
li.textContent = m;
ul.appendChild(li);
});
document.getElementById('modal').classList.add('show');
}

function closeModal() {
document.getElementById('modal').classList.remove('show');
}

load();
</script>
</body>
</html>`
fmt.Fprint(w, html)
}

func handleArtistsAPI(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
json := `[`
for i, a := range artists {
if i > 0 { json += `,` }
members := `[]`
if len(a.Members) > 0 {
members = `["`
for j, m := range a.Members {
if j > 0 { members += `","` }
members += strings.ReplaceAll(m, `"`, `\"`)
}
members += `"]`
}
json += fmt.Sprintf(`{"name":"%s","creationDate":%d,"firstAlbum":"%s","members":%s}`, 
strings.ReplaceAll(a.Name, `"`, `\"`), a.CreationDate, 
strings.ReplaceAll(a.FirstAlbum, `"`, `\"`), members)
}
json += `]`
fmt.Fprint(w, json)
}
