# Code of Conduct finder

This tool allows you to see which repositories in a GitHub org are missing a
Code of Conduct.

## How to use

First, download the package and build it:

```bash
go get github.com/jamiehannaford/coc-finder
go build
```

Then run it:

```bash
› ./coc-finder
GitHub auth token:
GitHub org:
````

You will be prompted for an auth token and the org you wish to search.

Although providing a token is optional, it's recommended to avoid rate limiting.
For more information on how to generate auth tokens, see [GitHub's docs](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/).

## Recommended COCs

If your project is missing one, consider using the [Contributor Covenant](http://contributor-covenant.org/).
