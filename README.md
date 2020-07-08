# gsc  ![Release](https://github.com/mskutin/gsc/workflows/Release/badge.svg) ![Linter](https://github.com/mskutin/gsc/workflows/Lint%20Code%20Base/badge.svg)
Github Stats Collector is a tiny cli that helps to collect stats for public repositories.

<img width="120" alt="Screenshot 2020-07-07 at 03 41 29" src="https://user-images.githubusercontent.com/11622907/86633978-eb19e880-c003-11ea-8bc8-fa4d6d797abb.png">

## Usage

```shell script
gsc get [flags]

Flags:
  -f, --format string   --format tsv (default "csv")
  -h, --help            help for get
  -r, --repos strings   Comma separated list of repositories, e.g. 'helm/charts,mskutin/gsc'
```
### Stats

#### get statistics for one repository:

`gsc get -r mskutin/gsc`

```csv
name,clone_url,last_commit_author,last_commit_date
mskutin/gsc,https://github.com/mskutin/gsc.git,Maksim Skutin,2020-07-07 05:43:31 +0000 UTC
```

#### get statistics for multiple repositories:

```shell script
1) gsc get --repos mskutin/gsc,helm/charts
2) gsc get \
    -r mskutin/gsc \
    -r helm/charts \
    -r github/hubot
3) echo "helm/charts,mskutin/gsc" | xargs gsc get -r
```

```csv
name,clone_url,last_commit_author,last_commit_date
mskutin/gsc,https://github.com/mskutin/gsc.git,Maksim Skutin,2020-07-07 05:43:31 +0000 UTC
helm/charts,https://github.com/helm/charts.git,Maxime Brunet,2020-07-07 11:51:58 +0000 UTC
```

### Formatters

```shell script
gsc get [flags]

Flags:
  -f, --format string   --format tsv (default "csv")
```

By default, gsc uses csv format output.  
The following types of output are supported:
- [x] csv
- [x] tsv
- [ ] table