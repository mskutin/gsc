package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	gh "github.com/mskutin/gsc/internal/github"
	"github.com/spf13/cobra"
)

var repos []string
var format string

type Repository struct {
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
		//TODO: Distinguish CSV & TSV
		var separator rune
		stats := getStats()
		switch format {
		case "tsv":
			separator = '\t'
		default:
			separator = ','
		}
		PrintCSV(stats, separator)
	},
}

func getStats() []Repository {
	var repositories []Repository
	for _, repo := range repos {
		params := strings.Split(repo, "/")

		head, err := gh.GetHead(params[0], params[1])
		if err != nil {
			log.Fatalln("Unable to retrieve head.", err)
		}
		details, err := gh.GetRepository(params[0], params[1])
		if err != nil {
			log.Fatalln("Unable to retrieve repository details.", err)
		}
		repositories = append(repositories, Repository{
			name:             details.FullName,
			cloneURL:         details.CloneURL,
			lastCommitAuthor: head.Commit.Author.Name,
			lastCommitDate:   head.Commit.Author.Date.UTC().String(),
			defaultBranch:    details.DefaultBranch,
		})
	}
	return repositories
}

func PrintCSV(repos []Repository, comma rune) {
	records := [][]string{{"name", "clone_url", "last_commit_author", "last_commit_date"}}
	for _, repo := range repos {
		row := []string{repo.name, repo.cloneURL, repo.lastCommitAuthor, repo.lastCommitDate}
		records = append(records, row)
	}
	w := csv.NewWriter(os.Stdout)
	w.Comma = comma
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
