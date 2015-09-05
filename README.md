# Charlie

A dead simple CI for projects built with Gradle.

Charlie runs as web service that accepts [Github's `PushEvent` webhook](https://developer.github.com/v3/activity/events/types/#pushevent). It will clone the repo, build the project, and return a 200 response if the build was succesful.

## Todos
Charlie is still a work in progress.
- [ ] Checkout/clone a specific commit
- [ ] Clone private repos
- [ ] Queue builds instead of blocking incoming requests
- [ ] Post build results to Github

## Installation
`go get github.com/f2prateek/charlie`

## Usage
```
charlie [--port=<port>]
charlie -h | --help
charlie --version

Options:
  --port=<port>    port [default: 3001].
  -h --help        Show this screen.
  --version        Show version.
```