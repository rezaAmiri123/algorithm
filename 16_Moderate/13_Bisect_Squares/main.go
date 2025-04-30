package main

// Bisect Squares: Given two squares on a two-dimensional plane, find a line that would cut these two
// squares in half. Assume that the top and the bottom sides of the square run parallel to the x-axis.

// SOLUTION
// Before we start, we should think about what exactly this problem means by a "line:' Is a line defined by a
// slope and a y-intercept? Or by any two points on the line? Or, should the line be really a line segment, which
// starts and ends at the edges of the squares?
// We will assume, since it makes the problem a bit more interesting, that we mean the third option: that the
// line should end at the edges of the squares. In an interview situation, you should discuss this with your
// interviewer.
// ::=:� .
// This line that cuts two squares in half must connect the two middles. We can easily calculate the slope,
// knowing that slope =(y1-y2)/(x1-x2)
// Once we calculate the slope using the two middles, we can use the same
// equation to calculate the start and end points of the line segment.
// In the below code, we will assume the origin ( 0, 0) is in the upper left-hand corner.

type Square struct{
	
}