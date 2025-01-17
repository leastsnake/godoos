package store

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"sync"
	"time"
)

type Process struct {
	Name         string    `json:"name"`
	Running      bool      `json:"running"`
	ExitCode     int       `json:"exitCode"`
	Cmd          *exec.Cmd `json:"cmd"`
	Pid          int       `json:"pid"`
	ProgressName string    `json:"progressName"`
	Waiting      bool      `json:"waiting"`
	IsOn         bool      `json:"isOn"`
	LastPing     time.Time `json:"lastPing"`
}

var (
	processesMu sync.RWMutex
	processes   = make(map[string]*Process)
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}
func RegisterProcess(name string, progressName string, isOn bool, cmdstr *exec.Cmd) {
	processesMu.Lock()
	defer processesMu.Unlock()
	processes[name] = &Process{
		Name:         name,
		Running:      true,
		Cmd:          cmdstr,
		Pid:          cmdstr.Process.Pid,
		IsOn:         isOn,
		ProgressName: progressName,
	}
}
func GetCmd(name string) *Process {
	processesMu.Lock()
	defer processesMu.Unlock()

	return processes[name]
}
func Status(w http.ResponseWriter, r *http.Request) {
	var ps []Process
	for name, cmd := range processes {
		ps = append(ps, Process{Name: name, Running: cmd.Running, Waiting: cmd.Waiting, Pid: cmd.Pid, LastPing: cmd.LastPing, ExitCode: cmd.Cmd.ProcessState.ExitCode()})
	}

	jsonBytes, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to encode process status: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("Error writing health check response: %v", err)
	}
}
