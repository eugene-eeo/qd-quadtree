build:
	cd qd && go test && go build

run: build
	cat points.json | qd/qd > data
	cat data | ./plot.py

install:
	pip install -R requirements.txt
	./genpoints.py > points.json
