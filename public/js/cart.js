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



let quantity = document.getElementById('quantity')
let productName = document.getElementById('product')
let productPrice = document.getElementById('price')
const renderCart = (id) => {
    quantity.value = cart[id]['count'];
    productName.value = cart[id]['name'];
    productPrice.value = cart[id]['price'] * cart[id]['count'];

    console.log(cart);
}

