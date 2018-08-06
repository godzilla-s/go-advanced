package network

func Run(typ int, edp, addr string) {
	if typ == 0 {
		if edp == "server" {
			Server(addr)
		}
		if edp == "client" {
			Client(addr)
		}
	}
	if typ == 1 {
		if edp == "server" {
			full_duplex_server(addr)
		}
		if edp == "client" {
			full_duplex_client(addr)
		}
	}
}
