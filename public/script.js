
const wrapper = document.querySelector('.wrapper');
const loginLink = document.querySelector('.login-link');
const registerLink = document.querySelector('.register-link');
const btnPopup=document.querySelector('.btnLogin-popup');
const iconClose=document.querySelector('.icon-close');

registerLink.addEventListener('click',()=>{
    wrapper.classList.add('active');
});

loginLink.addEventListener('click',()=>{
    wrapper.classList.remove('active');
});

btnPopup.addEventListener('click',()=>{
    wrapper.classList.add('active-popup');
});

iconClose.addEventListener('click',()=>{
    wrapper.classList.remove('active-popup');
});

document.getElementById("login").addEventListener("submit", function(event) {
    event.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const data = {
        username: username,
        password: password
    };

    fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    .then(function(response) {
        if (!response.ok) {
            throw new Error(response.statusText);
        }
        return response.json();
    })
    .then(function(result) {
        if (result.user_type === "admin") {
            // Redirect to admin dashboard or perform actions for admin users
        } else if (result.user_type === "Paciente") {
            window.location.href='/public/menupaciente.html'
        } else if (result.user_type === "Doctor") {
            window.location.href='/public/menudoctor.html'
        } else {
            // Handle unknown user types or show error message
        }
    })
    .catch(function(error) {
        console.error("Error:", error);
    });
});
document.getElementById("register").addEventListener("submit", function(e) {
    e.preventDefault();

    const fullname = document.getElementById("Full-Name").value;
    const email = document.getElementById("Email").value;
    const username1 = document.getElementById("user").value;
    const password1 = document.getElementById("pass").value;
    const userType = document.getElementById("user-type").value;
    const data = {
        fullname : fullname,
        email: email,
        username: username1,
        password: password1,
        user_type: userType
    };
    console.log(data)

    fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    .then(function(response) {
        if (!response.ok) {
            throw new Error(response.statusText);
        }
        return response.json();
    })
    .then(function(result) {
        // Handle registration success or show error message
    })
    .catch(function(error) {
        console.error("Error:", error);
    });
});


		