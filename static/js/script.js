document.addEventListener('DOMContentLoaded', () => {

    let scrollUpButton = document.querySelector('.roll-up');

    window.addEventListener('scroll', function (event) {
        let triggerScrollButton = document.querySelector('.first-section')
        let topTriggerScrollButton = triggerScrollButton.getBoundingClientRect().top
        let top = document.querySelector('.first-header')
        let topBottom = top.getBoundingClientRect().bottom
        let fixedHeader = document.querySelector('.second-header')
        let fixedHeaderHeight = fixedHeader.getBoundingClientRect().height
        let support = document.querySelector('.header-support')


        if (document.body.scrollTop >= topTriggerScrollButton) {
            scrollUpButton.classList.add('fixed');
            scrollUpButton.style.display = "block";
        } else {
            scrollUpButton.classList.remove('fixed');
            scrollUpButton.style.display = "none";
        }
        if (document.body.scrollTop >= topBottom) {
            support.style.cssText = `display: block; height: ` + fixedHeaderHeight + `px`;
            fixedHeader.classList.add('header-fixed');
        } else {
            fixedHeader.classList.remove('header-fixed');
            support.style.display = "none";

        }
    });
    scrollUpButton.onclick = function () {
        window.scrollTo(0, 0);
    };

    let field = document.querySelector('.field')
    window.addEventListener('load', () => {
        field.classList.add('hidden_field');
        setTimeout(() => {
            field.remove();
        }, 1000);
    })
    let logout = document.querySelector(".dropdown-item-logout")

    let login = document.querySelector(".dropdown-item-login")
    let allCookies = document.cookie;
    let nk = allCookies.split(";")
    for (let i =0; i < nk.length; i++){
        if (nk[i].trim().split("=")[0]==="User"){
            logout.classList.remove("item-hidden")
            login.classList.add("item-hidden")
            userGreeting.innerText = `Hi, ${nk[i].trim().split("=")[1]}!`
        }
    }


    let imgSrcArr = []
let suppliersImg = document.getElementsByClassName("suppliers-img")

        for (let i =0; i < suppliersImg.length; i++){
            imgSrcArr.push(suppliersImg[i].getAttribute("src"))
            suppliersImg[i] = new Image();
            console.log(suppliersImg[i].baseURI)
            suppliersImg[i].src = imgSrcArr[i];

            console.log(suppliersImg[i].src)

        }



});


