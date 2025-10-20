Go-Reloaded goal is to proccess the text("string") input that the user is giving 
in this specific case is text from a file.


1. What is the core of the problem
    1.1. The core of the problem is the edit rules that will or will not
        encounter inside the text.
    
    1.2. EDIT RULES:
            Here we group the "edit rules into 3 categories"

            1.2.1 Tags:
                Cases that starts and ends with a parenthesis containing key words like: "(hex),(bin),(up),(low),(cap)"

                hex: should replace the word before with the decimal version of the word

                bin:should replace the word before with the decimal version of the word.

                up:converts the word before with the Uppercase version of it.

                low:converts the word before with the Lowercase version of it.

                cap:converts the word before with the capitalized version of it.

                Also we have the diversion of (up,low,cap) can also be prosseced with a number inside the parenthesis indicating the number of words that needs to be modified.
            
            1.2.2 Punctuations:
                Cases like ('.' ',' '!' '?' ';')
                
                Should be close to the previous word and with space apart from the next one.

                Exception: 
                    if there are groups of punctuation like: ... or !?. In this case the program should format after realizing the size of the group implementing the same rules 
            
            1.2.3 punctuation Mark :
                We mean for single quotes (') when we encoter this symbol when want to find the closing one and after that to remove trailing sapce at the start and at the end of it 

                Exception:
                    All the other rules should applied inside the quatation space that include if inside the quotes are Tags or Punctuations or another pair of quotes
            
            1.2.4 from a to an :
                Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or an h
---------------------------------------------------------------------------------
    2. Examples:

        1. Tags:
            (HEX)
            "1E (hex) files were added" -> "30 files were added"
            (BIN)
            "It has been 10 (bin) years" -> "It has been 2 years"
            (UP)
            "Ready, set, go (up) !" -> "Ready, set, GO!"
            (LOW)
            "I should stop SHOUTING (low)" -> "I should stop shouting"
            (CAP)
            "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the  Brooklyn Bridge"
        
            OR
                
            (UP , LOW , CAP) -> * int
            "This is so exciting (up, 2)" -> "This is SO EXCITING"

        2.Punctuations:
            "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!"
            
            Exception:
                "I was thinking ... You were right" -> "I was thinking... You were right".
        
        3.Single Quotes:
            "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'"


        4.From a To an:
            "There it was. A amazing rock!" -> "There it was. An amazing rock!"


