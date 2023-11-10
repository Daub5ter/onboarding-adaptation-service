import {useState} from "react";

export const AuthUser = (email, password) => {
    const handleShow = async (email, password) => {
        const payload = {
            action: "auth_user",
            auth: {
                email: email,
                password: password,
            }
        }

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const response = await fetch("http:\/\/localhost:8080/handle", {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        });

        return await response.json()
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