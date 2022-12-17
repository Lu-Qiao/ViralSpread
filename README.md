# ViralSpread
Final project for the course 02-601 Programming for Scientists at CMU.

# Link to demo video: 
https://drive.google.com/file/d/1TDpHYgrfGso4W-efcdu4ACn3N0jYVgJv/view?usp=sharing

## Objectives
1. Simulate viral spread caused by one type of virus with applying different types of antiviral therapies.
2. Iterative simulation of antiviral therapy effectiveness and find the critical effectiveness that could terminate the viral spread.
3. Simulate viral spread caused by two types of virus with applying different types of antiviral therapies.

## How to start:
1. Download and unzip "ViralSpread.zip".
2. Move folder "gowut" into your own folder "src".
3. In terminal, change directory to the unzipped folder ViralSpread and execute command "go build".
4. For Mac users, please enter "./ViralSpread" to start the program; For PC user, please enter "./ViralSpread.exe" to start the program.
5. A web page should pop out automatically, and choose the available windows for further simulation.
6. Enter the parameters and click the button to start the simulation.
7. Simulations in Window 1 and 3 will produce a GIF for the simulation and will also export a csv file that contains cell counts over time.
8. Simulations in Window 2 will export a csv file that contains cell counts over different drug effectiveness after 30 days of infection.

## General parameters:

### Width
The size of tissue is width by width, and the resolution for the output GIF is also width(pixel) by width(pixel)

### Mode of initiation
1. assign - assign one infectious cell at a certain position along the diagonal of the board
2. random - randomly spread n infectious cells across the board

### Number of infectious point:
For random initiation only.

### Initial position
For assign initiation only.

### Number of generations
The number of generations in this simulation.

### Time steps
The duration for each generation in days.

## Cell parameters:

### λ
Target cells are produced at a constant rate λ.

### ω
Viruses from infectious cells have a cell-to-cell transmission rate constant ω.

### dT
Target cells have an average lifetime of 1/dT.

### δ
Infectious cells die at a rate δ

### State of the cells
#### Single Virus Infection
1. Uninfected
2. Infected
3. Infectious
4. Dead

#### Double Virus Coinfection
1. Uninfected
2. Infected1
3. Infectious1
4. Infected2
5. Infectious2
6. Dead1
6. Dead2

## Virus parameters:

### Threshold
Critical virus concentration in target cell that could change its state from "Infected" to "Infectious".

### Rcap
Carrying capacity of virus in each cell.

### ⍺
Maximal replication rate ⍺ of virus

### 𝛾
Positive-strand virus RNA is degraded with rate 𝛾

### ⍴
Virus exporting rate.

## Treatment parameters:

### Treatment options
1. no
2. blockcell
3. blockvirus
4. blockboth

### Antiviral therapy effectiveness
1. Effectiveness of blocking cell-to-cell transmistion
2. Effectiveness of blocking virus replication
