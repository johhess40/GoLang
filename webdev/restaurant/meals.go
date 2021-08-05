package main

import (
	"log"
	"os"
	"text/template"
)

type Menu struct {
	Dish   string
	Side   string
	Drinks []string
}

type Meals struct {
	Breakfast []Menu
	Lunch     []Menu
	Dinner    []Menu
}

type Restaurants struct {
	Restaurant map[string][]Meals
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("meals.gohtml"))
}

func main() {

	spots := Restaurants{
		map[string][]Meals{
			"TheEggAndI": {
				Meals{
					Breakfast: []Menu{
						{
							Dish: "Scrambled Eggs",
							Side: "Grits",
							Drinks: []string{
								"Coffee",
								"OJ",
								"Chocolate Milk",
							},
						},
						{
							Dish: "Eggs Benedict",
							Side: "MixedFruit",
							Drinks: []string{
								"Coffee",
								"OJ",
								"ChocolateMilk",
							}},
					},
					Lunch: []Menu{
						{
							Dish: "Grilled Cheese",
							Side: "Chips",
							Drinks: []string{
								"Sprite",
								"Coke",
								"Water",
							},
						},
						{
							Dish: "Carne Asada",
							Side: "Rice",
							Drinks: []string{
								"Sprite",
								"Coke",
								"Water",
							}},
					},
					Dinner: []Menu{
						{
							Dish: "Filet Mignon",
							Side: "Sweet Potatoe",
							Drinks: []string{
								"Beer",
								"Whiskey",
								"Water",
							},
						},
						{
							Dish: "Tuna Steak",
							Side: "Asparagus",
							Drinks: []string{
								"Beer",
								"Whiskey",
								"Water",
							}},
					},
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, spots)
	if err != nil {
		log.Fatalln(err)
	}
}
