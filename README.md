# ghrelease

A tool to download latest release archive for a GitHub project using GitHub API.

### Usage

```
http://ghrelease.abiosoft.com/<username>/<repository>/<filename>
```
##### Example
```
http://ghrelease.abiosoft.com/mholt/caddy/caddy_linux_amd64.zip
```
This will download the latest version of `caddy_linux_amd64.zip` for the project `github.com/mholt/caddy`.

### Building from source

Go is a prerequisite. Download it [here](http://golang.org/doc/install.html)

```shell
$ go get github.com/abiosoft/ghrelease
```
Start
```
$ ghrelease
2015/06/05 02:23:50 Waiting for requests on 8888
```
