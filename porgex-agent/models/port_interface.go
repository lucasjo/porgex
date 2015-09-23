package models

type PortInterface struct {
	CartridgeName   string   `json:"cartridge_name"`
	ExternalPort    string   `json:"external_port"`
	InternalAddress string   `json:"internal_address"`
	Protocols       []string `json:"proptocols"`
	Type            []string `json:"type"`
	Mappings        []string `json:"mappings"`
}
