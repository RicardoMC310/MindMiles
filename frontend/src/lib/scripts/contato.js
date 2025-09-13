export async function handlerContact(event) {
    event.preventDefault();

    const data = await fetch("https://api.chucknorris.io/jokes/random");

    if (data.ok) {
        alert("Responderemos assim que possível!");
    }else{
        alert("Erro ao enviar formulário!");
    }

    const inputName = document.querySelector("#form-container-contact input[type='text']");
    const inputEmail = document.querySelector("#form-container-contact input[type='email']");
    const inputContent = document.querySelector("#form-container-contact textarea");

    inputContent.value = inputEmail.value = inputName.value = "";
}