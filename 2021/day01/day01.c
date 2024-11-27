// Day 01 of Advent of Code 2021
//
// Read a text file of numbers, report the number of increases,
// first between subsequent lines, then between subsequent sums
// of 3 lines.
//
// To compile this C program: gcc day01.c -o day01
// To run: ./day01 (requires data in text file day01.data)
//
// Andreas Kaempf, 01/01/2021

#include <stdio.h>
#include <stdlib.h>

int main()
{
    // Read data file into an array of numbers
    const int max_rows = 5000;          // maximum data points
    int data[max_rows];                 // array of data points
    FILE *f = fopen("input.txt", "r");  // open text file
    char buf[32];                       // buffer for reading each line
    int n = 0;                          // number of data points
    while ( fgets(buf, 32, f) ) {       // read each line
        data[n] = atoi(buf);            // parse integer
        if ( ++n > max_rows ) {
            puts("File too big!");
            return(0);
        }
    }
    printf("%d data points found\n", n);

    // Part 1: count up differences to previous point
    int i, increases = 0;               // counters
    for ( int i = 1; i < n; ++i ) {     // each pair
        if ( data[i] > data[i - 1] )    // increase from previous?
            ++increases;                // incrememt counter
    }

    // Report the number found
    printf("Part 1: %d increases found\n", increases);

    // Part 2: compare windows of 3
    increases = 0;                      // reset counter
    int sum, prev_sum = 0;              // sums of sliding windows
    for ( int i = 2; i < n; ++i ) {     // start of each triplet
        sum = data[i] + data[i - 1] + data[i - 2];
        if ( prev_sum > 0 && sum > prev_sum ) 
            ++increases;                // incrememt counter if increase
        prev_sum = sum;                 // remember this window
    }

    // Report the number found
    printf("Part 2: %d increases found\n", increases);
}

