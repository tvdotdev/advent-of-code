package register

// map [year] -> map [day] -> fund (string)
var Solutions = map[int]map[int]func(string){}

func Register(year int, day int, solve func(string)) {
	if Solutions[year] == nil {
		Solutions[year] = map[int]func(string){}
	}
	Solutions[year][day] = solve
}
