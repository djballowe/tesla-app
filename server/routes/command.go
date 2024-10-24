package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tesla-app/server/common"
	"tesla-app/server/vehicle-state"
)

type IssueCommandResponse struct {
	Message string
	Success bool
}

func IssueCommand(writer http.ResponseWriter, req *http.Request) {
	var issueCommandResponse = IssueCommandResponse{}
	if req.Method != "POST" {
		http.Error(writer, "This method is not supported", http.StatusNotFound)
		return
	}

	req.ParseForm()
	command := req.Form.Get("command")
	vehicleState, err := vehicle.VehicleState()
	if err != nil {
		http.Error(writer, "Could not get vehicle state", http.StatusInternalServerError)
	}

	state := vehicleState.State
	vin := vehicleState.Vin
	fmt.Println("State: ", state)
	fmt.Println("Vin: ", vin)
	if state != "online" {
		// call wake endpoint
	}

	tokenStore, state := common.GetTokenStore()

	authToken := tokenStore[state].AccessToken
	var commandReq = common.CommandRequest{
		AuthToken: authToken,
		Vin:       vin,
		Command:   command,
	}

	commandResp := common.HandleCommand(commandReq)

	fmt.Printf("Command issue status: Success? %t Message: %s\n", commandResp.Success, commandResp.Message)

	if !commandResp.Success {
		http.Error(writer, "Command could not be issued", http.StatusInternalServerError)
		return
	}

	issueCommandResponse.Message = commandResp.Message
	issueCommandResponse.Success = commandResp.Success

	jsonResponse, err := json.Marshal(issueCommandResponse)
	if err != nil {
		http.Error(writer, "Could not marshal response body", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}
