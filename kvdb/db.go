package code_challenge

type Database interface {
	// Set the key to given value
	Set(key string, value int)

	// Get the value for the given key, set 'ok' to true if key exists
	Get(key string) (value int, ok bool)

	// Unset the key, making it just like that key was never set
	Unset(key string)

	// Begin opens a new transaction
	Begin()

	// Commit closes all open transaction blocks, permanently apply the
	// changes made in them.
	Commit() error

	// Rollback undoes all of the commands issued in the most recent
	// transaction block, and closes the block.
	Rollback() error
}

type cat struct {
	name  string
	value int
}

//receiver (c *cat)
func (c *cat) Begin() {

}

func (c *cat) Commit() error {
	panic("todo")
}

func (c *cat) Get(key string) (value int, ok bool) {
	return c.value
}

func (c *cat) Set(key string, value int) {
	c.name = key
	c.value = value
}

func (c *cat) Unset(key string) {
	panic("todo")
}

func (c *cat) Rollback() error {
	panic("todo")
}

func NewDatabase() Database {
	c := cat{}
	return &c
}
