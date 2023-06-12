# xk6-g0-figure

Example addon for [xk6-g0](https://github.com/szkiba/xk6-g0)

The xk6-g0-figure addon is an example xk6-g0 addon that demonstrates how to add third-party go packages. The example includes the [go-figure](https://github.com/common-nighthawk/go-figure) package, which is of course not useful in k6 tests.

This is a template repository from which a new addon repository can be easily created. [Create a repository based on this template](https://github.com/szkiba/xk6-g0-figure/generate), then:
 - change the `module` line of the `go.mod` file to the name of the new module
 - change the `package` line of the `exports.go` file to the name of the new package
 - run `go mod tidy` to fix the go.mod file
 - modify the `//go generate` line in the `exports.go` file to use the appropriate third-party package
 - add the corresponding third-party package using the `go get` command
 - generate the source code using the `go generate .` command
 - delete `github_com-common-nighthawk-go-figure.go` file
 - run `go mod tidy` to fix the go.mod file

After that, the new addon repository is ready to use.

> In the rest of the page, of course, the name of the new repository should be used instead of `szkiba/xk6-g0-figure`

## Build

You can build the k6 binary on various platforms, each with its requirements. The following shows how to build k6 binary with this extension on GNU/Linux distributions.

### Prerequisites

You must have the latest Go version installed to build the k6 binary. The latest version should match [k6](https://github.com/grafana/k6#build-from-source) and [xk6](https://github.com/grafana/xk6#requirements).

- [Git](https://git-scm.com/) for cloning the project
- [xk6](https://github.com/grafana/xk6) for building k6 binary with extensions

### Install and build the latest tagged version

1. Install `xk6`:

   ```shell
   go install go.k6.io/xk6/cmd/xk6@latest
   ```

2. Build the binary:

   ```shell
   xk6 build --with github.com/szkiba/xk6-g0-figure@latest
   ```

### Build for development

If you want to add a feature or make a fix, clone the project and build it using the following commands. The xk6 will force the build to use the local clone instead of fetching the latest version from the repository. This process enables you to update the code and test it locally.

```bash
git clone git@github.com:szkiba/xk6-g0.git && cd xk6-g0-figure
xk6 build --with github.com/szkiba/xk6-g0-figure@latest=.
```

## Run

There is an example in the [example](https://github.com/szkiba/xk6-g0-figure/tree/master/example) directory that show how to use the extension.

**command**
```bash
./k6 run example/script.go
```

**script.go**
```go
package main

import "github.com/common-nighthawk/go-figure"

func Default() {
  figure.NewFigure("Hello World!", "", true).Print()
}
```

**output**
```plain
          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: -
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

  _   _          _   _            __        __                 _       _   _
 | | | |   ___  | | | |   ___     \ \      / /   ___    _ __  | |   __| | | |
 | |_| |  / _ \ | | | |  / _ \     \ \ /\ / /   / _ \  | '__| | |  / _` | | |
 |  _  | |  __/ | | | | | (_) |     \ V  V /   | (_) | | |    | | | (_| | |_|
 |_| |_|  \___| |_| |_|  \___/       \_/\_/     \___/  |_|    |_|  \__,_| (_)

     data_received........: 0 B 0 B/s
     data_sent............: 0 B 0 B/s
     iteration_duration...: avg=1.46ms min=1.46ms med=1.46ms max=1.46ms p(90)=1.46ms p(95)=1.46ms
     iterations...........: 1   613.200866/s


running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU
```
