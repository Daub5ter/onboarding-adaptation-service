import {useState} from "react";

export const GetUserByID = () => {
    const [showed, setShowed] = useState();

    const handleShow = async () => {
        const payload = {
            action: "get_user_by_email",
            email: {
                email: "sysadmin@test.com",
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
            <button onClick={handleShow}>show user by id</button>

            {showed && (
                <div>
                    <h2>{showed.data.email}</h2>
                    <h2>{showed.data.first_name}</h2>
                </div>

            )}
        </>
    )
}