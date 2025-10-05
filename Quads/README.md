# 🧩 Quads Project (Go)

This project is part of my **Piscine Go journey** at [01Talent x Nextera],
focused on learning **loops**, **conditionals**, and **ASCII-based pattern generation** using the Go programming language.

---

## 🧠 Overview

The **Quads Project** includes five different functions (`QuadA` to `QuadE`) that draw rectangles with specific corner and border characters.  
Each function follows unique logic to print its corresponding pattern based on the provided width (`x`) and height (`y`).

If `x` or `y` are not positive integers, the function prints nothing.

This exercise helped reinforce **logical thinking**, **nested loops**, and **pattern recognition** in Go.

---

## 💻 Usage Examples

### QuadA
**Input:**
```go
piscine.QuadA(5, 3)
```
**Output:**
```
o---o
|   |
o---o
```

### QuadB
**Input:**
```go
piscine.QuadB(5, 3)
```
**Output:**
```
/***\
*   *
\***/
```

### QuadC
**Input:**
```go
piscine.QuadC(5, 3)
```
**Output:**
```
ABBBA
B   B
CBBBC
```

### QuadD
**Input:**
```go
piscine.QuadD(5, 3)
```
Output:
```
ABBBC
B   B
ABBBC
```

### QuadE
Input:
```go
piscine.QuadE(5, 3)
```
Output:
```
ABBBC
B   B
CBBBA
```
---

## 🧩 Skills Practiced

- Go basics (functions, loops, conditionals)
- Modular programming and file structure
- ASCII pattern generation
- Problem-solving and debugging
- Writing clean, reusable functions

---

## 🗂️ Project Structure

The project follows a simple and modular folder structure:  

```
Quads/
├── QuadA.go      
├── QuadB.go     
├── QuadC.go      
├── QuadD.go      
├── QuadE.go     
├── main.go       
├── go.mod       
└── go.sum       
```

Each file implements one specific **rectangle pattern logic**,  making it easy to test, maintain, and understand.

---


## 🚀 How to Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/EssamGamal88/piscine-go.git
   ```

2. **Navigate to the project folder:**
   ```bash
   cd piscine-go/Quads
   ```

3. **Run the main file:**
   ```bash
   go run .
   ```

---

## 🧠 What I Learned

This project was my **first step into Go programming** during the **01Talent x Nextera Piscine**.  
It helped me understand and apply core programming concepts, including:

- 🔁 How **loops** and **conditionals** interact to create patterns  
- 🧩 The importance of **modular** and **readable code**  
- 🧪 How to **test** and **debug** small components effectively  
- 💭 How logical structure improves code clarity and problem-solving  

---
⭐ *If you liked this project or found it useful, feel free to star the repo and follow my journey on [GitHub](https://github.com/EssamGamal88)!*

