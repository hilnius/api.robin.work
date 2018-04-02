// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

// [START import_statements]
import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine" // Required external AppEngine library
) // [END import_statements]

// [START main_func]
func main() {
	http.HandleFunc("/api/", indexHandler)
	http.HandleFunc("/api/blog/", blogHandler)
	appengine.Main() // Starts the server to receive requests
} // [END main_func]

// [START indexHandler]
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "Hello world!")
} // [END indexHandler]

func escapeQuotes(str string) string {
	r := strings.NewReplacer("\"", "\\\"")
	return r.Replace(str)
}

func returnJsonWithAllowedOrigins(w http.ResponseWriter, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "{\"content\":\""+escapeQuotes(message)+"\"}")
}

// [START indexHandler]
func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/blog/1" {
		returnJsonWithAllowedOrigins(w, "<h1>Hello world!</h1><p>Welcome to my blog. It exists so I can write stories from time to time. You'll find almost exclusively articles about programming and my experience, as I don't enjoy exposing my life - Facebook does that enough.</p><h2>A few things about me</h2><p>My name's Robin, 25, and I come from a small village lost somewhere in the Alps. I have graduated my masters degree in Telecom-Bretagne (now IMT Atlantique). I've worked at Content-Square as a fullstack developer intern in 2014, and then on a project called <a target=\"_blank\" href='http://www.sentimy.com'>Sentimy</a>. We pivoted to <a href='https://www.supermood.com' target=\"_blank\">Supermood</a> which is where I work as CTO since september 2015.</p><h2>Hobbies</h2><p>I have always been a big fan of sports, my preferred being Tennis, Badminton, and of course Skiing (cross-country and alpine). I'm very competitive and I've always been. I believe the reason I like competition so much is because it brings the best in me, and it gives me a measurable way to become better. I also enjoy hiking but I don't do that very often, there aren't any mountains here in Paris.</p><p>I enjoy movies - like everybody - and I learned how to use some tools like Adobe After Effects, Premiere Pro or Cinema4d for the fun of making my own small videos. Last but not least, as every developer, I like beer. So if you're up for a pint hit me up!</p>")
	} else if r.URL.Path == "/api/blog/2" {
		returnJsonWithAllowedOrigins(w, "<h1>Migrating from AWS to AppEngine</h1><p>For Sentimy, we originally an OVH shared hosting service inclusding one database. When we realized we didn't have enough control on the machine, and that it couldn't support big loads (surveys generate traffic peaks) we decided to migrate to AWS. I spent a lot of time administrating machines on AWS, and realized at some point that <strong>load balancing was a headache</strong>. Indeed, machines need around 2 minutes to get up and running, so you need to permanently have 2 running servers plus the load balancer plus the database, which <strong>costs a lot</strong>. Plus we had to create automatic deployments which adds to the length of the deployment process. I had been looking for quite some time <strong>Google AppEngine</strong>. To keep it simple, it's advantages are that it's way easier to setup, deploy and manage. It's drawbacks are that you are limited to some languages and you do not have access to machines making it impossible to install custom software (you won't be able to <code>ssh</code> into your instances, ever).</p><p>It was around the time we had decided to create Supermood so I decided to spend 3 days trying to setup Supermood on AppEngine. After half a day of work, I had my development environment working, and a demo version running online. I saw with great satisfaction that it took me less than a minute for a new version to be deployed. The main limitation, once again, is that you are limited to the sofware Google makes available in AppEngine. I saw quite happily that I <strong>didn't have to do anything to make load balancing work</strong>, and that when nobody was on the website there was no running instance at all! This has a simple yet important consequence: <strong>when nobody is on our website, it does not cost us anything</strong>. Database backups are easy to setup and the whole project costs me between 7 and 8 euros every month (small traffic, around 100 visits a day). This is nothing compared to the 40€ I was spending every month on Amazon Web Services. One thing that bugs me is that Google bills you 'instance-hours' which is a very unclear metric. Indeed, you don't know what an instance is and I didn't find to this day the definition of an instance. If Google decides to create a thousand instances to handle 10 people, I'll have to pay, which isn't so nice.</p><p>That said, I'm not saying that AWS is bad and GCP (Google Cloud Platform) is better. But my conclusion is that if you do not have a complicated stack, and if you can comply to the strict rules of Google AppEngine, it's worth using. <strong>I saved weeks in sys-admin tasks, and hundreds of euros</strong>. One thing's sure, it helped me a lot. As a proof, a week later, I did what was necessary to completely move Sentimy from AWS to GAE.</p><p>Guess where is this website hosted?</p>")
	} else if r.URL.Path == "/api/blog/3" {
		returnJsonWithAllowedOrigins(w, "<h1>ngRights: Angular library for ABAC policies</h1><p>During my time at Supermood, I found it difficult to <strong>conditionnaly show or hide DOM elements depending on the rights that the user had</strong>. Sometimes I wanted to have a manager display a list of employees. He could delete only the employees that belong to his team. That lead to an awfully complex and hard-to-read <code>ng-show</code> in my HTML code. The solution was to create a library that allows to define ABAC policies and use them easily directly inside the DOM. Basically, you write <strong>rules</strong> in JavaScript that decide wether to show the element or not, based on the <strong>User</strong>, some <strong>Object</strong>, and a <strong>Context</strong>.</p><p>The list of cool things that you can achieve with that library is just limited by imagination. A more thorough description on its usage is in our github. Don't hesitate to check it out, fork it, create issues and comment to help us make it even better!<br><br><a href='https://github.com/Supermood/ngRights' target=\"_blank\">ngRights on github</a></p>")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
} // [END indexHandler]
