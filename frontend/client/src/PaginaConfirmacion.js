import Navbar from "./components/Navbar";
import "./DatePicker.css";
export default function PaginaConfirmacion(){
    
    
    const navbarLinks = [

        { url: "/", title: "Home" },
        { url: "/about", title: "Sobre nosotros" },
        { url: "/reservas", title: "Reservar" },
        { url: "/login", title: "Login" },
      ];
      
    return (<div className="paginaConfirmacion">

        <Navbar navbarLinks={navbarLinks} />
         
         
         
         
         
         
        <div className="hotel-container"><h1>Se ha registrado su reserva en el Hotel (id hotel) !!!</h1></div>

    </div>
    )

}