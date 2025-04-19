package noip

func Call(username string, password string, hostname string) string {
	var baseURL = "@dynupdate.no-ip.com/nic/update?hostname="

	var url = "http://" + username +  ":" + password + baseURL + hostname + "&myip=ip"
	
	return url
}