package main

// Build Order: You are given a list of projects and a list of dependencies (which is a list of pairs of
// projects, where the second project is dependent on the first project). All of a project's dependencies
// must be built before the project is. Find a build order that will allow the projects to be built. If there
// is no valid build order, return an error.
// EXAMPLE
// Input:
// projects: a, b, c, d, e, f
// dependencies: (a, d), (f, b), (b, d), (f, a), (d, c)
// Output: f, e, a, b, d, c

// SOLUTION
// Visualizing the information as a graph probably works best. Be careful with the direction of the arrows. In
// the graph below, an arrow from d to g means that d must be compiled before g. You can also draw them
// in the opposite direction, but you need to consistent and clear about what you mean

// In drawing this example (which is not the example from the problem description), I looked for a few things.
// I wanted the nodes labeled somewhat randomly. If I had instead put a at the top, with b and c as chil­
// dren, then d and e, it could be misleading. The alphabetical order would match the compile order.
// I wanted a graph with multiple parts/components, since a connected graph is a bit of a special case.
// I wanted a graph where a node links to a node that cannot immediately follow it. For example, f links to
// a but a cannot immediately follow it (since b and c must come before a and after f).
// I wanted a larger graph since I need to figure out the pattern.
// I wanted nodes with multiple dependencies.
// Now that we have a good example, let's get started with an algorithm.

// Solution#l
// Where do we start? Are there any nodes that we can definitely compile immediately?
// Yes. Nodes with no incoming edges can be built immediately since they don't depend on anything. Let's
// add all such nodes to the build order. In the earlier example, this means we have an order of f, d (or d, f).
// Once we've done that, it's irrelevant that some nodes are dependent on d and f since d and f have already
// been built. We can reflect this new state by removing d and f's outgoing edges.
// build order: f, d

// Next, we know that c , b, and g are free to build since they have no incoming edges. Let's build those and
// then remove their outgoing edges.
// build order: f, d, c, b, g

// Project a can be built next, so let's do that and remove its outgoing edges. This leaves just e. We build that
// next, giving us a complete build order.
// build order: f, d, c, b, g, a, e

// Did this algorithm work, or did we just get lucky? Let's think about the logic.
// 1. We first added the nodes with no incoming edges. If the set of projects can be built, there must be some
// "first" project, and that project can't have any dependencies. If a project has no dependencies (incoming
// edges), then we certainly can't break anything by building it first.
// 2. We removed all outgoing edges from these roots. This is reasonable. Once those root projects were built,
// it doesn't matter if another project depends on them.
// 3. After that, we found the nodes that now have no incoming edges. Using the same logic from steps 1 and
// 2, it's okay if we build these. Now we just repeat the same steps: find the nodes with no dependencies,
// add them to the build order, remove their outgoing edges, and repeat.
// 4. What if there are nodes remaining, but all have dependencies (incoming edges)? This means there's no
// way to build the system. We should return an error.
// The implementation follows this approach very closely.

// Initialization and setup:
// 1. Build a graph where each project is a node and its outgoing edges represent the projects that depend
// on it. That is, if A has an edge to B (A-> B), it means B has a dependency on A and therefore A must be
// built before B. Each node also tracks the number of incoming edges.
// 2. Initialize a buildOrder array. Once we determine a project's build order, we add it to the array. We also
// continue to iterate through the array, using a toBeProcessed pointer to point to the next node to be
// fully processed

// 3. Find all the nodes with zero incoming edges and add those to a buildOrder array. Set a
// toBeProcessed pointer to the beginning of the array.
// Repeat until toBeProcessed is at the end of the buildOrder:
// 1. Read node at toBeProcessed.
// » If node is null, then all remaining nodes have a dependency and we have detected a cycle.
// 2. For each child of node:
// »Decrement child. dependencies (the number of incoming edges).
// »If child. dependencies is zero, add child to end of buildOrder.
// 3. Increment toBeProcessed.
// The code below implements this algorithm.

type Project struct {
	childern     []*Project
	projectMap   map[string]*Project
	name         string
	dependencies int
}
func (p *Project)addNeighbor(node *Project){
	_, exists :=p.projectMap[node.name]
	if !exists{
		p.childern = append(p.childern, node)
		p.projectMap[node.name]=node
		node.dependencies++
	} 
}

type Graph struct{
	nodes []*Project
	nodesSize int
	projectMap map[string]*Project
}

func (g *Graph)getOrCreateNode(name string)*Project{
	node,exists := g.projectMap[name]
	if exists{
		return node
	}
	node = &Project{name: name}
	g.nodes=append(g.nodes, node)
	g.nodesSize++
	g.projectMap[name]=node
	return node
}

func (g*Graph)addEdge(startName,endName string){
	start := g.getOrCreateNode(startName)
	end := g.getOrCreateNode(endName)
	start.addNeighbor(end)
}

/* A helper function to insert projects with zero dependencies into the order
* array, starting at index offset. */
func addNonDependent(order []*Project,projects []*Project, offset int)int{
	for _, project := range projects{
		if project.dependencies== 0{
			order[offset]=project
			offset++
		}
	}
	return offset
}

// Return a list of the projects a correct build order.
func orderProjects(projects []*Project)[]*Project{
	order := make([]*Project, len(projects))
	// Add "roots" to the build order first.
	endOfList := addNonDependent(order,projects,0)

	toBeProcessed :=0
	for toBeProcessed<len(order){
		current := order[toBeProcessed]
		/* We have a circular dependency since there are no remaining projects with
		* zero dependencies. */
		if current == nil{
			return nil
		}
		// Remove myself as a dependency.
		children := current.childern
		for _,child := range children{
			child.dependencies--
		}
		// Add children that have no one depending on them.
		endOfList = addNonDependent(order,children,endOfList)
		toBeProcessed++
	}
	return order
}

//  Build the graph, adding the edge (a, b) if b is dependent on a. Assumes a pair
// is listed in "build order". The pair (a, b) in dependencies indicates that b
// depends on a and a must be built before b. *
func buildGraph(projects []string, dependencies [][]string)*Graph{
	graph := &Graph{}
	for _, project := range projects{
		graph.getOrCreateNode(project)
	}

	for _, dependency := range dependencies{
		first := dependency[0]
		second := dependency[1]
		graph.addEdge(first,second)
	}

	return graph
}

// Find a correct build order.
func findBuildOrder(project []string, dependencies [][]string)[]*Project{
	graph := buildGraph(project, dependencies)
	return orderProjects(graph.nodes)
}

// This solution takes O ( P + D) time, where P is the number of projects and D is the number of dependency
// pairs.

// Note: You might recognize this as the topological sort algorithm on page 632. We've rederived
// this from scratch. Most people won't know this algorithm and it's reasonable for an interviewer
// to expect you to be able to derive it.

// to be continued at page 265
