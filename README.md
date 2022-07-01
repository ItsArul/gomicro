<h1 align="center"> GoMicro </h1>
<div align="center">
    <img src="https://wallpaperaccess.com/full/869910.gif" />
</div>

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif"></p>


## Feature

| Feature             | Status              | 
|---------------------|---------------------|
| Authentication      | Ongoing             |
| Message Broker      | Ongoing             |
| Product             | :heavy_check_mark:  |
| Order               | :heavy_check_mark:  |
| Dockerize           | :heavy_check_mark:  |
| API Gateway         | Ongoing             |

## Endpoint

| Endpoint                  | Method      | Descriptin                    | 
|---------------------- ----|-------------|-------------------------------|
|`/api/product`             | *GET*       | Get all product               |
|`/api/product`             | *POST*      | Post product / upload product |
|`/api/product/:id`         | *GET*       | Get product by id             |
|`/api/product/:id`         | *PUT*       | Update product by id          |
|`/api/product/:category`   | *GET*       | Get product by category       |
|`/api/product/order`       | *POST*      | Order product                 |
|`/api/product/:id`         | *DELETE*    | Delete product by id          |


## Installation

```
cp .env.example .env
```
> Note : Edit file .env
```bash
chmod +x install
```
> Note : change permission file
```bash
install app
or 
install requirepment
```
> Note : if dont have docker , you can `install requirepment`


## License
Distributed under the MIT License. See [`LICENSE`](https://github.com/ItsArul/gomicro/blob/main/LICENSE) for more information