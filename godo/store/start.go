package store

import (
	"fmt"
	"godo/libs"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/gorilla/mux"
)

func StartProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	processesMu.Lock()
	defer processesMu.Unlock()
	err := ExecuteScript(name)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Process %s started.", name)
}
func ExecuteStartAll() error {
	processesMu.Lock()
	defer processesMu.Unlock()

	for name, cmd := range processes {
		if err := cmd.Cmd.Start(); err != nil {
			return fmt.Errorf("failed to stop process %s: %v", name, err)
		}
	}

	return nil
}
func StartAll(w http.ResponseWriter, r *http.Request) {
	if err := ExecuteStartAll(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "All processes started.")
}

// ExecuteScript 执行指定名称的脚本。
// 参数：
// name - 脚本的名称。
// 返回值：
// 返回可能遇到的错误，如果执行成功，则返回nil。
func ExecuteScript(name string) error {
	storeInfo, err := GetStoreInfo(name)
	if err != nil {
		return err
	}
	err = runStart(storeInfo)
	if err != nil {
		return fmt.Errorf("failed to run script: %v", err)
	}
	return nil
}
func RunOutHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	fileName := filepath.Base(url)
	cacheDir := libs.GetCacheDir()
	filePath := filepath.Join(cacheDir, fileName)
	if !libs.PathExists(filePath) {
		libs.ErrorMsg(w, "file not found")
		return
	}
	cmd := exec.Command(filePath)
	if err := cmd.Start(); err != nil {
		libs.ErrorMsg(w, fmt.Sprintf("start error: %v", err))
		return
	}

	// 如果 cmd.Start() 成功，返回成功消息
	libs.SuccessMsg(w, "", "start success")

}
