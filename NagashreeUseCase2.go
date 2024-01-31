package main

import "fmt"

type NSPEvent struct {
	EventDate      string
	EventName      string
	PrimaryContact string
	Organizingteam OrganizingTeam
}

type OrganizingTeam struct {
	TeamMembers    []string
	PrimaryContact string
}

func main() {

	Team1 := OrganizingTeam{TeamMembers: []string{"shriti", "giri", "sahana"}, PrimaryContact: "shriti"}
	Team2 := OrganizingTeam{TeamMembers: []string{"umesan", "jubin", "sunny", "Deepak"}, PrimaryContact: "umesan"}
	//Team3 := OrganizingTeam{TeamMembers: []string{"rakesh", "yeswanth", "soundarya"}, PrimaryContact: "rakesh"}
	//Team4 := OrganizingTeam{TeamMembers: []string{"subba", "anubav"}, PrimaryContact: "subba"}
	//Team5 := OrganizingTeam{TeamMembers: []string{"Divya", "Harshita", "chaitanya"}, PrimaryContact: "Divya"}

	//Events iniatilization
	Event1 := NSPEvent{"jan-23", "NSP-sankranti", Team1.PrimaryContact, Team1}
	Event2 := NSPEvent{"apr-16", "NSP-Eid party", "Bean", Team2}
	Event3 := NSPEvent{"jun-27", "NSP-Team outing", " ", OrganizingTeam{[]string{"rakesh", "yeswanth", "soundarya"}, "rakesh"}}
	Event4 := NSPEvent{"july-1", "NSP-Hackaton", "Tom", OrganizingTeam{[]string{"subba", "anubav"}, "subba"}}
	Event5 := NSPEvent{"july-20", "NSP-Walkathon", "Jerry", OrganizingTeam{[]string{"Divya", "Harshita", "chaitanya"}, "Divya"}}

	fmt.Println("Event 1 details are :", Event1.EventDate, Event1.EventName, Event1.PrimaryContact, Event1.Organizingteam.TeamMembers)
	fmt.Println("Event 2 details are :", Event2.EventDate, Event2.EventName, Event2.PrimaryContact, Event2.Organizingteam.TeamMembers)
	fmt.Println("Event 3 details are :", Event3.EventDate, Event3.EventName, Event3.Organizingteam.PrimaryContact, Event3.Organizingteam.TeamMembers)
	fmt.Println("Event 4 details are :", Event4.EventDate, Event4.EventName, Event4.PrimaryContact, Event4.Organizingteam.TeamMembers)
	fmt.Println("Event 5 details are :", Event5.EventDate, Event5.EventName, Event5.Organizingteam.PrimaryContact, Event5.Organizingteam.TeamMembers)
}
