/*
Given a string of words, you need to find the highest scoring word.

Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.

For example, the score of abad is 8 (1 + 2 + 1 + 4).

You need to return the highest scoring word as a string.

If two words score the same, return the word that appears earliest in the original string.

All letters will be lowercase and all inputs will be valid.


*/
package kata
import ("strings")
func High(s string) string {
    words:= strings.Split(s," ")
    highestword:= ""
    highestscore := 0
   for _, word := range words {
     score:=0
     for _, char := range word {
       score += int(char)- 'a'+1
     }
     if score > highestscore{
        highestscore = score
       highestword = word
     }
   }
  return highestword
    
}