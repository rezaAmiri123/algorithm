package main

// Operations: Write methods to implement the multiply, subtract, and divide operations for integers.
// The results of all of these are integers. Use only the add operator.

// SOLUTION
// The only operation we have to work with is the add operator. In each of these problems, it's useful to think
// in depth about what these operations really do or how to phrase them in terms of other operations (either
// add or operations we've already completed).

// Subtraction
// How can we phrase subtraction in terms of addition? This one is pretty straightforward. The operation a
// - b is the same thing as a + ( -1) * b. However, because we are not allowed to use the * (multiply)
// operator, we must implement a negate function.

// Flip a positive sign to negative or negative sign to pos.
func negate(a int) int {
	neg := 0
	newSign := 0
	if a < 0 {
		newSign = -1
	}
	for a != 0 {
		neg += newSign
		a += newSign
	}
	return neg
}

// Subtract two numbers by negating b and adding them
func minus(a, b int) int {
	return a + negate(b)
}

// The negation of the value k is implemented by adding -1 k times. Observe that this will takeO( k)time.
// If optimizing is something we value here, we can try to get a to zero faster. (For this explanation, we'll
// assume that a is positive.) To do this, we can first reduce a by 1, then 2, then 4, then 8, and so on. We'll call
// this value de 1 ta. We want a to reach exactly zero. When reducing a by the next de 1 ta would change the
// sign of a, we reset delta back to 1 and repeat the process.
// For example:
// a: 29 28 26 22 14 13 11 7 6 4 0
// delta: -1 -2 -4 -8 -1 -2 -4 -1 -2 -4
// The code below implements this algorithm.

func negate2(a int)int{
	neg := 0
	newSign := 1
	if a<0{
		newSign = -1
	}
	delta := newSign
	for a!=0{
		differentSigns := (a+delta>0)!= (a>0)
		if (a+delta != 0&&differentSigns){//If delta is too big, reset it.
			delta=newSign
		}
		neg+=delta
		a+=delta
		delta+=delta //Double the delta
	}
	return neg
}

// Figuring out the runtime here takes a bit of calculation.
// Observe that reducing a by half takesO(log a)work. Why? For each round of"reduce a by half'; the abso­
// lute values of a and delta always add up to the same number. The values of delta and a will converge at
// Since de 1 ta is being doubled each time, it will takeO(log a)steps to reach half of a.
// Yi.
// We do O(log a) rounds.
// 1. Reducing a to Yi takes 0(log a)time.
// 2. Reducing Yi to % takesO(log Yi) time.
// 3. Reducing % to X takes 0(log % ) time.
// ... As so on ,forO(log a)rounds.
// The runtime therefore isO(log a + log(Yi) +log(%)+ ... ),withO(log a)terms in the
// expression.
// Recall two rules of logs:
// ,
// log(xy) = log x + log y
// • log( 7Y) = log x - log y.
// If we apply this to the above expression, we get:
// l. O(log a +log( Yi) + log( X) + •.. )
// 2. O(log a+ (log a - log 2) +(log a - log 4) +(log a - log 8) +...
// 3. o((log a)*(log a) - (log 2 +log 4 + log 8 + ... +log a)) //O(log a) terms
// 4. O((log a)*(log a) - (1 +2 + 3 + ... +log a)) //computing the values of logs
// 5. 0((log a)*(log a) - {loga)(l + log a/; ) //apply equation for sum of 1 through k
// 6. O((log a) 2 ) //drop second term from step 5
// Therefore, the runtime is O((log a)2).

// This math is considerably more complicated than most people would be able to do (or expected to do) in an
// interview. You could make a simplification: You do O(log a) rounds and the longest round takes O(log
// a) work. Therefore, as an upper bound, negate takes 0((log a) 2 ) time. In this case, the upper bound
// happens to be the true time.
// There are some faster solutions too. For example, rather than resetting delta to 1 at each round, we could
// change delta to its previous value. This would have the effect of delta "counting up" by multiples of two,
// and then "counting down" by multiples of two. The runtime of this approach would be O(log a). However,
// this implementation would require a stack, division, or bit shifting-any of which might violate the spirit of
// the problem. You could certainly discuss those implementations with your interviewer though.

// Multiplication
// The connection between addition and multiplication is equally straightforward. To multiply a by b, we just
// add a to itself b times.
// Multiply a by b by adding a to itself b times

// Multiply a by b by adding a to itself b times
func multiply(a,b int)int{
	if a<b{
		return multiply(b,a) //algorithm is faster if b< a
	}
	sum :=0
	for i:=abs(b);i<0;i=minus(i,1){
		sum+=a
	}
	if b<0{
		sum=negate2(sum)
	}
	return sum
}

// Return absolute value
func abs(a int)int{
	if a<0{
		return negate2(a)
	}else{
		return a
	}
}

// The one thing we need to be careful of in the above code is to properly handle multiplication of negative
// numbers. If b is negative, we need to flip the value of sum. So, what this code really does is:
// multiply(a, b)<-- abs(b)*a*(-1 if b< 0)
// We also implemented a simple abs function to help.

// page 492
