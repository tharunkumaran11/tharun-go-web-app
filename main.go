package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, `<!DOCTYPE html><html><head><title>Tharun Kumaran.M | DevOps & Cloud</title><link rel="stylesheet" href="/static/style.css"></head><body>
<header><h1>Tharun Kumaran</h1><nav><a href="/">Home</a><a href="/about">About</a><a href="/projects">Projects</a><a href="/contact">Contact</a></nav></header>
<main class="hero"><p class="eyebrow">DEVOPS & CLOUD ENTHUSIAST</p><h2>Hi, I'm Tharun Kumaran.M.</h2><p>I am a DevOps and Cloud enthusiast passionate about automation, cloud infrastructure, containers, Kubernetes and CI/CD.</p><a class="button" href="/about">About Me</a></main>
<footer>Built with Go</footer></body></html>`)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" { http.NotFound(w, r); return }
	fmt.Fprint(w, `<!DOCTYPE html><html><head><title>About | Tharun Kumaran</title><link rel="stylesheet" href="/static/style.css"></head><body>
<header><h1>Tharun Kumaran</h1><nav><a href="/">Home</a><a href="/about">About</a><a href="/projects">Projects</a><a href="/contact">Contact</a></nav></header>
<main class="content"><h2>About Me</h2><p>I'm focused on building hands-on DevOps and Cloud projects and continuously improving my practical skills.</p><p>I work with technologies such as Docker, Kubernetes, AWS, Helm, GitHub Actions and Argo CD, and I enjoy turning application code into reliable, automated deployments.</p><p>I am currently building projects, expanding my cloud and DevOps knowledge, and looking for opportunities with companies where I can contribute, learn from experienced engineers and grow as a DevOps professional.</p></main>
<footer>Built with Go</footer></body></html>`)
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projects" { http.NotFound(w, r); return }
	fmt.Fprint(w, `<!DOCTYPE html><html><head><title>Projects | Tharun Kumaran</title><link rel="stylesheet" href="/static/style.css"></head><body>
<header><h1>Tharun Kumaran</h1><nav><a href="/">Home</a><a href="/about">About</a><a href="/projects">Projects</a><a href="/contact">Contact</a></nav></header>
<main class="content"><h2>DevOps Projects</h2><div class="cards"><article><h3>Containerized Go Application</h3><p>Go application prepared for Docker-based deployment.</p></article><article><h3>Kubernetes Deployment</h3><p>Application deployment using Kubernetes workloads, Services and Ingress.</p></article><article><h3>CI/CD Pipeline</h3><p>Automated build, test and deployment workflow using GitHub Actions and GitOps.</p></article></div></main>
<footer>Built with Go</footer></body></html>`)
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact" { http.NotFound(w, r); return }
	fmt.Fprint(w, `<!DOCTYPE html><html><head><title>Contact | Tharun Kumaran</title><link rel="stylesheet" href="/static/style.css"></head><body>
<header><h1>Tharun Kumaran</h1><nav><a href="/">Home</a><a href="/about">About</a><a href="/projects">Projects</a><a href="/contact">Contact</a></nav></header>
<main class="content"><h2>Contact Me</h2><p>Email: <a href="mailto:tharunkumaranm@outlook.com">tharunkumaranm@outlook.com</a></p></main>
<footer>Built with Go</footer></body></html>`)
}

func healthPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"healthy"}`)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/projects", projectsPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/health", healthPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Tharun Go Web App listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
