package main

import (
	"fmt"
)

type OrganizingTeam struct {
	TeamMembers    []string
	PrimaryContact string
}

type NSPEvent struct {
	EventDate      string
	EventName      string
	PrimaryContact string
	OrganizingTeam OrganizingTeam
}

func CreateOrganizingTeam(eventDate string, eventName string, teamMembers []string, primaryContact string) NSPEvent {
	event := NSPEvent{
		EventDate: eventDate,
		EventName: eventName,
		OrganizingTeam: OrganizingTeam{
			TeamMembers:    teamMembers,
			PrimaryContact: primaryContact,
		},
	}
	return event
}

func ReadOrganizingTeam(event NSPEvent) {
	fmt.Println("Event Name:", event.EventName)
	fmt.Println("Event Date:", event.EventDate)
	fmt.Println("Organizing Team:")
	fmt.Println("Team Members:", event.OrganizingTeam.TeamMembers)
	fmt.Println("Primary Contact:", event.OrganizingTeam.PrimaryContact)
	fmt.Println()
}

func UpdateOrganizingTeam(event NSPEvent, teamMembers []string, primaryContact string) NSPEvent {
	if teamMembers != nil && len(teamMembers) != 0 {
		event.OrganizingTeam.TeamMembers = append(event.OrganizingTeam.TeamMembers, teamMembers...)
	}
	if len(primaryContact) != 0 {
		event.OrganizingTeam.PrimaryContact = primaryContact
	}
	return event
}

func RemoveTeamMembers(event NSPEvent, membersToRemove []string) NSPEvent {
	var updatedMembers []string

	for _, member := range event.OrganizingTeam.TeamMembers {
		shouldRemove := false
		for _, removeMember := range membersToRemove {
			if member == removeMember {
				shouldRemove = true
				break
			}
		}
		if !shouldRemove {
			updatedMembers = append(updatedMembers, member)
		}
	}

	event.OrganizingTeam.TeamMembers = updatedMembers
	return event
}

func DeleteOrganizingTeam(event NSPEvent) NSPEvent {
	event.OrganizingTeam = OrganizingTeam{}
	return event
}

func main() {

	event := CreateOrganizingTeam("1-Feb-2024", "Event1", []string{"Member1", "Member2"}, "TeamLead")
	ReadOrganizingTeam(event)

	event = UpdateOrganizingTeam(event, []string{"NewMember1", "NewMember2"}, "NewTeamLead")
	ReadOrganizingTeam(event)

	event = UpdateOrganizingTeam(event, []string{"NewMember3"}, event.OrganizingTeam.PrimaryContact)
	ReadOrganizingTeam(event)

	event = RemoveTeamMembers(event, []string{"Member1", "Member3"})
	ReadOrganizingTeam(event)

	event = DeleteOrganizingTeam(event)
	ReadOrganizingTeam(event)
}
