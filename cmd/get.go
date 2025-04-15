package cmd

import (
	"encoding/json"
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
	stars            int
	forks            int
	openIssues       int
	description      string
	language         string
	license          string
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get statistics for one or more repositories",
	Long: `
get statistics for one repository:
	1) gsc get -r mskutin/gsc
	2) gsc get --repos=mskutin/gsc

get statistics for multiple repositories:
	1) gsc get -f tsv \
		-r mskutin/gsc \
		-r mskutin/nginx-fluentd \
		-r helm/charts
	2) gsc get -r=helm/charts,mskutin/gsc
	3) echo "helm/charts,mskutin/gsc" | xargs gsc get -r
`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Validation
		var token, tokenIsPresent = os.LookupEnv("GITHUB_TOKEN")
		var username, usernameIsPresent = os.LookupEnv("GITHUB_USERNAME")
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
			log.Println("gsc: Github Authorization is not implemented yet. Unset GITHUB_TOKEN and GITHUB_USERNAME env variables.")
		default:
			client, err := github.New()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			minStars, _ := cmd.Flags().GetInt("min-stars")
			minForks, _ := cmd.Flags().GetInt("min-forks")
			minOpenIssues, _ := cmd.Flags().GetInt("min-open-issues")

			filteredStats := filterStats(getStats(client), minStars, minForks, minOpenIssues)

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
		stats := Stats{
			name:             details.FullName,
			cloneURL:         details.CloneURL,
			lastCommitAuthor: head.Commit.Author.Name,
			lastCommitDate:   head.Commit.Author.Date.UTC().String(),
			defaultBranch:    details.DefaultBranch,
			stars:            details.StargazersCount,
			forks:            details.ForksCount,
			openIssues:       details.OpenIssuesCount,
			description:      details.Description,
			language:         details.Language,
		}
		if details.License != nil {
			stats.license = details.License.SPDXID
		}
		repositories = append(repositories, stats)
	}
	return repositories
}

func filterStats(stats []Stats, minStars, minForks, minOpenIssues int) []Stats {
	var filtered []Stats
	for _, s := range stats {
		if s.stars >= minStars && s.forks >= minForks && s.openIssues >= minOpenIssues {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func printStats(repos []Stats, format string) {
	switch format {
	case "tsv":
		printCSV(repos, '\t')
	case "json":
		printJSON(repos)
	default: //csv
		printCSV(repos, ',')
	}
}

func printCSV(repos []Stats, separator rune) {
	records := [][]string{{"name", "clone_url", "last_commit_author", "last_commit_date", "stars", "forks", "open_issues", "description", "language", "license"}}
	for _, repo := range repos {
		row := []string{
			repo.name, repo.cloneURL, repo.lastCommitAuthor, repo.lastCommitDate,
			string(repo.stars), string(repo.forks), string(repo.openIssues), repo.description, repo.language, repo.license,
		}
		// Convert int fields to strings for CSV output
		row = convertIntToString(row, []int{4, 5, 6})
		records = append(records, row)

	}
	w := csv.NewWriter(os.Stdout)
	w.Comma = separator
	err := w.WriteAll(records)
	if err != nil {
		log.Fatalln("error writing csv", err)
	}
}

func printJSON(repos []Stats) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ") // Optional: for pretty printing
	if err := enc.Encode(repos); err != nil {
		log.Fatalf("Error encoding JSON: %s\n", err)
	}

}

func convertIntToString(row []string, indices []int) []string {
	for _, i := range indices {
		if i >= 0 && i < len(row) {
			row[i] = convertString(row[i])
		}
	}
	return row
}
func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(
		&format,
		"format",
		"f",
		"csv", // Default format is CSV
		"Output format. Choose from: csv, tsv, json")
	getCmd.Flags().
		StringSliceVarP(
			&repos,
			"repos",
			"r",
			[]string{},
			`One or more repositories: 'gsc get -r mskutin/gsc'
See help for more examples.`)
	getCmd.Flags().IntP(
		"min-stars",
		"",
		0,
		"Minimum number of stars a repository must have")
	getCmd.Flags().IntP(
		"min-forks",
		"",
		0,
		"Minimum number of forks a repository must have")
	getCmd.Flags().IntP(
		"min-open-issues",
		"",
		0,
		"Minimum number of open issues a repository must have")
	err := getCmd.MarkFlagRequired("repos")
	if err != nil {
		log.Fatalln("MarkFlagRequired is not set", err)
	}
}
