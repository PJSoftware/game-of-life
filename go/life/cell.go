package life

// Cell object represents individual cell in world grid
type Cell struct {
	currentState, futureState bool
}

// Init sets Cell to initial state
func (c *Cell) Init(state bool) {
	c.currentState = state
	c.futureState = c.currentState
}

// UpdateState sets future state of Cell
func (c *Cell) UpdateState(state bool) {
	c.futureState = state
}

// Refresh switches Cell to its future state
func (c *Cell) Refresh() {
	c.currentState = c.futureState
}

// IsAlive returns true if Cell is currently "alive"
func (c *Cell) IsAlive() bool {
	return c.currentState
}

// String provides automatic stringification for Cell
func (c *Cell) String() string {
	rv := "."
	if c.currentState {
		rv = "O"
	}
	return rv
}
