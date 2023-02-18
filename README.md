# Reddit Sync

A command-line tool to interact with Reddit's API and retrieve Posts and Comments data.

## Dependencies

- [Go](https://go.dev/) `v1.20.1`
  - [Cobra](https://github.com/spf13/cobra) `v1.6.1`

## Build

### Linux x64

```bash
env GOOS=linux GOARCH=amd64 go build
```

## Usage

After building the binary, execute the following command to confirm the compiled binary executes successfully and to get information on the commands available.

```bash
./reddit-sync --help
```
