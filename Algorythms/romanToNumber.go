// different map declerations
// METHOD 1
// symbols:= make(map[rune] int);
    
// symbols['I']=1;
// symbols['V']=5;
// symbols['X']=10;
// symbols['L']=50;
// symbols['C']=100;
// symbols['D']=500;
// symbols['M']=1000;


// METHOD 2
  
// var symbols= map[rune] int{
// 	'I': 1,
// 	'V': 5,
// 	'X': 10,
// 	'L': 50,
// 	'C': 100,
// 	'D': 500,
// 	'M': 1000,
// }



package main
import "fmt"

func main() {
  fmt.Println(romanToInt("IX"))
}

func romanToInt(s string) int {
    
    var symbols= map[rune] int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
    

   var sLength int=len(s);
    
    var total int;
    
    for i:=0; i< sLength; i++{
        
        var currentValue int= symbols[rune(s[i])];
        
        
        if i < (sLength-1){
            
            var nextValue int= symbols[rune( s[i+1] )];
            
            if(currentValue<nextValue){
                total-= currentValue + currentValue;
            }
            
        }
        
        total+=currentValue;
        
        // s[i] is represented as a byte to convert it into rune do rune(s[i]); 
        // fmt.Println(symbols[rune(s[i])]) // the individual char is passes as ASCCI
    }
    
    return total;
}