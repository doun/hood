package hood

type Dialect interface {
	Name() string          // dialect name
	Pk() string            // primary key
	Quote(s string) string // quote string
	Marker(pos int) string // marker for a prepared statement, e.g. $0 or ?
}