# gsc  ![Release](https://github.com/mskutin/gsc/workflows/Release/badge.svg) ![Linter](https://github.com/mskutin/gsc/workflows/Lint%20Code%20Base/badge.svg)
Github Stats Collector is a tiny cli that helps to collect stats for public repositories.

<img width="120" alt="Screenshot 2020-07-07 at 03 41 29" src="https://user-images.githubusercontent.com/11622907/86633978-eb19e880-c003-11ea-8bc8-fa4d6d797abb.png">

```bash
Collect statistics for a given repository:
  gsc -r mskutin/gsc

Collect statistics for a set of repositories:
	gsc --repos mskutin/gsc,helm/charts

Usage:
  gsc collect [flags]

Flags:
  -h, --help            help for collect
  -r, --repos strings   Comma separated list of repositories, e.g. 'helm/charts,mskutin/gsc'
  ```
```bash
gsc collect -r mskutin/gsc
```

```bash
name,clone_url,last_commit_author,last_commit_date
mskutin/gsc,https://github.com/mskutin/gsc.git,Maksim Skutin,2020-07-06 19:46:18 +0000 UTC
```
