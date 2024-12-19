// Find the smallest possitive integer not in a give array
// The Time Complexity O(n) and Space Complexity O(n)

// Actually this is the first time I encounted this kind
// So this is the work after asking ChatGPT
// prompt1 : give a list of integer, find the smalless positive integer not in the list without sort
// prompt2 : index-based placement strategy
// prompt3: explain: For an array of size n, the possible smallest missing positive integers are in the range [1, n+1]

package main

var (
	givenArr = []int { -5, 7, 1, 3} // expected result is 2
)

func findSmallestPossitiveInteger(arr []int) int {
	n := len(arr)

	// Place numbers in their correct index positions
	for i := 0; i < n; i++ {
		for arr[i] >= 1 && arr[i] <= n && arr[arr[i]-1] != arr[i] {
			// Swap arr[i] with arr[arr[i]-1]
			arr[arr[i]-1], arr[i] = arr[i], arr[arr[i]-1]
		}
	}

	// Find the first missing positive integer
	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}

	// If all are in place, return n + 1
	return n + 1
}