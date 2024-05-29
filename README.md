# Online Bookstore

This application allows a user to interact with an online bookstore. The following actions are supported:


- inventory actions
    - view all inventory `GET /inventory`
    - view inventory item `GET /inventory/{sku}`
- user actions
    - view all users `GET /users`
    - view a specific user `GET /users/{id}`
    - add a user `POST /users`
    - view a users cart `GET /users/{id}/cart`
    - add items to a users cart `PATCH /users/{id}/cart`
- cart actions
    - view all carts `GET /carts`
    - view a specific cart `GET /carts/{id}`
    - checkout a specific cart `POST /carts/{id}/checkout`
- order actions
- Add items to cart
- Checkout and confirm purchase



## Technical Choices

- Implemented using REST API design to give a realistic view of how such an application might be designed
- Used Gorilla Mux to support query path params

