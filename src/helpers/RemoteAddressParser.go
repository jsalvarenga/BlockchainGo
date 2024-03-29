package helper

import (
	"errors"
	"regexp"
	"strings"
)

var ipRegex = `(\[.*\])`
var ptRegex = `(\]:(.*))`

// ParseRemoteAddress func
func ParseRemoteAddress(s string) (string, string, error) {
	// vars
	var ip string
	var port string
	var err error
	// regex
	iprx, _ := regexp.Compile(ipRegex)
	ptrx, _ := regexp.Compile(ptRegex)
	// get id
	if iprx.MatchString(s) {
		ip = string(iprx.Find([]byte(s)))
		ip = strings.Trim(ip, "[")
		ip = strings.Trim(ip, "]")
	} else {
		err = errors.New("can't find ip")
	}
	// get port
	if ptrx.MatchString(s) {
		port = string(ptrx.Find([]byte(s)))
		port = strings.Trim(port, "]:")
	} else {
		err = errors.New("can't find port")
	}
	return ip, port, err
}
