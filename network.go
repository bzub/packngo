package packngo

type NetworkPort struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Type string `json:"type"`
	Name string `json:"name"`
	Data struct {
		Mac    string `json:"mac"`
		Bonded bool   `json:"bonded"`
	} `json:"data"`

	//TODO:
	// VirtualNetworks []*VirtualNetwork `json:"virtual_networks,omitempty"`
}
