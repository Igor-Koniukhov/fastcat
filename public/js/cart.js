let cart = {
    'sisls23': {
        "name" : "gamburger",
        "count":2,
        "price": 5,
    },
    'pqmry28': {
        "name" : "gamburger",
        "count":2,
        "price": 3,
    },
};

// increase goods

// decrease goods

// delete goods

document.onclick = event => {

    if (event.target.classList.contains('plus')) {
        plusFunction(event.target.dataset.id);
    }
    if (event.target.classList.contains('minus')) {
        minusFunction(event.target.dataset.id);
    }
}

const plusFunction = id => {
    cart[id]['count']++;
    renderCart(id);
}
const minusFunction = id => {
    if (cart[id]['count'] - 1 == 0) {
        deleteFunction(id);
        return true;
    }
    cart[id]['count']--;
    renderCart(id);
}

const deleteFunction = id => {
    delete cart[id];
    renderCart(id);

}




const renderCart = (id) => {
    console.log(cart.indexOf())
    let quantity = document.getElementsByClassName('quantity')
    console.log(quantity[id])
    let productName = document.getElementsByClassName('product')
    let productPrice = document.getElementsByClassName('price')

    quantity.value = cart[id]['count'];
    productName.value = cart[id]['name'];
    productPrice.value = cart[id]['price'] * cart[id]['count'];

    console.log(cart);
}

