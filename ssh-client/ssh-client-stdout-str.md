```
package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	//
	"bytes"
	"fmt"
	"strings"
)

var username = "root"
var password = "123456"
var host = "172.16.99.10:22"

//var commandToExecute = "hostname"
var commandToExecute = "find /opt/bozz/   -name pom.xml | grep parent"

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

	//
	var stdOut, stdErr bytes.Buffer

	session.Stdout = &stdOut
	session.Stderr = &stdErr

	session.Run(commandToExecute)
	//fmt.Println(stdOut.String())
	//fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
	if stdOut.String() == "" {
		fmt.Println("无法获取pom")
		return
	} else {
		fmt.Println(strings.TrimSpace(stdOut.String()))
	}

}
```
