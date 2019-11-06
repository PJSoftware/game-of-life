package life

type Cell struct {
	currentState, futureState bool
}

func (c *Cell) Init(state bool) {
	c.currentState = state
	c.futureState = c.currentState
}

func (c *Cell) UpdateState(state bool) {
	c.futureState = state
}

func (c *Cell) Refresh() {
	c.currentState = c.futureState
}

func (c *Cell) IsAlive() bool {
	return c.currentState
}

func (c *Cell) ToString() string {
	rv := "."
	if c.currentState {
		rv = "O"
	}
	return rv
}
