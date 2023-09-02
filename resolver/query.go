package resolver

type QueryResolver struct {}

func (r *QueryResolver) Fileinfo(args struct{ Name string }) *FileinfoResolver {
	// run os
	if args.Name == "aa" {
		return &FileinfoResolver{ name: "aa", description: "aa-description" }
	}
	return nil
}
