package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(teamName1 string, point1 uint, teamName2 string, point2 uint) {
	if _, ok := l.Teams[teamName1]; !ok {
		l.Teams[teamName1] = Team{Name: teamName1}
		l.Wins[teamName1] = 0
	}
	if _, ok := l.Teams[teamName2]; !ok {
		l.Teams[teamName2] = Team{Name: teamName2}
		l.Wins[teamName2] = 0
	}

	if point1 > point2 {
		l.Wins[teamName1]++
	} else if point1 < point2 {
		l.Wins[teamName2]++
	}
}

func (l *League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k, _ := range l.Teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": Team{
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": Team{
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": Team{
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": Team{
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()
	fmt.Println(results)
}
