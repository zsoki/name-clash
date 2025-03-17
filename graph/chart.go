package graph

import (
	"bufio"
	"github.com/yaricom/goGraphML/graphml"
	"name-clash/names"
	"os"
	"strconv"
)

func ExportGraphML(namesInput []*names.Name, edgesInput []*names.Edge) {
	gml := graphml.NewGraphML("name-clash")

	graphAttr := map[string]interface{}{
		"default_weight": 1.0,
		"acyclic":        false,
		"max_depth":      10,
	}

	graph, _ := gml.AddGraph("name-clash", graphml.EdgeDirectionUndirected, graphAttr)

	nameToGraphMlNode := make(map[*names.Name]*graphml.Node)

	for _, name := range namesInput {
		//if len(name.Edges) == 0 {
		//	continue
		//}
		nodeAttrs := map[string]interface{}{
			"Name": name.Text,
		}
		node, _ := graph.AddNode(nodeAttrs, name.Text)
		nameToGraphMlNode[name] = node
	}

	for _, edge := range edgesInput {
		edgeAttrs := map[string]interface{}{
			"Distance": edge.Distance,
		}
		n1 := nameToGraphMlNode[edge.Names[0]]
		n2 := nameToGraphMlNode[edge.Names[1]]

		_, _ = graph.AddEdge(n1, n2, edgeAttrs, graphml.EdgeDirectionDefault, strconv.Itoa(edge.Distance))
	}

	output, _ := os.Create("output.graphml")
	writer := bufio.NewWriter(output)

	err := gml.Encode(writer, true)
	if err != nil {
		panic(err)
	}

}
