package swisstable

type Entry struct {
	key   string
	value string
}
type Group struct {
	topHash []string
	hash    []string
	entries []Entry
}
type Table struct {
	entries []Group
}
type Directory struct {
	tables []Table
}

func main() {

}
