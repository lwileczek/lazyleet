go test -bench . 
go test -bench=searchInsert

go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
go test -cpuprofile cpu.prof -memprofile mem.prof -bench=searchInsert

go tool pprof cpu.prof
go tool pprof mem.prof
