package libs

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
)

// UserOsInfo 定义系统信息结构体
type UserOsInfo struct {
	MAC        string `json:"mac"`
	OS         string `json:"os"`
	Arch       string `json:"arch"`
	AppName    string `json:"app_name"`
	Hostname   string `json:"hostname"`
	TopType    string `json:"toptype"`
	UseType    string `json:"usetype"`
	SourceType string `json:"sourcetype"`
}

// GetSystemInfo 获取系统信息
func GenerateSystemInfo() UserOsInfo {
	info := UserOsInfo{}

	// 获取MAC地址
	mac, err := getMACAddress()
	if err == nil {
		info.MAC = mac
	} else {
		info.MAC = ""
	}

	// 获取主机名
	hostname, err := os.Hostname()
	if err == nil {
		info.Hostname = hostname
	} else {
		info.Hostname = ""
	}

	// 获取操作系统和架构信息
	info.OS = runtime.GOOS
	info.Arch = runtime.GOARCH

	info.AppName = "godoos"
	info.TopType = os.Getenv("GODOTOPTYPE")
	if info.TopType == "" {
		info.TopType = "web"
	}
	info.UseType = "person"
	info.SourceType = "open"
	return info
}

// GetIPAddress 获取本机IP地址
func GetIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("no valid IP address found")
}

// getMACAddress 获取MAC地址
func getMACAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var macs string
	// 记录所有接口的信息
	for _, iface := range ifaces {
		if iface.HardwareAddr != nil {
			macs += iface.HardwareAddr.String()
		}
	}
	//返回md5加密数据
	return Md5Encrypt(macs), nil
}
func Md5Encrypt(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GetSystemInfo 生成基于mac、os、arch信息的Base64编码字符串
func GetSystemInfo() (string, error) {
	// 构造系统信息对象
	lineseInfo, ok := GetConfig("osInfo")
	if !ok {
		return "", fmt.Errorf("未找到osInfo配置")
	}
	systemInfo, ok := lineseInfo.(UserOsInfo)
	if !ok {
		return "", fmt.Errorf("osInfo配置错误")
	}
	// 将系统信息序列化为JSON字符串
	jsonBytes, err := json.Marshal(systemInfo)
	if err != nil {
		return "", err
	}
	// 对JSON字符串进行Base64编码
	encodedInfo := base64.StdEncoding.EncodeToString(jsonBytes)

	return encodedInfo, nil
}
