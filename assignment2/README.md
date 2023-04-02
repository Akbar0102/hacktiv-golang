## Dokumentasi

#### Creating new order

<details>
 <summary><code>POST</code> <code><b>:</b></code> <code>localhost:8080/order</code></summary>
 ```
 {
    "orderedAt": "2023-04-02T16:45:51+07:00",
    "customerName": "Willi",
    "items": [
        {
            "itemCode": "C-220",
            "description": "Kabel HDMI",
            "quantity": 2
        },
        {
            "itemCode": "B-11",
            "description": "Power Bank",
            "quantity": 1
        }
    ]
}
 ```
</details>

#### Retrieve all orders

<details>
 <summary><code>GET</code> <code><b>:</b></code> <code>localhost:8080/orders</code></summary>
</details>

#### Update order by id

<details>
 <summary><code>PUT</code> <code><b>:</b></code> <code>localhost:8080/order/1</code></summary>
 ```
 {
    "customerName": "Serena",
    "items": [
        {
            "itemCode": "B-26",
            "description": "Mic",
            "quantity": 1
        }
    ]
}
 ```
</details>

#### Delete order by id

<details>
 <summary><code>DELETE</code> <code><b>:</b></code> <code>localhost:8080/order/1</code></summary>
</details>
