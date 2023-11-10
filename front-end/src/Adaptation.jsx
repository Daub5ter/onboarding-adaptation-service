import { useState } from "react";
import {AuthUser} from "./AuthUser";

export const Adaptation = () => {
    const [logined, setLogined] = useState();

    const handleMain = async () => {
        const email = document.getElementById("login-email").value;
        const password = document.getElementById("login-password").value;

        const userID = AuthUser(email, password);
        if (userID !== undefined) {
            setLogined(userID);
        }
    };
    return (
        <>
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

            <button onClick={handleMain}>Login</button>

            {logined && (
                <h2>Successful logined</h2>
            )}
        </>
    );
};