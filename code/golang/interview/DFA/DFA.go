package DFA

import "fmt"

type Status struct {
	status string
	rules  []*Rule
}

type Rule struct {
	input  string
	output string
	next   string
}

type DFA struct {
	statusList []*Status
	current    *Status
	output     string
}

func InitStatus(status, input, output, next string) *Status {
	return &Status{status: status, rules: []*Rule{
		{input, output, next},
	}}
}

func InitDFA(statusList []*Status) *DFA {
	return &DFA{statusList, statusList[0], ""}
}

func (dfa *DFA) Hit(input string) *Rule {
	for _, rule := range dfa.current.rules {
		if rule.input == input {
			return rule
		}
	}
	return nil
}

func (dfa *DFA) String() string {
	return fmt.Sprintf("Current: %s, Output: %s", dfa.current.status, dfa.output)
}

func (dfa *DFA) Set(status string, output string) {
	for _, st := range dfa.statusList {
		if st.status == status {
			dfa.current = st
			break
		}
	}
	dfa.output = output
}

func (dfa *DFA) Run(inputs []string) *Status {
	for _, input := range inputs {
		if rule := dfa.Hit(input); rule != nil {
			dfa.Set(rule.next, rule.output)
		} else {
			break
		}
	}
	return dfa.current
}
