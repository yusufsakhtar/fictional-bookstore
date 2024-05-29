# Online Bookstore

This application allows a user to interact with an online bookstore.

## Running Instructions

Start the server by running `go run cmd/onlinestore/main.go`

The application is designed as a series of REST API endpoints. You'll need to execute HTTP calls via Curl or Postman to run the application (sample Curl commands provided).

The application is pre-seeded with example user, inventory, and cart data for ease of demo use.

- inventory actions
    - view all inventory `GET /inventory`
        - `curl -X GET http://localhost:8080/inventory`
    - view inventory item `GET /inventory/{sku}`
        - `curl -X GET http://localhost:8080/inventory/sku1`
- user actions
    - view all users `GET /users`
        - `curl -X GET http://localhost:8080/users`
    - view a specific user `GET /users/{id}`
        - `curl -X GET http://localhost:8080/users/user1`
    - add a user `POST /users`
        - `curl -X POST -H "Content-Type: application/json" -d '{"first_name":"John","last_name":"Doe", "email":"john.doe@example.com"}' http://localhost:8080/users `
    - view a users cart `GET /users/{id}/cart`
        - `curl -X GET http://localhost:8080/users/user1/cart`
    - add items to a users cart `PATCH /users/{id}/cart`
        - `curl -X PATCH -H "Content-Type: application/json" -d '["sku1","sku2"]' http://localhost:8080/users/user1/cart`
- cart actions
    - view all carts `GET /carts`
        - `curl -X GET http://localhost:8080/carts`
    - view a cart `GET /carts/{id}`
        - `curl -X GET http://localhost:8080/carts/cart1`
    - checkout a cart `POST /carts/{id}/checkout`
        - `curl -X POST http://localhost:8080/carts/cart1/checkout`
- order actions
    - view all orders `GET /orders`
        - `curl -X GET http://localhost:8080/orders`
    - view an order `GET /orders/{id}`
        - `curl -X GET http://localhost:8080/orders/8e7512ec-ad5d-41d0-aa21-18a0eb48ec5f`
        - NOTE: This command won't work as is, as order IDs are generated during checkout and not pre-seeded. You'll need to copy the order ID from the output of the `POST /carts/{id}/checkout` command
    - confirm an order `POST /orders/{id}/confirm`
        - `curl -X POST http://localhost:8080/orders/8e7512ec-ad5d-41d0-aa21-18a0eb48ec5f/confirm`

## High Level Flow
- bootup
    - seeding user, inventory, and cart data using static files for demo convenience
    - alternatively, create a user
- user views inventory
- user adds inventory item(s) to cart
    - details
        - user cart is created
    - future improvements
        - update the inventory item status to reflect that it's in X # of carts
- user checks out cart
    - details
        - an order is created
        - cart items' stock is updated to reflect pending status
    - future improvements
        - an async system running on a schedule makes items available again to other users after a certain period of time in cart
- user confirms their order
    - details
        - order items' stock is updated
        - order status is updated
        - associated cart is deleted
    - future improvements
        - change flow to
            - update order status to partially complete
            - charge a hold on user payment for calculated order total
            - update order items' stock
                - if any items are OOS
                    - update order status to pending remediation
                    - send message to order remediation service
                - if not
                    - send message to order fulfillment service
                    - update order status to pending fulfillment
            - fulfillment service
                - figures out shipping for each item
                - updates order status to pending final payment
            - order remediation service
                - determines which items are still available
                    - sends message to order fulfillment
                - updates order total accordingly
            - scheduled job queries for orders pending final payment
                - payment service transacts with payment provider
                - order status updated to complete
                - message generated for receipt service


## High Level Design
- Implemented using REST API design to give a realistic view of how such an application might be designed
- 'services' are an approximation for actual microservices in some cases
    - Processes that require multiple services (e.g. cart checkout, order confirm) would be performed in steps
        - each service would handle as much as it could while accessing only its own data store
        - communication between services would frequently be through queuing systems
- 'repos' represent the data access layer
    - the interfaces defined in `repository.go` are useful for later swapping out the data store and using something like sqlite
    - to save time, I didn't implement a service for each repo, but that would be the ideal pattern

## Low Level Design
- cart is modeled separately from user since an anonymous user can add items to a cart
- orders are modeled separately from cart since order history needs to be persisted but carts are ephemeral

## Technical Choices
- Used Gorilla Mux to support query path params
