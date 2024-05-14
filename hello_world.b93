 >25*"!dlrow ,olleH":v
                  v:,_@
                  >  ^

This program is stolen from https://en.wikipedia.org/wiki/Befunge
and will print "Hello, world!" to standard output. Here is a full
explaination of the program (which is copied from wikipedia):

It adds the ASCII character 10 (a line feed character) to the stack,
and then pushes "!dlrow ,olleH" to the stack. Again, LIFO ordering
means that "H" is now the top of the stack and will be the first
printed, "e" is second, and so on. To print the characters, the
program enters a loop that first duplicates the top value on the
stack (so now the stack would look like "\n!dlrow ,olleHH"). Then
the "_" operation will pop the duplicated value, and go right if
it's a zero, left otherwise. (This assumes a compliant interpreter
that "returns" 0 when popping an empty stack.) When it goes left,
it pops and prints the top value as an ASCII character. It then
duplicates the next character and loops back to the "_" test,
continuing to print the rest of the stack until it is empty and so
the next value popped is 0, at which point "@" ends the program.
