import {useState} from "react";

export const GetAllUsers = () => {
    const [showed, setShowed] = useState();

    const handleShow = async () => {
        const payload = {
            action: "get_all_user",
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
            <button onClick={handleShow}>show all users</button>

            {showed && (
                <div>
                    <h2>{showed.data[2].email}</h2>
                    <h2>{showed.data[2].first_name}</h2>
                </div>

            )}
        </>
    )
}