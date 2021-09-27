# go-sorting-benckmark

## Visão geral

Benchmark de funções de ordenação utilizando a linguagem Go (golang)

As funções testadas são: BubbleSort, SelectionSort, ShellSort, MergeSort, QuickSort e RadixSort.

## Utilização

### Dependências

- [Go](https://golang.org/doc/install) (1.16)
- [gnuplot](http://gnuplot.sourceforge.net/download.html) (5.2)
- [GNU Make](https://www.gnu.org/software/make/) (4.3)

### Uso 

```bash
# Executa o programa e gera um arquivo csv (out/results.csv) com os resultados
make generate-results 

# Gera um arquivo PDF (out/graph.pdf) contendo o gráfico do tempo de execução 
# das funções a partir do arquivo csv gerado no passo anterior
make generate-graph
```
