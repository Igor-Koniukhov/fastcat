

let scrollUpButton = document.querySelector('.roll-up');

window.addEventListener('scroll', function (event) {
    let triggerScrollButton = document.querySelector('.first-section')
    let topTriggerScrollButton = triggerScrollButton.getBoundingClientRect().top

    if (document.body.scrollTop >= topTriggerScrollButton) {
        scrollUpButton.classList.add('fixed');
        scrollUpButton.style.display = "block";
    } else {
        scrollUpButton.classList.remove('fixed');
        scrollUpButton.style.display = "none";
    }
});
scrollUpButton.onclick = function () {
    window.scrollTo(0, 0);
};


