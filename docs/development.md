# Develop Environment

We develop inside a container by using [VS Code Remote Container](https://code.visualstudio.com/docs/remote/containers).

## Prerequirements

- [VSCode](https://github.com/microsoft/vscode) Version: 1.63.2, or later.
- [VSCode Remote Container](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) Version: Depends on VSCode. Latest version recommended
- [Docker desktop](https://www.docker.com/products/docker-desktop): 2.3.0.4, or later.

## How to use VSCode.

1. Startup VSCode.
2. Open Command Palette(Menu [View]->[Command Palette])
3. Select `Remote-COntainers: Open Folder in Container`
4. Select cloned repository directory.

# Edit Code

Ensure that no error and no warning.

```bash
go build . 
go fmt ./...  
golangci-lint run 
go test -v ./...
```

# Testing

#### e2e Test

before running,you should mount passenger directory at `/sock` directory.

```
root@842ff56c6f6f:/workspace# go build .
root@842ff56c6f6f:/workspace# E2E=true go test ./test/e2e/
ok      github.com/rakutentech/passenger-go-exporter/test/e2e   0.151s
root@842ff56c6f6f:/workspace# 
```

#### All Test Caverage

```
root@842ff56c6f6f:/workspace# go test -coverprofile=cover.out -cover ./...  && go tool cover -html=cover.out -o cover.html
ok      github.com/rakutentech/passenger-go-exporter    0.039s  coverage: 90.5% of statements
ok      github.com/rakutentech/passenger-go-exporter/logging    0.032s  coverage: 100.0% of statements
ok      github.com/rakutentech/passenger-go-exporter/metric     0.010s  coverage: 100.0% of statements
ok      github.com/rakutentech/passenger-go-exporter/passenger  0.011s  coverage: 67.7% of statements
ok      github.com/rakutentech/passenger-go-exporter/test/e2e   0.007s  coverage: [no statements]
root@842ff56c6f6f:/workspace# 
```

If passenger application started and mounted /sock directory,please set USE_PASSENGER variable.

```
root@842ff56c6f6f:/workspace# USE_PASSENGER=true go test -coverprofile=cover.out -cover ./...  && go tool cover -html=cover.out -o cover.html
ok      github.com/rakutentech/passenger-go-exporter    0.047s  coverage: 90.5% of statements
ok      github.com/rakutentech/passenger-go-exporter/logging    0.028s  coverage: 100.0% of statements
ok      github.com/rakutentech/passenger-go-exporter/metric     0.022s  coverage: 100.0% of statements
ok      github.com/rakutentech/passenger-go-exporter/passenger  0.017s  coverage: 95.2% of statements
ok      github.com/rakutentech/passenger-go-exporter/test/e2e   0.007s  coverage: [no statements]
root@842ff56c6f6f:/workspace# 
```

#### HTML Report

```bash
go test -coverprofile=cover.out -cover ./... \
 && go tool cover -html=cover.out -o cover.html
```

