### benchmark and profile commands
## go test -bench . -benchtime=10x -run ^$
## go test -bench . -benchtime=10x -run ^$ | tee benchresults00.txt
## go test -bench . -benchtime=10x -run ^$ -cpuprofile cpu00.pprof
## go tool pprof cpu00.pprof
## go test -bench . -benchtime=10x -run ^$ -memprofile mem00.pprof
## go tool pprof -alloc_space mem00.pprof
## benchstat file1.txt file2.txt
