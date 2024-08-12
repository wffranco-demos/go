#!/bin/bash

# ejecutar tests:
go test

# ejecutar code coverage:
go test -cover

# ejecutar code coverage y exportar a un archivo:
go test -coverprofile=coverage.out

# abrir el coverage en el navegador:
go tool cover -html=coverage.out

