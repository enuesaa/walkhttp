package resolver

type FileinfoResolver struct {
	name string
	description string
}

func (r *FileinfoResolver) Name() string {
	return r.name
}

func (r *FileinfoResolver) Description() string {
	return r.description
}
