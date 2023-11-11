import { useState } from "react";
import {AuthUser} from "../Auth/AuthUser";
import {GetAllKnowledge} from "./GetAllKnowledge";
import {GetPercentKnowledge} from "./GetPercentKnowledge";

export const Onboarding = () => {
    const [logined, setLogined] = useState();
    const [knowledgeState, setKnowledge] = useState();
    const [percentState, setPercent] = useState();
    const [passState, setPass] = useState();
    const handleMain = async () => {
        const email = document.getElementById("login-email").value;
        const password = document.getElementById("login-password").value;

        const user = await AuthUser(email, password);
        if (user.data.id != null) {
            setLogined(user);
            const l = document.getElementById('login');
            l.style.display = 'none';

            const knowledge = await GetAllKnowledge(user.data.id);
            setKnowledge(knowledge);

            const percent = await GetPercentKnowledge(user.data.id);
            setPercent(percent);
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

            {knowledgeState && (
                    <div>
                        <h2>{knowledgeState.data[1].title}</h2>
                        <h2>{knowledgeState.data[1].description}</h2>
                        <p id="knowledge-id" style={{ display: 'none' }}>{knowledgeState.data[1].id}</p>
                    </div>
            )}

            {percentState && (
                <h1>percent = {percentState.data}%</h1>
            )}
        </>
    );
};