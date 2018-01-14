package main

import (
	"fmt"

	mealy "github.com/evnix/mealy-fsm/mealy"
)

func onTransition(previousState string, currentState string, input string) {

	fmt.Println("previous state: "+previousState, "input: "+input, "current state: "+currentState)
}

func main() {

	st := mealy.CreateStateTransitionTable()
	st.AddRule("Q0", "0", "Q0", onTransition)
	st.AddRule("Q0", "1", "Q1", onTransition)
	st.AddRule("Q1", "1", "Q0", onTransition)
	st.AddRule("Q1", "0", "Q1", onTransition)

	st.SetInitialState("Q0")
	ipString := "10110011"

	for i := 0; i < len(ipString); i++ {
		nextState := st.GetNextState(string(ipString[i]))
		st.SetState(nextState, string(ipString[i]))
	}

	if st.GetCurrentState() == "Q0" {
		fmt.Println("There are Even number of 1s")
	} else {
		fmt.Println("There are Odd number of 1s")
	}

}
