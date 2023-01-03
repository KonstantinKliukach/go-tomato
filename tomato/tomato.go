package tomato

import (
	"fmt"
	"pomidoro/utils"
	"strings"
	"time"
)

type Tomato struct {
	Name     string
	Duration int
}

func (t Tomato) Start() {
	fmt.Printf("Time to %s\n", strings.ToLower(t.Name))
	seconds := 3
	time.Sleep(time.Duration(seconds) * time.Second)
	utils.Timer(t.Duration)
}

func GetListOfCommands(t map[string]Tomato) string {
	var r strings.Builder
	for key, el := range t {
		r.WriteString(fmt.Sprintf("\"%s\" for running %s with Duration of %d minutes\n", key, el.Name, el.Duration))
	}
	return r.String()
}

func GetAllTomatoes() map[string]Tomato {
	t := map[string]Tomato{
		"work": {
			Name:     "Work",
			Duration: 25,
		},
		"sr": {
			Name:     "Short rest",
			Duration: 5,
		},
		"lr": {
			Name:     "Long rest",
			Duration: 15,
		},
	}
	return t
}
