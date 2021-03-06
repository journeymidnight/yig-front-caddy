package lib

import "fmt"

func NewCustomDomain(domainHost string) (int, error) {
	url := "http://" + End_Point + "/?x-oss-action=NewCustomDomain"
	token, err := GetJwtForCustomDomain(domainHost)
	if err != nil {
		return 0, err
	}
	status, err := SetRequest("PUT", url, token)
	return status, err
}

func GetCustomDomain(domainHost string) (int, error) {
	url := "http://" + End_Point + "/?x-oss-action=GetCustomDomain"
	token, err := GetJwtForCustomDomain(domainHost)
	if err != nil {
		return 0, err
	}
	status, err := SetRequest("GET", url, token)
	return status, err
}

func DelCustomDomain(domainHost string) (int, error) {
	url := "http://" + End_Point + "/?x-oss-action=DelCustomDomain"
	token, err := GetJwtForCustomDomain(domainHost)
	if err != nil {
		return 0, err
	}
	status, err := SetRequest("DELETE", url, token)
	return status, err
}

func CustomDomainAccess(domainHost string) (int, error) {
	url := "http://" + domainHost + "/" + TEST_KEY
	fmt.Println("Custom domain access path", url)
	status, err := SetRequestWithDomain("GET", url, "")
	return status, err
}

func CustomDomainAccessWithcert(domainHost string) (int, error) {
	url := "https://" + domainHost + "/" + TEST_KEY
	fmt.Println("Custom domain access path", url)
	status, err := SetRequestWithDomain("GET", url, "")
	return status, err
}
