const container = document.querySelector(".container1"),
      pwShowHide = document.querySelectorAll(".showHidePw"),
      pwFields = document.querySelectorAll(".password"),
      signUp = document.querySelector(".signup-link"),
      login = document.querySelector(".login-link");

    //   js code to show/hide password and change icon
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

    // js code to appear signup and login form
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
            // Redirect to admin dashboard or perform actions for admin users
        } else if (result.user_type === "Paciente") {
            window.location.href='/public/dashboard.html'
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

document.getElementById("Reserva").addEventListener("submit", function(e) {
    e.preventDefault();

    const Descrip2 = document.getElementById("Descrip").value;
    const Esp = document.getElementById("especialidades").value;
    const List_M = document.getElementById("listMedicos").value;
    const fechas = document.getElementById("fecha1").value;
    const Horaa = document.getElementById("hora").value;

    const data = {
        paciente: username,
        Descripcion: Descrip2 ,
        Esecialidad: Esp,
        Lista_Med: List_M,
        Fech1: fechas,
        Hor: Horaa
    };
    console.log(data)

    fetch("http://localhost:8080/Reservar", {
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
        
    })
    .catch(function(error) {
        console.error("Error:", error);
    });
});






