const ul = document.querySelector("#header-container #links-tabs ul");
const menuButton = document.querySelector("#header-container #menu-container #menu-button");

menuButton.addEventListener("click", () => {
    if (ul.style.display == "none") {
        ul.style.display = "inline-flex";
        ul.classList.add("menu-tabs-links-show");
        ul.classList.remove("menu-tabs-links-hidden");
        menuButton.classList.add("menu-button-active");
    }else {
        ul.classList.remove("menu-tabs-links-show");
        ul.classList.add("menu-tabs-links-hidden");
        setTimeout(() => ul.style.display = "none", 250);
        menuButton.classList.remove("menu-button-active");
    }
});