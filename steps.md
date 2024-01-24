1. Project directories

```sh
mkdir project name
cd project name
code .
```

2. Install package(gin, gorm(db, driver), daemon,)

- Iniatilize package

```sh
#initialize package first
go mod init example.com/projectName
```

- Install compileDameon for automatic rebuild

```sh
go get github.com/githubnemo/CompileDaemon

go install github.com/githubnemo/CompileDaemon
```
