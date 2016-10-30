build:
	cd qd && go test && go build

compute:
	cat points.json | qd/qd > data

plot:
	cat data | ./proc/plot.py

run: build compute plot

install:
	pip install -r requirements.txt
	go get -u github.com/paulmach/go.geo
	./proc/genpoints.py > points.json
