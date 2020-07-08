### goCYOA

A simple Golang implementation of a text-based adventure game:

* in which a player is able to choose alternative paths
* that renders to a webserver base UI

***

The code leverages the following packages:

* [encoding/json](https://golang.org/pkg/encoding/json/)
* [json-to-Go](https://mholt.github.io/json-to-go/)
* `flags`
* `net/http`
8 `io`
* `fmt`
* `strings`
* `log`
* `error`
* `reflect`
* `errors`
* `os`
* `html/template`
* `bytes`
* `css`

***

#### Example
To run the code nominally, first go to `/cmd` within the root folder [on your terminal]. Then start the local webserver by running:
```bash
    $ ./cmd -switch=true
```
Then got to either `http://127.0.0.1:8085/story` or `http://127.0.0.1:8085` depending on if customPathFunction is used [or otherwise] respectively. The webserver listening port can be changed via using the flag `-port=<port-integer-value>`.
```bash
    $ ./cmd -switch=true -port=3031
```
e.g. running `./cmd -switch=true -port=3031` allows you to access the same game at `http://127.0.0.1:3031/story`.