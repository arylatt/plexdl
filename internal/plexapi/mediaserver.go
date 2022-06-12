package plexapi

import "fmt"

func (c *Client) Library() (m *MediaContainer, err error) {
	resp, err := c.MediaServerGet("/library")
	if err != nil {
		return
	}

	err = DecodeXMLResponse(resp, &m)
	return
}

func (c *Client) Sections() (m *MediaContainer, err error) {
	resp, err := c.MediaServerGet("/library/sections")
	if err != nil {
		return
	}

	err = DecodeXMLResponse(resp, &m)
	return
}

func (c *Client) Section(id string) (m *MediaContainer, err error) {
	resp, err := c.MediaServerGet(fmt.Sprintf("/library/sections/%s", id))
	if err != nil {
		return
	}

	err = DecodeXMLResponse(resp, &m)
	return
}

func (c *Client) SectionAll(id string) (m *MediaContainer, err error) {
	resp, err := c.MediaServerGet(fmt.Sprintf("/library/sections/%s/all", id))
	if err != nil {
		return
	}

	err = DecodeXMLResponse(resp, &m)
	return
}
