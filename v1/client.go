package proxmox

import (
    "crypto/tls"
    "fmt"
    "net/http"
    "strings"

    validator "gopkg.in/validator.v2"
)

var Log Logger

type Client struct {
    ticket   Ticket
    config   *Config
    client   http.Client
    Simulate bool
}
type Config struct {
    ApiHost  string
    Username string
    Password string
    OTP      string
    TLS      *tls.Config
}

func (c *Client) Ping() error {
    var ticket Ticket
    err := c.Request(http.MethodGet, "/access/ticket", nil, &ticket)
    if err != nil {
        return fmt.Errorf("client.ping.Request: %v", err)
    }

    return nil
}

func (c *Client) validate(req interface{}) error {
    if err := validator.Validate(req); err != nil {
        // values not valid, deal with errors here
        return fmt.Errorf("client.validator.Error: %v", err)
    }
    return nil
}

func (c *Client) replaceBool(input string) string {
    payload := strings.ReplaceAll(input, "true", "1")
    payload = strings.ReplaceAll(payload, "false", "0")

    return payload
}

func NewClient(cfg *Config) (*Client, error) {
    client := http.Client{
        Transport: &http.Transport{
            TLSClientConfig:    cfg.TLS,
            DisableCompression: true,
            Proxy:              nil,
        },
        // CheckRedirect: nil,
        // Jar: nil,
        // Timeout: 0,
    }

    cfg.ApiHost = strings.TrimRight(cfg.ApiHost, "/") + "/api2/json"

    cli := Client{client: client, config: cfg}
    err := cli.authenticate()
    if err != nil {
        return nil, err
    }

    return &cli, nil
}
