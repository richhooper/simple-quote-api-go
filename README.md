# Simple Quote API created in Go

I am learning the Go language and how to create things using Go.

The purpose of this repo is so I can compare what is required to create a functional API and then compare to a .net minial api, which I already have experience with.

Target is to build something that can:

- integrate with other components (e.g. a data store)
- have some unit tests
- built in a way that could sustainably be grown
- have a build and deploy process
- test out linting tools and cli

## Call examples

    curl  -X POST \
    'http://localhost:8080/quote' \
    --header 'Content-Type: application/json' \
    --data-raw '{"age": 30, "coverStartDate": "2023-01-01T00:00:00Z"}'
