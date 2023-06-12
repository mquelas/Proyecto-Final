import Navbar from "./components/Navbar";
import loginphoto from "./assets/loginphoto.jpg";
import { FormularioReg } from "./components/FormularioReg";
import {Home} from './components/Home';
import { useState } from "react";
import HeroLogin from "./components/HeroLogin";


export default function PaginaLogin(){

    const [user, setUser] = useState([])

    const navbarLinks = [

        { url: "/", title: "Home" },
        { url: "/about", title: "Sobre nosotros" },
        { url: "/reservas", title: "Reservar" },
        { url: "/login", title: "Login" },
      ];

      

    return <div className="PaginaLogin">
        
        <Navbar navbarLinks={navbarLinks} />
        <HeroLogin imageSrc={loginphoto} />

        {
            !user.length > 0
            ?<FormularioReg setUser={setUser}/>
            :<Home user={user}/>
        }

        
        </div>
}