package pkg

import (
	"encoding/json"
	"os"
)

const (
	assetsDataLocation  = "../assets/"
	accountDataLocation = assetsDataLocation + "account_data.json"
	abbreviationLen     = 3
)

func mockBasicAccountData(id string, name string) (ad *AccountData, e error) {
	bs, e := os.ReadFile(accountDataLocation)
	if e != nil {
		return nil, e
	}
	r := request[AccountData]{}
	if e = json.Unmarshal(bs, &r); e != nil {
		return nil, e
	}
	ad = &r.Data
	ad.ID = id
	if ad.Attributes != nil && len(name) > 0 {
		overwriteNames(name, ad.Attributes.Name)
		if abbreviationLen < len(name) {
			// mock of alternative name if applicable using prefix:
			overwriteNames(name[:abbreviationLen], ad.Attributes.AlternativeNames)
		} else {
			overwriteNames(name, ad.Attributes.AlternativeNames)
		}
	}
	return ad, e
}

func overwriteNames(name string, names []string) {
	if names != nil && len(names) > 0 {
		names[0] = name
	}
}
