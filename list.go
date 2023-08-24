package masterserver

import (
	"bufio"
	"net/http"
	"strings"
)

func ListServers() (Servers, error) {
	url := buildUrl()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := Servers{}

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			continue
		}
		res = append(res, ServerLine{
			parts[0], parts[1], parts[2],
		})
	}

	return res, nil
}
