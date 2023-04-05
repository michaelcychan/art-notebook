# Project Setup

## Go Fiber

First initialise Go module:  

```bash
go mod init <project name on github>
```

Get Go Fiber:  

```bash
go get -u github.com/gofiber/fiber/v2
```

### Hot reload while developing

Get Go package named Air.  

```bash
go get github.com/cosmtrek/air
```

If Air is installed, we can start dev env with:  

```bash
go run github.com/cosmtrek/air
```

## Using .ENV

[Use Environment Variable in your next Golang Project](https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66)  
The above link shows how to use `godotenv` and `os` packages to read configuration from `.env` file.  

### Set up

Run the following command:  

```bash
go get github.com/joho/godotenv
```
