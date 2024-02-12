package agg

type Input1 struct {
	ID      int
	Name    string
	Root    string
	Admin   bool
	Color   string
	Tax     rune
	Phone   string
	Address string
}

type Output1 struct {
	ID      int
	Name    string
	Root    string
	Admin   bool
	Color   string
	Tax     rune
	Phone   string
	Address []string
}
