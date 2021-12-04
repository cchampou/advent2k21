# Advent of code 2021. Go edition.

You can read more about the Advent of Code on the [website](https://adventofcode.com/2021/about).

## Motivations

There is no language restriction, so I choose to do it Go this year. Reason
of this choice is that I like the language for being efficient, minimalistic in
its syntax, performant, and strongly typed. I learned a bit about this language,
but I really need to practice, in order to get more used to it.

## Code structure

I had no clue how to organize the code when I started the project, so I created 1
executable program in each "day" folder, then I switched to one main executable
at the repos root. The input data are raw pasted in "input" text files, in each
"day" folder.

To execute the program, call the root executable with one argument, the number
of the day.

```shell
go run . 4
```
*This will execute program for day4, using data in the file `./day4/input`*