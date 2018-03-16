docker:
	docker build -t navikt/muckscan .

truffletool: truffletool.go
	go build truffletool.go
