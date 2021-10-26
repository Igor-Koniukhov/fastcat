document.addEventListener('DOMContentLoaded', () => {
    const productsBtn = document.querySelectorAll('.product__btn');
    const cartProductsList = document.querySelector('.cart-content__list');
    const cart = document.querySelector('.cart');
    const cartQuantity = cart.querySelector('.cart__quantity');
    const fullPrice = document.querySelector('.fullprice');
    const orderModalOpenProd = document.querySelector('.order-modal__btn');
    const orderModalList = document.querySelector('.order-modal__list');
    let price = 0.00;
    let randomId = 0;
    let productArray = [];


    const stringWithoutSpaces = (str) => {
        return str.replace(/\s/g, '');
    };

    const normalPrice = (str) => {
        return String(str).replace(/(\d)(?=(\d\d\d)+([^\d]|$))/g, '$1 ');
    };

    function roundPlus(x, n) { //x - число, n - количество знаков
        if (isNaN(x) || isNaN(n)) return false;
        let m = Math.pow(10, n);
        return Math.round(x * m) / m;
    }

    const plusFullPrice = (currentPrice) => {
        return  price += currentPrice;

    };

    const minusFullPrice = (currentPrice) => {
         return price -= currentPrice;

    };

    const printQuantity = () => {
        let productsListLength = cartProductsList.querySelector('.simplebar-content').children.length;
        cartQuantity.textContent = productsListLength;
        productsListLength > 0 ? cart.classList.add('active') : cart.classList.remove('active');
    };

    const printFullPrice = () => {
        fullPrice.textContent = `${price.toFixed(2)} `;
    };

    const generateCartProduct = (img, title, price, id, idProd, idSupp) => {
        return `
			<li class="cart-content__item">
				<article class="cart-content__product cart-product" data-id="${id}">
					<img src="${img}" alt="" class="cart-product__img">
					<div class="cart-product__text">
						<h3 class="cart-product__title">${title}</h3>
						<span class="cart-product__price">${normalPrice(price)}</span>
						<span class="product-id" hidden>${idProd}</span>
						<span class="supplier-id" hidden>${idSupp}</span>												
					</div>
					<button class="cart-product__delete" aria-label="Delete order"></button>
				</article>
			</li>
		`;
    };

    const deleteProducts = (productParent) => {
        let id = productParent.querySelector('.cart-product').dataset.id;
        document.querySelector(`.product[data-id="${id}"]`).querySelector('.product__btn').disabled = false;

        let currentPrice = parseFloat(stringWithoutSpaces(productParent.querySelector('.cart-product__price').innerHTML));

        minusFullPrice(currentPrice);
        printFullPrice();
        productParent.remove();

        printQuantity();
        updateStorage();
    };

    productsBtn.forEach(el => {

        el.closest('.product').setAttribute('data-id', randomId++);

        el.addEventListener('click', (e) => {

            let self = e.currentTarget;
            let parent = self.closest('.product');
            let id = parent.dataset.id;
            let img = parent.querySelector('.image-switch__img img').getAttribute('src');
            let title = parent.querySelector('.product__title').textContent;
            let idProd = parent.querySelector('.product-id').textContent;
            let idSupp = parent.querySelector('.supplier-id').textContent;
            let priceString = stringWithoutSpaces(parent.querySelector('.product-price__current').textContent);
            let priceNumber = parseFloat(parent.querySelector('.product-price__current').textContent);

            plusFullPrice(priceNumber);
            printFullPrice();

            cartProductsList.querySelector('.simplebar-content')
                .insertAdjacentHTML('afterbegin', generateCartProduct(img, title, priceString, id, idProd, idSupp));
            printQuantity();
            updateStorage();
            self.disabled = true;
        });
    });


    cartProductsList.addEventListener('click', (e) => {
        if (e.target.classList.contains('cart-product__delete')) {
            console.log('click')
            deleteProducts(e.target.closest('.cart-content__item'));
        }
    });
    let flag = 0;
    orderModalOpenProd.addEventListener('click', (e) => {
        if (flag == 0) {
            orderModalOpenProd.classList.add('open');
            orderModalList.style.display = 'block';
            flag = 1;
        } else {
            orderModalOpenProd.classList.remove('open');
            orderModalList.style.display = 'none';
            flag = 0;
        }
    });

    const generateModalProduct = (img, title, price, id, idProd, idSupp, prodInfo) => {
        return `
			<div class="form-group order-modal__item">
				<article class="order-modal__product order-product" data-id="${id}">				
					<img src="${img}" alt="" class="order-product__img">
					
					<div class="order-product__text row">					
					<div class="col-md-6 col-sm-12">
					<input type="text" class="order-product__title"  value="${title.trim()}" name="prodName" readonly >
                    </div>
                    <div class="col-md-3 col-sm-12">
                    <input type="number" class="order-product__price" value="${normalPrice(price)}" name="price" readonly >
                    </div>                   
                    <input type="number" class="product-id" value='${idProd}' name="idProd" readonly hidden>             
                    <input type="number" class="supplier-id" value='${idSupp}' name="idSupp" readonly hidden>
                    <input type="text" class="prod-info" value= '${prodInfo}' name="prodInfo" readonly hidden>                   
                    
                    <button type="button" class="order-product__delete product__btn btn btn-success btn-small col-3" hidden disabled>remove</button>
                                       						
                    </div>						
				</article>
			</div>
		`;
    };

    let prodInfo = () => {
        return stringWithoutSpaces(JSON.stringify(productArray))
    }


    const modalCart = new CartModal({
        isOpen: (modalCart) => {
            let array = cartProductsList.querySelector('.simplebar-content').children;
            let fullprice = fullPrice.textContent;
            let length = array.length;

            document.querySelector('.order-modal__quantity span').textContent = `${length} шт`;
            document.querySelector('.order-modal__summ input').value = `${fullprice}`;

            for (item of array) {
                let img = item.querySelector('.cart-product__img').getAttribute('src');
                let title = item.querySelector('.cart-product__title').textContent;
                let priceString = stringWithoutSpaces(item.querySelector('.cart-product__price').textContent);
                let idProdString = stringWithoutSpaces(item.querySelector('.product-id').textContent);
                let idSuppString = stringWithoutSpaces(item.querySelector('.supplier-id').textContent);
                let id = item.querySelector('.cart-product').dataset.id;

                let obj = {};
                obj.product_id = idProdString
                obj.supplier_id = idSuppString
                obj.title = stringWithoutSpaces(title);
                obj.price = priceString;
                productArray.push(obj);
                let pi = prodInfo()
                orderModalList.insertAdjacentHTML('afterbegin', generateModalProduct(
                    img,
                    title,
                    priceString,
                    id,
                    idProdString,
                    idSuppString,
                    pi
                ));
            }

        },
        isClose: () => {
            location.reload()
            console.log('closed');
        }
    });

    document.querySelector('.order__btn').addEventListener('click', (e) => {
        localStorage.removeItem('products');
    });

    document.querySelector('.order__btn').addEventListener('submit', (e) => {
        e.preventDefault();
        let self = e.currentTarget;

        let formData = new FormData(self);
        let name = self.querySelector('[name="name"]').value;
        let tel = self.querySelector('[name="phone"]').value;
        let mail = self.querySelector('[name="email"]').value;
        formData.append('goods', JSON.stringify(productArray));
        formData.append('name', name);
        formData.append('phone', tel);
        formData.append('email', mail);

        let xhr = new XMLHttpRequest();

        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    console.log('Sent');
                }
            }
        }

        xhr.open('POST', 'mail.php', true);
        xhr.send(formData);

        self.reset();
    });


    const countSumm = () => {
        document.querySelectorAll('.cart-content__item').forEach(el => {
            price += parseInt(stringWithoutSpaces(el.querySelector('.cart-product__price').textContent));
        });
    };

    const initialState = () => {
        if (localStorage.getItem('products') !== null) {
            cartProductsList.querySelector('.simplebar-content').innerHTML = localStorage.getItem('products');
            printQuantity();
            countSumm();
            printFullPrice();
            document.querySelectorAll('.cart-content__product').forEach(el => {
                let id = el.dataset.id;
                console.log(id)
                document.querySelector(`.product[data-id="${id}"]`).querySelector('.product__btn').disabled = false;
            });
        }
    };

    initialState();

    const updateStorage = () => {
        let parent = cartProductsList.querySelector('.simplebar-content');
        let html = parent.innerHTML;
        html = html.trim();

        if (html.length) {
            localStorage.setItem('products', html);
        } else {
            localStorage.removeItem('products');
        }
    };

    document.querySelector('.modal-cart').addEventListener('click', (e) => {

        if (e.target.classList.contains('order-product__delete')) {

            let id = e.target.closest('.order-modal__product').dataset.id;
            let cartProduct = document.querySelector(`.cart-content__product[data-id="${id}"]`).closest('.cart-content__item');
            deleteProducts(cartProduct)
            e.target.closest('.order-modal__product').remove();
            updateStorage();
        }
    });

    document.querySelector('.order__btn').addEventListener('click', (e) => {
        let productInfo = JSON.stringify(productArray)
        document.cookie = "Product=" + encodeURIComponent(productInfo) + ";path=/cart/create";
        localStorage.removeItem('products');
    });

});
