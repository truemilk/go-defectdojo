package defectdojo

type DojoGroupsService struct {
	client *Client
}

type DojoGroup struct {
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Users       []*int  `json:"users,omitempty"`
}

type DojoGroups struct {
	Count    *int         `json:"count,omitempty"`
	Next     *string      `json:"next,omitempty"`
	Previous *string      `json:"previous,omitempty"`
	Results  []*DojoGroup `json:"results,omitempty"`
}
