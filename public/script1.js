const container = document.querySelector(".container1"),
      pwShowHide = document.querySelectorAll(".showHidePw"),
      pwFields = document.querySelectorAll(".password"),
      signUp = document.querySelector(".signup-link"),
      login = document.querySelector(".login-link");

    
    pwShowHide.forEach(eyeIcon =>{
        eyeIcon.addEventListener("click", ()=>{
            pwFields.forEach(pwField =>{
                if(pwField.type ==="password"){
                    pwField.type = "text";

                    pwShowHide.forEach(icon =>{
                        icon.classList.replace("uil-eye-slash", "uil-eye");
                    })
                }else{
                    pwField.type = "password";

                    pwShowHide.forEach(icon =>{
                        icon.classList.replace("uil-eye", "uil-eye-slash");
                    })
                }
            }) 
        })
    })

    // 
    signUp.addEventListener("click", ( )=>{
        container.classList.add("active");
    });
    login.addEventListener("click", ( )=>{
        container.classList.remove("active");
    });


const body = document.querySelector("body"),
    sidebar = body.querySelector("nav");
    sidebarToggle = body.querySelector(".sidebar-toggle");



sidebarToggle.addEventListener("click", () => {
    sidebar.classList.toggle("close");
    if(sidebar.classList.contains("close")){
        localStorage.setItem("status", "close");
    }else{
        localStorage.setItem("status", "open");
    }
})



document.getElementById("login").addEventListener("submit", function(event) {
    event.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    

    localStorage.setItem('nombreUsuario', username);
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
            
        } else if (result.user_type === "Paciente") {
            window.location.href='/public/dashboard.html'
        } else if (result.user_type === "Doctor") {
            window.location.href='/public/my_appointments.html'
        } else {
            Swal.fire({
                icon: 'error',
                title: 'Error',
                text: "Correo y contraseña incorrecta"
            })
        }
    })
    .catch(function(error) {
        Swal.fire({
            icon: 'error',
            title: 'Intente de nuevo',
            text: "Username o Contraseña incorrecta"
        })
    });
});








