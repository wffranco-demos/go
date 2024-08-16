#!/bin/bash

# ejecutar tests:
go test

# ejecutar profile:
go test -cpuprofile=cpu.out

# ejecutar profile y abrir en el navegador:
go tool pprof cpu.out
