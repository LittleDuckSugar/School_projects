
window.onload = function () {
    let button = document.querySelector('.btn-light')
    console.log(button)
    button.addEventListener('click', tag)


    function tag() {
        let li = ""
        let inputHidden = document.querySelector('.hidden').value
        tagImage = inputHidden.split(',')
        let select = document.querySelector('.inner')
        li = select.childNodes[0].childNodes

        li.forEach(element => {
            console.log(element)
            tag = element.childNodes[0].childNodes[1].innerHTML
            a = element.childNodes[0]
            
            tagImage.forEach(value => {
                if (value === tag){
                    a.setAttribute("aria-selected", "true")
                    element.classList.add("selected")
                }
            });
        });
    }
}

