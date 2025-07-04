package thirdparty

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "pharmacy/consumer/internal/models"
)

type Client struct {
    baseURL string
    client  *http.Client
}

func NewClient(baseURL string) *Client {
    return &Client{
        baseURL: baseURL,
        client:  &http.Client{},
    }
}

func (c *Client) SendTransaction(tx models.Transaction) error {
    data, err := json.Marshal(tx)
    if err != nil {
        return err
    }

    resp, err := c.client.Post(c.baseURL, "application/json", bytes.NewBuffer(data))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("third party API returned status: %d", resp.StatusCode)
    }

    return nil
}