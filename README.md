# Cracking the coding interview

The current repo represents the solutions to some of the exercises, using Go.

## Chapter 1 - Arrays and Strings

1.1 [Is Unique](pkg/chapter_01/01_is_unique.go)             
: Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data 
structures.

1.2 [Check Permutation](pkg/chapter_01/02_check_permutation.go)    
: Given two strings, write a method to decide if one is a permutation of the other.

1.3 [URLify](pkg/chapter_01/03_urlify.go)
: Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at 
the end to the hold the additional characters, and that you are given the "true" length of the string. (Note: If 
implementing in Java, please use a character array so that you can perform this operation in place.)

```
EXAMPLE
Input: "Mr John Smith    ", 13
Output: "Mr%20John%20Smith"
```

1.4 [Palindrome Permutation](pkg/chapter_01/04_palindromePermutation.go)
: PalindromePermutation Given a string, write a function to check if it is a permutation of a palindrome. A palindrome 
is a word of phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The 
palindrome does not need to be limited to just dictionary words.

```
EXAMPLE
Input: Tact Coa
Output: True (permutations: "taco cat", "atco cta", etc.)
```

1.5 [One Away](pkg/chapter_01/05_one_away.go)
: There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a 
character. Given two strings, write a function to check if they are one edit (or zero edits) away.

```
EXAMPLE
pale, ple -› true 
pales, pale -› true 
pale, bale -> true 
pale, bake -› false
```

1.6 [String Compression](pkg/chapter_01/06_string_compression.go)
: Implement a method to perform basic string compression using the counts of repeated characters. 
For example, the string aabccccaaa would become a2b1c5a3. If the "compressed" string would not become smaller than the 
original string, your method should return the original string. You can assume the string has only uppercase and 
lowercase letters (a - z). 

1.7 [Rotate Matrix](pkg/chapter_01/07_rotate_matrix.go)
: Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method to rotate the 
image by 90 degrees. Can you do this in place?
    
1.8 [Zero Matrix](pkg/chapter_01/08_zero_matrix.go)
: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.

## Docker support
It has support for Docker, to use it do as following:

1. **Build** the docker compose tool to build the image
: ```bash
  docker compose -f deployments/docker-compose.yml build
```

2. **Run** the docker compose tool to run the cli utility:
: ```bash
   docker compose -f deployments/docker-compose.yml run cli cli IsUnique -entry uniqe
```
