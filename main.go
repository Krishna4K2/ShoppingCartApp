package main

import (
	"fmt"
)

type (
	product struct {
		name  string
		Id    int
		price float64
	}
	cartItem struct {
		product  product
		quantity int
	}
)

var (
	products = []product{
		{name: "Apple", Id: 1, price: 0.5},
		{name: "Banana", Id: 2, price: 0.3},
		{name: "Orange", Id: 3, price: 0.8},
		{name: "Grapes", Id: 4, price: 2.0},
		{name: "Mango", Id: 5, price: 1.5},
		{name: "Pineapple", Id: 6, price: 3.0},
		{name: "Strawberry", Id: 7, price: 2.5},
		{name: "Blueberry", Id: 8, price: 4.0},
		{name: "Watermelon", Id: 9, price: 1.0},
		{name: "Peach", Id: 10, price: 1.2},
	}
	cart []cartItem
)

func main() {
	var (
		option   int
		shopping string
	)

	fmt.Println("===================================================")
	fmt.Println("\tWelcome to Shopping Cart APP\t")
	fmt.Println("===================================================")
	for shopping != "Exit" {
		fmt.Println("===================================================")
		fmt.Println("\t\tMain Menu\t\t")
		fmt.Print("1. List Products\n2.Add Product to cart\n3.Remove Product From Cart\n4.View Cart\n5.Checkout\n6.Exit\n")
		fmt.Print("Please select an option: ")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("Listing all products...")
			for _, p := range products {
				fmt.Println("Name:", p.name, "| ID:", p.Id, "| Price: $", p.price)
			}
		case 2:
			fmt.Println("Add product to cart...")
			var productId int
			fmt.Print("Enter the product ID to add to cart: ")
			fmt.Scanln(&productId)
			found := false
			for i, item := range cart {
				if item.product.Id == productId {
					cart[i].quantity++
					fmt.Printf("Increased quantity of %s to %d.\n", item.product.name, cart[i].quantity)
					found = true
					break
				}
			}
			if !found {
				for _, p := range products {
					if p.Id == productId {
						cart = append(cart, cartItem{product: p, quantity: 1})
						fmt.Printf("Added %s to cart.\n", p.name)
						found = true
						break
					}
				}
			}
			if !found {
				fmt.Println("Product not found. Please try again.")
			}
		case 3:
			fmt.Println("Remove product from cart...")
			var productId int
			fmt.Print("Enter the product ID to remove from cart: ")
			fmt.Scanln(&productId)
			found := false
			for i, p := range cart {
				if p.product.Id == productId {
					if cart[i].quantity > 1 {
						cart[i].quantity--
						fmt.Printf("Decreased quantity of %s to %d.\n", p.product.name, cart[i].quantity)
					} else {
						cart = append(cart[:i], cart[i+1:]...)
						fmt.Println("Product removed from cart:", p.product.name)
					}
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Product not found in cart. Please try again.")
			}
		case 4:
			fmt.Println("View cart...")
			if len(cart) == 0 {
				fmt.Println("Your cart is empty.")
			} else {
				fmt.Println("Your cart contains:")
				var total float64
				for _, p := range cart {
					fmt.Println("Name:", p.product.name, "| ID:", p.product.Id, "| Price: $", p.product.price, "| Quantity:", p.quantity)
					total += p.product.price * float64(p.quantity)
				}
				fmt.Printf("Your total amount is: $%.2f\n", total)
			}
		case 5:
			fmt.Println("Check out...")
			if len(cart) == 0 {
				fmt.Println("Your cart is empty. Please add products to your cart before checking out.")
			} else {
				var total float64
				fmt.Println("Calculating total price...")
				for _, p := range cart {
					total += p.product.price * float64(p.quantity)
				}
				fmt.Printf("Your total amount is: $%.2f\n", total)
				cart = nil // Clear the cart after checkout
				fmt.Println("Thank you for your purchase!")
			}
		case 6:
			fmt.Println("Exiting the application...")
			shopping = "Exit"
			fmt.Println("Thank you for Shopping! Have a great day!")
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}
