/**/

#include <stdio.h>
#include <string.h>
#define smax 20

int main() {

    int i=0, scount=0;
    char word[smax];
    printf("Please enter a string > ");
    fgets(word,smax,stdin);
    while (i < 18) {
        if ((word[i] == 'a') || (word[i] == 'e') ||
            (word[i] == 'i') || (word[i] == 'o') ||
            (word[i] == 'u') || (word[i] == 'y'))
           {
            scount=scount+1;
           }
    i=i+1;
    }
    printf("The number of syllables is %d.\n", scount);
    return 0;
}
