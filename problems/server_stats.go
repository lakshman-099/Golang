package main

import "fmt"

// Constants representing server statuses
const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// serverInfo function counts and prints server statuses
func serverInfo(statusMap map[string]int) {
	fmt.Println("\nTotal Servers: ", len(statusMap))

	// Initialize a map to count the number of servers in each status
	statusCount := make(map[int]int)

	for _, status := range statusMap {
		switch status {
		case Online:
			statusCount[Online]++
		case Offline:
			statusCount[Offline]++
		case Maintenance:
			statusCount[Maintenance]++
		case Retired:
			statusCount[Retired]++
		}
	}

	fmt.Println(statusCount[Online], " Servers are online")
	fmt.Println(statusCount[Offline], " Servers are offline")
	fmt.Println(statusCount[Maintenance], " Servers are in maintenance")
	fmt.Println(statusCount[Retired], " Servers are retired")
}

func main() {
	// Initialize a list of server names
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	// Initialize a map to store the status of each server
	serverStatus := make(map[string]int)

	// Set the initial status of all servers to Online
	for _, item := range servers {
		serverStatus[item] = Online
	}

	serverInfo(serverStatus)
	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline
	serverInfo(serverStatus)
	for key := range serverStatus {
		serverStatus[key] = Maintenance
	}

	serverInfo(serverStatus)
}
