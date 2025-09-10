const ul = document.querySelector("#header-container #links-tabs ul");
const menuButton = document.querySelector("#header-container #menu-container #menu-button");

ul.style.display = "none";

menuButton.addEventListener("click", () => {
    if (ul.style.display == "none") {
        ul.style.display = "inline-flex";
        ul.classList.add("menu-tabs-links-show");
        ul.classList.remove("menu-tabs-links-hidden");
    } else {
        ul.classList.remove("menu-tabs-links-show");
        ul.classList.add("menu-tabs-links-hidden");
        setTimeout(() => ul.style.display = "none", 250);
    }
});