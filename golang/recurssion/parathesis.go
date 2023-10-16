package main

func generateParenthesis(n int) []string {
	res := []string{}
	stack := ""
	backtrack(n, 0, 0, &res, stack)
	return res
}

func backtrack(n, openN, closedN int, res *[]string, stack string) {

	if openN == n && closedN == n{
		(*res) = append((*res), stack)
		return 
	}


	if openN < n {
		stack += "("
		backtrack(n, openN+1, closedN, res, stack)
		stack = stack[:len(stack)-1]
	}

	if closedN < openN {
		stack += ")"
		backtrack(n, openN, closedN+1, res, stack)
	}

	

}