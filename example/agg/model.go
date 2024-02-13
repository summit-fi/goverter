package agg

type Input1 struct {
	ID      int
	Name    string
	Root    string
	Admin   string
	Color   string
	Tax     string
	Phone   string
	Address string
}
type Output1 struct {
	ID      int
	Name    string
	Root    string
	Admin   string
	Color   string
	Tax     string
	Phone   string
	Address []string
}
