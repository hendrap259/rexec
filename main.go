package main

import (
	"flag"
	"fmt"
	"github.com/hendrap259/rexec/consul"
	"github.com/hendrap259/rexec/log"
	"github.com/hendrap259/rexec/util"
	"os"
	"os/exec"
	"strings"
)

var (
	hosts  = flag.String("h", "", "comma separated host")
	ghosts = flag.String("g", "", "specify group name from configuration to run")
	edit   = flag.Bool("e", false, "edit config")
	csl    = flag.String("consul", "", "put hostgroup name")
)

func main() {

	rexec := NewRexec()

	//parse flags
	args, err := rexec.ParseParameter()
	if err != nil {
		os.Exit(0)
	}

	//init configuration

	//configuration files
	if *edit {
		err := editConfig()
		if err != nil {
			log.Error(err.Error())
		}
		os.Exit(0)
	}

	//exec
	if err := MainFunction(args); err != nil {
		os.Exit(0)
	}
}

func NewRexec() MainModules {

	err := editConfig()

	return &Main{
		Consul: consul.NewConsul(),
	}
}

type Main struct {
	Consul consul.ConsulInterface
}

type MainModules interface {
	ParseParameter() (flags []string, err error)
}

func (m *Main) ParseParameter() (flags []string, err error) {
	//validation input
	if len(os.Args) < 2 {
		log.Error("usage: rexec [-e | -h <hosts> | -g <group> | -consul <hostgroupName>] <command>")
		return flags, err
	}

	//validate parameter after flags
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 && !*edit {
		log.Error("invalid arguments, not enough arguments")
		return flags, err
	}

	return args, nil
}

func MainFunction(args []string) error {

	var err error
	var IPAddressList []string

	//searching list of IP Address from user input
	switch {
	case *hosts != "":
		IPAddressList = strings.Split(*hosts, ",")

	case *ghosts != "":
		IPAddressList, err = readHostConfig(*ghosts)
		if err != nil {
			return err
		}

	case *csl != "":
		con := consul.NewConsul("")
		IPAddressList, err = con.GetListIPHostgroup(*csl)
		if err != nil {
			return err
		}
	}

	//execute command
	var grCount int
	errChan := make(chan error)
	for _, ip := range IPAddressList {
		go run(fmt.Sprintf("root@%s", ip), args, errChan)
		grCount++
	}
	for grCount != 0 {
		err := <-errChan
		if err != nil {
			log.Error(err.Error())
		}
		grCount--
	}

	return nil
}

func run(server string, command []string, err chan error) {
	cmds := []string{"tsh", "ssh", server, strings.Join(command, " ")}
	fmt.Println("Executing : ", cmds)

	cmd := exec.Command(cmds[0], cmds[1:]...)

	cmd.Stdout = newWriter(util.RandomizeColor(fmt.Sprintf("[%s] ", server)))
	cmd.Stderr = newWriter(util.ErrorColor(fmt.Sprintf("[%s] ERR : ", server)))

	if errno := cmd.Run(); errno != nil {
		err <- fmt.Errorf("[%s] %s", server, errno.Error())
		return
	}

	err <- fmt.Errorf("[%s] %s", server, "session closed")
}

type writer struct {
	prefix string
	pipe   chan string
}

func newWriter(prefix string) *writer {
	w := &writer{
		prefix: prefix,
		pipe:   make(chan string),
	}
	go w.run()
	return w
}

func (c *writer) run() {
	for {
		fmt.Print(<-c.pipe)
	}
}

func (c *writer) Write(b []byte) (int, error) {
	c.pipe <- c.prefix + string(b)
	return len(b), nil
}
