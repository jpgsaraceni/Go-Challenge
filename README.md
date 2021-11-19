# The Challange

Create a program to:

1. Receive a list of items (with unitary prices and amounts for each product).
2. Receive a list of emails (one email per person).
3. Devide equally the total price among the people.
3.1. (if the division is not exact, one person must pay the minimum amount more than the others so that the sum of what each person pays equals the sum of the prices times amounts).
4. Return a map with the keys being the person's email and the value how much they should pay.

## Status

I've created:

1. A `struct` to model each item, with its name, unitary price and amount
2. An Email type that is a slice of strings.
3. A List type that is a slice of items (structs).
4. A SplitBill method binded to the List type that recieves an Email type.

### SplitBill method

* Iterates over the List to calculate the price of each item (unitary price times amount);
* Sums the prices of each item;
* Divides the sum by the amount of people;
* If the avarage calculated for each person to pay is not an integer, adds the rest of the division to a random person from the list before dividing, to always have exact values.

## To do

* Receive input (list of items (with prices in cents) and list of emails);
* Convert the prices to int type;
* Refactor to receive prices as decimals or integers;
* Write unitary tests.
