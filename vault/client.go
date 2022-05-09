package vault

import (
	"errors"
	"fmt"

	"github.com/hashicorp/vault/api"
)

type VaultClientService interface {
	GetInstance() *api.Logical
	GetValue(string) (string, error)
	SetPath(string)
	SetPathSecret(string)
}

type VaultClient struct {
	client *api.Logical
	path string
	pathSecret string
	results map[string]map[string]string
}

func NewVaultClient(addr, path, pathSecret, token string) (VaultClientService, error){
	config := api.Config {
		Address: addr,
	}

	client, err := api.NewClient(&config)
	if err != nil {
		return nil, err
	}

	client.SetToken(token)

	return &VaultClient{
		client: client.Logical(), 
		path: path, 
		pathSecret: pathSecret,
		results: make(map[string]map[string]string),
	}, nil
}

func (c *VaultClient) GetInstance() *api.Logical {
	return c.client
}

// <path>/data/<path-secret>/key
func (c *VaultClient) GetValue(key string) (string, error) {
	res, ok := c.results[c.pathSecret]
	if ok {
		val, ok := res[key]
		if !ok {
			return "", errors.New("Key not found in cached data")
		}

		return val, nil
	}

	secret, err := c.client.Read(fmt.Sprintf("%s/data/%s", c.path, c.pathSecret))
	if err != nil {
		return "", fmt.Errorf("Error reading secret path and path-secret: %w", err)
	}

	if secret == nil {
		return "", errors.New("Secret not found")
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", errors.New("Invalid data in secret")
	}

	secrets := make(map[string]string)

	for k, v := range data {
		val, ok := v.(string)
		if !ok {
			return "", errors.New("Secret value in data is not string")
		}

		secrets[k] = val
	}

	val, ok := secrets[key]
	if !ok {
		return "", errors.New("Key not found in retrieved data")
	}

	c.results[c.pathSecret] = secrets

	return val, nil
}

func (c *VaultClient) SetPath(path string) {
	c.path = path
}

func (c *VaultClient) SetPathSecret(pathSecret string) {
	c.pathSecret = pathSecret
}