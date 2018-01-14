package mealy

type StateTable struct {
	currentState string
	callback     map[string]OnNextStateEnter
	nextState    map[string]string
}

func CreateStateTransitionTable() *StateTable {
	st := &StateTable{}
	st.callback = make(map[string]OnNextStateEnter)
	st.nextState = make(map[string]string)
	return st
}

func (self *StateTable) SetInitialState(state string) {
	self.currentState = state
}

func (self *StateTable) GetNextState(input string) string {
	return self.nextState[self.currentState+input]
}

func (self *StateTable) GetCurrentState() string {
	return self.currentState
}

type OnNextStateEnter func(previousState string, currentState string, input string)

func (self *StateTable) AddRule(currentState string, input string, nextState string, callback OnNextStateEnter) {
	key := currentState + input
	self.callback[key] = callback
	self.nextState[key] = nextState
}

func (self *StateTable) SetState(state string, input string) {
	previousState := self.currentState
	self.currentState = state
	self.callback[self.currentState+input](previousState, self.currentState, input)
}
