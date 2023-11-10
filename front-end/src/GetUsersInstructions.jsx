import {useState} from "react";

export const GetUsersInstructions = () => {
    const [showed, setShowed] = useState();

    const handleShow = async () => {
        const payload = {
            action: "get_users_instructions",
            id: {
                id: 1,
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
            <button onClick={handleShow}>show all instruction</button>

            {showed && (
                <div>
                    <h2>test</h2>
                    <h2>{showed.data[0].title}</h2>
                    <h2>{showed.data[0].description}</h2>
                </div>

            )}
        </>
    )

}