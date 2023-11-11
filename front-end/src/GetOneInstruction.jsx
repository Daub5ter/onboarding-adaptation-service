import {useState} from "react";

export const GetOneInstruction = () => {
    const [showed, setShowed] = useState();

    const handleShow = async () => {
        const payload = {
            action: "get_instruction",
            id: {
                id: 3,
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
            <button onClick={handleShow}>show one instruction</button>

            {showed && (
                <div>
                    <h2>{showed.data.title}</h2>
                    <h2>{showed.data.description}</h2>
                </div>

            )}
        </>
    )
}