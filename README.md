# howbob

Generate a Brewfile from a manifest.k [kcl](https://www.kcl-lang.io) file.

Motivation: [victor](https://www.youtube.com/watch?v=Gn6btuH3ULw) prompted me to want to learn what kcl offers compared to using yaml.

```bash
howbob run --path=homebrew.k --output=Brewfile
```

For example:

```
# file: homebrew.k

{packages = [
    {
        name = "golang"
        version_check = "go version"
    }
    {
        name = "awscli@2"
        version_check = "aws --version"
    }
```

```
# file: Brewfile
brew "golang"
brew "awscli@2"
```
