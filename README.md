# The Challange

Create a program to:

1. Receive a list of items (with unitary prices and amounts for each product).
2. Receive a list of emails (one email per person).
3. Devide equally the total price among the people.
3.1. (if the division is not exact, one person must pay the minimum amount more than the others so that the sum of what each person pays equals the sum of the prices times amounts).
4. Return a map with the keys being the person's email and the value how much they should pay.

## Run the program

You will need [Golang](https://go.dev/doc/install) installed locally.

Run

```shell
go run split/split.go
```

in your terminal.

This program does not accept input, so it will run with the values stored in variables in the `split.go` file. You may change them, as long as: the UnitPrice is a comma or period separated decimal (or whole) number string.

To run the automated tests, execute

```shell
go test
```

in your terminal, in the directories `/brlParser` and `/split`

## Status

I've created:

1. A `struct` to model each item, with its name, unitary price and amount
2. An Email type that is a slice of strings.
3. A List type that is a slice of items (structs).
4. A SplitBill method binded to the List type that recieves an Email type and returns a map with email/value owed key/value pairs.
5. A brlParser to receive value in BRL (X.XX, X,XX, X,X, X.X or X formats), parse to integer (value in cents) and parse back to R$X,XX format.
6. Automated tests.

### SplitBill method

* Iterates over the List to calculate the price of each item (unitary price times amount);
* Sums the prices of each item;
* Divides the sum by the amount of people;
* If the avarage calculated for each person to pay is not an integer, adds the x rest of the division to x random people (1 cent each) from the list after dividing the whole part, to always have exact values and the total sum be equal the sum of each individual payment.

## To do

* Improve test coverage (cover cases when division is not exact, repeated emails, and error handling);
