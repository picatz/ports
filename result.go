package ports

type Result struct {
	IP    string
	Port  int
	Open  bool
	Error error
}
