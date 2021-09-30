default: generate-results generate-graph

create-dirs:
	@mkdir -p out 
generate-results: create-dirs
	@go run . > out/results.csv
generate-graph: out/results.csv
	@gnuplot plot.gnu > out/graph.pdf
generate-docker: create-dirs
	@sudo docker build -t go-sorting-benchmark .
	@sudo docker run -v ${PWD}/out:/app/out:z go-sorting-benchmark