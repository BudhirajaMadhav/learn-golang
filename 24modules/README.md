[Blog](https://ckakos.medium.com/go-modules-485510e3737f#8f8a)

## Initialising a module
`go mod init github.com/budhirajamadhav/praujact`

### Importing a Dependency
You could run `go get package_name` from your command line to import a dependency, however I chose to be more explicit in my example and added it to my main.go file.

## Go mod tidy
This will cache the package being imported — and all of its dependencies — into your GOPATH. If the package has already been cached, it skips that process. If the package has been deleted and/or no longer being used, it will be removed.

## go.sum
In addition to go.mod, the go command maintains *a file named go.sum containing **cryptographic hashes** which represents the contents of a specific version of the package*.
The go.sum file to ensure that future downloads of these modules retrieve the same bits as the first download and that the modules — which your project depends on — do not change unexpectedly, whether for malicious or accidentally. Both go.mod and go.sum should be checked into version control.

## Vendoring dependencies
Vendoring is a process which downloads a local copy of all dependencies into your repository. This is most useful whenever you have an absorbent amount of dependencies that need to be imported without having to continually pull dependencies upon every build. You can vendor a package by running the command *go mod vender*