.PHONY:
.SILENT:
.DEFAULT_GOAL := lint

gdraft:
	git add .
	git commit -m "${msg}"

git:
	git add .
	git commit -m "${msg}"
	git push

release:
	git tag -a "${ver}" -m "${ver}"
	git push origin "${ver}"

lint:
	golangci-lint run

build:
	go build -o ./.dist/prp -ldflags="-X 'github.com/liopun/prp/cmd/prp.version=${ver}'" main.go