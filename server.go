package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var adjectives = []string{
	"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
	"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient", "twilight", "dawn",
	"crimson", "wispy", "weathered", "blue", "billowing", "broken", "cold", "damp", "falling",
	"frosty", "green", "long", "late", "lingering", "bold", "little", "morning", "muddy", "old",
	"red", "rough", "still", "small", "sparkling", "throbbing", "shy", "wandering", "withered",
	"wild", "black", "young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
	"restless", "divine", "polished", "ancient", "purple", "lively", "nameless",
}

var nouns = []string{
	"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning", "snow",
	"lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest", "hill", "cloud",
	"meadow", "sun", "glade", "bird", "brook", "butterfly", "bush", "dew", "dust", "field", "fire",
	"flower", "firefly", "feather", "grass", "haze", "mountain", "night", "pond", "darkness",
	"snowflake", "silence", "sound", "sky", "shape", "surf", "thunder", "violet", "water",
	"wildflower", "wave", "water", "resonance", "sun", "wood", "dream", "cherry", "tree", "fog",
	"frost", "voice", "paper", "frog", "smoke", "star",
}

var adjectivesLength = len(adjectives)
var nounsLength = len(nouns)

const maxID = 9999
const minID = 1000

const defaultPort = "5000"

// NameServer starts a haiku name handlers. GET /names
func NameServer(w http.ResponseWriter, r *http.Request) {
	adjective := adjectives[rand.Intn(adjectivesLength)]
	noun := nouns[rand.Intn(nounsLength)]
	id := rand.Intn(maxID-minID) + minID

	haiku := fmt.Sprintf("%s-%s-%d", adjective, noun, id)
	fmt.Fprint(w, haiku)
}

func main() {
	port := fmt.Sprintf(":%s", getPort())
	http.HandleFunc("/names", NameServer)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("could not listen on port %s %v", port, err)
	}
}

func getPort() string {
	value := os.Getenv("NAMES_PORT")
	if len(value) == 0 {
		return defaultPort
	}

	return value
}
