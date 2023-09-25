let copyright = document.querySelector(".isVisible-js")
let rightOfUse = document.querySelector(".rightOfUse-js")
console.log(copyright.children)
rightOfUse.onchange = function(){
    copyright.style.display = (this.selectedIndex == 0) ? "block" : "none";
    if(rightOfUse.value == "oui"){
      copyright.children[1].setAttribute("required", "")
    }else{
      copyright.children[1].removeAttribute("required")
    }
  }
