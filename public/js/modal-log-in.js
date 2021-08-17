const logIn = document.querySelector('.log-in');
const modal = document.querySelector('.modals');
const closeBtn = document.querySelector('.mod-btn');


logIn.addEventListener('click', () => {
	openModalDesktop();

});

modal.addEventListener('click', (e) => {
	if (e.target == modal) {
		closeModal();
	}
});

closeBtn.addEventListener('click', () => {
	closeModal();
});

const openModalDesktop = () => {
	modal.classList.add('is-open');

}

const closeModal = () => {
	modal.classList.remove('is-open');

}


