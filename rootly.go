package main

import (
	"log"
	"os"
	"samfaina/dev/rootly_fetch/database"
	"samfaina/dev/rootly_fetch/jobs"
	"samfaina/dev/rootly_fetch/logger"
)

func init() {
	logger.SetupLogger()
}

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 || len(argsWithoutProg) > 1 || (!contains(argsWithoutProg, "sow") && !contains(argsWithoutProg, "harvest") && !contains(argsWithoutProg, "all")) {
		log.Println("Invalid arguments: Must be 'sow' OR 'harvest' OR 'all'")
		os.Exit(1)
	}

	arg := argsWithoutProg[0]
	switch arg {
	case "sow":
		log.Println("Starting the application into SOW mode")
		plants := jobs.FetchPlantsIds()
		database.SaveIDSToDatabase(plants)

	case "harvest":
		log.Println("Starting the application into HARVEST mode")
		jobs.FetchPlants()

	case "all":
		plants := jobs.FetchPlantsIds()
		database.SaveIDSToDatabase(plants)
		jobs.FetchPlants()
	}

}

// contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
