import Navbar from "./components/Navbar";

export default function PaginaLogin(){

    const navbarLinks = [

        { url: "/", title: "Home" },
        { url: "/about", title: "Sobre nosotros" },
        { url: "/reservas", title: "Reservar" },
        { url: "/login", title: "Login" },
      ];

    return <div className="PaginaLogin">
        
        <Navbar navbarLinks={navbarLinks} />
        
        
        </div>
}