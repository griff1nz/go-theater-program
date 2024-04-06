# Theater Program

## Description
This simple CLI application allows the user to track the occupancy of a theater. The user can use actions randomize the seats, check if a seat is occupied, reserve two seats next to each other.

## Installation
Go 1.22.1 or later must be installed on the computer before running the program. If the proper Go version is installed, the program can be started entering typing ```go run .``` in the terminal.

## Usage
Upon starting the program using ```go run .```, the user will be prompted to enter the full names and heights of up to 60 customers. Once the maximum number of customers has been reached, or if the user chooses to not add any more customers, they will be provided a list of actions, including:
- Randomizing the seating arrangement
- Checking if a seat is occupied
- Finding the most occupied row
- Finding the tallest customer
- Finding customers who are seated directly behind someone more than 3 inches taller than them
- Reserving two seats next to each other
- Exiting the application.

## License
Please refer to the LICENSE in the repo.

## Credits
Explanation of bufio: https://stackoverflow.com/questions/34647039/how-to-use-fmt-scanln-read-from-a-string-separated-by-spaces   
Golang tutorial that I watched: https://www.youtube.com/watch?v=yyUHQIec83I

This was originally a project that I wrote in Java for AP Computer Science A, so thank you to my teacher Mrs. Stoudt for the project idea!