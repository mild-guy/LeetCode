/*
905. Sort Array By Parity
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
You may return any answer array that satisfies this condition.
Example 1:

Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 
Note:

1 <= A.length <= 5000
0 <= A[i] <= 5000
*/
/**
 * @param {number[]} A
 * @return {number[]}
 */
var sortArrayByParity = function(A) {
    var i=0, j=A.length-1;
    
    while( i < j ){
        if(A[i]%2!=0 && A[j]%2!=1) A[i]=A[j]+(A[j]=A[i])-A[i];
        if(A[i]%2 == 0) i++;
        if(A[j]%2 == 1) j--;   
    }
    return A;
};
/*
result:
Runtime: 72 ms, faster than 91% of JavaScript online submissions for Sort Array By Parity.
Memory Usage: 37.3 MB, less than 20.00% of JavaScript online submissions for Sort Array By Parity.

*/
