package main

import (
	"fmt"
	"sort"
	"strings"
)

// State definition
// which contains an array of 9 integers. (0 for blank)
// and an integer variable, h that describes
// admissible heuristic value
type State struct {
	array [9]int
	cost  [2]int
}

var goalState State = State{
	array: [9]int{1, 2, 3, 4, 5, 6, 7, 8, 0},
	cost:  [2]int{0, 0},
}

var initialStates [30]State

func calculateHeuristic(s State) (h int) {
	for i := 0; i < len(s.array)-1; i++ {
		if (s.array[i] - 1) != i {
			h++
		}
	}

	return
}

//func puzzleGenerator() {
// random number generator
//s1 := rand.NewSource(time.Now().UnixNano())
//r1 := rand.New(s1)
// generated rand-num
//randNum := r1.Intn(27)
//randInit := goalState

//for i := 0; i < randNum; i++ {

//}
//}

var initialState State = State{
	array: [9]int{1, 8, 2, 0, 4, 3, 7, 6, 5},
	cost:  [2]int{0, 0},
}

func printState(state State) {
	fmt.Printf("+------+")
	for i, v := range state.array {
		if i%3 == 0 {
			fmt.Printf("\n|")
		}
		if v == 0 {
			fmt.Printf("  ")
		} else {
			fmt.Printf("%d ", v)
		}
		if (i+1)%3 == 0 {
			fmt.Printf("|")
		}
	}
	fmt.Printf("\n+------+\n")
}

func printOp(op Operator) {
	fmt.Printf("\n-> Operation: ")
	switch {
	case op.dir == 0:
		fmt.Printf("%d UP", op.piece)
	case op.dir == 1:
		fmt.Printf("%d RIGHT", op.piece)
	case op.dir == 2:
		fmt.Printf("%d DOWN", op.piece)
	case op.dir == 3:
		fmt.Printf("%d LEFT", op.piece)
	}
}

// Operator definition
// which is applied to the state.
// The struct "Operator" describes which the piece
// moved which direction.
// it contains two int variables 'piece' and 'dir'
// e.g. Operator {6, 0} means that piece 6 moved up.
// dir: 0 - up, 1 - right, 2 - down, 3 - left
type Operator struct {
	piece, dir int
}

// Operators is an array of Operator
// contains the possible operations on a state
// a piece has four degrees of freedom
var Operators = [32]Operator{
	{1, 0},
	{1, 1},
	{1, 2},
	{1, 3},

	{2, 0},
	{2, 1},
	{2, 2},
	{2, 3},

	{3, 0},
	{3, 1},
	{3, 2},
	{3, 3},

	{4, 0},
	{4, 1},
	{4, 2},
	{4, 3},

	{5, 0},
	{5, 1},
	{5, 2},
	{5, 3},

	{6, 0},
	{6, 1},
	{6, 2},
	{6, 3},

	{7, 0},
	{7, 1},
	{7, 2},
	{7, 3},

	{8, 0},
	{8, 1},
	{8, 2},
	{8, 3},
}

func arrayToString(a [9]int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func findPos(e int, state State) int {
	pos := -1
	for i, v := range state.array {
		if v == e {
			pos = i
			break
		}
	}
	return pos
}

// valid is the validator for the given state
// it ensures that the given state is "legal"
func valid(state State, op Operator) bool {
	blankPos := findPos(0, state)
	dir := op.dir
	piecePos := findPos(op.piece, state)
	if blankPos == -1 {
		fmt.Println("Something terrible happened. -0")
		return false
	}

	switch {
	case dir == 0 && piecePos-3 == blankPos:
		return true
	case dir == 1 && piecePos+1 == blankPos && blankPos%3 != 0:
		return true
	case dir == 2 && piecePos+3 == blankPos:
		return true
	case dir == 3 && piecePos-1 == blankPos && (blankPos+1)%3 != 0:
		return true
	default:
		return false
	}

}

func stateTransition(currentState State, op Operator) (nextState State, ok bool) {

	blankPos := findPos(0, currentState)
	piecePos := findPos(op.piece, currentState)
	nextState.array = currentState.array
	nextState.array[blankPos] = op.piece
	nextState.array[piecePos] = 0

	ok = valid(currentState, op)

	return // dry return
}

func main() {

	var q []State
	//printState(initialState)
	initState := initialState
	initialState.cost[0] = calculateHeuristic(initialState)
	initialState.cost[1] = 0
	fmt.Println(initialState.array, initialState.cost)
	fmt.Println(initState.array, initState.cost)
	q = append(q, initialState)
	initStr := arrayToString(initialState.array, ",")
	//i1Str := strings.Trim(strings.Join(strings.Split(fmt.Sprint(initState.array), " "), ","), "[]")
	history := map[string]bool{initStr: true}
	moves := 0
	for len(q) != 0 {
		s := q[0]
		q = q[1:]
		fmt.Println("State: ")
		printState(s)
		fmt.Printf("h: %d, g: %d f: %d\n", s.cost[0], s.cost[1], s.cost[0]+s.cost[1])

		if s.array == goalState.array {
			fmt.Printf("Goal Found!\n")
			break
		}
		var tmpStr string
		moves++
		for _, op := range Operators {
			nextState, ok := stateTransition(s, op)
			if ok { // legal state
				tmpStr = arrayToString(nextState.array, ",")
				if !history[tmpStr] { // new state
					nextState.cost[0] = calculateHeuristic(nextState)
					nextState.cost[1] = s.cost[1] + 1
					history[tmpStr] = true
					//printOp(op)
					//fmt.Println("Child:")
					//printState(nextState)
					//fmt.Printf("h: %d, g: %d f: %d\n", nextState.cost[0], nextState.cost[1], nextState.cost[0]+nextState.cost[1])
					// add it to queue
					q = append(q, nextState)
				} else { // already discovered state
					// no need to consider it since if this state is already visited,
					// then there is no way a repated state can be less than previous f
					// since h is constant and every level g increases by 1.
				}
			}
		}
		// all immediate child nodes discovered & they are in the queue
		// sort the queue for minimum f
		fmt.Println()
		sort.Slice(q, func(i, j int) bool {
			f1 := q[i].cost[1] + q[i].cost[0]
			f2 := q[j].cost[1] + q[j].cost[0]
			if f1 < f2 {
				return true
			}
			return false
		})
		/*reader := bufio.NewReader(os.Stdin)
		fmt.Print("\n<continue: >")
		text, _ := reader.ReadString('\n')
		text += ""*/
	}
	fmt.Println("moves", moves)
}
