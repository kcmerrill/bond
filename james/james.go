package james

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/kcmerrill/common.go/config"
	yaml "gopkg.in/yaml.v2"
)

var lock *sync.Mutex

type action struct {
	re  *regexp.Regexp
	cmd string
}

var actions map[string]*action

type actionConfig struct {
	actions map[string]string `yaml:",inline"`
}

// Execute a given command, and watch for a regular expression
func Execute(line string) bool {
	executed := false
	for _, act := range actions {
		if matches := act.re.FindStringSubmatch(line); matches != nil {
			cParsed := strings.Replace(act.cmd, ":match", matches[1], -1)
			cmd := exec.Command("bash", "-c", cParsed)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			executed = true
		}
	}
	return executed
}

// LookFor will take in a regular expression, and if found, will perform an action
func LookFor(expr string, cmd string) {
	lock.Lock()
	actions[expr] = &action{re: regexp.MustCompile(expr), cmd: cmd}
	lock.Unlock()
}

// Read will read from stdin
func Read() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		var wg sync.WaitGroup
		for scanner.Scan() {
			wg.Add(1)
			go func() {
				Execute(scanner.Text())
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// ReloadConfig will load LookFor at certain intervals
func ReloadConfig(configFile string, interval time.Duration) {
	for {
		if contents, err := config.Find(configFile); err == nil {
			ac := &actionConfig{
				actions: make(map[string]string),
			}
			if err := yaml.Unmarshal(contents, &ac.actions); err == nil {
				for expr, cmd := range ac.actions {
					LookFor(expr, cmd)
				}
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
		<-time.After(interval)
	}
}

func init() {
	// init our actions map
	actions = make(map[string]*action)
	// init our mutex
	lock = &sync.Mutex{}
}
