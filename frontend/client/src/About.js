import Navbar from "./components/Navbar";
import HeroAbout from "./components/HeroAbout";
import aboutphoto from "./assets/aboutphoto.jpg";


export default function About(){

    const navbarLinks = [

        { url: "/", title: "Home" },
        { url: "/about", title: "Sobre nosotros" },
        { url: "/reservas", title: "Reservar" },
        { url: "/login", title: "Login" },
      ];

    return <div className="About">
        
        <Navbar navbarLinks={navbarLinks} />
        <HeroAbout imageSrc={aboutphoto} />
        
        
        
        </div>
}