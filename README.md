# Cracking the coding interview

The current repo represents the solutions to some of the exercises, using Go.
## Chapter 1 - Arrays and Strings

1.1 [Is Unique](pkg/chapter_01/01_is_unique.go)             
: Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data 
structures.

1.2 [Check Permutation](pkg/chapter_01/02_check_permutation.go)    
: Given two strings, write a method to decide if one is a permutation of the other.

## Docker support
It has support for Docker, to use it do as following:

1. **Build** Run the docker compose tool to build the image
: ```bash
  docker compose -f deployments/docker-compose.yml build
```

2. **Run** Run the docker compose tool to run the cli utility:
: ```bash
   docker compose -f deployments/docker-compose.yml run cli cli IsUnique -entry uniqe
```
