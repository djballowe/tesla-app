package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CommandStatus struct {
	Status string `json:"status"`
}

type IssueCommandResponse struct {
	Body       CommandStatus
	StatusCode int
}

func CallIssueCommand(command string) (IssueCommandResponse, error) {
	commandApiResponse := &IssueCommandResponse{}
	var commandStatus CommandStatus

	reqUrl := fmt.Sprintf("http://localhost:8080/command?command=%s", command)
	req, err := http.NewRequest(http.MethodPost, reqUrl, nil)
	// TODO
	// this needs to be reworked the err doesnt throw if there is an http error
	// this shouldnt throw server esc errors on client code

	if err != nil {
		return handleCommandReturn(500, commandStatus, err)
	}

	req.Header.Add("Accept", "aplication/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if resp == nil || err != nil {
		return handleCommandReturn(500, commandStatus, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return handleCommandReturn(500, commandStatus, err)
	}

	if resp.StatusCode != 200 {
		return handleCommandReturn(500, commandStatus, err)
	}

	err = json.Unmarshal(body, &commandStatus)
	if err != nil {
		return handleCommandReturn(500, commandStatus, err)
	}

	commandApiResponse.Body = commandStatus
	commandApiResponse.StatusCode = resp.StatusCode

	return *commandApiResponse, nil
}

func handleCommandReturn(status int, body CommandStatus, err error) (IssueCommandResponse, error) {
	return IssueCommandResponse{
		StatusCode: status,
		Body:       body,
	}, err
}
