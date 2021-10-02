package main

import (
	Domain "server/domain"
	Lib "server/lib"
)

func main() {
	Lib.Start(Domain.ROUTES)
}
