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


