package initialize

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"me-test/config"
	"regexp"
)

var (
	Client  *ssh.Client
	Session *ssh.Session
	err     error
)

func init() {
	cfg := &ssh.ClientConfig{
		User: config.SSHConfig["user"],
		Auth: []ssh.AuthMethod{
			ssh.Password(config.SSHConfig["pwd"]),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	Client, err = ssh.Dial("tcp", config.SSHConfig["addr"], cfg)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
}

func executeCmd(cmd string) (string, error) {
	Session, err = Client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer Session.Close()
	output, err := Session.Output(cmd)
	if err != nil {
		panic("Failed to execute command: " + err.Error())
	}
	return string(output), nil
}

func GetValidatorPubKey(nodeID string) (string, error) {
	showValidator := config.SSHConfig["home"] + "./me-chaind tendermint show-validator --home=" + nodeID

	res, err := executeCmd(showValidator)
	if err != nil {
		return "", err
	}
	zap.S().Info("showValidator: ", res)

	re := regexp.MustCompile(`"key":"([^"]+)"`)
	match := re.FindStringSubmatch(res)
	if len(match) < 2 {
		zap.S().Error("Failed to extract key")
		return "", nil
	}
	key := match[1]

	tmPubK := "{\"type\": \"tendermint/PubKeyEd25519\",\"value\": \"${key}\"}"
	re2 := regexp.MustCompile(`\$\{key\}`)
	tmPubK = re2.ReplaceAllString(tmPubK, key)

	return tmPubK, nil
}
