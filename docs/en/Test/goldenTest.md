### Rule 1: Special Commands (Tags)
The text may contain special commands in parentheses, called tags, which modify the preceding word or words.

#### Number Conversion Tags
- `(hex)`: Converts the preceding hexadecimal number (base-16) to a decimal number (base-10).
   - **Before:** 1E (hex) files were added
   - **After:** 30 files were added
- `(bin)`: Converts the preceding binary number (base-2) to a decimal number.
   - **Before:** It has been 10 (bin) years
   - **After:** It has been 2 years

#### Text Case Tags
- `(up)`: Changes the preceding word to UPPERCASE.
   - **Before:** Ready, set, go (up) !
   - **After:** Ready, set, GO!
- `(low)`: Changes the preceding word to lowercase.
   - **Before:** I should stop SHOUTING (low)
   - **After:** I should stop shouting
- `(cap)`: Capitalizes the preceding word.
   - **Before:** Welcome to the brooklyn bridge (cap)
   - **After:** Welcome to the Brooklyn Bridge

#### Multi-Word Tags
The case-changing tags (up, low, cap) can specify a number of words to modify.
- **Example:** `(up, 2)` applies the uppercase rule to the two preceding words.
   - **Before:** This is so exciting (up, 2)
   - **After:** This is SO EXCITING

### Rule 2: Punctuation Spacing
Ensure correct spacing for punctuation marks (., ,, !, ?, ;).
- **The Rule:** Punctuation must be attached to the end of the preceding word with no space, followed by a single space.
   - **Before:** I was sitting over there ,and then BAMM !!
   - **After:** I was sitting over there, and then BAMM!!

Special Case (Groups of Punctuation): Treat grouped punctuation (e.g., ..., !?) as a single unit, applying the same spacing rule.
   - **Before:** I was thinking ... You were right
   - **After:** I was thinking... You were right

### Rule 3: Single Quotes (' ')
Clean up spacing within single quotes.
- **The Rule:** Remove any spaces immediately following the opening quote (') and preceding the closing quote (').
   - **Before:** As Elton John said: ' I am the most well-known homosexual in the world '
   - **After:** As Elton John said: 'I am the most well-known homosexual in the world'

### Rule 4: The "a" to "an" Transformation
Automate the "a" vs. "an" grammar rule.
- **The Rule:** Change "a" to "an" if the next word begins with a vowel (a, e, i, o, u) or the letter h.
   - **Before:** There it was. A amazing rock!
   - **After:** There it was. An amazing rock!

### Tricky Tests & Edge Cases
A robust formatter must handle complex scenarios where rules interact.
- **Combining Multiple Rules at Once:**
   - **Before:** ' The word is awesome (up, 2) ... ' said the programmer
   - **After:** 'The word is AWESOME...' said the programmer
  
- **Invalid Tag Usage:**
   - **Before:** The number is 1F (bin) which is wrong.
   - **After:** The number is 1F (bin) which is wrong.

- **"a" to "an" with Tags in Between:**
   - **Before:** It was a apple (low) a day
   - **After:** It was an apple a day

- **Punctuation Attached to Tags:**
   - **Before:** This is the final word (up).
   - **After:** This is the final WORD.

- **Nested and Sequential Quotes:**
   - **Before:** He said ' hello ' and she replied ' HI (low) ' .
   - **After:** He said 'hello' and she replied 'hi'.

- **Multi-word Tag Exceeding Word Count:**
   - **Before:** One two three (up, 5).
   - **After:** ONE TWO THREE.


### Multi Rule Test
- **Test case that combines rules 1, 2, 3, and 4:**

    - **Before:** It was a 111 (bin) hour wait... ' so I read a book (up) ' .

    - **After:** It was an 7 hour wait... 'so I read a BOOK'.

### Complex Multi-Rule Scenarios

**Advanced Test Cases:**

1. **Hexadecimal conversion with punctuation and quotes:**
   - **Before:** ' The result was A (hex) , which surprised everyone ' !
   - **After:** 'The result was 10, which surprised everyone'!

2. **Binary conversion with multi-word capitalization and article correction:**
   - **Before:** It took a 1010 (bin) amazing developers (cap, 2) to finish .
   - **After:** It took an 10 Amazing Developers to finish.

3. **Nested quotes with case changes and punctuation:**
   - **Before:** She whispered ' he said HELLO (low) to me ' , then smiled ...
   - **After:** She whispered 'he said hello to me', then smiled...

4. **Multiple conversions with grouped punctuation:**
   - **Before:** The values were 1A (hex) and 101 (bin) respectively !!!
   - **After:** The values were 26 and 5 respectively!!!

5. **Complex article correction with case modification:**
   - **Before:** A elephant (up) walked to a orange (cap) tree .
   - **After:** An ELEPHANT walked to an Orange tree.

6. **All rules combined in one sentence:**
   - **Before:** ' It was a FF (hex) hour journey (up, 2) , and we saw a eagle flying overhead ' !!!
   - **After:** 'It was an 255 HOUR JOURNEY, and we saw an eagle flying overhead'!!!

7. **Sequential tag applications with punctuation:**
   - **Before:** The binary 1111 (bin) became FIFTEEN (low) , then fifteen (up) again .
   - **After:** The binary 15 became fifteen, then FIFTEEN again.

8. **Edge case with invalid conversion and valid formatting:**
   - **Before:** ' The code XYZ (hex) failed , but ABC (hex) worked perfectly (cap) ' .
   - **After:** 'The code XYZ (hex) failed, but 2748 worked Perfectly'.

### Ultimate Multi-Rule Test

**The Complete Challenge - All Rules in One Paragraph:**

- **Before:** ' Yesterday , a engineer told me that the project took 1A (hex) intensive weeks (up, 2) to complete . He said the team converted 1101 (bin) binary files and worked on a apple (low) optimization algorithm . The final result was AMAZING (cap) , and everyone was shouting HOORAY (low) !!! He whispered ' it was a honor to work on this ' , then smiled ... '

- **After:** 'Yesterday, an engineer told me that the project took 26 INTENSIVE WEEKS to complete. He said the team converted 13 binary files and worked on an apple optimization algorithm. The final result was Amazing, and everyone was shouting hooray!!! He whispered 'it was an honor to work on this', then smiled...'