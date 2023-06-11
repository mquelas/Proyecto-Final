import React from "react";
import "./HeroAbout.css";

const HeroAbout = ({ imageSrc }) => {
  return (
    <div className="hero">
      <img src={imageSrc} alt=" MyStay" className="hero__image" />
      <h1 className="hero__titleAbout">Sobre Nosotros:</h1>
      <h1 className="descripcion">
      Descubre un mundo de excelencia en nuestra página de hoteles, donde el lujo, la elegancia y la comodidad se entrelazan 
      para brindarte una experiencia de alojamiento inigualable. Nuestra cuidada selección de propiedades, acompañada de un servicio personalizado, te garantiza momentos memorables. 
      Explora nuestras descripciones detalladas y fotografías cautivadoras para encontrar el lugar perfecto para tus necesidades y preferencias. Únete a nosotros en un viaje hacia la excelencia en hospitalidad, donde cada experiencia es cuidadosamente diseñada para superar tus expectativas. 
      Bienvenido a un mundo de lujo y sofisticación.
      </h1>
    </div>
  );
};

export default HeroAbout;