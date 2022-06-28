package proxmox

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "regexp"
    "strings"
)

func (c *Client) authenticate() error {
    data := url.Values{
        "username": {c.config.Username},
        "password": {c.config.Password},
        "otp":      {c.config.OTP},
    }

    err := c.Request(http.MethodPost, "/access/ticket", strings.NewReader(data.Encode()), &c.ticket)
    if err != nil {
        return fmt.Errorf("client.authenticate.Request: %v", err)
    }

    jar, err := cookiejar.New(nil)
    if err != nil {
        log.Fatalf("Got error while creating cookie jar %s", err.Error())
    }
    c.client.Jar = jar

    cookie := &http.Cookie{
        Name:   "PVEAuthCookie",
        Value:  c.ticket.Ticket,
        MaxAge: 300,
    }

    // fmt.Println("setting cookie for:", fmt.Sprintf("%s", strings.TrimRight(c.config.ApiHost, "/")))
    urlObj, _ := url.Parse(fmt.Sprintf("%s", strings.TrimRight(c.config.ApiHost, "/")))
    // urlObj, _ := url.Parse("https://dc1.r3d.lab:8006")
    c.client.Jar.SetCookies(urlObj, []*http.Cookie{cookie})

    // fmt.Println("authenticate:", c.ticket.Data.Ticket)

    return nil
}

func (c *Client) Request(method, url string, body io.Reader, resp interface{}) error {
    if body != nil {
        buf, err := ioutil.ReadAll(body)
        if err != nil {
            return fmt.Errorf("client.request.ReadAll: %v", err)
        }
        m1 := regexp.MustCompile(`password=[^&$]+`)

        fmt.Println("request", method, url, string(m1.ReplaceAll(buf, []byte("password=***MASKED***"))))

        body = bytes.NewReader(buf)
    } else {
        fmt.Println("request", method, url)
    }
    if c.Simulate {
        return nil
    }

    req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", strings.TrimRight(c.config.ApiHost, "/"), strings.TrimLeft(url, "/")), body)
    if err != nil {
        return fmt.Errorf("client.request.NewRequest: %v", err)
    }

    if strings.ToUpper(method) == http.MethodPost {
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    }
    if c.ticket.CSRFPreventionToken != "" {
        req.Header.Add("CSRFPreventionToken", c.ticket.CSRFPreventionToken)
    }

    res, err := c.client.Do(req)
    if err != nil {
        return fmt.Errorf("client.request.Do: %v", err)
    }
    defer res.Body.Close()

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return fmt.Errorf("client.request.ReadAll: %v", err)
    }

    // empty response, skip unmarshalling
    if len(b) == 0 {
        fmt.Println("   response body was empty, returned status code", res.StatusCode)
        return nil
    }

    unwrap := struct {
        Data   interface{}
        Errors map[string]string
    }{Data: resp, Errors: map[string]string{}}

    err = json.Unmarshal(b, &unwrap)
    if err != nil {
        return fmt.Errorf("client.request.Unmarshal: %v", err)
    }

    if res.StatusCode != 200 {
        return fmt.Errorf("client.request.Error: %v", unwrap.Errors)
    }

    // fmt.Println("RESULT BODY:", res.StatusCode, res.Header, string(b))

    return err
}
