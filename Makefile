all:
	docker build -t ambientsound/muckscan .

parsetool:
	go build parsetool.go
