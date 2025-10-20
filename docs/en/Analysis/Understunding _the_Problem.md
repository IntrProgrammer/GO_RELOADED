## **Project Goal: Text Formatter**

Our project is to build a program that acts like a "text cleaner." It will read text from a file, look for specific patterns and commands, and then apply a set of rules to format the text correctly. Think of it as an automated proofreader that follows a very specific style of rules and formats the text.

The core of the challenge is correctly identifying and applying these editing rules. Let's break them down.

---

### **Rule 1: Special Commands (Tags)**

Sometimes, the text will contain special commands enclosed in parentheses, which we'll call **tags**. These tags tell our program to modify the word (or words) that came just before them.

* **Number Conversion Tags:**
    * `(hex)`: Finds the word before it and converts it from a **hexadecimal** number (base-16) to a regular **decimal** number (base-10).
        * **Before:** `1E (hex) files were added`
        * **After:** `30 files were added`
    * `(bin)`: Finds the word before it and converts it from a **binary** number (base-2) to a **decimal** number.
        * **Before:** `It has been 10 (bin) years`
        * **After:** `It has been 2 years`

* **Text Case Tags:**
    * `(up)`: Changes the word before it to **UPPERCASE**.
        * **Before:** `Ready, set, go (up) !`
        * **After:** `Ready, set, GO!`
    * `(low)`: Changes the word before it to **lowercase**.
        * **Before:** `I should stop SHOUTING (low)`
        * **After:** `I should stop shouting`
    * `(cap)`: **Capitalizes** the word before it (makes the first letter uppercase and the rest lowercase).
        * **Before:** `Welcome to the brooklyn bridge (cap)`
        * **After:** `Welcome to the Brooklyn Bridge`

* **Multi-Word Tags:**
    The case-changing tags (`up`, `low`, `cap`) can also include a number. This tells our program **how many words** before the tag to modify.
    * **Example:** `(up, 2)` means "change the 2 words before this tag to uppercase."
        * **Before:** `This is so exciting (up, 2)`
        * **After:** `This is SO EXCITING`

---

### **Rule 2: Punctuation Spacing**

We need to make sure punctuation marks like periods (`.`), commas (`,`), exclamation marks (`!`), question marks (`?`), and semicolons (`;`) are spaced correctly.

* **The Rule:** Punctuation should be attached to the end of the previous word (no space) and have exactly one space after it before the next word begins.
    * **Before:** `I was sitting over there ,and then BAMM !!`
    * **After:** `I was sitting over there, and then BAMM!!`

* **Special Case (Groups of Punctuation):**
    Sometimes punctuation comes in groups, like `...` or `!?.` Our program should treat the entire group as a single unit and apply the same rule: attach the whole group to the previous word and put one space after it.
    * **Before:** `I was thinking ... You were right`
    * **After:** `I was thinking... You were right`

---

### **Rule 3: Single Quotes (' ')**

When our program finds text wrapped in single quotes, it needs to clean up the spacing inside them.

* **The Rule:** Remove any extra spaces immediately after the opening quote (`'`) and immediately before the closing quote (`'`).
    * **Before:** `As Elton John said: ' I am the most well-known homosexual in the world '`
    * **After:** `As Elton John said: 'I am the most well-known homosexual in the world'`

* **Important Note:** All the other rules (like tags and punctuation spacing) should still be applied to the text *inside* the single quotes!

---

### **Rule 4: The "a" to "an" Transformation**

This is a classic English grammar rule we need to automate.

* **The Rule:** The word "**a**" should be changed to "**an**" if the very next word starts with a vowel (`a`, `e`, `i`, `o`, `u`) or the letter `h`.
    * **Before:** `There it was. A amazing rock!`
    * **After:** `There it was. An amazing rock!`

By tackling each of these rules, we can build a program that correctly reformats any text we give it.