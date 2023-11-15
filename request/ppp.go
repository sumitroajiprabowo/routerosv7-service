package request

type CreatePPPoERequest struct {
	RouterIpAddr       string `json:"router_ip_addr" validate:"required,ipv4"`
	RouterUsername     string `json:"router_username" validate:"required"`
	RouterPassword     string `json:"router_password" validate:"required"`
	UsernamePPoE       string `json:"username_pppoe" validate:"required"`
	PasswordPPPoE      string `json:"password_pppoe" validate:"required"`
	ProfilePPPoE       string `json:"profile_pppoe" validate:"required"`
	RemoteAddressPPPoE string `json:"remote-address_pppoe" validate:"required,ipv4"`
}

type DeletePPPoERequest struct {
	RouterIpAddr       string `json:"router_ip_addr" validate:"required,ipv4"`
	RouterUsername     string `json:"router_username" validate:"required"`
	RouterPassword     string `json:"router_password" validate:"required"`
	RemoteAddressPPPoE string `json:"remote-address_pppoe" validate:"required,ipv4"`
}
