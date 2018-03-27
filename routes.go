package main

type Routes []Route

type Route struct {
	URI    string
	Type   string
	Params []string
}

func (r *Routes) SetPostRoutes() {
	*r = append(*r, Route{
		URI:  "/foo",
		Type: "POST",
		Params: []string{
			"id",
		},
	})
}

func (r *Routes) SetGetRoutes() {
	*r = append(*r, Route{
		URI:  "/bar",
		Type: "GET",
		Params: []string{
			"id",
		},
	})
}
