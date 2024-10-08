/*
Your task is to find the next higher number (int) with the same number of '1'- Bits.

I.e. as many 1 bits as before and output next higher than input. Input is always an int between 1 and 1<<30 (possibly inclusive). No bad cases or special tricks...

Some easy examples:
Input: 129  => Output: 130 (10000001 => 10000010)
Input: 127 => Output: 191 (01111111 => 10111111)
Input: 1 => Output: 2 (01 => 10)
Input: 323423 => Output: 323439 (1001110111101011111 => 1001110111101101111)
First some static tests, later on many random tests too;-)!

Hope you have fun! :-)
*/
package kata

func NextHigher(n int) int {
    c := n & -n
    r := n + c
    return ((((r ^ n) >> 2) / c) | r)
}