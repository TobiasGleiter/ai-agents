package main

import ()

func main() {
	// LATs (https://arxiv.org/abs/2310.04406)

	// 1. Starting at root node either an answer or a tool execution
	// 2. Reflection LLM generates a reflection on the output from 1. and creates a score (to determine if a solution is found)
	// - {reflection, score, found_solution}
	// 3. n candiates are generated with the context and the prior output and reflexion (expand the tree)
	// 4. Select the Node with the best score
	// Start again with 1. step until max depth is reached.

}