default:
generate-results:
	@mkdir -p out && go run . > out/results.csv
generate-graph: out/results.csv
	@gnuplot plot.gnu > out/graph.pdf