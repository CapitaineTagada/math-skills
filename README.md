# math-skills

## Overview
This Go program processes a file containing integers, performs various statistical calculations, and outputs the results. The calculations include:

- Average
- Median
- Variance
- Standard Deviation

## Project Structure
The project is divided into the following components:

1. **Main Program (`main.go`)**:
   - Reads integers from a file
   - Calculates statistical measures such as average, median, variance, and standard deviation
   - Prints the results to the console

2. **Shared Package (`shared`)**:
   - Provides shared global variables used across the program for calculations

3. **Utils Package (`utils`)**:
   - Provides utility functions such as reading files and checking errors

## Usage

### Prerequisites
- Go installed on your system (version 1.18 or later)
- A file containing integers separated by spaces or new lines

### Steps to Run

1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```

2. Navigate to the project directory:
   ```bash
   cd math-skills
   ```

3. Build the program:
   ```bash
   go build -o math-skills
   ```

4. Run the program:
   ```bash
   ./math-skills <path-to-your-file>
   ```
   Example:
   ```bash
   ./math-skills numbers.txt
   ```

### Output
The program outputs the following statistics:
- **Average**: The mean of the integers.
- **Median**: The middle value of the sorted integers.
- **Variance**: The measure of spread in the integers.
- **Standard Deviation**: The square root of the variance.

## Example Input File
`numbers.txt`:
```
1 2 3 4 5 6 7 8 9
```

### Example Output
```bash
Average: 5
Median: 5
Variance: 6
Standard Deviation: 2
```

## Code Overview

### Main Functions

- `Median(ints []int) float64`:
  - Calculates the median of a sorted slice of integers.

- `Variance(ints []int, average float64) float64`:
  - Calculates the variance by summing the squared differences from the mean.

- `utils.CheckArgs()`:
  - Validates command-line arguments.

- `utils.ReadFile(filename string) string`:
  - Reads the content of a file.

### Shared Variables
Defined in the `shared` package:
- `Average` (int): Stores the calculated average.
- `Median` (int): Stores the calculated median.
- `Variance` (int): Stores the calculated variance.
- `Deviation` (int): Stores the calculated standard deviation.

## License
This project is licensed under the GPL 3.0 License.
