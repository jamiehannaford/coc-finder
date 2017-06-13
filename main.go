package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Printf("GitHub auth token: ")
	byteAuthToken, _ := terminal.ReadPassword(int(syscall.Stdin))
	authToken := string(byteAuthToken)

	ctx := context.Background()

	var client *github.Client
	if authToken == "" {
		client = github.NewClient(nil)
	} else {
		tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: authToken},
		))
		client = github.NewClient(tc)
	}

	fmt.Printf("\nGitHub org: ")
	orgName, _ := r.ReadString('\n')
	orgName = strings.TrimSpace(orgName)

	if authToken == "" {
		fmt.Println("No OAuth token provided. You might receive rate limits")
	}

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(ctx, orgName, opt)
	if err != nil {
		log.Fatalf("Error listing repos: %v", err)
	}

	fmt.Printf("\nScanning %s repos!\n\n", orgName)

	for _, repo := range repos {
		coc, _, err := client.Repositories.GetCodeOfConduct(ctx, orgName, *repo.Name)
		if err != nil {
			log.Fatalf("Error returning COC for %s/%s: %v", orgName, *repo.Name, err)
		} else if coc.URL == nil {
			fmt.Printf("github.com/%s/%s has no code of conduct!\n", orgName, *repo.Name)
		}
	}
}
