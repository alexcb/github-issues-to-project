package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v51/github"
	"github.com/machinebox/graphql"
	"golang.org/x/oauth2"
)

func die(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func GetIssues(ctx context.Context, client *github.Client, owner, repo string) ([]*github.Issue, error) {
	var allIssues []*github.Issue

	var opt github.IssueListByRepoOptions
	for {
		fmt.Printf("fetching page %d\n", opt.Page)
		issues, resp, err := client.Issues.ListByRepo(ctx, owner, repo, &opt)
		if err != nil {
			return nil, err
		}
		for _, issue := range issues {
			allIssues = append(allIssues, issue)
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage

		// temporary while doing dev work
		if opt.Page > 2 {
			break
		}
	}
	return allIssues, nil
}

//type User struct {
//	Name     string `json:"name"`
//	Location string `json:"location"`
//}
//
//func projects(githubToken, owner string) {
//	client := gql.NewClient(
//		"https://api.github.com/graphql",
//		gql.Header("Authorization", "bearer "+githubToken),
//	)
//
//	var me User
//	meField := gql.Field("organization",
//
//		gql.Arg("login", owner),
//
//		gql.Field("name"),
//		gql.Field("location"),
//		gql.Dest(&me),
//	)
//	err := client.Query(meField)
//	fmt.Println(err, me)
//}

func getProjectID(githubToken, owner string, projectNumber int) string {
	req := graphql.NewRequest(
		`query($organization: String!, $projectNumber: Int!) {
			organization(login: $organization){
				projectV2(number: $projectNumber) {
					id
				}
			}
		}`,
	)

	req.Var("organization", owner)
	req.Var("projectNumber", projectNumber)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubToken))

	client := graphql.NewClient("https://api.github.com/graphql")

	var res map[string]map[string]map[string]string // the highway to hell is paved with good intentions and graphql
	err := client.Run(context.Background(), req, &res)
	if err != nil {
		die("failed %s\n", err)
	}
	projectID := res["organization"]["projectV2"]["id"] // TODO error handling

	fmt.Printf("projectID is %s\n", projectID)
	return projectID
}

func listProjectViews(githubToken, owner string, projectNumber int) {
	req := graphql.NewRequest(
		`query($organization: String!, $projectNumber: Int!) {
			organization(login: $organization){
				projectV2(number: $projectNumber) {
					id
					views(first: 100) {
						nodes {
							id
							name
							number
							databaseId
						}
					}
				}
			}
		}`,
	)

	req.Var("organization", owner)
	//req.Var("repo", repo)
	req.Var("projectNumber", projectNumber)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubToken))

	client := graphql.NewClient("https://api.github.com/graphql")

	var res any
	err := client.Run(context.Background(), req, &res)
	if err != nil {
		die("failed %s\n", err)
	}

	fmt.Printf("got is %v\n", res)

}

func getProjectsLinkedToIssue(githubToken, owner, repo string, issueNumber int) ([]string, error) {
	req := graphql.NewRequest(
		`query($owner: String!, $name: String!, $issue: Int!) {
			repository(owner: $owner, name: $name) {
				issue(number: $issue) {
					projectItems(first: 100) {
						nodes {
							id
							project {
								id
								title
							}
						}
					}
				}
			}
		}`,
	)

	req.Var("owner", owner)
	req.Var("name", repo)
	req.Var("issue", issueNumber)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubToken))

	client := graphql.NewClient("https://api.github.com/graphql")

	var res any
	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return nil, err
	}

	projects := []string{}
	return projects, nil
}

func main() {
	progName := "gh-issue-projector"
	if len(os.Args) > 0 {
		progName = os.Args[0]
	}

	if len(os.Args) != 4 {
		die("usage: %s <owner> <repo> <project-ID>")
	}
	owner := os.Args[1]
	repo := os.Args[2]
	projectNumber, err := strconv.Atoi(os.Args[3])
	if err != nil {
		die("failed to convert project ID: %s", err)
	}
	_ = projectID

	fmt.Printf("hello world from %s\n", progName)

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		die("environment variable GITHUB_TOKEN not set (or empty)")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	getProjectID(githubToken, owner, repo, projectNumber)

	getProjectsLinkedToIssue(githubToken, owner, repo, 1808)
	getProjectsLinkedToIssue(githubToken, owner, repo, 2867)
	die("done")

	issues, err := GetIssues(ctx, client, owner, repo)
	if err != nil {
		die("failed to get issues for %s/%s: %s", owner, repo, err)
	}
	for _, issue := range issues {
		if issue.ID == nil {
			die("issue ID is nil; full issue: %+v", issue)
		}
		if issue.Number == nil {
			die("issue Number is nil; full issue: %+v", issue)
		}
		fmt.Printf("ID: %d; issue number %d\n", *issue.ID, *issue.Number)
	}
}
