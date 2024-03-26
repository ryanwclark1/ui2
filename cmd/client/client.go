package client

import (
    "fmt"

    "net/http"
    "os"
    "time"
    "io"
    "crypto/tls"
)

type InvalidArgumentError struct {
    argumentName string
}

func (e *InvalidArgumentError) Error() string {
    return fmt.Sprintf("Invalid value for argument \"%s\"", e.argumentName)
}

type BaseClient struct {
    Host              string
    Port              int
    Version           string
    Token             *string
    Tenant            *string
    HTTPS             bool
    Timeout           time.Duration
    VerifyCertificate bool
    Prefix            string
    UserAgent         string
    // Plugins would need a different approach in Go
}

func NewBaseClient(host string, port int, version string, token, tenant *string, https bool, timeout int, verifyCertificate bool, prefix, userAgent string) (*BaseClient, error) {
    if host == "" {
        return nil, &InvalidArgumentError{"host"}
    }
    if userAgent == "" {
        userAgent = os.Args[0]
    }
    return &BaseClient{
        Host:              host,
        Port:              port,
        Version:           version,
        Token:             token,
        Tenant:            tenant,
        HTTPS:             https,
        Timeout:           time.Duration(timeout) * time.Second,
        VerifyCertificate: verifyCertificate,
        Prefix:            prefix,
        UserAgent:         userAgent,
    }, nil
}

// Implement _build_prefix as a method in Go
func (c *BaseClient) buildPrefix(prefix string) string {
    if prefix == "" {
        return ""
    }
    if prefix[0] != '/' {
        prefix = "/" + prefix
    }
    return prefix
}

// Similar methods for session setup, URL building, etc., would be needed
// Example: URL building
func (c *BaseClient) URL(fragments ...string) string {
    scheme := "http"
    if c.HTTPS {
        scheme = "https"
    }
    port := ""
    if c.Port != 0 {
        port = fmt.Sprintf(":%d", c.Port)
    }
    version := ""
    if c.Version != "" {
        version = "/" + c.Version
    }
    base := fmt.Sprintf("%s://%s%s%s%s", scheme, c.Host, port, c.Prefix, version)
    if len(fragments) > 0 {
        path := "/" + join(fragments, "/")
        base += path
    }
    return base
}

func join(elements []string, sep string) string {
    // Implement join logic similar to strings.Join
    fmt.Println("Joining elements", elements, sep)
    x:= ""
    return x
}

// func (c *BaseClient) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
//     req, err := http.NewRequest(method, url, body)
//     if err != nil {
//         return nil, err
//     }
//     // Set headers based on BaseClient fields (Token, UserAgent, etc.)
//     return req, nil
// }

// NewRequest now also accepts custom headers to allow per-request customization.
func (c *BaseClient) NewRequest(method, url string, body io.Reader, headers map[string]string) (*http.Request, error) {
    req, err := http.NewRequest(method, url, body)
    if err != nil {
        return nil, err
    }

    // Set default headers here, for instance, User-Agent, X-Auth-Token, etc.
    if c.Token != nil {
        req.Header.Set("X-Auth-Token", *c.Token)
    }
    if c.Tenant != nil {
        req.Header.Set("Accent-Tenant", *c.Tenant)
    }
    req.Header.Set("User-Agent", c.UserAgent)

    // Apply any additional headers passed to the function.
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    return req, nil
}

// Example method to perform a GET request
func (c *BaseClient) Get(url string) (*http.Response, error) {
    req, err := c.NewRequest("GET", url, nil, map[string]string{})
    if err != nil {
        return nil, err
    }
    client := &http.Client{
        Timeout: c.Timeout,
    }
    return client.Do(req)
}

// Commands

// This method mimics the session behavior by configuring a custom HTTP client.
func (c *BaseClient) NewHTTPClient() *http.Client {
    // Customize the TLS configuration to match your verification requirements.
    tlsConfig := &tls.Config{
        InsecureSkipVerify: !c.VerifyCertificate,
    }
    transport := &http.Transport{
        TLSClientConfig: tlsConfig,
    }
    client := &http.Client{
        Timeout:   c.Timeout,
        Transport: transport,
    }
    return client
}



// HTTPCommander interface defines behavior for HTTP commands.
type HTTPCommander interface {
    Do() (*http.Response, error)
}

// HTTPCommand provides a base struct that other commands can embed.
type HTTPCommand struct {
    Client *BaseClient
}

// RESTCommand represents a specific type of HTTP command.
type RESTCommand struct {
    HTTPCommand
    Resource string
}

// NewRESTCommand creates a new instance of RESTCommand.
func NewRESTCommand(client *BaseClient, resource string) *RESTCommand {
    return &RESTCommand{
        HTTPCommand: HTTPCommand{Client: client},
        Resource:    resource,
    }
}

// Do executes the REST command. This is a placeholder for actual implementation.
func (cmd *RESTCommand) Do() (*http.Response, error) {
    // Construct the URL from the Resource and any necessary parameters.
    url := cmd.Client.URL(cmd.Resource)

    // Create a new request with any required headers.
    req, err := cmd.Client.NewRequest("GET", url, nil, map[string]string{
        "Accept": "application/json",
    })
    if err != nil {
        return nil, err
    }

    // Use the client's configured HTTP client to send the request.
    response, err := cmd.Client.NewHTTPClient().Do(req)
    if err != nil {
        return nil, err
    }

    // Additional error handling or response processing can go here.

    return response, nil
}
