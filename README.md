# gsc  ![Release](https://github.com/mskutin/gsc/workflows/Release/badge.svg) ![Linter](https://github.com/mskutin/gsc/workflows/Lint%20Code%20Base/badge.svg)
Github Stats Collector is a tiny cli that helps to collect stats for public repositories.

<img width="120" alt="Screenshot 2020-07-07 at 03 41 29" src="https://user-images.githubusercontent.com/11622907/86633978-eb19e880-c003-11ea-8bc8-fa4d6d797abb.png">

## Install

### Use ready binaries
Download [latest binaries](https://github.com/mskutin/gsc/releases/tag/v0.2-alpha) for your platform.

##### Mac OS X

1. `wget -c https://github.com/mskutin/gsc/releases/download/v0.2-alpha/gsc-darwin-amd64.tar.gz -O - | tar -xz`
2. `mv gsc /usr/local/bin/gsc`

### Build from source

1. `git clone https://github.com/mskutin/gsc`
2. `go build`
3. `./gsc help`

## Usage

Repository stats can be retrieved by calling a `get` command:

```shell script
Usage:
  gsc get [flags]

Flags:
  -f, --format string   --format tsv (default "csv")
  -h, --help            help for get
  -r, --repos strings   One or more repositories: 'gsc get -r mskutin/gsc'
                        See help for more examples.
```
### Authorization

By default `gsc` interacts with github api anonymously. 
Since GitHub's default request rate is only 60 req/h you may run out of this limit shortly.
You can set `GITHUB_TOKEN` and `GITHUB_USERNAME` environment variables to authorize gsc in GitHub. 

### Stats

#### get statistics for one repository:

```shell script
1) gsc get -r mskutin/gsc
2) gsc get --repos=mskutin/gsc
```

```csv
name,clone_url,last_commit_author,last_commit_date
mskutin/gsc,https://github.com/mskutin/gsc.git,Maksim Skutin,2020-07-07 05:43:31 +0000 UTC
```

#### get statistics for multiple repositories:

```shell script
1) gsc get -f tsv \
    -r mskutin/gsc \
    -r mskutin/nginx-fluentd \
    -r helm/charts
2) gsc get -r=helm/charts,mskutin/gsc
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

```shell script
./gsc get -f tsv \
    -r mskutin/gsc \
    -r helm/charts \
    -r github/hubot 
```

```shell script
name	clone_url	last_commit_author	last_commit_date
mskutin/gsc	https://github.com/mskutin/gsc.git	Maksim Skutin	2020-07-10 20:10:52 +0000 UTC
helm/charts	https://github.com/helm/charts.git	jabdoa2	2020-07-11 17:39:23 +0000 UTC
hubotio/hubot	https://github.com/hubotio/hubot.git	Misty De Meo	2019-04-17 04:49:47 +0000 UTC
```

By default, gsc uses csv format output.  
The following types of output are supported:
- [x] csv
- [x] tsv
- [ ] table