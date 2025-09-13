export async function handlerLogin(event) {
    try {
        let inputEmail = document.querySelector("#container-form-login input[type='email']");
        let inputPasswd = document.querySelector("#container-form-login input[type='password']");

        if (inputEmail.value.trim() == "" || inputPasswd.value.trim() == "") {
            return {error: "Dados vazios!"};
        }

        const data = await fetch("https://mindmiles.onrender.com/login", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                email: inputEmail.value,
                password: inputPasswd.value
            })
        });

        inputEmail.value = inputPasswd.value = "";
        return await data.json();
    } catch (error) {
        return error;
    }
}