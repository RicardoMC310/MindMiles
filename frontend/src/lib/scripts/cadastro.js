export async function handlerCadastro(event) {
    try {
        let inputName = document.querySelector("#container-form-cadastro input[type='text']");
        let inputEmail = document.querySelector("#container-form-cadastro input[type='email']");
        let inputPasswd = document.querySelector("#container-form-cadastro input[type='password']");

        const data = await fetch("https://indiles-ricardomc3107728-byk1p9cq.leapcell.dev/createUser", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                name: inputName.value,
                email: inputEmail.value,
                password: inputPasswd.value,
                rules: "student"
            })
        });

        inputName.value = inputEmail.value = inputPasswd.value = "";
        return data;
    } catch (error) {
        return error;
    }
}