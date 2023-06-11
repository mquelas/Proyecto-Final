import "./App.css";
import travel_01 from "./assets/travel-01.jpg";
import travel_02 from "./assets/travel-02.jpg";
import travel_03 from "./assets/travel-03.jpg";
import Hero from "./components/Hero";
import Navbar from "./components/Navbar";
import Slider from "./components/Slider";
import React, { Component }  from 'react';
import {
  createBrowserRouter,
  RouterProvider,
}from "react-router-dom";
import About from "./components/pages/About";
import PaginaReseva from "./components/pages/PaginaReserva";
import Home from "./components/pages/Home";

function App() {

  let component

  // switch(window.location.pathname){

  //     case "/home":
  //       component = <Home/>
  //     break
  //     case "/about":
  //       component = <About/>
  //     break
  //     case "/paginaReserva":
  //       component = <PaginaReseva/>
  //     break

  // }

  // const router = createBrowserRouter([
  //   {
  //     path: "/Reservasfront.js",
  //     element: <div>Hello</div>
  //   }


  
  
  // ]);
    
  const navbarLinks = [

    { url: "#", title: "Home" },
    { url: "#", title: "About" },
    { url: "#", title: "Reservar" },
  ];

  return (
    <div className="App">
      <Navbar navbarLinks={navbarLinks} />
      {component}
      <Hero imageSrc={travel_01} />
      <Slider
        imageSrc={travel_02}
        title={"Be an explorer."}
        subtitle={
          "Our platform offers a wide variety of unique travel locations!"
        }
      />
      <Slider
        imageSrc={travel_03}
        title={"Memories for a lifetime."}
        subtitle={"Your dream vacation is only a few clicks away."}
        flipped={true}
      />
    </div>
  );
}

export default App;
