package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

//连接的配置
type ClientConfig struct {
	Host       string      //ip address
	Port       int64       // port
	Username   string      //username
	Password   string      //password
	Client     *ssh.Client //ssh client
	LastResult string      //lastest result
}

func (cliConf *ClientConfig) createClient(flags input) {
	var (
		client *ssh.Client
		err    error
	)
	cliConf.Host = flags.ip
	cliConf.Port = flags.port
	cliConf.Username = flags.username
	cliConf.Password = flags.password
	cliConf.Port = flags.port

	var auth []ssh.AuthMethod
	if flags.keyfile != "" {
		key, err := ioutil.ReadFile(flags.keyfile)
		checkError(err)
		var signer ssh.Signer
		if flags.password == "" {
			signer, err = ssh.ParsePrivateKey(key)
			checkError(err)
		} else if flags.password != "" {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte("password"))
			checkError(err)
		}
		auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	} else if flags.keyfile == "" {
		auth = []ssh.AuthMethod{ssh.Password(flags.password)}
	}

	config := ssh.ClientConfig{
		User: cliConf.Username,
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", cliConf.Host, cliConf.Port)

	if client, err = ssh.Dial("tcp", addr, &config); err != nil {
		checkError(err)
	}

	cliConf.Client = client
}

func (cliConf *ClientConfig) RunShell(shell string) string {
	var (
		session *ssh.Session
		err     error
	)

	if session, err = cliConf.Client.NewSession(); err != nil {
		checkError(err)
	}

	if output, err := session.CombinedOutput(shell); err != nil {
		checkError(err)
	} else {
		cliConf.LastResult = string(output)
	}
	return cliConf.LastResult
}

type input struct {
	ip       string
	port     int64
	username string
	password string
	keyfile  string
}

func dealFlag() (flags input, cmd string) {
	flag.StringVar(&flags.ip, "ip", "", "target address")
	flag.Int64Var(&flags.port, "port", 22, "target port")
	flag.StringVar(&flags.username, "u", "root", "ssh username")
	flag.StringVar(&flags.password, "p", "", "ssh password or password of private key")
	flag.StringVar(&flags.keyfile, "k", "", "private key filename")
	flag.StringVar(&cmd, "c", "", "command")
	flag.Parse()
	return flags, cmd
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Something was wrong", err)
	}
}

func banner() {
	name := os.Args[0]
	fmt.Println("#######################################")
	fmt.Println("#    SSHPASSWD    |   Author:huy4ng   #")
	fmt.Println("#######################################")
	fmt.Println("Useage Example:")
	fmt.Println("\nUser username and password to connect SSH and execute command")
	fmt.Printf("\n\t %s -ip=127.0.0.1 -port=22 -u=root -p=root -c=id\n", name)
	fmt.Println("\nUser username and password to connect SSH and execute command in an interactive method")
	fmt.Printf("\n\t %s -ip=127.0.0.1 -port=22 -u=root -p=root\n", name)
	fmt.Println("\nUser username and private key to connect SSH and execute command")
	fmt.Printf("\n\t %s -ip=127.0.0.1 -port=22 -u=root -k=key.pem -c=id\n", name)
	fmt.Println("\nUser username and private key to connect SSH and execute command in an interactive method")
	fmt.Printf("\n\t %s -ip=127.0.0.1 -port=22 -u=root -k=key.pem\n", name)
	fmt.Println("\nUser username and encryptedprivate key to connect SSH and execute command")
	fmt.Printf("\n\t %s -ip=127.0.0.1 -port=22 -u=root -k=key.pem -p=123456 -c=id\n", name)
	fmt.Println("\nUser username and encrypted private key to connect SSH and execute command in an interactive method")
	fmt.Printf("\n\t %s -ip=127.0.0.1 -port=22 -u=root -k=key.pem -p=123456\n", name)
	fmt.Printf("\nUse \"%s -h\" for more help\n", name)
}
func main() {
	flags, cmd := dealFlag()
	//cliConf := new(ClientConfig)
	if len(os.Args) > 1 {
		//cliConf.createClient(flags)
		fmt.Println(flags.port)
		if cmd == "" {
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("$ ")
				command, err := reader.ReadString('\n')
				checkError(err)
				if command == "exit" {
					os.Exit(0)
				}
				//fmt.Println(cliConf.RunShell(command))
			}
		} else {
			//fmt.Println(cliConf.RunShell(cmd))
		}

	} else {
		banner()
	}

}
