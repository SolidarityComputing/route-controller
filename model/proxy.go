package model

type RoutingBody struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Rule string `json:"rule"`
}

type AddRoutingBody struct {
	RoutingList []*RoutingBody `json:"rules"`
}

type DelRoutingBody struct {
	RoutingList []string `json:"rules"`
}
