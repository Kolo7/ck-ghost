package common

import "fmt"

func BuildDsn(host, port, username, password, option string) string {
	return fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&%s", host, port, username, password, option)
}
