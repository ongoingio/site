# Site

## Run

	go run site.go

	// http://localhost:8080/

### Auto-Restart

	go get github.com/codegangsta/gin
	gin --port 8000 --appPort 8080

	// http://localhost:8080/


## Example Ideas

* Layout/parent template
* Decode JSON starting with an array, when the number of objects is unknown.
* Models in Go
* Unmarshal a JSON array or unknown length into a slice.
* Database connection best practice.
* Testing an external API (like Github).