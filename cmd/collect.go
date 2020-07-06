package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"

	gh "github.com/mskutin/gsc/internal/github"
	"github.com/spf13/cobra"
)

var repos []string

type Repository struct {
	name             string
	cloneURL         string
	lastCommitDate   string
	lastCommitAuthor string
	defaultBranch    string
}

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Collect statistics for one or more repositories",
	Long: `
Collect statistics for a given repository:
	gsc -r mskutin/gsc

Collect statistics for a set of repositories:
	gsc --repos mskutin/gsc,helm/charts
`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Validation
		//TODO: Distinguish CSV & TSV
		var repositories []Repository
		for _, repo := range repos {
			params := strings.Split(repo, "/")

			head, err := gh.GetHead(params[0], params[1])
			if err != nil {
				errors.Wrap(err, "Unable to retrieve head.")
			}
			details, err := gh.GetRepository(params[0], params[1])
			if err != nil {
				errors.Wrap(err, "Unable to retrieve repository details.")
			}
			repositories = append(repositories, Repository{
				name:             details.FullName,
				cloneURL:         details.CloneURL,
				lastCommitAuthor: head.Commit.Author.Name,
				lastCommitDate:   head.Commit.Author.Date.UTC().String(),
				defaultBranch:    details.DefaultBranch,
			})
		}
		PrintCSV(repositories)
	},
}

func PrintCSV(repos []Repository) {
	records := [][]string{{"name", "clone_url", "last_commit_author", "last_commit_date"}}
	for _, repo := range repos {
		row := []string{repo.name, repo.cloneURL, repo.lastCommitAuthor, repo.lastCommitDate}
		records = append(records, row)
	}
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
func init() {
	rootCmd.AddCommand(collectCmd)
	collectCmd.Flags().
		StringSliceVarP(
			&repos,
			"repos",
			"r",
			[]string{},
			"Comma separated list of repositories, e.g. 'helm/charts,mskutin/gsc'")
	collectCmd.MarkFlagRequired("repos")
}
