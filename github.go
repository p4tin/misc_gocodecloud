package main

import (
	"context"
	"log"
	"time"
	"encoding/json"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"net/http"
	//gojenkins "github.com/yosida95/golang-jenkins"
	//"net/url"
	//"strconv"
	//"fmt"
)

var client *github.Client

var (
	pending = "pending"
	success = "success"
	targetUrl = "https://XX.ngrok.com/status/"
	pendingDesc = "Build/testing in progress, please wait."
	successDesc = "Build/testing successful."
	appName = "URBN-CI-GO"
	jenkinsBaseUrl = "http://jenkins.urbn.com/"
	jenkinsUsername = "fortinp1"
	jenkinsJob = "PaulTestJob"
)

//Sensitive info
var (
	githubToken = "d6adafd20b851fff256c1cf86d86bd1af169ab6e"
	jenkinsApiToken = "c3ce006f2323ebb9a887e3e4652a6731"
)


func main() {
	glo := github.ListOptions{
		Page: 1,
		PerPage: 9999,
	}
	rlo := github.RepositoryListOptions{
		Visibility: "all",
		ListOptions: glo,
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},		//pfortin-urbn token
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client = github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List("urbn", &rlo)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("URBN has %d repositories.", len(repos))
	for _, repo := range repos {
		log.Println(*repo.Name)
	}


	//hook_name := "web"
	//True := true
	//config := make(map[string]interface{})
	//config["url"] = "http://3b549b58.ngrok.io/event"
	//config["content_type"] = "json"
	//hook, _, err := client.Repositories.CreateHook(
	//	"p4tin",
	//	"GoAws",
	//	&github.Hook{
	//		Name: &hook_name,
	//		Active: &True,
	//		Events: []string{"pull-request"},
	//		Config: config,
	//	},
	//)
	//if err != nil {
	//	log.Panic(err)
	//}
	//log.Println(hook)



	//http.HandleFunc("/event", EventHandler)
	//
	////ngrok http -subdomain="p4tin" 8081
	//
	//log.Println("Listening...")
	//log.Fatal(http.ListenAndServe(":8082", nil))
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Message: %+v\n", r.Body)
	v := new(github.PullRequest)
	json.NewDecoder(r.Body).Decode(v)

	if v.Number != nil {
		pr, _, err := client.PullRequests.Get("pfortin-urbn", "hooktesting", *v.Number)
		if err != nil {
			panic(err)
		}
		if *pr.State == "open" {
			log.Println(*pr.State)
			status1 := &github.RepoStatus{State: &pending, TargetURL: &targetUrl, Description: &pendingDesc, Context: &appName}
			client.Repositories.CreateStatus("pfortin-urbn", "hooktesting", *pr.Head.SHA, status1)
			log.Println(pr)

			//jauth := &gojenkins.Auth{ApiToken: jenkinsApiToken, Username: jenkinsUsername}
			//jenkins := gojenkins.NewJenkins(jauth, jenkinsBaseUrl)
			//job, err := jenkins.GetJob("PaulTestJob")
			//if err != nil {
			//	log.Printf("Jenkins ERROR: %+v\n", err)
			//} else {
			//	fmt.Printf("%s\n", job)
			//	params := url.Values{
			//		"PR": []string{strconv.Itoa(*v.Number)},
			//		"SHA": []string{"test1"},
			//	}
			//	status := jenkins.Build(job, params)
			//	log.Printf("Status: %+v\n", status)
			//}


			time.Sleep(20 * time.Second)
			status2 := &github.RepoStatus{State: &success, TargetURL: &targetUrl, Description: &successDesc, Context: &appName}
			client.Repositories.CreateStatus("pfortin-urbn", "hooktesting", *pr.Head.SHA, status2)
		} else {
			log.Printf("PR event but not open state. (%s)", *pr.State)
		}
	} else {
		log.Printf("Hook sent event that was not a PR - %+v\n", v)
	}
}
