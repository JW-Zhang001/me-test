package initialize

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"regexp"
)

var (
	Client *ssh.Client
	err    error
	home   = "cd /home/meuser/me-test/deploy && "
)

func init() {
	config := &ssh.ClientConfig{
		User: "meuser",
		Auth: []ssh.AuthMethod{
			ssh.Password("12345678"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	Client, err = ssh.Dial("tcp", "192.168.0.207:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
}

func executeCmd(cmd string) (string, error) {
	// 执行命令
	defer Client.Close()
	session, err := Client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	output, err := session.Output(cmd)
	if err != nil {
		panic("Failed to execute command: " + err.Error())
	}
	return string(output), nil
}

func GetValidatorPubKey(nodeID string) (string, error) {
	showValidator := home + "./me-chaind tendermint show-validator --home=" + nodeID

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
