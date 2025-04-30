package main

// Intersection: Given two straight line segments (represented as a start point and an end point),
// compute the point of intersection, if any.

// SOLUTION
// We first need to think about what it means for two line segments to intersect.
// For two infinite lines to intersect, they only have to have different slopes. If they have the same slope, then
// they must be the exact same line (same y-intercept). That is:
// slope 1 != slope 2
// OR
// slope 1 slope 2 AND intersect 1 == intersect 2
// For two straight lines to intersect, the condition above must be true, plus the point of intersection must be
// within the ranges of each line segment.
// extended infinite segments intersect
// AND
// intersection is within line segment 1 (x and y coordinates)
// AND
// intersection is within line segment 2 (x and y coordinates)
// What if the two segments represent the same infinite line? In this case, we have to ensure that some portion
// of their segments overlap. If we order the line segments by their x locations (start is before end, point 1 is
// before point 2), then an intersection occurs only if:
// Assume:
// start1.x < start2.x && start1.x < end1.x && start2.x < end2.x
// Then intersection occurs if:
// start2 is between startl and endl
// We can now go ahead and implement this algorithm.

type Point struct{
	x,y float64
}

type Line struct{
	slope,yintercept float64
}

func(l *Line)Line(start,end *Point){
	deltaY:=end.y-start.y
	deltaX:=end.x-start.x
	l.slope=deltaY/deltaX // Will be Infinity (not exception) when deltaX=0
	l.yintercept=end.y-l.slope*end.x
}

// Swap coordinates of point one and two.
func swap(one,two *Point){
	x:=one.x
	y :=one.y
	one.x = two.x
	one.y=two.y
	two.x=x
	two.y=y
}

// Checks if middle is between start and end.
func isBetweenPoint(start, middle,end *Point)bool{
	return isBetween(start.x, middle.x, end.x)&&
	isBetween(start.y, middle.y, end.y)
}

// Checks if middle is between start and end.
func isBetween(start,middle,end float64)bool{
	if start>end{
		return end<=middle&&middle<=start
	}else{
		return start<=middle&&middle<=end
	}
}

func intersection(start1,end1,start2,end2*Point)*Point{
	// Rearranging these so that, in order of x values: start is before end and
	// point 1 is before point 2. This will make some of the later logic simpler
	if start1.x>end1.x{
		swap(start1,end1)
	}
	if start2.x>end2.x{
		swap(start2,end2)
	}
	if start1.x>start2.x{
		swap(start1,start2)
		swap(end1,end2)
	}

	// Compute lines (including slope and y-intercept).
	line1 :=&Line{}
	line1.Line(start1,end1)
	line2:=&Line{}
	line2.Line(start2,end2)
	
	// If the lines are parallel, they intercept only if they have the same y
	// intercept and start 2 is on line 1

	if line1.slope == line2.slope{
		if line1.yintercept==line2.yintercept&&
		isBetweenPoint(start1,start2,end1){
			return start2
		}
		return nil
	}

	// Get intersection coordinate.
	x := (line2.yintercept-line1.yintercept)/(line1.slope-line2.slope)
	y := x*line1.slope+line1.yintercept

	intersection := &Point{x: x,y: y}

	// Check if within line segment range.
	if isBetweenPoint(start1,intersection,end1)&&
	isBetweenPoint(start2,intersection,end2){
		return intersection
	}
	return nil
}

// For simplicity and compactness (it really makes the code easier to read), we've chosen to make the variables
// within Point and Line public. You can discuss with your interviewer the advantages and disadvantages
// of this choice.
