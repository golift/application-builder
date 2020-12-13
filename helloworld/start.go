package helloworld

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	flag "github.com/spf13/pflag"
	"golift.io/version"
	yaml "gopkg.in/yaml.v3"
)

// Flags defines our application's CLI arguments.
type Flags struct {
	VersionReq bool
	ConfigFile string
}

// Config defines our applications's config file parameters.
type Config struct {
	Worlds int `json:"worlds" xml:"worlds" toml:"worlds" yaml:"worlds"`
	Hellos int `json:"hellos" xml:"hellos" toml:"hellos" yaml:"hellos"`
}

// HelloWorld is the main application struct.
type HelloWorld struct {
	Flag *flag.FlagSet
	*Flags
	*Config
}

// Binary is the app name.
const Binary = "hello-world"

const (
	defaultConfFile = "/etc/hello-world/helloworld.conf"
	defaultWorlds   = 2
	defaultHellos   = 1
)

// Start begins the application from a CLI.
// Parses flags, parses config and executes Run().
func Start() error {
	hw := &HelloWorld{
		Config: &Config{},
		Flags:  &Flags{},
		Flag:   flag.NewFlagSet(Binary, flag.ExitOnError),
	}

	log.SetFlags(log.LstdFlags)
	hw.ParseFlags(os.Args[1:])

	if hw.VersionReq {
		fmt.Println(version.Print(Binary))

		return nil // don't run anything else w/ version request.
	}

	if err := hw.GetConfig(); err != nil {
		hw.Flag.Usage()

		return err
	}

	return hw.Run()
}

// ParseFlags runs the parser for CLI arguments.
func (u *HelloWorld) ParseFlags(args []string) {
	u.Flag.Usage = func() {
		fmt.Printf("Usage: %s [--config=filepath] [--version]", Binary)
		u.Flag.PrintDefaults()
	}

	u.Flag.StringVarP(&u.ConfigFile, "config", "c", defaultConfFile, "Config File (TOML Format)")
	u.Flag.BoolVarP(&u.VersionReq, "version", "v", false, "Print the version and exit")

	_ = u.Flag.Parse(args)
}

// GetConfig parses and returns our configuration data.
// Supports any format for config file: xml, yaml, json, toml.
func (u *HelloWorld) GetConfig() error {
	// Preload our defaults.
	u.Config = &Config{
		Hellos: defaultHellos,
		Worlds: defaultWorlds,
	}

	log.Printf("Loading Configuration File: %s", u.ConfigFile)

	switch buf, err := ioutil.ReadFile(u.ConfigFile); {
	case err != nil:
		return fmt.Errorf("ioutil.ReadFile: %w", err)
	case strings.Contains(u.ConfigFile, ".json"):
		return json.Unmarshal(buf, u.Config)
	case strings.Contains(u.ConfigFile, ".xml"):
		return xml.Unmarshal(buf, u.Config)
	case strings.Contains(u.ConfigFile, ".yaml"):
		return yaml.Unmarshal(buf, u.Config)
	default:
		return toml.Unmarshal(buf, u.Config)
	}
}

// Run starts doing things.
func (u *HelloWorld) Run() error {
	log.Printf("[INFO] Hello World v%v Starting Up! PID: %d", version.Version, os.Getpid())
	time.Sleep(time.Second)

	for i := 1; i <= u.Config.Hellos; i++ {
		fmt.Println(i, "hello")
	}

	for i := 1; i <= u.Config.Worlds; i++ {
		fmt.Println(i, "world")
	}

	return nil
}
