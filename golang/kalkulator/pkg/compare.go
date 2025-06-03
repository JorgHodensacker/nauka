package pkg

func Compare(a int, b int) {
	if a > b {
		println(a, " jest wieksze od b")
	} else if a < b {
		println("a jest mniejsze od b")
	} else {
		println("a jest rowne b")
	}
}
