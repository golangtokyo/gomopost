package gomopost

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GomoPoster interface {
	Post(name, msg string) error
}

type gomopost struct {
	address string
}

func NewClient(address string) GomoPoster {
	return &gomopost{
		address: address,
	}
}

type payload struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (g *gomopost) Post(name, msg string) error {
	p := payload{
		Name:    name,
		Message: msg,
	}
	pb, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("failed to marshal parameters: %s", err.Error())
	}

	buf := bytes.NewBuffer(pb)
	req, err := http.NewRequest("POST", g.address, buf)
	if err != nil {
		return fmt.Errorf("failed to create post request: %s", err.Error())
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("post request failed: %s", err.Error())
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %s", err.Error())
	}

	return nil
}
