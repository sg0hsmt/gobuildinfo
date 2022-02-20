# Go Build Info

Print build info from binary, using buildinfo package.

https://pkg.go.dev/debug/buildinfo@go1.18rc1

Note: This was created to help me learn about go 1.18.

Japanese article:
https://qiita.com/sg0hsmt/items/6d852c50baa37a0c957e

## Requirements

- Go 1.18-rc.1 or later

## Output Examples

```console
$ go run main.go -path testdata/hello.go117.exe
go      go1.17.5
path    github.com/sg0hsmt/try-goapp/hello
mod     github.com/sg0hsmt/try-goapp    (devel)
dep     github.com/google/uuid  v1.3.0  h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=
dep     github.com/maxence-charriere/go-app/v9  v9.2.1  h1:wP+A7zMDc1DIxXvhZQvWVq+vzlb8Hmo75AEVMChoVRg=
```

```console
$ go run main.go -path testdata/hello.go118.exe
go      go1.18beta1
path    github.com/sg0hsmt/try-goapp/hello
mod     github.com/sg0hsmt/try-goapp    (devel)
dep     github.com/google/uuid  v1.3.0  h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=
dep     github.com/maxence-charriere/go-app/v9  v9.2.1  h1:wP+A7zMDc1DIxXvhZQvWVq+vzlb8Hmo75AEVMChoVRg=
build   -compiler=gc
build   CGO_ENABLED=1
build   CGO_CFLAGS=
build   CGO_CPPFLAGS=
build   CGO_CXXFLAGS=
build   CGO_LDFLAGS=
build   GOARCH=amd64
build   GOOS=windows
build   GOAMD64=v1
build   vcs=git
build   vcs.revision=619f4a18485903816d91d53a6d1618d5024d47d1
build   vcs.time=2021-12-19T07:51:57Z
build   vcs.modified=true
```

```console
$ go run main.go
go      go1.18rc1
path    command-line-arguments
build   -compiler=gc
build   CGO_ENABLED=1
build   CGO_CFLAGS=
build   CGO_CPPFLAGS=
build   CGO_CXXFLAGS=
build   CGO_LDFLAGS=
build   GOARCH=amd64
build   GOOS=windows
build   GOAMD64=v1
```

## Develop Setting (VSCode)

Download and install go1.18rc1.

https://pkg.go.dev/golang.org/dl/go1.18rc1

```console
$ go install golang.org/dl/go1.18rc1@latest
$ go1.18rc1 download
```

Change vscode settings.

https://github.com/golang/vscode-go/blob/master/docs/settings.md#goalternatetools

```json
{
  "go.alternateTools": {
    "go": "go1.18rc1"
  }
}
```
