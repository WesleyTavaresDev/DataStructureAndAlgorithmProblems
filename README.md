# Data Strucures and Algorithms

## List and Strings

1. Write a function that returns the largest element in a list.

```go

input: arr := [3]int{1, 3, 2}

output: 3

```

***

2. Write function that reverses a list, preferably in place.

```go

input: arr := [3]int{5, 2, 4}

output: [4, 2, 5]

```

***

3. Write a function that checks whether an element occurs in a list.

```go

inputs: arr := []int{10, 20, 30, 40, 50, 65, 3}, target := 30

output: true

```

***

4. Write a function that computes the running total of a list.

```go

inputs: arr := []int{2, 4, 6, 8, 20}

output: [2, 6, 12, 20, 40]


```

**Check it out-> <https://visualgo.net/en>**

## Sorting

1. Selection Sort

        Iterate a list of elements separing it on two portions: sorted and unsorted. 
        Traversing, comparing and swapping the elements

<p align="center"><img src="https://github.com/WesleyTavaresDev/DataStructureAndAlgorithmProblems/blob/main/assets/Selection-Sort-example.gif" align="center"/></p>

***

2. Bubble Sort

        Bubble sort is a sorting algorithm that compares two adjacent elements and 
        swaps them until they are in the intended order.
       
<p align="center"><img src="https://github.com/WesleyTavaresDev/DataStructureAndAlgorithmProblems/blob/main/assets/Bubble-sort-example.gif"/>
</p>

***

3. Insertion Sort

        Insertion sort  works similar to the way you sort playing cards in your hands. 
        The array is splited into a sorted and an unsorted part. Values from the unsorted part 
        are picked and placed at the correct position in the sorted part.
        
<p align="center"><img src="https://github.com/WesleyTavaresDev/DataStructureAndAlgorithmProblems/blob/main/assets/Insertion-sort-example.gif"/></p>

***

4. Merge Sort 

        Merge sort virtually breaks an array recursively in halves
        until it cannot be further divided, then Merge.
        Merge takes two small pieces to our array sorting and merging them until make a large one.
        
<p align="center"><img src="https://github.com/WesleyTavaresDev/DataStructureAndAlgorithmProblems/blob/main/assets/Merge-sort-example.gif"/></p>        

***

4. Quick Sort

        QuickSort is a Divide and Conquer algorithm. 
        It picks an element as a pivot and partitions the given array around the picked pivot.
