const form = document.querySelector("#form-contact-container form");

form.addEventListener("submit", async event => {
    event.preventDefault();

    const inputName = document.querySelector("#form-contact-container form input[type='text']");
    const inputEmail = document.querySelector("#form-contact-container form input[type='email']");
    const inputContent = document.querySelector("#form-contact-container form textarea");

    const data = await fetch("https://api.chucknorris.io/jokes/random");

    if (data.ok) {
        alert("Enviado, esponderemos assim que possível!");
        inputName.value = inputEmail.value = inputContent.value = "";
    }else{
        alert("Erro ao enviar!");
    }

});