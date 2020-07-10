package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/mskutin/gsc/pkg/github"
	"github.com/spf13/cobra"
)

var repos []string
var format string

type Stats struct {
	name             string
	cloneURL         string
	lastCommitDate   string
	lastCommitAuthor string
	defaultBranch    string
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get statistics for one or more repositories",
	Long: `
get statistics for one repository:
	gsc get -r mskutin/gsc

get statistics for multiple repositories:
	1) gsc get --repos mskutin/gsc,helm/charts
	2) gsc get \
		-r mskutin/gsc \
		-r helm/charts \
		-r github/hubot
	3) echo "helm/charts,mskutin/gsc" | xargs gsc get -r

`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Validation
		var token, tokenIsPresent = os.LookupEnv("GITHUB_TOKEN")
		var username, usernameIsPresent = os.LookupEnv("GITHUB_USERNAME")
		log.Println(username, token)
		switch {
		case tokenIsPresent && usernameIsPresent:
			//TODO: Authorization
			client, err := github.NewWithAuth(username, token)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			if !client.IsTokenValid() {
				log.Println("Username or token is invalid")
			}
			log.Println("gsc: Github Authorization is not implemented yet. Unset GITHUB_TOKEN env variable.")
		default:
			client, err := github.New()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			stats := getStats(client)
			printStats(stats, format)
		}
	},
}

func getStats(github *github.Client) []Stats {
	var repositories []Stats

	for i := 0; i < len(repos); i++ {
		repo := repos[i]
		params := strings.Split(repo, "/")
		head, err := github.GetHead(params[0], params[1])
		if err != nil {
			log.Println(err, repo)
			continue
		}
		details, err := github.GetRepository(params[0], params[1])
		if err != nil {
			log.Println(err, repo)
			continue
		}
		repositories = append(repositories, Stats{
			name:             details.FullName,
			cloneURL:         details.CloneURL,
			lastCommitAuthor: head.Commit.Author.Name,
			lastCommitDate:   head.Commit.Author.Date.UTC().String(),
			defaultBranch:    details.DefaultBranch,
		})
	}
	return repositories
}
func printStats(repos []Stats, format string) {
	var separator rune
	switch format {
	case "tsv":
		separator = '\t'
	default:
		separator = ','
	}
	records := [][]string{{"name", "clone_url", "last_commit_author", "last_commit_date"}}
	for _, repo := range repos {
		row := []string{repo.name, repo.cloneURL, repo.lastCommitAuthor, repo.lastCommitDate}
		records = append(records, row)
	}
	w := csv.NewWriter(os.Stdout)
	w.Comma = separator
	err := w.WriteAll(records)
	if err != nil {
		log.Fatalln("error writing csv", err)
	}
}
func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(
		&format,
		"format",
		"f",
		"csv",
		"--format tsv")
	getCmd.Flags().
		StringSliceVarP(
			&repos,
			"repos",
			"r",
			[]string{},
			"Comma separated list of repositories, e.g. 'helm/charts,mskutin/gsc'")
	err := getCmd.MarkFlagRequired("repos")
	if err != nil {
		log.Fatalln("MarkFlagRequired is not set", err)

	}
}
