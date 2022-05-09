package vault

import (
	"fmt"
	"testing"

	"github.com/mushoffa/go-library/vault"
)

const (
	VAULT_HOST_URL = "YOUR_VAULT_HOST_URL"
	VAULT_PATH = "YOUR_VAULT_PATH"
	VAULT_PATH_SECRET = "YOUR_VAULT_SECRET"
	VAULT_TOKEN = "YOUR_VAULT_TOKEN"
)

func TestNewVaultClient_Success(t *testing.T) {
	vault, err := vault.NewVaultClient(VAULT_HOST_URL, VAULT_PATH, VAULT_PATH_SECRET, VAULT_TOKEN)
	if err != nil {
		t.Errorf("Error creating vault client: %v", err)
	} else {
		dbhost, _ := vault.GetValue("host")
		dbport, _ := vault.GetValue("port")
		dbname, _ := vault.GetValue("name")
		dbuser, _ := vault.GetValue("user")
		dbpassword, _ := vault.GetValue("password")

		fmt.Println("DATABASE - Host    : ", dbhost)
		fmt.Println("DATABASE - Port    : ", dbport)
		fmt.Println("DATABASE - Name    : ", dbname)
		fmt.Println("DATABASE - Username: ", dbuser)
		fmt.Println("DATABASE - Password: ", dbpassword)

		// Change path-secret to retrieve other secret value
		vault.SetPathSecret("api")

		apiHost, _ := vault.GetValue("host")
		apiSecret, _ := vault.GetValue("secret")
		apiEndpoint, _ := vault.GetValue("endpoint")

		fmt.Println("")
		fmt.Println("API - Host    : ", apiHost)
		fmt.Println("API - Secret  : ", apiSecret)
		fmt.Println("API - Endpoint: ", apiEndpoint)
	}
}