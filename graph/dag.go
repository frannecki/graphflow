package graph

type DAG struct {
	Vertices []string
	Edges map[string][]string
}

func NewDAG() *DAG {
	return &DAG{Edges: make(map[string][]string)}
}

func (graph *DAG) Prune(endpoints ...string) *DAG {
	// TODO
	return graph
}

func (graph *DAG) Select(endpoints ...string) *DAG {
	// TODO
	return graph
}

func (graph *DAG) IsValid() bool {
	if graph == nil {
		return false
	}
	// Check for duplicates
	vertexMap := make(map[string]struct{})
	for _, vertex := range graph.Vertices {
		if _, ok := vertexMap[vertex]; ok {
			return false
		}
		vertexMap[vertex] = struct{}{}
	}

	// Check for cycles
	_, sorted := graph.topologicalSort()
	return len(sorted) == len(graph.Vertices)
}

func (graph *DAG) topologicalSort() (layers [][]string, sorted []string) {
	var (
		dependencyMap = make(map[string][]string)
		inDegreeMap = make(map[string]int)
	)
	for k, vs := range graph.Edges {
		for _, v := range vs {
			dependencyMap[v] = append(dependencyMap[v], k)
		}
	}
	for k, vs := range dependencyMap {
		inDegreeMap[k] = len(vs)
	}

	// topological sort
	entries := make([]string, 0, len(graph.Vertices))
	for _, vertex := range graph.Vertices {
		if inDegreeMap[vertex] == 0 {
			entries = append(entries, vertex)
		}
	}

	var layer, nextLayer []string
	layer = entries
	for len(layer) > 0 {
		sorted = append(sorted, layer...)
		layers = append(layers, layer)
		for _, vertex := range layer {
			for _, v := range  graph.Edges[vertex] {
				inDegreeMap[v] -= 1
				if inDegreeMap[v] > 0 {
					continue
				}
				nextLayer = append(nextLayer, v)
			}
		}
		layer = nextLayer
		nextLayer = nil
	}

	return
}
