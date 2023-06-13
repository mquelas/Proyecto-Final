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
import { useState } from "react";
import { useNavigate } from "react-router"

export default function PaginaReserva() {
  const navbarLinks = [
    { url: "/", title: "Home" },
    { url: "/about", title: "Sobre nosotros" },
    { url: "/reservas", title: "Reservar" },
    { url: "/misreservas", title: "Mis Reservas" },
    { url: "/login", title: "Login" },

  ];

  const navigate = useNavigate();
  
  return (
    <div className="PaginaReserva">
      <Navbar navbarLinks={navbarLinks} />
      <div className="date-container">
        <h1>Seleccione las fechas</h1>
        <DatePicker className="DatePicker" />
        <button>Enviar</button>
      </div>
      
      <div  className="Buttom-container">
        <Slider imageSrc={Hotel1} title={"Hotel 1"} />
        <button onClick={()=> navigate("/confirmacion")}>Reservar</button>
      </div>
      <div className="Buttom-container">
        <Slider imageSrc={Hotel2} title={"Hotel 2"} />
        <button onClick={()=> navigate("/confirmacion")}>Reservar</button>
      </div>
      <div className="Buttom-container">
        <Slider imageSrc={Hotel3} title={"Hotel 3"} />
        <button onClick={()=> navigate("/confirmacion")}>Reservar</button>
      </div>
      <div className="Buttom-container">
        <Slider imageSrc={Hotel4} title={"Hotel 4"} />
        <button onClick={()=> navigate("/confirmacion")}>Reservar</button>
      </div>
    </div>
  );
}
