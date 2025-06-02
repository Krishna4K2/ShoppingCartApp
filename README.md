# 🛒 Shopping Cart App (Go)

A beginner-friendly, shopping cart application written in **Golang**, designed to help new developers practice and apply **core Go concepts** through a realistic mini project.

---

## 📚 Project Overview

This app simulates a simplified shopping cart experience. Users can:
- Browse a list of available products
- Add items to their shopping cart
- Remove items from the cart
- View the current cart and total price
- Checkout to finalize their order

It mirrors the basic backend logic of an online store — without any web or database dependencies.

---

## 🎯 Learning Objectives

By building this project, you'll reinforce the following **Go fundamentals**:

- ✅ Basic syntax & structure  
- ✅ Variables and declarations (`var`, `:=`)  
- ✅ Data types: `string`, `int`, `float64`, `bool`  
- ✅ Arrays and slices (products and cart)  
- ✅ Maps (inventory or product lookup)  
- ✅ Structs (Product, CartItem)  
- ✅ Functions (modular logic like add/remove/view)  
- ✅ Loops (`for`, `range`)  
- ✅ Conditionals (`if`, `else`, `switch`)  
- ✅ Error handling (`if err != nil`)  
- ✅ Packages: `fmt`, `os`, `bufio`, `strings`, `strconv`

---

## 🏗️ Project Structure

```bash
shopping-cart/
├── main.go         # Main application logic
├── product.go      # Product struct and product list
├── cart.go         # Cart operations (add, remove, view)
└── README.md       # Project documentation
