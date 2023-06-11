import "./Home.css";

export function Home({ user }){
    return (
        <section>
            <h1 className="bienvenido">Bienvenido</h1>
            <h2 className="username">{user}</h2>
        </section>
    )

}