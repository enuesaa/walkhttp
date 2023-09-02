package resolver

type FileinfoResolver struct {
	path string
	name string
	description string
	isdir bool
	created string
	modified string
}

func (r *FileinfoResolver) Path() string {
	return r.path
}

func (r *FileinfoResolver) Name() string {
	return r.name
}

func (r *FileinfoResolver) Description() string {
	return r.description
}

func (r *FileinfoResolver) Isdir() bool {
	return r.isdir
}

func (r *FileinfoResolver) Created() string {
	return r.created
}

func (r *FileinfoResolver) Modified() string {
	return r.modified
}
