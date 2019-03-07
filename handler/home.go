package handler

import (
	"net/http"
)

// Home 首页
func Home(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("action") {
	case "user":
		data := []struct {
			Y     string  `json:"y"`
			Item1 float64 `json:"item1"`
			Item2 float64 `json:"item2"`
		}{
			{"2011 Q1", 2666, 2666},
			{"2011 Q2", 2778, 2294},
			{"2011 Q3", 4912, 1969},
			{"2011 Q4", 3767, 3597},
			{"2012 Q1", 6810, 1914},
			{"2012 Q2", 5670, 4293},
			{"2012 Q3", 4820, 3795},
			{"2012 Q4", 15073, 5967},
			{"2013 Q1", 10687, 4460},
			{"2013 Q2", 8432, 5713},
		}
		jSuccess(w, data)
	default:
		rLayout(w, r, "index.tpl", nil)
	}
}
