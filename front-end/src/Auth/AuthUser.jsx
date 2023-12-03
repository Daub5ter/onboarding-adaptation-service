import {useState} from "react";

export const AuthUser = (email, password) => {
    const [data, setData] = useState([]);
    const handleShow = (email, password) => {
        const payload = {
            action: "auth_user",
            auth: {
                email: email,
                password: password,
            }
        }

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        /*const response = await fetch("http:\/\/localhost:8080/handle", {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        });*/



        const fetchInfo = () => {
            return fetch("http:\/\/localhost:8080/handle", {
                method: 'POST',
                body: JSON.stringify(payload),
                headers: headers,
            }).then((res) => res.json()).then((d) => setData(d))
        }

        return data
    }

    return handleShow(email, password);
       /* <>
            <div>
                <h1>Login</h1>
                <input
                    type="text"
                    placeholder="Type email"
                    id="login-email"
                />
                <input
                    type="text"
                    placeholder="Type password"
                    id="login-password"
                />
            </div>

            <button onClick={handleShow}>Login</button>

            {showed && (
                <div>
                    <label>logged</label>
                </div>
            )}
        </>
    )*/
}