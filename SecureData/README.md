# Main Code

- isFeasible Function: This function checks if it's possible to distribute the fragments such that no data center exceeds the given maxRisk. It iterates over each data center and counts how many fragments can be stored without exceeding the maxRisk.

- distributeFragments Function: This function performs a binary search to find the minimum possible maxRisk. It starts with left as 1 and right as the maximum possible risk (calculated as the highest base risk raised to the power of the number of fragments). It narrows down the range by checking feasibility using the isFeasible function.

## Efficiency

- Binary Search: The binary search efficiently narrows down the range of possible maximum risks, significantly reducing the number of checks needed compared to a brute-force approach.
- Feasibility Check: The feasibility check ensures that we are only counting valid fragment allocations, making the algorithm faster.

# Test

- The `TestDistributeFragments` function defines a set of test cases. Each test case includes input data centers, the number of fragments, and the expected minimized maximum risk.