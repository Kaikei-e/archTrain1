package authAnatomy

type HashedPass struct {
	Type string `json:"type",gqlgen:"type",required:"true"`
	Data []byte `json:"data",gqlgen:"data",required:"true"`
}
