package electionday
import (
    "fmt"
)


// ElectionResult represents an election result.
type ElectionResult struct {
	// Name is the name of the candidate.
	Name string
	// Number is the total number of votes the candidate had.
	Votes int
}




// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
    var counter *int
    counter = &initialVotes
	return counter
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if counter != nil{
    return *counter
    }
    return 0
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {

    p := ElectionResult{Name:candidateName, Votes:votes}
    return &p
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
    name := result.Name
    vote := result.Votes
	return fmt.Sprintf("%s (%d)",name, vote) 
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {

    results[candidate] -=1
}
