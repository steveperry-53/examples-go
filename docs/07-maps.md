# maps.go

Demonstrates how to use maps.

## Building and running the program
    
The source file, slices.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/07-maps
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/07-maps
    
The executable file, 07-maps, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/07-maps

A map is a set of key-value pairs. Maps are also called associative arrays,
dictionaries, hash tables, and ...?

For example, we could create a map where
each key is the name of a dog, and each value is the dog's age. The keys are
strings, and the values are integers.

    dogs = make(map[string]int)
      
Use [key] to set or read the value associated with key.

    dogs["Fluffy"] = 12
    dogs["Rover"] = 5
    dogs["Spot"] = 3

    fmt.Println("Fluffy is", dogs["Fluffy"], "years old.")
    
Create a map where each key is the name of a house plant, and each
value is a HousePlant struct.

    type HousePlant struct {
      location string
      cups int  // Cups of water per day needed by the plant
    }
    
    plants := map[string]HousePlant{
      "Spider": HousePlant{"corner window", 1},
      "Bush": HousePlant{"side window", 5},
    }
    
Then plants["Spider"] is a HousePlant. The location of the spider plant is
plants["Spider"].location.
    
    fmt.Println("The spider plant is near the", plants["Spider"].location + ".")
    fmt.Println("The bush requires", plants["Bush"].cups, "cups of water per day.")

Create a map where each key is the name of a motorcycle racer, and each value
has type Motorcycle.

    type Motorcycle struct {
      make string
      model string
      cc int  // engine size in cubic centimeters
    }

    racers := map[string]Motorcycle{
      "Bob": {"Suzuki", "Vstrom", 650},
      "Fred": {"Honda", "CBR", 900},
      "Harry": {"Kawasaki", "Ninja", 250},
    }

    fmt.Println("Bob rides a", racers["Bob"].make, racers["Bob"].model, "with a", racers["Bob"].cc, "cc engine.")
