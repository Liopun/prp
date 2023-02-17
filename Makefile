.PHONY:
.SILENT:
.DEFAULT_GOAL := lint

gdraft:
	git add .
	git commit -m "${msg}"

git: gdraft
	git push

release:
	git tag -a "${ver}" -m "${ver}"
	git push origin "${ver}"

lint:
	golangci-lint run

build:
	go env -w CGO_ENABLED=0
	go build -o ./.dist/prp -ldflags="-X 'github.com/liopun/prp/cmd/prp.version=${ver}'" main.go