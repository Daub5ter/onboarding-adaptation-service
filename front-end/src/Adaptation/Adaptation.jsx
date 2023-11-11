import { useState } from "react";
import {AuthUser} from "../Auth/AuthUser";
import {GetPercentInstructions} from "./GetPercentInstructions";
import {GetUsersInstructions} from "./GetUsersInstructions";

export const Adaptation = () => {
    const [logined, setLogined] = useState();
    const [instructionsState, setInstructions] = useState();
    const [percentState, setPercent] = useState();

    const handleMain = async () => {
        const email = document.getElementById("login-email").value;
        const password = document.getElementById("login-password").value;

        const user = await AuthUser(email, password);
        if (user.data.id != null) {
            setLogined(user);
            const l = document.getElementById('login');
            l.style.display = 'none';

            if (user.data.profession === "employer") {
                const instructions = await GetUsersInstructions(user.data.id);
                setInstructions(instructions);

                const percent = await GetPercentInstructions(user.data.id);
                setPercent(percent);
            }
        }
    };
    return (
        <>
            <div id="login" style={{ display: 'block' }}>
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
                <button onClick={handleMain}>Login</button>
            </div>

            {logined && (
                <div>
                    <h2>Successful logined</h2>
                    <p id="user-id" style={{ display: 'none' }}>{logined.data.id}</p>
                </div>
            )}

            {instructionsState && (
                <div>
                    <h2>{instructionsState.data[1].title}</h2>
                    <h2>{instructionsState.data[1].description}</h2>
                    <p id="knowledge-id" style={{ display: 'none' }}>{instructionsState.data[1].id}</p>
                </div>
            )}

            {percentState && (
                <h1>percent = {percentState.data}%</h1>
            )}
        </>
    );
};