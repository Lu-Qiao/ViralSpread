# ViralSpread
Final project for the course 02:601 Programming for Scientists at CMU.

ViralSpread simulates the spread of virus in cells on 2D board based on the following parameters.

## Mode of initiation
1. assign - assign one infectious cell at a certain position along the diagonal of the board
2. random - randomly spread n infectious cells across the board

## State of the cells
### Single Virus Infection
1. Uninfected
2. Infected
3. Infectious
4. Dead
### Double Virus Coinfection
1. Uninfected
2. Infected1
3. Infectious1
4. Infected2
5. Infectious2
6. Dead1
6. Dead2

## Treatment options
1. no
2. blockcell
3. blockvirus
4. blockboth
