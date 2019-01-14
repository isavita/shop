# Shop checkout
Implementation of [checkout task](https://gist.github.com/gustavolobo/7a7753a49779955a8e36d60c81f8f765) in Go.

## Setup and run

  * Compile the project with `go install`
  * Run the shop with `go run shop.go`
  * Run the tests with `go test ./...`

## Design (packages)

  * **promotion package** - defines promoter intefaces with two methods (`apply?` and `discount`) that each type of
    promotion (**GeneralPromotion** and **ProductPromotion**) should implement.
  * **product package** - defines basic struct for **Product** and **Products**.
  * **checkout package** - defines basic struct for **Checkout** with **Scan** and **Total** methods.
  * **shop.go** - simple file that uses the promotions and all of the scanned items to calculate the
    total amount that the client needs to pay.

