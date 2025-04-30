package main

// Diving Board: You are building a diving board by placing a bunch of planks of wood end-to-end.
// There are two types of planks, one of length shorter and one of length longer. You must use
// exactly K planks of wood. Write a method to generate all possible lengths for the diving board.

// SOLUTION
// One way to approach this is to think about the choices we make as we're building a diving board. This leads
// us to a recursive algorithm.
// Recursive Solution
// For a recursive solution, we can imagine ourselves building a diving board. We make K decisions, each time
// choosing which plank we will put on next. Once we've put on K planks, we have a complete diving board
// and we can add this to the list (assuming we haven't seen this length before).

// We can follow this logic to write recursive code. Note that we don't need to track the sequence of planks. All
// we need to know is the current length and the number of planks remaining.
// page 497
