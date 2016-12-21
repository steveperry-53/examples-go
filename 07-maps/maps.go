package main

import "fmt"

var dogs map[string]int

func main() {
  fmt.Println("Hello maps")

  dogs = make(map[string]int)

  dogs["Fluffy"] = 12
  dogs["Rover"] = 5
  dogs["Spot"] = 3

  fmt.Println(dogs)

  fmt.Println("Fluffy is", dogs["Fluffy"], "years old.")

  type HousePlant struct {
    location string
    cups int
  }

  plants := map[string]HousePlant{
    "Spider": HousePlant{"corner window", 1},
    "Bush": HousePlant{"side window", 5},
  }

  fmt.Println(plants)

  fmt.Println("The spider information is", plants["Spider"])
  fmt.Println("The bush information is", plants["Bush"])

  fmt.Println("The spider plant is near the", plants["Spider"].location + ".")
  fmt.Println("The bush requires", plants["Bush"].cups, "cups of water per day.")

  type Motorcycle struct {
    make string
    model string
    cc int
  }

  racers := map[string]Motorcycle{
    "Bob": {"Suzuki", "Vstrom", 650},
    "Fred": {"Honda", "CBR", 900},
    "Harry": {"Kawasaki", "Ninja", 250},
  }

  fmt.Println(racers)

  fmt.Println("Bob rides a", racers["Bob"].make, racers["Bob"].model, "with a", racers["Bob"].cc, "cc engine.")

}

// map, associative array, dictionary, hash table, set of key-value pairs






