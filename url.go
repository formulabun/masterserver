package masterserver

import "net/url"

func buildUrl() string {
	// TODO might make this more dynamic, like getting the version from the api
	raw, err := url.JoinPath("https://ms.kartkrew.org", "ms", "api", "games", "SRB2Kart", "10", "servers")
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}

	q, err := url.ParseQuery("v=2.2")
	if err != nil {
		panic(err)
	}
	u.RawQuery = q.Encode()

	return u.String()
}
