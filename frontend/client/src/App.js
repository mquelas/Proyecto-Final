import "./App.css";
import Home1 from "./assets/Home 1.jpg";
import Home2 from "./assets/Home 2.jpg";
import Home3 from "./assets/Home 3.jpg";
import ny from "./assets/ny.jpg";
import iceland from "./assets/iceland.jpg";
import Hero from "./components/Hero";
import Navbar from "./components/Navbar";
import Slider from "./components/Slider";
import {
  createBrowserRouter,
  RouterProvider,
}from "react-router-dom";
import About from "./About";
import PaginaReseva from "./PaginaReserva";
import PaginaLogin from "./PaginaLogin";


function App() {

  
  const navbarLinks = [

    { url: "/", title: "Home" },
    { url: "/about", title: "Sobre nosotros" },
    { url: "/reservas", title: "Reservar" },
    { url: "/login", title: "Login" },

  ];

  return (
    <div className="App">
      <Navbar navbarLinks={navbarLinks} />
      
      <Hero imageSrc={iceland} />
      <Slider
        imageSrc={Home2}
        title={"Experiencias Inolvidables"}
        subtitle={"Sumérgete en una experiencia inolvidable en nuestros hoteles de ensueño"}
        
      />
      <Slider
        imageSrc={Home3}
        title={"Lujo y comodidad en nuestros hoteles"}
        subtitle={"Un oasis de lujo y comodidad te espera en nuestros hoteles"}
        flipped={true}
      />
    </div>
  );
}

export default App;
