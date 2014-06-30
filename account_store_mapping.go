package stormpath

const (
	AccountStoreMappingBaseUrl = "https://api.stormpath.com/v1/accountStoreMappings"
)

type AccountStoreMapping struct {
	Href                  string `json:"href,omitempty"`
	ListIndex             *int   `json:"listIndex,omitempty"`
	IsDefaultAccountStore *bool  `json:"isDefaultAccountStore,omitempty"`
	IsDefaultGroupStore   *bool  `json:"isDefaultGroupStore,omitempty"`
	Application           Link   `json:"application"`
	AccountStore          Link   `json:"accountStore"`
}

func NewAccountStoreMapping(applicationHref string, accountStoreHref string) *AccountStoreMapping {
	return &AccountStoreMapping{Application: Link{Href: applicationHref}, AccountStore: Link{Href: accountStoreHref}}
}

func (mapping *AccountStoreMapping) Save() error {
	return Client.DoWithResult(&StormpathRequest{
		Method:  POST,
		URL:     AccountStoreMappingBaseUrl,
		Payload: mapping,
	}, mapping)
}