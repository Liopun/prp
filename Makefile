go run main.go --version

go build -o ./.dist/prp -ldflags="-X 'github.com/liopun/prp/cmd/prp.version=1.0.0'" main.go