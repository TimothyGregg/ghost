package field

import (
	ghost "github.com/TimothyGregg/ghost/graph"
)

type field struct {
	backing_graph ghost.Graph
}

type node struct {
	vertex ghost.Vertex
}

type connection struct {
	edge ghost.Edge
}