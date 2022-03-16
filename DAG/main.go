
package main

type Node struct {
    V     string
    Edges []*Node
}

func lengthOfLongestSubstring(s string) int {
    var head *Node
    var node *Node
    nodes := make(map[string]*Node)
    for _, v := range s {
        if head == nil {
            head = &Node{V: v, Edges: make([]*Node, 0)}
            nodes[v] = head
            node = head 
        } else {
            n, ok := nodes[v]
            if !ok {
                n = &Node{V: v, Edges: make([]*Node, 0)} 
                nodes[v] = n
            }
            node.Edges = append(node.Edges, n)
            node = n
        }
    }
    
}
