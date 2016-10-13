package main

// Commands is a collection of Command structs
type Commands []Command

// SortByMatch returns the best matches for the given args
func (commands *Commands) SortByMatch(args []string) Commands {
	matches := make(map[int]Commands)
	strongest := 0

	for _, command := range *commands {
		strength := command.MatchStrength(args)
		matches[strength] = append(matches[strength], command)
		if (strength > strongest) {
			strongest = strength
		}
	}

	return matches[strongest]
}

// FilterExecutables returns list of only executable commands
func (commands *Commands) FilterExecutables(args []string) Commands {	
	var executables Commands

	for _, command := range *commands {
		if (command.IsExecutable(args)) {
			executables = append(executables, command)
		}
	}

	return executables
}
