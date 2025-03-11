package main

type People struct {
	Name string
	Age  int
}

func AddNames(names []string, peoples []People) []string {
	index := len(names)

	names = append(names, make([]string, len(peoples))...)

	for _, people := range peoples {
		names[index] = people.Name
		index++
	}
	return names
}

func main() {

}
