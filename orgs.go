package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Org struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (o Org) DataSources(c *Client) ([]*DataSource, error) {
	return c.DataSourcesByOrgId(o.Id)
}

func (c *Client) Org(id int64) (Org, error) {
	org := Org{}

	req, err := c.newRequest("GET", fmt.Sprintf("/api/orgs/%d", id), nil)
	if err != nil {
		return org, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return org, err
	}
	if resp.StatusCode != 200 {
		return org, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return org, err
	}
	err = json.Unmarshal(data, &org)
	return org, err
}

func (c *Client) OrgByName(name string) (Org, error) {
	org := Org{}

	req, err := c.newRequest("GET", fmt.Sprintf("/api/orgs/name/%s", name), nil)
	if err != nil {
		return org, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return org, err
	}
	if resp.StatusCode != 200 {
		return org, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return org, err
	}
	err = json.Unmarshal(data, &org)
	return org, err
}

func (c *Client) Orgs() ([]Org, error) {
	orgs := make([]Org, 0)

	req, err := c.newRequest("GET", "/api/orgs/", nil)
	if err != nil {
		return orgs, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return orgs, err
	}
	if resp.StatusCode != 200 {
		return orgs, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orgs, err
	}
	err = json.Unmarshal(data, &orgs)
	return orgs, err
}

func (c *Client) NewOrg(name string) (int64, error) {
	settings := map[string]string{
		"name": name,
	}
	data, err := json.Marshal(settings)
	req, err := c.newRequest("POST", "/api/orgs", bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, errors.New(resp.Status)
	}

	body := struct {
		ID int64 `json:"orgId"`
	}{0}

	data, err = ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &body)

	return body.ID, err
}

func (c *Client) DeleteOrg(id int64) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/api/orgs/%d", id), nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}
