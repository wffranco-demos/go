#!/bin/bash

# detectar si hay riesgo de carrera en el código (no compatible con go 1.22)
go build --race main.go

# Verificar condicion de carrera para go 1.22+
go build -race main.go
