import React from "react";
import "./Hero.css";

const HeroAbout = ({ imageSrc }) => {
  return (
    <div className="hero">
      
      <h1 className="hero__titleAbout">Travel made simple.</h1>
      <img src={imageSrc} alt=" MyStay" className="hero__image" />
    </div>
  );
};

export default HeroAbout;