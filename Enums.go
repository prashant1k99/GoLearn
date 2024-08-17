package main

import "fmt"

/*
	Enumerated types(enums) are a special case of sum types.
	An enum is a type that has a fixed number of possible values, each with a distant name.
	Go doesn't have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.
*/

// Our enum type ServerState has an underlying int type.
type ServerState int

// The possible values for ServerState are defined as constants. The special keyword iota generates successive constant values automatically; in this case 0, 1,2 and so on.
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// By implementing the fmt.Stringer interface, values of ServerState can be printed out or converted to strings.
var stateName = map[ServerState]string{
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}
// This can get cumbersome if there are many possible values. In such cases the stringer tool can used in conjunction with go:generate to automate the process.

func (ss ServerState) String() string {
	return stateName[ss]
}

// If we have a value of type int, we cannot pass it to transition - the compiler will complain about type mismathc. This provides some degree of compile-time type safety for enums.
func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// connected

	ns2 := transition(ns)
	fmt.Println(ns2)
	// idle
}

// transition emulates a state transition for a server; it takes the existing state and returns a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Suppose we check some predicated here to determine the next state...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}