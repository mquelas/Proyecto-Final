import "./FormularioLogin.css";
import { useState } from "react";
import { FiMenu, FiX } from "react-icons/fi";
import { Link } from "react-router-dom";


export function FormularioLogin ({setUser}){

    const [nombre, setNombre] = useState("")
    const [contrasena, setContrasena] = useState("")
    const [error, setError] = useState(false)

    const handleSubmit = (e) =>{

        e.preventDefault()
        if (nombre == "" || contrasena == ""){
            setError(true)
            return
        }

        setError(false)

        setUser([nombre])

    }

    return(

        <section>

            <h1 className="Log">Login</h1>

            <form 
                className="formulario"
                onSubmit={handleSubmit}
            
            >

                <input  
                placeholder="Usuario"
                type="text"
                value={nombre}
                onChange={e => setNombre(e.target.value)}
                ></input>

                <input 
                placeholder="Contraseña"
                type="password"
                value={contrasena}
                onChange={e => setContrasena(e.target.value)}
                ></input>

                <button>Iniciar Sesión</button>

                <button className="botonRegistro">No tiene cuenta? Registrese!
                <li>
                    <Link className="navbar__link" to={"/register"}>
                        
                     </Link>
                 </li>
            </button>

                
            </form>
            {error && <p className="camposOb">Todos los campos son obligatorios</p>}
        </section>
    )
}