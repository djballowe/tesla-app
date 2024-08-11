package getdata

import "net/http"

func GetChargeState(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(writer, "404 not found", http.StatusNotFound)
	}

	response := "80%"

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(response))
}
