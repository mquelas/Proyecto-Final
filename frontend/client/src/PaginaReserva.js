import Navbar from "./components/Navbar";



export default function PaginaReserva(){

    const navbarLinks = [

        { url: "/", title: "Home" },
        { url: "/about", title: "About" },
        { url: "/reservas", title: "Reservar" },
      ];

    return <div className="PaginaReserva">
        
        <Navbar navbarLinks={navbarLinks} />
        
        
        
        </div>
}