var burger = document.querySelector('.navBurgerMenu');
var menu = document.querySelector('.navMenu');

burger.addEventListener('click', () => {
    burger.classList.toggle("change");
    menu.classList.toggle("navMenuActive"); 
});

