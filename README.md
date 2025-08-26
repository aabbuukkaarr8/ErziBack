 #ErziBack
1.  in the bash docker-compose down -v                                            
    docker-compose up --build
2. start your docker container 
3. for user register URl POST /user/register
4. Body for register {"username":,"email":, "password":}
5. for user login URL POST /user/login
6. for user login body {
   "email":"somemail@gmail.com",
   "password": "somepassword"
   }
7. for get products GET /products
8. for get by id GET /product/:id
9. for add product you need add auth token for header " Bearer "your token" "
10. URL for add product /products/create (only for admins)
11. body for add product {
    "title": "Product 1",
    "description": "Описание товара",
    "price": 99.99,
    "quantity": 50,
    "category": "equipment"
    }
12. adasdasdasdasda
