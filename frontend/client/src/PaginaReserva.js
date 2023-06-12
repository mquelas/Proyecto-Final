import React from "react";
import Navbar from "./components/Navbar";
import Slider from "./components/Slider";
import DatePicker from "./components/DatePicker";
import Hotel4 from "./assets/Hotel 4.jpg";
import Hotel3 from "./assets/Hotel 3.jpg";
import Hotel2 from "./assets/Hotel 2.jpg";
import Hotel1 from "./assets/Hotel 1.jpg";
import "./components/Hero";
import "./DatePicker.css";
import "./components/HeroAbout.css";
import "./components/Buttom.css";

export default function PaginaReserva() {
  const navbarLinks = [
    { url: "/", title: "Home" },
    { url: "/about", title: "Sobre nosotros" },
    { url: "/reservas", title: "Reservar" },
    { url: "/login", title: "Login" },
  ];

  return (
    <div className="PaginaReserva">
      <Navbar navbarLinks={navbarLinks} />
      <div className="date-container">
        <h1>Seleccione las fechas</h1>
        <DatePicker className="DatePicker" />
        <button>Enviar</button>
      </div>
      <Slider imageSrc={Hotel1} title={"Hotel 1"} />
      <button>Reservar</button>
      <Slider imageSrc={Hotel2} title={"Hotel 2"} />
      <button>Reservar</button>
      <Slider imageSrc={Hotel3} title={"Hotel 3"} />
      <button>Reservar</button>
      <Slider imageSrc={Hotel4} title={"Hotel 4"} />
      <button>Reservar</button>
    </div>
  );
}
