import {useState} from "react";

export const AuthUser = () => {
    const [showed, setShowed] = useState();

    const handleShow = async () => {
        const payload = {
            action: "auth_user",
            email: {
                email: "sysadmin@test.com",
                password: "sysadmin",
            }
        }

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const response = await fetch("http:\/\/localhost:8080/handle", {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        });

        const data = await response.json();

        setShowed(data);
    }


    return (
        <>
            <button onClick={handleShow}>auth user</button>

            {showed && (
                <div>
                    <h2>authed</h2>
                </div>

            )}
        </>
    )
}