# 🧩 Sudoku Solver (Go)

This project is part of my **Piscine Go journey** at 01Talent × Nextera.  
It implements a backtracking-based solver for **9×9 Sudoku** with strict input validation.

---


## 🧠 Overview

- Input: **9 arguments**, each is a string of length **9**, containing digits `1..9` or `.` for empty cells.  
- Output: prints the solved grid as **9 lines**, numbers separated by a **single space**, no extra spaces at line end.  
- On any invalid input or unsolvable Sudoku → prints **`Error`**.

A valid Sudoku has only one solution; this program solves a valid one or reports `Error` otherwise.

---


## 💻 Usage

**Example (valid Sudoku):**

**Input:**
```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
```
**Output:**
```
3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

**Input:**
```bash
go run . 1 2 3 4 | cat -e
```
**Output:**
```
Error$
```

**Input:**
```bash
go run . | cat -e
```
**Output:**
```
Error$
```
**Input:**
```bash
go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
```
**Output:**
```
Error$
```
---


## 🛠 Tech & Approach

- Language: Go
- Algorithm: Backtracking with safety checks (row/column/3×3 box)
- Input validation: length, allowed characters, and initial-state consistency

---


## 🌿 Skills Practiced

- Problem solving & recursion
- Input validation & error handling
- Clean code & modular functions
- CLI applications in Go
---

## 🚀 How to Run

1. **Clone the repository:**
```bash
git clone https://github.com/EssamGamal88/piscine-go.git
```
2. **Navigate to the Sudoku folder**
```bash
cd piscine-go/Sudoku
```
3. **Run**
```bash
go run .
```

---

## 🧠 What I Learned

This project was a great opportunity to apply **logical reasoning** and **algorithmic thinking** using Go.  
I learned how to:

- 🧩 Break down complex problems into smaller, solvable steps  
- 🔁 Implement **recursive algorithms** with **backtracking**  
- ⚙️ Validate input and handle errors gracefully  
- 🧠 Write clean, modular, and testable Go code  
- 🚀 Build and run real **CLI applications** with user input  

---
⭐ *If you liked this project or found it useful, feel free to star the repo and follow my journey on [GitHub](https://github.com/EssamGamal88)!*
