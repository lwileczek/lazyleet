# Permutations
url: https://leetcode.com/problems/permutations/
difficulty: medium

## Details
Given some list of numbers, return all the permutations of that array.

## Notes
Turns out, it's the easier if you build up each permutation 
from scratch instead of trying to reorder the elements.

I could use DFS regardless of the length of the array here, but i chose
to return simplified versions if the length is 2 or fewer since they are extremely simple.
