# Binary Search and Graph Traversal Algorithms

This Go program implements binary search along with Depth-First Search (DFS) and Breadth-First Search (BFS) algorithms for graph traversal.

## How to Run the Code

To run this console app, follow these steps:

1. Clone this repository to your local machine:

```
git clone <repository_url>
```

2. Navigate to the directory containing the Go code:

```
cd algorithms
```

3. Ensure you have Go installed on your system. If not, you can download it from [here](https://golang.org/dl/).

4. Execute the Go code by running the following command:

```
go run main.go
```

## Binary Search
The binarySearch function implements the binary search algorithm to search for a target element in a sorted array.
Function Signature:
  - mas: The sorted array to be searched.
  - data: The target element to search for.
  - left: The left index of the sub-array to search.
  - right: The right index of the sub-array to search.
Returns the index of the target element if found; otherwise, returns -1.

## Depth-First Search (DFS)
The dfs function performs Depth-First Search traversal on a given adjacency matrix representation of a graph.
Function Signature:
  - matrix: The adjacency matrix representing the graph.
  - masIndex: An array indicating visited nodes.
  - point: The starting node for DFS traversal.
## Breadth-First Search (BFS)
The bfs function performs Breadth-First Search traversal on a given adjacency matrix representation of a graph.
Function Signature:
  - matrix: The adjacency matrix representing the graph.
  - masIndex: An array indicating visited nodes.
  - point: The starting node for BFS traversal.
## Conclusion

This Go program provides implementations of binary search and two graph traversal algorithms, DFS and BFS, offering solutions for various search and traversal tasks. Feel free to modify and extend the code as needed for your projects.

Feel free to reach out if you have any questions or encounter any issues while running the code.
