* howbob

Generate a Brewfile from a manifest.k kcl file.

```bash
howbob run --path homebrew.k --output=Brewfile
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