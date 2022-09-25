package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to ecr registry",
	Long:  `Login to ecr registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login() {
	r, err := client.GetAuthorizationToken(context.TODO(), &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		log.Fatal(err)
	}
	for _, auth := range r.AuthorizationData {
		user, pass := parseToken(auth.AuthorizationToken)
		dockerLogin(&user, &pass, auth.ProxyEndpoint)
	}
}

// parse base64 token into user and password
func parseToken(token *string) (string, string) {
	data, err := base64.StdEncoding.DecodeString(*token)
	if err != nil {
		log.Fatal(err)
	}
	tok := strings.Split(string(data), ":")
	return tok[0], tok[1]
}

// run docker login shell command
// TODO: use --password-stdin
func dockerLogin(user, pass, endpoint *string) {
	out, err := exec.Command("docker", "login", "-u", *user, "-p", *pass, *endpoint).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}
