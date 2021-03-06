package core

type Info struct {
	Version string `json:"version,omitempty"`
	IP      string `json:"ip,omitempty"`
	FQDN    string `json:"fqdn,omitempty"`
	Env     string `json:"env,omitempty"`
	Color   string `json:"color,omitempty"`
	MOTD    string `json:"motd,omitempty"`

	API int `json:"api"`
}

func (core *Core) checkInfo() Info {
	return Info{
		IP:    core.ip,
		FQDN:  core.fqdn,
		MOTD:  core.motd,
		Color: core.color,
		Env:   core.env,
		API:   2,
	}
}
