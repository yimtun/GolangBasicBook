```
package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

var username = "root"
var password = "123456"
var host = "192.168.20.133:22"

var commandToExecute = "hostname"
//var commandToExecute = "find /opt/bozz/   -name pom.xml | grep parent"

func main() {

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server. ", err)
	}

	// Multiple sessions per client are allowed
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Pipe the session output directly to standard output
	// Thanks to the convenience of writer interface
	session.Stdout = os.Stdout

	session.Run(commandToExecute)
	if err != nil {
		log.Fatal("Error executing command. ", err)
	}
}
```
