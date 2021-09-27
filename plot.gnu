set terminal pdf
set datafile separator ','
set key outside
set ylabel "Tempo (segundos)"
set xlabel "Tamanho da entrada (n)" 
plot for [col=2:7] 'out/results.csv' using 1:col with lines title columnhead